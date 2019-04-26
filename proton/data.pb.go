// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package proton

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type DietType int32

const (
	DietType_BALANCED     DietType = 0
	DietType_HIGH_PROTEIN DietType = 1
	DietType_HIGH_FIBER   DietType = 2
	DietType_LOW_FAT      DietType = 3
	DietType_LOW_CARB     DietType = 4
	DietType_LOW_SODIUM   DietType = 5
)

var DietType_name = map[int32]string{
	0: "BALANCED",
	1: "HIGH_PROTEIN",
	2: "HIGH_FIBER",
	3: "LOW_FAT",
	4: "LOW_CARB",
	5: "LOW_SODIUM",
}

var DietType_value = map[string]int32{
	"BALANCED":     0,
	"HIGH_PROTEIN": 1,
	"HIGH_FIBER":   2,
	"LOW_FAT":      3,
	"LOW_CARB":     4,
	"LOW_SODIUM":   5,
}

func (x DietType) String() string {
	return proto.EnumName(DietType_name, int32(x))
}

func (DietType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type HealthType int32

const (
	HealthType_PEANUT_FREE    HealthType = 0
	HealthType_TREE_NUT_FREE  HealthType = 1
	HealthType_SOY_FREE       HealthType = 2
	HealthType_FISH_FREE      HealthType = 3
	HealthType_SHELLFISH_FREE HealthType = 4
)

var HealthType_name = map[int32]string{
	0: "PEANUT_FREE",
	1: "TREE_NUT_FREE",
	2: "SOY_FREE",
	3: "FISH_FREE",
	4: "SHELLFISH_FREE",
}

var HealthType_value = map[string]int32{
	"PEANUT_FREE":    0,
	"TREE_NUT_FREE":  1,
	"SOY_FREE":       2,
	"FISH_FREE":      3,
	"SHELLFISH_FREE": 4,
}

func (x HealthType) String() string {
	return proto.EnumName(HealthType_name, int32(x))
}

func (HealthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

type QueryRequest struct {
	Q                    string     `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
	R                    string     `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	From                 int32      `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   int32      `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Diet                 DietType   `protobuf:"varint,5,opt,name=Diet,proto3,enum=proton.DietType" json:"Diet,omitempty"`
	Health               HealthType `protobuf:"varint,6,opt,name=Health,proto3,enum=proton.HealthType" json:"Health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *QueryRequest) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *QueryRequest) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *QueryRequest) GetTo() int32 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *QueryRequest) GetDiet() DietType {
	if m != nil {
		return m.Diet
	}
	return DietType_BALANCED
}

func (m *QueryRequest) GetHealth() HealthType {
	if m != nil {
		return m.Health
	}
	return HealthType_PEANUT_FREE
}

type Recipe struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Label                string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Yield                float64  `protobuf:"fixed64,4,opt,name=yield,proto3" json:"yield,omitempty"`
	Calories             string   `protobuf:"bytes,5,opt,name=calories,proto3" json:"calories,omitempty"`
	TotalTime            float64  `protobuf:"fixed64,6,opt,name=total_time,json=totalTime,proto3" json:"total_time,omitempty"`
	TotalWeight          float64  `protobuf:"fixed64,7,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Recipe) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Recipe) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Recipe) GetYield() float64 {
	if m != nil {
		return m.Yield
	}
	return 0
}

func (m *Recipe) GetCalories() string {
	if m != nil {
		return m.Calories
	}
	return ""
}

func (m *Recipe) GetTotalTime() float64 {
	if m != nil {
		return m.TotalTime
	}
	return 0
}

func (m *Recipe) GetTotalWeight() float64 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

type Result struct {
	Entries              []*Recipe `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetEntries() []*Recipe {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("proton.DietType", DietType_name, DietType_value)
	proto.RegisterEnum("proton.HealthType", HealthType_name, HealthType_value)
	proto.RegisterType((*QueryRequest)(nil), "proton.QueryRequest")
	proto.RegisterType((*Recipe)(nil), "proton.Recipe")
	proto.RegisterType((*Result)(nil), "proton.Result")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xda, 0x4c,
	0x14, 0x85, 0x33, 0x18, 0x08, 0xbe, 0x10, 0x7e, 0xff, 0x57, 0x51, 0x65, 0x51, 0x55, 0xa2, 0x51,
	0x17, 0x08, 0xa9, 0x8e, 0x44, 0x77, 0xdd, 0x99, 0x60, 0x6a, 0x24, 0x0a, 0xe9, 0xd8, 0x51, 0xd4,
	0x15, 0x32, 0x70, 0x03, 0xae, 0x0c, 0x06, 0x7b, 0xac, 0x8a, 0xb7, 0xe9, 0x53, 0xf4, 0xf9, 0xaa,
	0x99, 0xc1, 0x49, 0xd4, 0x5d, 0x57, 0xcc, 0xf9, 0xe6, 0x0c, 0x73, 0xe6, 0xc8, 0x17, 0x60, 0x1d,
	0x89, 0xc8, 0x39, 0x64, 0xa9, 0x48, 0xb1, 0xae, 0x7e, 0xf6, 0x9d, 0xb7, 0x9b, 0x34, 0xdd, 0x24,
	0x74, 0xab, 0xe4, 0xb2, 0x78, 0xba, 0xa5, 0xdd, 0x41, 0x9c, 0xb4, 0xe9, 0xe6, 0x17, 0x83, 0xd6,
	0xb7, 0x82, 0xb2, 0x13, 0xa7, 0x63, 0x41, 0xb9, 0xc0, 0x16, 0xb0, 0xa3, 0xcd, 0xba, 0xac, 0x67,
	0x72, 0x76, 0x94, 0x2a, 0xb3, 0x2b, 0x5a, 0x65, 0x88, 0x50, 0x7d, 0xca, 0xd2, 0x9d, 0x6d, 0x74,
	0x59, 0xaf, 0xc6, 0xd5, 0x1a, 0xdb, 0x50, 0x11, 0xa9, 0x5d, 0x55, 0xa4, 0x22, 0x52, 0xfc, 0x00,
	0xd5, 0x51, 0x4c, 0xc2, 0xae, 0x75, 0x59, 0xaf, 0x3d, 0xb0, 0xf4, 0x35, 0x7b, 0x47, 0xb2, 0xf0,
	0x74, 0x20, 0xae, 0x76, 0xb1, 0x0f, 0x75, 0x9f, 0xa2, 0x44, 0x6c, 0xed, 0xba, 0xf2, 0x61, 0xe9,
	0xd3, 0x54, 0x39, 0xcf, 0x8e, 0x9b, 0xdf, 0x0c, 0xea, 0x9c, 0x56, 0xf1, 0x81, 0xd0, 0x02, 0xa3,
	0xc8, 0xe2, 0x73, 0x3c, 0xb9, 0xd4, 0x24, 0x39, 0x47, 0x94, 0x4b, 0xbc, 0x86, 0x5a, 0x12, 0x2d,
	0x29, 0x51, 0x29, 0x4d, 0xae, 0x85, 0xa4, 0xa7, 0x98, 0x92, 0xb5, 0x4a, 0xca, 0xb8, 0x16, 0xd8,
	0x81, 0xc6, 0x2a, 0x4a, 0xd2, 0x2c, 0xa6, 0x5c, 0x05, 0x36, 0xf9, 0xb3, 0xc6, 0x77, 0x00, 0x22,
	0x15, 0x51, 0xb2, 0x10, 0xf1, 0x8e, 0x54, 0x4c, 0xc6, 0x4d, 0x45, 0xc2, 0x78, 0x47, 0xf8, 0x1e,
	0x5a, 0x7a, 0xfb, 0x27, 0xc5, 0x9b, 0xad, 0xb0, 0x2f, 0x95, 0xa1, 0xa9, 0xd8, 0xa3, 0x42, 0x37,
	0xbe, 0xcc, 0x9d, 0x17, 0x89, 0xc0, 0x1e, 0x5c, 0xd2, 0x5e, 0xa8, 0x6b, 0x58, 0xd7, 0xe8, 0x35,
	0x07, 0xed, 0xf2, 0xbd, 0xfa, 0x61, 0xbc, 0xdc, 0x96, 0x39, 0x29, 0xcb, 0xd2, 0xb2, 0x74, 0x2d,
	0xfa, 0x04, 0x8d, 0xb2, 0x40, 0x6c, 0x41, 0x63, 0xe8, 0x4e, 0xdd, 0xd9, 0x9d, 0x37, 0xb2, 0x2e,
	0xd0, 0x82, 0x96, 0x3f, 0xf9, 0xe2, 0x2f, 0xee, 0xf9, 0x3c, 0xf4, 0x26, 0x33, 0x8b, 0x61, 0x1b,
	0x40, 0x91, 0xf1, 0x64, 0xe8, 0x71, 0xab, 0x82, 0x4d, 0xb8, 0x9c, 0xce, 0x1f, 0x17, 0x63, 0x37,
	0xb4, 0x0c, 0x79, 0x58, 0x8a, 0x3b, 0x97, 0x0f, 0xad, 0xaa, 0xb4, 0x4a, 0x15, 0xcc, 0x47, 0x93,
	0x87, 0xaf, 0x56, 0xad, 0x1f, 0x01, 0xbc, 0xf4, 0x8f, 0xff, 0x41, 0xf3, 0xde, 0x73, 0x67, 0x0f,
	0xe1, 0x62, 0xcc, 0x3d, 0xcf, 0xba, 0xc0, 0xff, 0xe1, 0x2a, 0xe4, 0x9e, 0xb7, 0x78, 0x46, 0x4c,
	0xfe, 0x5f, 0x30, 0xff, 0xae, 0x55, 0x05, 0xaf, 0xc0, 0x1c, 0x4f, 0x02, 0x5f, 0x4b, 0x03, 0x11,
	0xda, 0x81, 0xef, 0x4d, 0xa7, 0x2f, 0xac, 0x3a, 0xf8, 0x0c, 0x8d, 0x61, 0x91, 0xc7, 0x7b, 0xca,
	0x73, 0x74, 0xa0, 0x1e, 0x50, 0x94, 0xad, 0xb6, 0x78, 0x5d, 0xd6, 0xf1, 0xfa, 0x53, 0xec, 0xbc,
	0x2a, 0x49, 0xb6, 0x38, 0xf8, 0x01, 0xe6, 0x28, 0x12, 0x51, 0x20, 0xd2, 0x8c, 0xfe, 0xf5, 0x30,
	0x7e, 0x04, 0xc3, 0x5d, 0xaf, 0xf1, 0x2f, 0xdc, 0x79, 0xe3, 0xe8, 0xe9, 0x70, 0xca, 0xe9, 0x70,
	0x3c, 0x39, 0x1d, 0x4b, 0x3d, 0x3c, 0x9f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xb9, 0x32,
	0xf9, 0x51, 0x03, 0x00, 0x00,
}
