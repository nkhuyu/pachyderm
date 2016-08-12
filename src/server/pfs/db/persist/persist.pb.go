// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	Repo
	BlockRef
	Diff
	Commit
	ProvenanceCommit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Repo struct {
	Name       string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	Size       uint64                     `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Provenance []string                   `protobuf:"bytes,4,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Diff struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs []*BlockRef                `protobuf:"bytes,4,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete    bool                       `protobuf:"varint,5,opt,name=delete" json:"delete,omitempty"`
	Size      uint64                     `protobuf:"varint,6,opt,name=size" json:"size,omitempty"`
	Clock     *Clock                     `protobuf:"bytes,7,opt,name=clock" json:"clock,omitempty"`
	FileType  FileType                   `protobuf:"varint,8,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
	Modified  *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=modified" json:"modified,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetClock() *Clock {
	if m != nil {
		return m.Clock
	}
	return nil
}

func (m *Diff) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Commit struct {
	ID         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo       string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	FullClock  []*Clock                   `protobuf:"bytes,3,rep,name=full_clock,json=fullClock" json:"full_clock,omitempty"`
	Started    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled  bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance []*ProvenanceCommit        `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
	Size       uint64                     `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Commit) GetFullClock() []*Clock {
	if m != nil {
		return m.FullClock
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func (m *Commit) GetProvenance() []*ProvenanceCommit {
	if m != nil {
		return m.Provenance
	}
	return nil
}

type ProvenanceCommit struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
}

func (m *ProvenanceCommit) Reset()                    { *m = ProvenanceCommit{} }
func (m *ProvenanceCommit) String() string            { return proto.CompactTextString(m) }
func (*ProvenanceCommit) ProtoMessage()               {}
func (*ProvenanceCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterType((*ProvenanceCommit)(nil), "ProvenanceCommit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x92, 0x38, 0x8e, 0x3d, 0x95, 0xaa, 0x7c, 0x2b, 0x84, 0xa2, 0x0a, 0x15, 0x64, 0x04,
	0x54, 0x5c, 0xd8, 0xa2, 0x40, 0x1f, 0x80, 0xa6, 0x95, 0x82, 0x50, 0x41, 0xab, 0xde, 0x71, 0x11,
	0xf9, 0x67, 0xb6, 0x59, 0x61, 0xc7, 0xd6, 0xae, 0xd3, 0x0a, 0xae, 0x79, 0x02, 0x1e, 0x82, 0xe7,
	0x64, 0x77, 0xbd, 0x9b, 0x9a, 0x0a, 0x29, 0xb9, 0xf2, 0xec, 0xf8, 0x9c, 0x99, 0x33, 0x67, 0x06,
	0x9e, 0x4b, 0x14, 0xb7, 0x28, 0x92, 0x86, 0xc9, 0xa4, 0xc8, 0x92, 0x06, 0x85, 0xe4, 0xb2, 0x75,
	0xdf, 0xb8, 0x11, 0x75, 0x5b, 0x1f, 0x3d, 0xbd, 0xa9, 0xeb, 0x9b, 0x12, 0x13, 0xf3, 0xca, 0x36,
	0x2c, 0x69, 0x79, 0x85, 0xb2, 0x4d, 0xab, 0xc6, 0x02, 0x8e, 0x1f, 0x02, 0xee, 0x44, 0xda, 0xe8,
	0x1a, 0xdd, 0xff, 0xe8, 0x3d, 0x8c, 0xcf, 0xcb, 0x3a, 0xff, 0x46, 0x1e, 0x83, 0x9f, 0x89, 0x74,
	0x9d, 0xaf, 0x66, 0x83, 0x67, 0x83, 0x93, 0x90, 0xda, 0x17, 0x79, 0x04, 0xe3, 0x5c, 0x03, 0x66,
	0x43, 0x95, 0xf6, 0x68, 0xf7, 0x88, 0xbe, 0xc2, 0xc4, 0xd0, 0x16, 0x73, 0x72, 0x08, 0x43, 0x5e,
	0x58, 0x92, 0x8a, 0x08, 0x01, 0x4f, 0x60, 0x53, 0x1b, 0x7c, 0x48, 0x4d, 0xdc, 0x2b, 0x3e, 0xfa,
	0x77, 0x71, 0xaf, 0x5f, 0xfc, 0xe7, 0x00, 0x3c, 0xaa, 0x69, 0xaa, 0xd4, 0x3a, 0xad, 0xd0, 0x16,
	0x37, 0x31, 0x79, 0x07, 0x93, 0x5c, 0x60, 0xda, 0x62, 0x61, 0x3a, 0x1c, 0x9c, 0x1e, 0xc5, 0xdd,
	0x88, 0xb1, 0x1b, 0x31, 0xbe, 0x76, 0x1e, 0x50, 0x07, 0xd5, 0x95, 0x24, 0xff, 0x81, 0xa6, 0xbd,
	0x47, 0x4d, 0x4c, 0x8e, 0x01, 0x14, 0xe5, 0x16, 0xd7, 0x4a, 0x0a, 0x2a, 0x05, 0x23, 0xd5, 0xa3,
	0x97, 0x89, 0x3e, 0x42, 0xf0, 0x41, 0xeb, 0xa1, 0xc8, 0x34, 0x7f, 0x95, 0x4a, 0xe7, 0x8d, 0x89,
	0xb5, 0xf8, 0xb2, 0xbe, 0x43, 0xe1, 0x9c, 0x31, 0x0f, 0x9d, 0xdd, 0x68, 0x83, 0x6d, 0xab, 0xee,
	0x11, 0xfd, 0x1a, 0x82, 0x37, 0xe7, 0x8c, 0xed, 0xe5, 0x96, 0xca, 0x35, 0x69, 0xeb, 0xbc, 0x32,
	0x31, 0x39, 0x01, 0xc8, 0xb4, 0x98, 0xa5, 0x40, 0x26, 0x8d, 0xd8, 0x83, 0xd3, 0x30, 0x76, 0xfa,
	0x68, 0x98, 0xd9, 0x48, 0x6a, 0xaf, 0x0b, 0x2c, 0xb1, 0xc5, 0xd9, 0x58, 0xf1, 0x03, 0x6a, 0x5f,
	0x5b, 0x0b, 0xfc, 0x9e, 0x05, 0x4f, 0x9c, 0xff, 0x13, 0x63, 0xa5, 0x1f, 0x9b, 0xa5, 0xda, 0x3d,
	0x90, 0x97, 0x10, 0x32, 0x5e, 0xe2, 0xb2, 0xfd, 0xde, 0xe0, 0x2c, 0x50, 0x88, 0x43, 0xd5, 0xf2,
	0x52, 0x65, 0xae, 0x55, 0x82, 0x06, 0xcc, 0x46, 0xe4, 0x0c, 0x82, 0xaa, 0x2e, 0x38, 0xe3, 0x6a,
	0x27, 0xe1, 0xce, 0x9d, 0x6c, 0xb1, 0xd1, 0xef, 0x21, 0xf8, 0xe7, 0x75, 0x55, 0xf1, 0x76, 0x2f,
	0x5b, 0x5e, 0x00, 0xb0, 0x4d, 0x59, 0x2e, 0x3b, 0xc5, 0x23, 0x63, 0x81, 0x53, 0x1c, 0xea, 0x3f,
	0xdd, 0x21, 0xab, 0x03, 0x51, 0x8d, 0x84, 0x3e, 0x10, 0x6f, 0xf7, 0x81, 0x58, 0xa8, 0x9e, 0x81,
	0xf1, 0x35, 0x97, 0x2b, 0x45, 0x1b, 0xef, 0x9e, 0xc1, 0x61, 0x95, 0x83, 0x61, 0xae, 0xaf, 0xa5,
	0x2c, 0x15, 0xd1, 0x37, 0x86, 0xdf, 0x27, 0xc8, 0x9b, 0xbf, 0x4e, 0x6c, 0x62, 0x24, 0xff, 0x1f,
	0x7f, 0xd9, 0xa6, 0xba, 0xe9, 0xfb, 0x57, 0xb7, 0x5d, 0x53, 0x70, 0xbf, 0xa6, 0xe8, 0x0c, 0xa6,
	0x0f, 0x39, 0xfb, 0x38, 0xf6, 0xfa, 0x15, 0x04, 0x6e, 0x5d, 0x24, 0x00, 0xef, 0xea, 0xf3, 0xd5,
	0xc5, 0xf4, 0x3f, 0x1d, 0x5d, 0x2e, 0x3e, 0x5d, 0x4c, 0x07, 0x64, 0x02, 0xa3, 0xf9, 0x82, 0x4e,
	0x87, 0x99, 0x6f, 0x66, 0x7c, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xe3, 0x0d, 0xed, 0x74,
	0x04, 0x00, 0x00,
}