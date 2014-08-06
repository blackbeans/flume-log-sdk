package consumer

import (
	"encoding/json"
	"flume-log-sdk/config"
	"fmt"
	"github.com/blackbeans/redigo/redis"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

type counter struct {
	lastSuccValue int64

	currSuccValue int64

	lastFailValue int64

	currFailValue int64
}

// 用于向flume中作为sink 通过thrift客户端写入日志

type SinkServer struct {
	redisPool       map[string][]*redis.Pool
	flumeClientPool []*FlumePoolLink
	isStop          bool
	monitorCount    counter
	business        string
}

func newSinkServer(business string, redisPool map[string][]*redis.Pool, flumePool []*FlumePoolLink) (server *SinkServer) {
	sinkserver := &SinkServer{business: business, redisPool: redisPool, flumeClientPool: flumePool}
	go sinkserver.monitorFlume()
	return sinkserver
}

func (self *SinkServer) monitorFlume() {
	for !self.isStop {
		time.Sleep(1 * time.Second)
		currSucc := self.monitorCount.currSuccValue
		currFail := self.monitorCount.currFailValue
		log.Printf("succ-send:%d,fail-send:%d",
			(currSucc - self.monitorCount.lastSuccValue),
			(currFail - self.monitorCount.lastFailValue))
		self.monitorCount.lastSuccValue = currSucc
		self.monitorCount.lastFailValue = currFail
	}
}

//启动pop
func (self *SinkServer) start() {

	self.isStop = false

	var count = 0
	for k, v := range self.redisPool {

		log.Println("start redis queueserver succ " + k)
		for _, pool := range v {
			count++

			go func(queuename string, pool *redis.Pool) {
				conn := pool.Get()
				defer pool.Release(conn)
				for !self.isStop {

					// log.Println("pool active count :", strconv.Itoa(pool.ActiveCount()))
					reply, err := conn.Do("LPOP", queuename)
					if nil != err || nil == reply {
						if nil != err {
							log.Printf("LPOP|FAIL|%s", err)
							conn.Close()
							conn = pool.Get()
						} else {
							time.Sleep(100 * time.Millisecond)
						}

						continue
					}

					resp := reply.([]byte)
					var cmd config.Command
					err = json.Unmarshal(resp, &cmd)

					if nil != err {
						log.Printf("command unmarshal fail ! %s | error:%s\n", resp, err.Error())
						continue
					}
					//
					momoid := cmd.Params["momoid"].(string)

					businessName := cmd.Params["businessName"].(string)

					action := cmd.Params["type"].(string)

					bodyContent := cmd.Params["body"]

					//将businessName 加入到body中
					bodyMap := bodyContent.(map[string]interface{})
					bodyMap["business_type"] = businessName

					body, err := json.Marshal(bodyContent)

					if nil != err {
						log.Printf("marshal log body fail %s", err.Error())
						continue
					}

					//这里需要优化一下body,需要采用其他的方式定义Body格式，写入

					// log.Printf("%s,%s,%s,%s", momoid, businessName, action, string(body))

					//启动处理任务
					go self.innerSend(momoid, businessName, action, string(body))

				}
			}(k, pool)
		}
	}

}

func (self *SinkServer) innerSend(momoid, businessName, action string, body string) {

	for i := 0; i < 3; i++ {
		pool := self.getFlumeClientPool(businessName, action)
		flumeclient, err := pool.Get(5 * time.Second)
		if nil != err || nil == flumeclient {
			log.Fatalf("log_sink|fail get flumeclient from pool|%s\n", body)
			continue
		}
		//拼装头部信息
		header := make(map[string]string, 1)
		header["businessName"] = businessName
		header["type"] = action

		//拼Body
		flumeBody := fmt.Sprintf("%s\t%s\t%s", momoid, action, body)
		err = flumeclient.Append(header, []byte(flumeBody))
		defer func() {
			if err := recover(); nil != err {
				//回收这个坏的连接
				pool.ReleaseBroken(flumeclient)
			} else {
				pool.Release(flumeclient)
			}
		}()

		if nil != err {
			atomic.AddInt64(&self.monitorCount.currFailValue, 1)
			log.Printf("send 2 flume fail %s \t err:%s\n", body, err.Error())

		} else {
			atomic.AddInt64(&self.monitorCount.currSuccValue, 1)
			if rand.Int()%100 == 0 {
				log.Println("trace|send 2 flume succ|%s|%s", flumeclient.HostPort(), flumeBody)
			}

			break
		}

	}
}

//仅供测试使用推送数据
func (self *SinkServer) testPushLog(queuename, logger string) {

	for _, v := range self.redisPool {
		for _, pool := range v {
			conn := pool.Get()
			defer pool.Release(conn)

			reply, err := conn.Do("RPUSH", queuename, logger)
			log.Printf("%s|err:%s", reply, err)
			break

		}
	}

}

func (self *SinkServer) stop() {
	self.isStop = true
	time.Sleep(5 * time.Second)

	//遍历所有的flumeclientlink，将当前Business从该链表中移除
	for _, v := range self.flumeClientPool {
		v.mutex.Lock()
		for e := v.businessLink.Back(); nil != e; e = e.Prev() {
			if e.Value.(string) == self.business {
				//将自己从该flumeclientpoollink种移除
				v.businessLink.Remove(e)
				break
			}
		}
		v.mutex.Unlock()
	}

}

func (self *SinkServer) getFlumeClientPool(businessName, action string) *flumeClientPool {

	//使用随机算法直接获得

	idx := rand.Intn(len(self.flumeClientPool))
	return self.flumeClientPool[idx].flumePool

}
