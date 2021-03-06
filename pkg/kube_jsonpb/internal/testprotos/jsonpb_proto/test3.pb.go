// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: jsonpb_proto/test3.proto

package jsonpb_proto

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

// Enum value maps for Numeral.
var (
	Numeral_name = map[int32]string{
		0: "UNKNOWN",
		1: "ARABIC",
		2: "ROMAN",
	}
	Numeral_value = map[string]int32{
		"UNKNOWN": 0,
		"ARABIC":  1,
		"ROMAN":   2,
	}
)

func (x Numeral) Enum() *Numeral {
	p := new(Numeral)
	*p = x
	return p
}

func (x Numeral) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Numeral) Descriptor() protoreflect.EnumDescriptor {
	return file_jsonpb_proto_test3_proto_enumTypes[0].Descriptor()
}

func (Numeral) Type() protoreflect.EnumType {
	return &file_jsonpb_proto_test3_proto_enumTypes[0]
}

func (x Numeral) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Numeral.Descriptor instead.
func (Numeral) EnumDescriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{0}
}

type Simple3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dub float64 `protobuf:"fixed64,1,opt,name=dub,proto3" json:"dub,omitempty"`
}

func (x *Simple3) Reset() {
	*x = Simple3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonpb_proto_test3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple3) ProtoMessage() {}

func (x *Simple3) ProtoReflect() protoreflect.Message {
	mi := &file_jsonpb_proto_test3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple3.ProtoReflect.Descriptor instead.
func (*Simple3) Descriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{0}
}

func (x *Simple3) GetDub() float64 {
	if x != nil {
		return x.Dub
	}
	return 0
}

type SimpleSlice3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slices []string `protobuf:"bytes,1,rep,name=slices,proto3" json:"slices,omitempty"`
}

func (x *SimpleSlice3) Reset() {
	*x = SimpleSlice3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonpb_proto_test3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleSlice3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleSlice3) ProtoMessage() {}

func (x *SimpleSlice3) ProtoReflect() protoreflect.Message {
	mi := &file_jsonpb_proto_test3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleSlice3.ProtoReflect.Descriptor instead.
func (*SimpleSlice3) Descriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleSlice3) GetSlices() []string {
	if x != nil {
		return x.Slices
	}
	return nil
}

