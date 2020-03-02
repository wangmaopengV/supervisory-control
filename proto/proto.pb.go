// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package supervisory_controller_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SCErrorCode int32

const (
	SCErrorCode_SC_SUCCESS             SCErrorCode = 0
	SCErrorCode_SC_FAIL                SCErrorCode = 1
	SCErrorCode_SC_INSERT_DB_FAIL      SCErrorCode = 11
	SCErrorCode_SC_QUERY_DB_FAIL       SCErrorCode = 12
	SCErrorCode_SC_TRANSFORM_DATA_FAIL SCErrorCode = 41
)

var SCErrorCode_name = map[int32]string{
	0:  "SC_SUCCESS",
	1:  "SC_FAIL",
	11: "SC_INSERT_DB_FAIL",
	12: "SC_QUERY_DB_FAIL",
	41: "SC_TRANSFORM_DATA_FAIL",
}

var SCErrorCode_value = map[string]int32{
	"SC_SUCCESS":             0,
	"SC_FAIL":                1,
	"SC_INSERT_DB_FAIL":      11,
	"SC_QUERY_DB_FAIL":       12,
	"SC_TRANSFORM_DATA_FAIL": 41,
}

func (x SCErrorCode) String() string {
	return proto.EnumName(SCErrorCode_name, int32(x))
}

