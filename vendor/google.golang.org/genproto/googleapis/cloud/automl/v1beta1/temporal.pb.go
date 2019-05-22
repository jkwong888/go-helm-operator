// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/temporal.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A time period inside of an example that has a time dimension (e.g. video).
type TimeSegment struct {
	// Start of the time segment (inclusive), represented as the duration since
	// the example start.
	StartTimeOffset *duration.Duration `protobuf:"bytes,1,opt,name=start_time_offset,json=startTimeOffset,proto3" json:"start_time_offset,omitempty"`
	// End of the time segment (exclusive), represented as the duration since the
	// example start.
	EndTimeOffset        *duration.Duration `protobuf:"bytes,2,opt,name=end_time_offset,json=endTimeOffset,proto3" json:"end_time_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TimeSegment) Reset()         { *m = TimeSegment{} }
func (m *TimeSegment) String() string { return proto.CompactTextString(m) }
func (*TimeSegment) ProtoMessage()    {}
func (*TimeSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_temporal_20cabaedddc21db8, []int{0}
}
func (m *TimeSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSegment.Unmarshal(m, b)
}
func (m *TimeSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSegment.Marshal(b, m, deterministic)
}
func (dst *TimeSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSegment.Merge(dst, src)
}
func (m *TimeSegment) XXX_Size() int {
	return xxx_messageInfo_TimeSegment.Size(m)
}
func (m *TimeSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSegment.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSegment proto.InternalMessageInfo

func (m *TimeSegment) GetStartTimeOffset() *duration.Duration {
	if m != nil {
		return m.StartTimeOffset
	}
	return nil
}

func (m *TimeSegment) GetEndTimeOffset() *duration.Duration {
	if m != nil {
		return m.EndTimeOffset
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeSegment)(nil), "google.cloud.automl.v1beta1.TimeSegment")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/temporal.proto", fileDescriptor_temporal_20cabaedddc21db8)
}

var fileDescriptor_temporal_20cabaedddc21db8 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4a, 0x34, 0x31,
	0x14, 0xc5, 0xc9, 0x16, 0x5f, 0x91, 0xe5, 0x63, 0x71, 0x2a, 0xdd, 0x15, 0x15, 0x2b, 0xb1, 0x48,
	0x58, 0x2d, 0xad, 0xc6, 0x3f, 0x58, 0x89, 0xa2, 0x62, 0x21, 0x03, 0x4b, 0x66, 0xe7, 0x4e, 0x08,
	0x24, 0xb9, 0x43, 0xe6, 0xc6, 0x27, 0xf0, 0x1d, 0x7c, 0x27, 0x9f, 0x4a, 0x36, 0x89, 0x88, 0x20,
	0x5a, 0x1e, 0xce, 0xef, 0xfc, 0x48, 0x2e, 0x3f, 0xd6, 0x88, 0xda, 0x82, 0x5c, 0x5b, 0x8c, 0x9d,
	0x54, 0x91, 0xd0, 0x59, 0xf9, 0xb2, 0x6c, 0x81, 0xd4, 0x52, 0x12, 0xb8, 0x01, 0x83, 0xb2, 0x62,
	0x08, 0x48, 0x58, 0x2d, 0x32, 0x2b, 0x12, 0x2b, 0x32, 0x2b, 0x0a, 0x3b, 0xdf, 0x2d, 0x22, 0x35,
	0x18, 0xa9, 0xbc, 0x47, 0x52, 0x64, 0xd0, 0x8f, 0x79, 0x3a, 0xdf, 0x2b, 0x6d, 0x4a, 0x6d, 0xec,
	0x65, 0x17, 0x43, 0x02, 0x72, 0x7f, 0xf8, 0xc6, 0xf8, 0xf4, 0xd1, 0x38, 0x78, 0x00, 0xed, 0xc0,
	0x53, 0x75, 0xc5, 0xb7, 0x46, 0x52, 0x81, 0x56, 0x64, 0x1c, 0xac, 0xb0, 0xef, 0x47, 0xa0, 0x6d,
	0x76, 0xc0, 0x8e, 0xa6, 0x27, 0x3b, 0xa2, 0x3c, 0xe3, 0xd3, 0x25, 0x2e, 0x8b, 0xeb, 0x7e, 0x96,
	0x36, 0x1b, 0xcf, 0x6d, 0x5a, 0x54, 0x35, 0x9f, 0x81, 0xef, 0xbe, 0x49, 0x26, 0x7f, 0x49, 0xfe,
	0x83, 0xef, 0xbe, 0x14, 0xe7, 0xaf, 0x8c, 0xef, 0xaf, 0xd1, 0x89, 0x5f, 0xfe, 0x7e, 0xc7, 0x9e,
	0xeb, 0x52, 0x6b, 0xb4, 0xca, 0x6b, 0x81, 0x41, 0x4b, 0x0d, 0x3e, 0xc9, 0x65, 0xae, 0xd4, 0x60,
	0xc6, 0x1f, 0xaf, 0x7c, 0x96, 0xe3, 0xfb, 0x64, 0x71, 0x9d, 0xc0, 0xe6, 0x62, 0x03, 0x35, 0x75,
	0x24, 0xbc, 0xb1, 0xcd, 0x53, 0x86, 0xda, 0x7f, 0xc9, 0x75, 0xfa, 0x11, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x83, 0x10, 0xd5, 0xb0, 0x01, 0x00, 0x00,
}
