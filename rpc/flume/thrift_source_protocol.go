// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package flume

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type ThriftSourceProtocol interface {
	// Parameters:
	//  - Event
	Append(event *ThriftFlumeEvent) (r Status, err error)
	// Parameters:
	//  - Events
	AppendBatch(events []*ThriftFlumeEvent) (r Status, err error)
}

type ThriftSourceProtocolClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewThriftSourceProtocolClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ThriftSourceProtocolClient {
	return &ThriftSourceProtocolClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewThriftSourceProtocolClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ThriftSourceProtocolClient {
	return &ThriftSourceProtocolClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Event
func (p *ThriftSourceProtocolClient) Append(event *ThriftFlumeEvent) (r Status, err error) {
	if err = p.sendAppend(event); err != nil {
		return
	}
	return p.recvAppend()
}

func (p *ThriftSourceProtocolClient) sendAppend(event *ThriftFlumeEvent) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("append", thrift.CALL, p.SeqId)
	args2 := NewAppendArgs()
	args2.Event = event
	err = args2.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *ThriftSourceProtocolClient) recvAppend() (value Status, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result3 := NewAppendResult()
	err = result3.Read(iprot)
	iprot.ReadMessageEnd()
	value = result3.Success
	return
}

// Parameters:
//  - Events
func (p *ThriftSourceProtocolClient) AppendBatch(events []*ThriftFlumeEvent) (r Status, err error) {
	if err = p.sendAppendBatch(events); err != nil {
		return
	}
	return p.recvAppendBatch()
}

func (p *ThriftSourceProtocolClient) sendAppendBatch(events []*ThriftFlumeEvent) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("appendBatch", thrift.CALL, p.SeqId)
	args6 := NewAppendBatchArgs()
	args6.Events = events
	err = args6.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *ThriftSourceProtocolClient) recvAppendBatch() (value Status, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error8 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error9 error
		error9, err = error8.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error9
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result7 := NewAppendBatchResult()
	err = result7.Read(iprot)
	iprot.ReadMessageEnd()
	value = result7.Success
	return
}

type ThriftSourceProtocolProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ThriftSourceProtocol
}

func (p *ThriftSourceProtocolProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ThriftSourceProtocolProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ThriftSourceProtocolProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewThriftSourceProtocolProcessor(handler ThriftSourceProtocol) *ThriftSourceProtocolProcessor {

	self10 := &ThriftSourceProtocolProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self10.processorMap["append"] = &thriftSourceProtocolProcessorAppend{handler: handler}
	self10.processorMap["appendBatch"] = &thriftSourceProtocolProcessorAppendBatch{handler: handler}
	return self10
}

func (p *ThriftSourceProtocolProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x11 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x11.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x11

}

type thriftSourceProtocolProcessorAppend struct {
	handler ThriftSourceProtocol
}

func (p *thriftSourceProtocolProcessorAppend) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewAppendArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewAppendResult()
	if result.Success, err = p.handler.Append(args.Event); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing append: "+err.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("append", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type thriftSourceProtocolProcessorAppendBatch struct {
	handler ThriftSourceProtocol
}

func (p *thriftSourceProtocolProcessorAppendBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewAppendBatchArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("appendBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewAppendBatchResult()
	if result.Success, err = p.handler.AppendBatch(args.Events); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing appendBatch: "+err.Error())
		oprot.WriteMessageBegin("appendBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("appendBatch", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type AppendArgs struct {
	Event *ThriftFlumeEvent `thrift:"event,1"`
}

func NewAppendArgs() *AppendArgs {
	return &AppendArgs{}
}

func (p *AppendArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AppendArgs) readField1(iprot thrift.TProtocol) error {
	p.Event = NewThriftFlumeEvent()
	if err := p.Event.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Event)
	}
	return nil
}

func (p *AppendArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AppendArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Event != nil {
		if err := oprot.WriteFieldBegin("event", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:event: %s", p, err)
		}
		if err := p.Event.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Event)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:event: %s", p, err)
		}
	}
	return err
}

func (p *AppendArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AppendArgs(%+v)", *p)
}

type AppendResult struct {
	Success Status `thrift:"success,0"`
}

func NewAppendResult() *AppendResult {
	return &AppendResult{
		Success: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *AppendResult) IsSetSuccess() bool {
	return int64(p.Success) != math.MinInt32-1
}

func (p *AppendResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AppendResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = Status(v)
	}
	return nil
}

func (p *AppendResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AppendResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *AppendResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AppendResult(%+v)", *p)
}

type AppendBatchArgs struct {
	Events []*ThriftFlumeEvent `thrift:"events,1"`
}

func NewAppendBatchArgs() *AppendBatchArgs {
	return &AppendBatchArgs{}
}

func (p *AppendBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AppendBatchArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Events = make([]*ThriftFlumeEvent, 0, size)
	for i := 0; i < size; i++ {
		_elem12 := NewThriftFlumeEvent()
		if err := _elem12.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem12)
		}
		p.Events = append(p.Events, _elem12)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *AppendBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("appendBatch_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AppendBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Events != nil {
		if err := oprot.WriteFieldBegin("events", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:events: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Events)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Events {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:events: %s", p, err)
		}
	}
	return err
}

func (p *AppendBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AppendBatchArgs(%+v)", *p)
}

type AppendBatchResult struct {
	Success Status `thrift:"success,0"`
}

func NewAppendBatchResult() *AppendBatchResult {
	return &AppendBatchResult{
		Success: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *AppendBatchResult) IsSetSuccess() bool {
	return int64(p.Success) != math.MinInt32-1
}

func (p *AppendBatchResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AppendBatchResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = Status(v)
	}
	return nil
}

func (p *AppendBatchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("appendBatch_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AppendBatchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *AppendBatchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AppendBatchResult(%+v)", *p)
}
