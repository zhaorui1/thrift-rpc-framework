// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package addrservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Response struct {
	Code int32  `thrift:"code,1" json:"code"`
	Desc string `thrift:"desc,2" json:"desc"`
	Data string `thrift:"data,3" json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (p *Response) GetCode() int32 {
	return p.Code
}

func (p *Response) GetDesc() string {
	return p.Desc
}

func (p *Response) GetData() string {
	return p.Data
}
func (p *Response) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
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
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
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

func (p *Response) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *Response) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Desc = v
	}
	return nil
}

func (p *Response) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Data = v
	}
	return nil
}

func (p *Response) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Response"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:code: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Code)); err != nil {
		return fmt.Errorf("%T.code (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:code: %s", p, err)
	}
	return err
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("desc", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:desc: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Desc)); err != nil {
		return fmt.Errorf("%T.desc (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:desc: %s", p, err)
	}
	return err
}

func (p *Response) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("data", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:data: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Data)); err != nil {
		return fmt.Errorf("%T.data (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:data: %s", p, err)
	}
	return err
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}

type Request struct {
	Id string `thrift:"id,1" json:"id"`
}

func NewRequest() *Request {
	return &Request{}
}

func (p *Request) GetId() string {
	return p.Id
}
func (p *Request) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
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
			if err := p.ReadField1(iprot); err != nil {
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

func (p *Request) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Id = v
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Request"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}