type SimpleMap3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stringy map[string]string `protobuf:"bytes,1,rep,name=stringy,proto3" json:"stringy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SimpleMap3) Reset() {
	*x = SimpleMap3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonpb_proto_test3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMap3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMap3) ProtoMessage() {}

func (x *SimpleMap3) ProtoReflect() protoreflect.Message {
	mi := &file_jsonpb_proto_test3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMap3.ProtoReflect.Descriptor instead.
func (*SimpleMap3) Descriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleMap3) GetStringy() map[string]string {
	if x != nil {
		return x.Stringy
	}
	return nil
}

type SimpleNull3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Simple *Simple3 `protobuf:"bytes,1,opt,name=simple,proto3" json:"simple,omitempty"`
}

func (x *SimpleNull3) Reset() {
	*x = SimpleNull3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonpb_proto_test3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleNull3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleNull3) ProtoMessage() {}

func (x *SimpleNull3) ProtoReflect() protoreflect.Message {
	mi := &file_jsonpb_proto_test3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleNull3.ProtoReflect.Descriptor instead.
func (*SimpleNull3) Descriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleNull3) GetSimple() *Simple3 {
	if x != nil {
		return x.Simple
	}
	return nil
}

type Mappy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nummy    map[int64]int32    `protobuf:"bytes,1,rep,name=nummy,proto3" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Strry    map[string]string  `protobuf:"bytes,2,rep,name=strry,proto3" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Objjy    map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy,proto3" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Buggy    map[int64]string   `protobuf:"bytes,4,rep,name=buggy,proto3" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Booly    map[bool]bool      `protobuf:"bytes,5,rep,name=booly,proto3" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Enumy    map[string]Numeral `protobuf:"bytes,6,rep,name=enumy,proto3" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=jsonpb_test.Numeral"`
	S32Booly map[int32]bool     `protobuf:"bytes,7,rep,name=s32booly,proto3" json:"s32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	S64Booly map[int64]bool     `protobuf:"bytes,8,rep,name=s64booly,proto3" json:"s64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U32Booly map[uint32]bool    `protobuf:"bytes,9,rep,name=u32booly,proto3" json:"u32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U64Booly map[uint64]bool    `protobuf:"bytes,10,rep,name=u64booly,proto3" json:"u64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Mappy) Reset() {
	*x = Mappy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonpb_proto_test3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mappy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mappy) ProtoMessage() {}

func (x *Mappy) ProtoReflect() protoreflect.Message {
	mi := &file_jsonpb_proto_test3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mappy.ProtoReflect.Descriptor instead.
func (*Mappy) Descriptor() ([]byte, []int) {
	return file_jsonpb_proto_test3_proto_rawDescGZIP(), []int{4}
}

func (x *Mappy) GetNummy() map[int64]int32 {
	if x != nil {
		return x.Nummy
	}
	return nil
}

func (x *Mappy) GetStrry() map[string]string {
	if x != nil {
		return x.Strry
	}
	return nil
}

func (x *Mappy) GetObjjy() map[int32]*Simple3 {
	if x != nil {
		return x.Objjy
	}
	return nil
}

func (x *Mappy) GetBuggy() map[int64]string {
	if x != nil {
		return x.Buggy
	}
	return nil
}

func (x *Mappy) GetBooly() map[bool]bool {
	if x != nil {
		return x.Booly
	}
	return nil
}

func (x *Mappy) GetEnumy() map[string]Numeral {
	if x != nil {
		return x.Enumy
	}
	return nil
}

func (x *Mappy) GetS32Booly() map[int32]bool {
	if x != nil {
		return x.S32Booly
	}
	return nil
}

func (x *Mappy) GetS64Booly() map[int64]bool {
	if x != nil {
		return x.S64Booly
	}
	return nil
}

func (x *Mappy) GetU32Booly() map[uint32]bool {
	if x != nil {
		return x.U32Booly
	}
	return nil
}

func (x *Mappy) GetU64Booly() map[uint64]bool {
	if x != nil {
		return x.U64Booly
	}
	return nil
}

var File_jsonpb_proto_test3_proto protoreflect.FileDescriptor

var file_jsonpb_proto_test3_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6a, 0x73, 0x6f, 0x6e,
	0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x07, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x64, 0x75, 0x62, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a,
	0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x33, 0x12, 0x3e, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a,
	0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4d, 0x61, 0x70, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4e, 0x75, 0x6c, 0x6c, 0x33, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x33, 0x52, 0x06, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0xb9, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x12, 0x33,
	0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70,
	0x79, 0x2e, 0x4e, 0x75, 0x6d, 0x6d, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x75,
	0x6d, 0x6d, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x72, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x73, 0x74, 0x72, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x6a,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x4f, 0x62, 0x6a, 0x6a,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x6a, 0x79, 0x12, 0x33, 0x0a,
	0x05, 0x62, 0x75, 0x67, 0x67, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a,
	0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79,
	0x2e, 0x42, 0x75, 0x67, 0x67, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62, 0x75, 0x67,
	0x67, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x79,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x79, 0x12, 0x3c, 0x0a, 0x08,
	0x73, 0x33, 0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70,
	0x70, 0x79, 0x2e, 0x53, 0x33, 0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x73, 0x33, 0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x36,
	0x34, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a,
	0x73, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79,
	0x2e, 0x53, 0x36, 0x34, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x73, 0x36, 0x34, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x75, 0x33, 0x32, 0x62,
	0x6f, 0x6f, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x55,
	0x33, 0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x33,
	0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x75, 0x36, 0x34, 0x62, 0x6f, 0x6f,
	0x6c, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70,
	0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x79, 0x2e, 0x55, 0x36, 0x34,
	0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x36, 0x34, 0x62,
	0x6f, 0x6f, 0x6c, 0x79, 0x1a, 0x38, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x6d, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x6a,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x33, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x42, 0x75, 0x67, 0x67,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0a,
	0x45, 0x6e, 0x75, 0x6d, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d,
	0x53, 0x33, 0x32, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x36, 0x34,
	0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x55, 0x33, 0x32, 0x62, 0x6f, 0x6f,
	0x6c, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x55, 0x36, 0x34, 0x62, 0x6f, 0x6f, 0x6c, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x2d, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x52, 0x41, 0x42,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x42,
	0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x70, 0x62,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jsonpb_proto_test3_proto_rawDescOnce sync.Once
	file_jsonpb_proto_test3_proto_rawDescData = file_jsonpb_proto_test3_proto_rawDesc
)

func file_jsonpb_proto_test3_proto_rawDescGZIP() []byte {
	file_jsonpb_proto_test3_proto_rawDescOnce.Do(func() {
		file_jsonpb_proto_test3_proto_rawDescData = protoimpl.X.CompressGZIP(file_jsonpb_proto_test3_proto_rawDescData)
	})
	return file_jsonpb_proto_test3_proto_rawDescData
}

var file_jsonpb_proto_test3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jsonpb_proto_test3_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_jsonpb_proto_test3_proto_goTypes = []interface{}{
	(Numeral)(0),         // 0: jsonpb_test.Numeral
	(*Simple3)(nil),      // 1: jsonpb_test.Simple3
	(*SimpleSlice3)(nil), // 2: jsonpb_test.SimpleSlice3
	(*SimpleMap3)(nil),   // 3: jsonpb_test.SimpleMap3
	(*SimpleNull3)(nil),  // 4: jsonpb_test.SimpleNull3
	(*Mappy)(nil),        // 5: jsonpb_test.Mappy
	nil,                  // 6: jsonpb_test.SimpleMap3.StringyEntry
	nil,                  // 7: jsonpb_test.Mappy.NummyEntry
	nil,                  // 8: jsonpb_test.Mappy.StrryEntry
	nil,                  // 9: jsonpb_test.Mappy.ObjjyEntry
	nil,                  // 10: jsonpb_test.Mappy.BuggyEntry
	nil,                  // 11: jsonpb_test.Mappy.BoolyEntry
	nil,                  // 12: jsonpb_test.Mappy.EnumyEntry
	nil,                  // 13: jsonpb_test.Mappy.S32boolyEntry
	nil,                  // 14: jsonpb_test.Mappy.S64boolyEntry
	nil,                  // 15: jsonpb_test.Mappy.U32boolyEntry
	nil,                  // 16: jsonpb_test.Mappy.U64boolyEntry
}
var file_jsonpb_proto_test3_proto_depIdxs = []int32{
	6,  // 0: jsonpb_test.SimpleMap3.stringy:type_name -> jsonpb_test.SimpleMap3.StringyEntry
	1,  // 1: jsonpb_test.SimpleNull3.simple:type_name -> jsonpb_test.Simple3
	7,  // 2: jsonpb_test.Mappy.nummy:type_name -> jsonpb_test.Mappy.NummyEntry
	8,  // 3: jsonpb_test.Mappy.strry:type_name -> jsonpb_test.Mappy.StrryEntry
	9,  // 4: jsonpb_test.Mappy.objjy:type_name -> jsonpb_test.Mappy.ObjjyEntry
	10, // 5: jsonpb_test.Mappy.buggy:type_name -> jsonpb_test.Mappy.BuggyEntry
	11, // 6: jsonpb_test.Mappy.booly:type_name -> jsonpb_test.Mappy.BoolyEntry
	12, // 7: jsonpb_test.Mappy.enumy:type_name -> jsonpb_test.Mappy.EnumyEntry
	13, // 8: jsonpb_test.Mappy.s32booly:type_name -> jsonpb_test.Mappy.S32boolyEntry
	14, // 9: jsonpb_test.Mappy.s64booly:type_name -> jsonpb_test.Mappy.S64boolyEntry
	15, // 10: jsonpb_test.Mappy.u32booly:type_name -> jsonpb_test.Mappy.U32boolyEntry
	16, // 11: jsonpb_test.Mappy.u64booly:type_name -> jsonpb_test.Mappy.U64boolyEntry
	1,  // 12: jsonpb_test.Mappy.ObjjyEntry.value:type_name -> jsonpb_test.Simple3
	0,  // 13: jsonpb_test.Mappy.EnumyEntry.value:type_name -> jsonpb_test.Numeral
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_jsonpb_proto_test3_proto_init() }
func file_jsonpb_proto_test3_proto_init() {
	if File_jsonpb_proto_test3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jsonpb_proto_test3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonpb_proto_test3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleSlice3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonpb_proto_test3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMap3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonpb_proto_test3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleNull3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonpb_proto_test3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mappy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jsonpb_proto_test3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jsonpb_proto_test3_proto_goTypes,
		DependencyIndexes: file_jsonpb_proto_test3_proto_depIdxs,
		EnumInfos:         file_jsonpb_proto_test3_proto_enumTypes,
		MessageInfos:      file_jsonpb_proto_test3_proto_msgTypes,
	}.Build()
	File_jsonpb_proto_test3_proto = out.File
	file_jsonpb_proto_test3_proto_rawDesc = nil
	file_jsonpb_proto_test3_proto_goTypes = nil
	file_jsonpb_proto_test3_proto_depIdxs = nil
}
