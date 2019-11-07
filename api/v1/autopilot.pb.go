// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilot.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The AutoPilotProject file is the root configuration file for the project itself
// this file will be used to build and deploy the autopilot operator
// Loaded automatically by the autopilot CLI
// Default location is 'autopilot.yaml'
type AutoPilotProject struct {
	// the name (kubernetes Kind) of the top-level
	// CRD for the operator
	// Specified via the `ap init <Kind>` command
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// the ApiVersion of the top-level
	// CRD for the operator
	ApiVersion string `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	// Each phase represents a different
	// stage in the lifecycle of the CRD (e.g. Pending/Succeeded/Failed).
	//
	// Each phase specifies a unique name
	// and its own set of inputs and outputs.
	Phases []*Phase `protobuf:"bytes,3,rep,name=phases,proto3" json:"phases,omitempty"`
	// enable use of a Finalizer to handle object deletion
	EnableFinalizer      bool     `protobuf:"varint,4,opt,name=enableFinalizer,proto3" json:"enableFinalizer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoPilotProject) Reset()         { *m = AutoPilotProject{} }
func (m *AutoPilotProject) String() string { return proto.CompactTextString(m) }
func (*AutoPilotProject) ProtoMessage()    {}
func (*AutoPilotProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{0}
}

func (m *AutoPilotProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoPilotProject.Unmarshal(m, b)
}
func (m *AutoPilotProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoPilotProject.Marshal(b, m, deterministic)
}
func (m *AutoPilotProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoPilotProject.Merge(m, src)
}
func (m *AutoPilotProject) XXX_Size() int {
	return xxx_messageInfo_AutoPilotProject.Size(m)
}
func (m *AutoPilotProject) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoPilotProject.DiscardUnknown(m)
}

var xxx_messageInfo_AutoPilotProject proto.InternalMessageInfo

func (m *AutoPilotProject) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *AutoPilotProject) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *AutoPilotProject) GetPhases() []*Phase {
	if m != nil {
		return m.Phases
	}
	return nil
}

func (m *AutoPilotProject) GetEnableFinalizer() bool {
	if m != nil {
		return m.EnableFinalizer
	}
	return false
}

// MeshProviders provide an interface to monitoring and managing a specific
// mesh.
// AutoPilot does not abstract the mesh API - AutoPilot developers must
// still reason able about Provider-specific CRDs. AutoPilot's job is to
// abstract operational concerns such as discovering control plane configuration
// and monitoring metrics.
type Phase struct {
	// name of the phase. must be unique
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// description of the phase. used for comments and docs
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// indicates whether this is the initial phase of the system.
	// exactly one phase must be the initial phase
	Initial bool `protobuf:"varint,3,opt,name=initial,proto3" json:"initial,omitempty"`
	// the set of inputs for this phase
	// the inputs will be retrieved by the scheduler
	// and passed to the worker as input parameters
	//
	// custom inputs can be defined in the
	// autopilot.yaml
	Inputs []string `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// the set of outputs for this phase
	// the inputs will be propagated to k8s storage (etcd) by the scheduler.
	//
	// custom outputs can be defined in the
	// autopilot.yaml
	Outputs              []string `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phase) Reset()         { *m = Phase{} }
func (m *Phase) String() string { return proto.CompactTextString(m) }
func (*Phase) ProtoMessage()    {}
func (*Phase) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{1}
}

func (m *Phase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phase.Unmarshal(m, b)
}
func (m *Phase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phase.Marshal(b, m, deterministic)
}
func (m *Phase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phase.Merge(m, src)
}
func (m *Phase) XXX_Size() int {
	return xxx_messageInfo_Phase.Size(m)
}
func (m *Phase) XXX_DiscardUnknown() {
	xxx_messageInfo_Phase.DiscardUnknown(m)
}

var xxx_messageInfo_Phase proto.InternalMessageInfo

func (m *Phase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Phase) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Phase) GetInitial() bool {
	if m != nil {
		return m.Initial
	}
	return false
}

func (m *Phase) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Phase) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoPilotProject)(nil), "autopilot.AutoPilotProject")
	proto.RegisterType((*Phase)(nil), "autopilot.Phase")
}

func init() { proto.RegisterFile("autopilot.proto", fileDescriptor_f7c7e86e2b87635e) }

var fileDescriptor_f7c7e86e2b87635e = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0x95, 0x2f, 0x69, 0x3e, 0x72, 0x1d, 0x5a, 0x79, 0x40, 0x9e, 0x50, 0x54, 0x84, 0xe4,
	0x85, 0x44, 0xc0, 0x13, 0xc0, 0xc0, 0x1c, 0x65, 0x60, 0x60, 0x73, 0x52, 0x8b, 0x1e, 0xa4, 0x3e,
	0x2b, 0xbe, 0x30, 0xf0, 0x06, 0xbc, 0x02, 0x4f, 0x8b, 0x6c, 0x35, 0x6d, 0xc5, 0x76, 0xff, 0xdf,
	0x4f, 0xa7, 0xfb, 0xdb, 0xb0, 0xd2, 0x13, 0x93, 0xc3, 0x81, 0xb8, 0x72, 0x23, 0x31, 0x89, 0xe2,
	0x08, 0x36, 0x3f, 0x09, 0xac, 0x1f, 0x27, 0xa6, 0x26, 0xa4, 0x66, 0xa4, 0x77, 0xd3, 0xb3, 0x10,
	0x90, 0x7d, 0xa0, 0xdd, 0xca, 0xa4, 0x4c, 0x54, 0xd1, 0xc6, 0x59, 0x5c, 0x01, 0x68, 0x87, 0x2f,
	0x66, 0xf4, 0x48, 0x56, 0xfe, 0x8b, 0xe6, 0x8c, 0x08, 0x05, 0xb9, 0xdb, 0x69, 0x6f, 0xbc, 0x4c,
	0xcb, 0x54, 0x2d, 0xef, 0xd7, 0xd5, 0xe9, 0x6a, 0x13, 0x44, 0x7b, 0xf0, 0x42, 0xc1, 0xca, 0x58,
	0xdd, 0x0d, 0xe6, 0x19, 0xad, 0x1e, 0xf0, 0xcb, 0x8c, 0x32, 0x2b, 0x13, 0x75, 0xd1, 0xfe, 0xc5,
	0x9b, 0xef, 0x04, 0x16, 0x71, 0x37, 0x34, 0xb2, 0x7a, 0x6f, 0xe6, 0x46, 0x61, 0x16, 0x25, 0x2c,
	0xb7, 0xc6, 0xf7, 0x23, 0x3a, 0x3e, 0x55, 0x3a, 0x47, 0x42, 0xc2, 0x7f, 0xb4, 0xc8, 0xa8, 0x07,
	0x99, 0xc6, 0x0b, 0x73, 0x14, 0x97, 0x90, 0xa3, 0x75, 0x13, 0x7b, 0x99, 0x95, 0xa9, 0x2a, 0xda,
	0x43, 0x0a, 0x1b, 0x34, 0x71, 0x14, 0x8b, 0x28, 0xe6, 0xf8, 0x74, 0xf3, 0x7a, 0xfd, 0x86, 0xbc,
	0x9b, 0xba, 0xaa, 0xa7, 0x7d, 0xed, 0x69, 0xa0, 0x5b, 0xa4, 0xfa, 0xf8, 0xc6, 0x5a, 0x3b, 0xac,
	0x3f, 0xef, 0xba, 0x3c, 0xfe, 0xf0, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x7d, 0xca,
	0xb9, 0x74, 0x01, 0x00, 0x00,
}