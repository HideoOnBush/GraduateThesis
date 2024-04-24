// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package neonode

import (
	"GraduateThesis/biz/model/base"
	"GraduateThesis/biz/model/line"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type NeoNode struct {
	Name   string `thrift:"name,1" form:"name" json:"name" query:"name"`
	IsCore bool   `thrift:"isCore,2" form:"isCore" json:"isCore" query:"isCore"`
	Scene  string `thrift:"scene,3" form:"scene" json:"scene" query:"scene"`
}

func NewNeoNode() *NeoNode {
	return &NeoNode{}
}

func (p *NeoNode) GetName() (v string) {
	return p.Name
}

func (p *NeoNode) GetIsCore() (v bool) {
	return p.IsCore
}

func (p *NeoNode) GetScene() (v string) {
	return p.Scene
}

var fieldIDToName_NeoNode = map[int16]string{
	1: "name",
	2: "isCore",
	3: "scene",
}

func (p *NeoNode) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NeoNode[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NeoNode) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Name = v
	}
	return nil
}
func (p *NeoNode) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.IsCore = v
	}
	return nil
}
func (p *NeoNode) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Scene = v
	}
	return nil
}

func (p *NeoNode) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("NeoNode"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NeoNode) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *NeoNode) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("isCore", thrift.BOOL, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.IsCore); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *NeoNode) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("scene", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Scene); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *NeoNode) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NeoNode(%+v)", *p)

}

type NeoNodeService interface {
	EsToNeo(ctx context.Context, request *line.LineReq) (r *base.SampleResp, err error)

	EmptyNeo(ctx context.Context) (r *base.SampleResp, err error)
}

type NeoNodeServiceClient struct {
	c thrift.TClient
}

func NewNeoNodeServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *NeoNodeServiceClient {
	return &NeoNodeServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewNeoNodeServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *NeoNodeServiceClient {
	return &NeoNodeServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewNeoNodeServiceClient(c thrift.TClient) *NeoNodeServiceClient {
	return &NeoNodeServiceClient{
		c: c,
	}
}

func (p *NeoNodeServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *NeoNodeServiceClient) EsToNeo(ctx context.Context, request *line.LineReq) (r *base.SampleResp, err error) {
	var _args NeoNodeServiceEsToNeoArgs
	_args.Request = request
	var _result NeoNodeServiceEsToNeoResult
	if err = p.Client_().Call(ctx, "EsToNeo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *NeoNodeServiceClient) EmptyNeo(ctx context.Context) (r *base.SampleResp, err error) {
	var _args NeoNodeServiceEmptyNeoArgs
	var _result NeoNodeServiceEmptyNeoResult
	if err = p.Client_().Call(ctx, "EmptyNeo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type NeoNodeServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      NeoNodeService
}

func (p *NeoNodeServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *NeoNodeServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *NeoNodeServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewNeoNodeServiceProcessor(handler NeoNodeService) *NeoNodeServiceProcessor {
	self := &NeoNodeServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("EsToNeo", &neoNodeServiceProcessorEsToNeo{handler: handler})
	self.AddToProcessorMap("EmptyNeo", &neoNodeServiceProcessorEmptyNeo{handler: handler})
	return self
}
func (p *NeoNodeServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type neoNodeServiceProcessorEsToNeo struct {
	handler NeoNodeService
}

func (p *neoNodeServiceProcessorEsToNeo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NeoNodeServiceEsToNeoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("EsToNeo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := NeoNodeServiceEsToNeoResult{}
	var retval *base.SampleResp
	if retval, err2 = p.handler.EsToNeo(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing EsToNeo: "+err2.Error())
		oprot.WriteMessageBegin("EsToNeo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("EsToNeo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type neoNodeServiceProcessorEmptyNeo struct {
	handler NeoNodeService
}

func (p *neoNodeServiceProcessorEmptyNeo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NeoNodeServiceEmptyNeoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("EmptyNeo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := NeoNodeServiceEmptyNeoResult{}
	var retval *base.SampleResp
	if retval, err2 = p.handler.EmptyNeo(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing EmptyNeo: "+err2.Error())
		oprot.WriteMessageBegin("EmptyNeo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("EmptyNeo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type NeoNodeServiceEsToNeoArgs struct {
	Request *line.LineReq `thrift:"request,1"`
}

func NewNeoNodeServiceEsToNeoArgs() *NeoNodeServiceEsToNeoArgs {
	return &NeoNodeServiceEsToNeoArgs{}
}

var NeoNodeServiceEsToNeoArgs_Request_DEFAULT *line.LineReq

func (p *NeoNodeServiceEsToNeoArgs) GetRequest() (v *line.LineReq) {
	if !p.IsSetRequest() {
		return NeoNodeServiceEsToNeoArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_NeoNodeServiceEsToNeoArgs = map[int16]string{
	1: "request",
}

func (p *NeoNodeServiceEsToNeoArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *NeoNodeServiceEsToNeoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NeoNodeServiceEsToNeoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = line.NewLineReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *NeoNodeServiceEsToNeoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EsToNeo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NeoNodeServiceEsToNeoArgs(%+v)", *p)

}

type NeoNodeServiceEsToNeoResult struct {
	Success *base.SampleResp `thrift:"success,0,optional"`
}

func NewNeoNodeServiceEsToNeoResult() *NeoNodeServiceEsToNeoResult {
	return &NeoNodeServiceEsToNeoResult{}
}

var NeoNodeServiceEsToNeoResult_Success_DEFAULT *base.SampleResp

func (p *NeoNodeServiceEsToNeoResult) GetSuccess() (v *base.SampleResp) {
	if !p.IsSetSuccess() {
		return NeoNodeServiceEsToNeoResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_NeoNodeServiceEsToNeoResult = map[int16]string{
	0: "success",
}

func (p *NeoNodeServiceEsToNeoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NeoNodeServiceEsToNeoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NeoNodeServiceEsToNeoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = base.NewSampleResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *NeoNodeServiceEsToNeoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EsToNeo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *NeoNodeServiceEsToNeoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NeoNodeServiceEsToNeoResult(%+v)", *p)

}

type NeoNodeServiceEmptyNeoArgs struct {
}

func NewNeoNodeServiceEmptyNeoArgs() *NeoNodeServiceEmptyNeoArgs {
	return &NeoNodeServiceEmptyNeoArgs{}
}

var fieldIDToName_NeoNodeServiceEmptyNeoArgs = map[int16]string{}

func (p *NeoNodeServiceEmptyNeoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NeoNodeServiceEmptyNeoArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("EmptyNeo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NeoNodeServiceEmptyNeoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NeoNodeServiceEmptyNeoArgs(%+v)", *p)

}

type NeoNodeServiceEmptyNeoResult struct {
	Success *base.SampleResp `thrift:"success,0,optional"`
}

func NewNeoNodeServiceEmptyNeoResult() *NeoNodeServiceEmptyNeoResult {
	return &NeoNodeServiceEmptyNeoResult{}
}

var NeoNodeServiceEmptyNeoResult_Success_DEFAULT *base.SampleResp

func (p *NeoNodeServiceEmptyNeoResult) GetSuccess() (v *base.SampleResp) {
	if !p.IsSetSuccess() {
		return NeoNodeServiceEmptyNeoResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_NeoNodeServiceEmptyNeoResult = map[int16]string{
	0: "success",
}

func (p *NeoNodeServiceEmptyNeoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NeoNodeServiceEmptyNeoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NeoNodeServiceEmptyNeoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NeoNodeServiceEmptyNeoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = base.NewSampleResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *NeoNodeServiceEmptyNeoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EmptyNeo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NeoNodeServiceEmptyNeoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *NeoNodeServiceEmptyNeoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NeoNodeServiceEmptyNeoResult(%+v)", *p)

}
