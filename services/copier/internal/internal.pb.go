// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Request
	Response
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	ShardID          *uint64 `protobuf:"varint,1,req,name=ShardID" json:"ShardID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetShardID() uint64 {
	if m != nil && m.ShardID != nil {
		return *m.ShardID
	}
	return 0
}

type Response struct {
	Error            *string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}