func (SCErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

type SCDeviceStatus int32

const (
	SCDeviceStatus_SC_STATUS_UNKNOWN SCDeviceStatus = 0
	SCDeviceStatus_SC_STATUS_ENABLE  SCDeviceStatus = 1
	SCDeviceStatus_SC_STATUS_DISABLE SCDeviceStatus = 2
	SCDeviceStatus_SC_STATUS_RESTART SCDeviceStatus = 3
)

var SCDeviceStatus_name = map[int32]string{
	0: "SC_STATUS_UNKNOWN",
	1: "SC_STATUS_ENABLE",
	2: "SC_STATUS_DISABLE",
	3: "SC_STATUS_RESTART",
}

var SCDeviceStatus_value = map[string]int32{
	"SC_STATUS_UNKNOWN": 0,
	"SC_STATUS_ENABLE":  1,
	"SC_STATUS_DISABLE": 2,
	"SC_STATUS_RESTART": 3,
}

func (x SCDeviceStatus) String() string {
	return proto.EnumName(SCDeviceStatus_name, int32(x))
}

func (SCDeviceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

type SCDeviceType int32

const (
	SCDeviceType_SC_Type_UNKNOWN SCDeviceType = 0
	SCDeviceType_SC_Type_LOCK    SCDeviceType = 1
)

var SCDeviceType_name = map[int32]string{
	0: "SC_Type_UNKNOWN",
	1: "SC_Type_LOCK",
}

var SCDeviceType_value = map[string]int32{
	"SC_Type_UNKNOWN": 0,
	"SC_Type_LOCK":    1,
}

func (x SCDeviceType) String() string {
	return proto.EnumName(SCDeviceType_name, int32(x))
}

func (SCDeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

type SCResult struct {
	Msg                  string      `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Code                 SCErrorCode `protobuf:"varint,2,opt,name=Code,proto3,enum=supervisory_controller_service.SCErrorCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SCResult) Reset()         { *m = SCResult{} }
func (m *SCResult) String() string { return proto.CompactTextString(m) }
func (*SCResult) ProtoMessage()    {}
func (*SCResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *SCResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCResult.Unmarshal(m, b)
}
func (m *SCResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCResult.Marshal(b, m, deterministic)
}
func (m *SCResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCResult.Merge(m, src)
}
func (m *SCResult) XXX_Size() int {
	return xxx_messageInfo_SCResult.Size(m)
}
func (m *SCResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SCResult.DiscardUnknown(m)
}

var xxx_messageInfo_SCResult proto.InternalMessageInfo

func (m *SCResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SCResult) GetCode() SCErrorCode {
	if m != nil {
		return m.Code
	}
	return SCErrorCode_SC_SUCCESS
}

type SCDeviceMeta struct {
	UID                  int64                `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Status               SCDeviceStatus       `protobuf:"varint,2,opt,name=Status,proto3,enum=supervisory_controller_service.SCDeviceStatus" json:"Status,omitempty"`
	Type                 SCDeviceType         `protobuf:"varint,3,opt,name=Type,proto3,enum=supervisory_controller_service.SCDeviceType" json:"Type,omitempty"`
	Meta                 string               `protobuf:"bytes,4,opt,name=Meta,proto3" json:"Meta,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,11,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SCDeviceMeta) Reset()         { *m = SCDeviceMeta{} }
func (m *SCDeviceMeta) String() string { return proto.CompactTextString(m) }
func (*SCDeviceMeta) ProtoMessage()    {}
func (*SCDeviceMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *SCDeviceMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCDeviceMeta.Unmarshal(m, b)
}
func (m *SCDeviceMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCDeviceMeta.Marshal(b, m, deterministic)
}
func (m *SCDeviceMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCDeviceMeta.Merge(m, src)
}
func (m *SCDeviceMeta) XXX_Size() int {
	return xxx_messageInfo_SCDeviceMeta.Size(m)
}
func (m *SCDeviceMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SCDeviceMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SCDeviceMeta proto.InternalMessageInfo

func (m *SCDeviceMeta) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *SCDeviceMeta) GetStatus() SCDeviceStatus {
	if m != nil {
		return m.Status
	}
	return SCDeviceStatus_SC_STATUS_UNKNOWN
}

func (m *SCDeviceMeta) GetType() SCDeviceType {
	if m != nil {
		return m.Type
	}
	return SCDeviceType_SC_Type_UNKNOWN
}

func (m *SCDeviceMeta) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *SCDeviceMeta) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

//
type QueryDeviceRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDeviceRequest) Reset()         { *m = QueryDeviceRequest{} }
func (m *QueryDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDeviceRequest) ProtoMessage()    {}
func (*QueryDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

func (m *QueryDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDeviceRequest.Unmarshal(m, b)
}
func (m *QueryDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDeviceRequest.Marshal(b, m, deterministic)
}
func (m *QueryDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeviceRequest.Merge(m, src)
}
func (m *QueryDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_QueryDeviceRequest.Size(m)
}
func (m *QueryDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeviceRequest proto.InternalMessageInfo

func (m *QueryDeviceRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type QueryDeviceResponse struct {
	Result               *SCResult     `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Device               *SCDeviceMeta `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueryDeviceResponse) Reset()         { *m = QueryDeviceResponse{} }
func (m *QueryDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDeviceResponse) ProtoMessage()    {}
func (*QueryDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{3}
}

func (m *QueryDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDeviceResponse.Unmarshal(m, b)
}
func (m *QueryDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDeviceResponse.Marshal(b, m, deterministic)
}
func (m *QueryDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeviceResponse.Merge(m, src)
}
func (m *QueryDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_QueryDeviceResponse.Size(m)
}
func (m *QueryDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeviceResponse proto.InternalMessageInfo

func (m *QueryDeviceResponse) GetResult() *SCResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryDeviceResponse) GetDevice() *SCDeviceMeta {
	if m != nil {
		return m.Device
	}
	return nil
}

//
type CensusDeviceRequest struct {
	StartTime            int64    `protobuf:"varint,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              int64    `protobuf:"varint,2,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CensusDeviceRequest) Reset()         { *m = CensusDeviceRequest{} }
func (m *CensusDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*CensusDeviceRequest) ProtoMessage()    {}
func (*CensusDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{4}
}

func (m *CensusDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CensusDeviceRequest.Unmarshal(m, b)
}
func (m *CensusDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CensusDeviceRequest.Marshal(b, m, deterministic)
}
func (m *CensusDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CensusDeviceRequest.Merge(m, src)
}
func (m *CensusDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_CensusDeviceRequest.Size(m)
}
func (m *CensusDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CensusDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CensusDeviceRequest proto.InternalMessageInfo

func (m *CensusDeviceRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *CensusDeviceRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type CensusDeviceResponse struct {
	Result               *SCResult       `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	TotalDevice          int32           `protobuf:"varint,2,opt,name=TotalDevice,proto3" json:"TotalDevice,omitempty"`
	Devices              []*SCDeviceMeta `protobuf:"bytes,3,rep,name=Devices,proto3" json:"Devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CensusDeviceResponse) Reset()         { *m = CensusDeviceResponse{} }
func (m *CensusDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*CensusDeviceResponse) ProtoMessage()    {}
func (*CensusDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{5}
}

func (m *CensusDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CensusDeviceResponse.Unmarshal(m, b)
}
func (m *CensusDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CensusDeviceResponse.Marshal(b, m, deterministic)
}
func (m *CensusDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CensusDeviceResponse.Merge(m, src)
}
func (m *CensusDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_CensusDeviceResponse.Size(m)
}
func (m *CensusDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CensusDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CensusDeviceResponse proto.InternalMessageInfo

func (m *CensusDeviceResponse) GetResult() *SCResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CensusDeviceResponse) GetTotalDevice() int32 {
	if m != nil {
		return m.TotalDevice
	}
	return 0
}

func (m *CensusDeviceResponse) GetDevices() []*SCDeviceMeta {
	if m != nil {
		return m.Devices
	}
	return nil
}

//
type QueryDeviceMonitorRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDeviceMonitorRequest) Reset()         { *m = QueryDeviceMonitorRequest{} }
func (m *QueryDeviceMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDeviceMonitorRequest) ProtoMessage()    {}
func (*QueryDeviceMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{6}
}

func (m *QueryDeviceMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDeviceMonitorRequest.Unmarshal(m, b)
}
func (m *QueryDeviceMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDeviceMonitorRequest.Marshal(b, m, deterministic)
}
func (m *QueryDeviceMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeviceMonitorRequest.Merge(m, src)
}
func (m *QueryDeviceMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_QueryDeviceMonitorRequest.Size(m)
}
func (m *QueryDeviceMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeviceMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeviceMonitorRequest proto.InternalMessageInfo

func (m *QueryDeviceMonitorRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type QueryDeviceMonitorResponse struct {
	Result               *SCResult              `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Device               *SCDeviceMeta          `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
	Login                []*timestamp.Timestamp `protobuf:"bytes,3,rep,name=Login,proto3" json:"Login,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *QueryDeviceMonitorResponse) Reset()         { *m = QueryDeviceMonitorResponse{} }
func (m *QueryDeviceMonitorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDeviceMonitorResponse) ProtoMessage()    {}
func (*QueryDeviceMonitorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{7}
}

func (m *QueryDeviceMonitorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDeviceMonitorResponse.Unmarshal(m, b)
}
func (m *QueryDeviceMonitorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDeviceMonitorResponse.Marshal(b, m, deterministic)
}
func (m *QueryDeviceMonitorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeviceMonitorResponse.Merge(m, src)
}
func (m *QueryDeviceMonitorResponse) XXX_Size() int {
	return xxx_messageInfo_QueryDeviceMonitorResponse.Size(m)
}
func (m *QueryDeviceMonitorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeviceMonitorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeviceMonitorResponse proto.InternalMessageInfo

func (m *QueryDeviceMonitorResponse) GetResult() *SCResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryDeviceMonitorResponse) GetDevice() *SCDeviceMeta {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *QueryDeviceMonitorResponse) GetLogin() []*timestamp.Timestamp {
	if m != nil {
		return m.Login
	}
	return nil
}

func init() {
	proto.RegisterEnum("supervisory_controller_service.SCErrorCode", SCErrorCode_name, SCErrorCode_value)
	proto.RegisterEnum("supervisory_controller_service.SCDeviceStatus", SCDeviceStatus_name, SCDeviceStatus_value)
	proto.RegisterEnum("supervisory_controller_service.SCDeviceType", SCDeviceType_name, SCDeviceType_value)
	proto.RegisterType((*SCResult)(nil), "supervisory_controller_service.SCResult")
	proto.RegisterType((*SCDeviceMeta)(nil), "supervisory_controller_service.SCDeviceMeta")
	proto.RegisterType((*QueryDeviceRequest)(nil), "supervisory_controller_service.QueryDeviceRequest")
	proto.RegisterType((*QueryDeviceResponse)(nil), "supervisory_controller_service.QueryDeviceResponse")
	proto.RegisterType((*CensusDeviceRequest)(nil), "supervisory_controller_service.CensusDeviceRequest")
	proto.RegisterType((*CensusDeviceResponse)(nil), "supervisory_controller_service.CensusDeviceResponse")
	proto.RegisterType((*QueryDeviceMonitorRequest)(nil), "supervisory_controller_service.QueryDeviceMonitorRequest")
	proto.RegisterType((*QueryDeviceMonitorResponse)(nil), "supervisory_controller_service.QueryDeviceMonitorResponse")
}

func init() {
	proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8)
}

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x6b, 0xdb, 0x48,
	0x14, 0x8e, 0x6c, 0xc7, 0xd9, 0x3c, 0x85, 0xac, 0x76, 0x92, 0x5d, 0xb4, 0x22, 0xec, 0x1a, 0x1d,
	0x16, 0x6f, 0x76, 0x63, 0xef, 0x3a, 0xed, 0xa1, 0xb9, 0x34, 0x8e, 0x2c, 0x83, 0x89, 0xed, 0x90,
	0x19, 0x99, 0xd2, 0x43, 0x11, 0x4a, 0x32, 0x35, 0xa6, 0xb6, 0xc6, 0xd5, 0x8c, 0x03, 0x26, 0xe4,
	0xd2, 0x5f, 0x50, 0xe8, 0xb9, 0xa7, 0x9e, 0xfb, 0x27, 0xfa, 0x17, 0x0a, 0xfd, 0x05, 0xfd, 0x15,
	0x85, 0x42, 0xd1, 0x8c, 0xec, 0xc8, 0x24, 0x8d, 0xe3, 0xd2, 0x43, 0x2f, 0x66, 0xfc, 0xde, 0xf7,
	0xbd, 0xf7, 0xbd, 0x6f, 0x34, 0x0f, 0xf4, 0x61, 0xc4, 0x04, 0x2b, 0xc9, 0x5f, 0xf4, 0x07, 0x1f,
	0x0d, 0x69, 0x74, 0xde, 0xe3, 0x2c, 0x1a, 0xfb, 0xa7, 0x2c, 0x14, 0x11, 0xeb, 0xf7, 0x69, 0xe4,
	0xf3, 0x38, 0x78, 0x4a, 0xad, 0x3f, 0xbb, 0x8c, 0x75, 0xfb, 0xb4, 0x2c, 0xd1, 0x27, 0xa3, 0xa7,
	0x65, 0xd1, 0x1b, 0x50, 0x2e, 0x82, 0xc1, 0x50, 0x15, 0xb0, 0xb6, 0x12, 0x40, 0x30, 0xec, 0x95,
	0x83, 0x30, 0x64, 0x22, 0x10, 0x3d, 0x16, 0x72, 0x95, 0xb5, 0x9f, 0xc0, 0x4f, 0xc4, 0xc1, 0x94,
	0x8f, 0xfa, 0x02, 0x19, 0x90, 0x6d, 0xf1, 0xae, 0xa9, 0x15, 0xb4, 0xe2, 0x2a, 0x8e, 0x8f, 0xe8,
	0x21, 0xe4, 0x1c, 0x76, 0x46, 0xcd, 0x4c, 0x41, 0x2b, 0xae, 0x57, 0xfe, 0x29, 0xdd, 0xae, 0xa5,
	0x44, 0x1c, 0x37, 0x8a, 0x58, 0x14, 0x53, 0xb0, 0x24, 0xda, 0x9f, 0x34, 0x58, 0x23, 0x4e, 0x8d,
	0xc6, 0xe9, 0x16, 0x15, 0x41, 0xdc, 0xa3, 0xd3, 0xa8, 0xc9, 0x1e, 0x59, 0x1c, 0x1f, 0x51, 0x1d,
	0xf2, 0x44, 0x04, 0x62, 0xc4, 0x93, 0x2e, 0xa5, 0xf9, 0x5d, 0x54, 0x3d, 0xc5, 0xc2, 0x09, 0x1b,
	0xed, 0x43, 0xce, 0x1b, 0x0f, 0xa9, 0x99, 0x95, 0x55, 0xfe, 0xbd, 0x6b, 0x95, 0x98, 0x83, 0x25,
	0x13, 0x21, 0xc8, 0xc5, 0x1a, 0xcd, 0x9c, 0x34, 0x40, 0x9e, 0xd1, 0x1e, 0x80, 0x13, 0xd1, 0x40,
	0x50, 0xaf, 0x37, 0xa0, 0xa6, 0x5e, 0xd0, 0x8a, 0x7a, 0xc5, 0x2a, 0x29, 0x4b, 0x4b, 0x13, 0xcf,
	0x4b, 0xde, 0xc4, 0x73, 0x9c, 0x42, 0xdb, 0x7f, 0x01, 0x3a, 0x1e, 0xd1, 0x68, 0xac, 0x1a, 0x61,
	0xfa, 0x7c, 0x44, 0xb9, 0xb8, 0xee, 0x80, 0xfd, 0x5a, 0x83, 0x8d, 0x19, 0x20, 0x1f, 0xb2, 0x90,
	0x53, 0xb4, 0x0f, 0x79, 0x75, 0x33, 0x12, 0xac, 0x57, 0x8a, 0xf3, 0x67, 0x52, 0x78, 0x9c, 0xf0,
	0x50, 0x0d, 0xf2, 0xaa, 0xa6, 0xf4, 0x56, 0xbf, 0xbb, 0x2b, 0xf1, 0xec, 0x38, 0xe1, 0xda, 0x2d,
	0xd8, 0x70, 0x68, 0xc8, 0x47, 0x7c, 0x76, 0x90, 0x2d, 0x58, 0x25, 0x22, 0x88, 0x84, 0x74, 0x46,
	0x8d, 0x73, 0x15, 0x40, 0x26, 0xac, 0xb8, 0xe1, 0x99, 0xcc, 0x65, 0x64, 0x6e, 0xf2, 0xd7, 0x7e,
	0xa7, 0xc1, 0xe6, 0x6c, 0xbd, 0xef, 0x36, 0x6f, 0x01, 0x74, 0x8f, 0x89, 0xa0, 0x9f, 0x1a, 0x7a,
	0x19, 0xa7, 0x43, 0xa8, 0x0e, 0x2b, 0xea, 0xc4, 0xcd, 0x6c, 0x21, 0xbb, 0xb0, 0x25, 0x13, 0xb2,
	0xbd, 0x03, 0xbf, 0xa7, 0xae, 0xac, 0xc5, 0xc2, 0x9e, 0x60, 0xd1, 0xd7, 0xaf, 0xf8, 0x83, 0x06,
	0xd6, 0x4d, 0xf8, 0x1f, 0xeb, 0xa6, 0xd1, 0x7f, 0xb0, 0xdc, 0x64, 0xdd, 0x5e, 0x98, 0x78, 0x73,
	0xdb, 0x87, 0xae, 0x80, 0xdb, 0x1c, 0xf4, 0xd4, 0xab, 0x47, 0xeb, 0x00, 0xc4, 0xf1, 0x49, 0xc7,
	0x71, 0x5c, 0x42, 0x8c, 0x25, 0xa4, 0xc3, 0x0a, 0x71, 0xfc, 0x7a, 0xb5, 0xd1, 0x34, 0x34, 0xf4,
	0x2b, 0xfc, 0x42, 0x1c, 0xbf, 0xd1, 0x26, 0x2e, 0xf6, 0xfc, 0xda, 0x81, 0x0a, 0xeb, 0x68, 0x13,
	0x0c, 0xe2, 0xf8, 0xc7, 0x1d, 0x17, 0x3f, 0x9e, 0x46, 0xd7, 0x90, 0x05, 0xbf, 0x11, 0xc7, 0xf7,
	0x70, 0xb5, 0x4d, 0xea, 0x47, 0xb8, 0xe5, 0xd7, 0xaa, 0x5e, 0x55, 0xe5, 0xfe, 0xde, 0x7e, 0x06,
	0xeb, 0xb3, 0x4b, 0x20, 0x29, 0x4d, 0xbc, 0xaa, 0xd7, 0x21, 0x7e, 0xa7, 0x7d, 0xd8, 0x3e, 0x7a,
	0xd4, 0x36, 0x96, 0x92, 0xd2, 0x49, 0xd8, 0x6d, 0x57, 0x0f, 0x9a, 0xee, 0x54, 0x47, 0x12, 0xad,
	0x35, 0x88, 0x0c, 0x67, 0x66, 0xc3, 0xd8, 0x25, 0x5e, 0x15, 0x7b, 0x46, 0x76, 0xfb, 0xfe, 0xd5,
	0x06, 0x93, 0x5b, 0x62, 0x03, 0x7e, 0x8e, 0x85, 0x8d, 0x87, 0x34, 0xd5, 0xc8, 0x88, 0x41, 0x2a,
	0xd8, 0x3c, 0x72, 0x0e, 0x0d, 0xad, 0xf2, 0x39, 0x0b, 0x88, 0x38, 0x53, 0xdf, 0x89, 0xb2, 0x1d,
	0xbd, 0xd4, 0x40, 0x4f, 0x7d, 0x08, 0xa8, 0x32, 0xef, 0x9e, 0xae, 0x6f, 0x10, 0x6b, 0x77, 0x21,
	0x8e, 0xfa, 0xc4, 0x6c, 0xf3, 0xc5, 0xfb, 0x8f, 0xaf, 0x32, 0x08, 0x19, 0xe5, 0xf3, 0xff, 0xcb,
	0x67, 0x32, 0x57, 0xbe, 0xe8, 0x34, 0x6a, 0x97, 0xe8, 0x8d, 0x06, 0x6b, 0xe9, 0xf7, 0x88, 0xe6,
	0xd6, 0xbf, 0x61, 0x1b, 0x58, 0xf7, 0x16, 0x23, 0x25, 0xaa, 0x8a, 0x52, 0x95, 0x8d, 0x0a, 0x69,
	0x55, 0xd3, 0x25, 0x72, 0x59, 0xbe, 0x48, 0x96, 0xc6, 0x25, 0x7a, 0xab, 0xcd, 0x6c, 0xd3, 0xe4,
	0x05, 0xa1, 0x07, 0x0b, 0x78, 0x31, 0xfb, 0x4a, 0xad, 0xbd, 0x6f, 0xa1, 0x26, 0xba, 0x0b, 0x52,
	0xb7, 0x85, 0xcc, 0x2b, 0xdd, 0x3b, 0x03, 0x85, 0x51, 0xae, 0x9e, 0xe4, 0xe5, 0x9b, 0xd9, 0xfd,
	0x12, 0x00, 0x00, 0xff, 0xff, 0xce, 0x48, 0x7d, 0xd0, 0xcd, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SControllerServiceClient is the client API for SControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SControllerServiceClient interface {
	QueryDevice(ctx context.Context, in *QueryDeviceRequest, opts ...grpc.CallOption) (*QueryDeviceResponse, error)
	CensusDevice(ctx context.Context, in *CensusDeviceRequest, opts ...grpc.CallOption) (*CensusDeviceResponse, error)
	QueryDeviceMonitor(ctx context.Context, in *QueryDeviceMonitorRequest, opts ...grpc.CallOption) (*QueryDeviceMonitorResponse, error)
}

type sControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSControllerServiceClient(cc grpc.ClientConnInterface) SControllerServiceClient {
	return &sControllerServiceClient{cc}
}

func (c *sControllerServiceClient) QueryDevice(ctx context.Context, in *QueryDeviceRequest, opts ...grpc.CallOption) (*QueryDeviceResponse, error) {
	out := new(QueryDeviceResponse)
	err := c.cc.Invoke(ctx, "/supervisory_controller_service.SControllerService/QueryDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sControllerServiceClient) CensusDevice(ctx context.Context, in *CensusDeviceRequest, opts ...grpc.CallOption) (*CensusDeviceResponse, error) {
	out := new(CensusDeviceResponse)
	err := c.cc.Invoke(ctx, "/supervisory_controller_service.SControllerService/CensusDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sControllerServiceClient) QueryDeviceMonitor(ctx context.Context, in *QueryDeviceMonitorRequest, opts ...grpc.CallOption) (*QueryDeviceMonitorResponse, error) {
	out := new(QueryDeviceMonitorResponse)
	err := c.cc.Invoke(ctx, "/supervisory_controller_service.SControllerService/QueryDeviceMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SControllerServiceServer is the server API for SControllerService service.
type SControllerServiceServer interface {
	QueryDevice(context.Context, *QueryDeviceRequest) (*QueryDeviceResponse, error)
	CensusDevice(context.Context, *CensusDeviceRequest) (*CensusDeviceResponse, error)
	QueryDeviceMonitor(context.Context, *QueryDeviceMonitorRequest) (*QueryDeviceMonitorResponse, error)
}

// UnimplementedSControllerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSControllerServiceServer struct {
}

func (*UnimplementedSControllerServiceServer) QueryDevice(ctx context.Context, req *QueryDeviceRequest) (*QueryDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDevice not implemented")
}
func (*UnimplementedSControllerServiceServer) CensusDevice(ctx context.Context, req *CensusDeviceRequest) (*CensusDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CensusDevice not implemented")
}
func (*UnimplementedSControllerServiceServer) QueryDeviceMonitor(ctx context.Context, req *QueryDeviceMonitorRequest) (*QueryDeviceMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceMonitor not implemented")
}

func RegisterSControllerServiceServer(s *grpc.Server, srv SControllerServiceServer) {
	s.RegisterService(&_SControllerService_serviceDesc, srv)
}

func _SControllerService_QueryDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SControllerServiceServer).QueryDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisory_controller_service.SControllerService/QueryDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SControllerServiceServer).QueryDevice(ctx, req.(*QueryDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SControllerService_CensusDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CensusDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SControllerServiceServer).CensusDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisory_controller_service.SControllerService/CensusDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SControllerServiceServer).CensusDevice(ctx, req.(*CensusDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SControllerService_QueryDeviceMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SControllerServiceServer).QueryDeviceMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisory_controller_service.SControllerService/QueryDeviceMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SControllerServiceServer).QueryDeviceMonitor(ctx, req.(*QueryDeviceMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SControllerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisory_controller_service.SControllerService",
	HandlerType: (*SControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryDevice",
			Handler:    _SControllerService_QueryDevice_Handler,
		},
		{
			MethodName: "CensusDevice",
			Handler:    _SControllerService_CensusDevice_Handler,
		},
		{
			MethodName: "QueryDeviceMonitor",
			Handler:    _SControllerService_QueryDeviceMonitor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
