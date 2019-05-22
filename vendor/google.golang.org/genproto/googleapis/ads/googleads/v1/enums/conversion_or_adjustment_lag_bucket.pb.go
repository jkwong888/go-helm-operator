// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/conversion_or_adjustment_lag_bucket.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Enum representing the number of days between the impression and the
// conversion or between the impression and adjustments to the conversion.
type ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket int32

const (
	// Not specified.
	ConversionOrAdjustmentLagBucketEnum_UNSPECIFIED ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionOrAdjustmentLagBucketEnum_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 1
	// Conversion lag bucket from 0 to 1 day. 0 day is included, 1 day is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_LESS_THAN_ONE_DAY ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 2
	// Conversion lag bucket from 1 to 2 days. 1 day is included, 2 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_ONE_TO_TWO_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 3
	// Conversion lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWO_TO_THREE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 4
	// Conversion lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THREE_TO_FOUR_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 5
	// Conversion lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FOUR_TO_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 6
	// Conversion lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FIVE_TO_SIX_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 7
	// Conversion lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SIX_TO_SEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 8
	// Conversion lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SEVEN_TO_EIGHT_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 9
	// Conversion lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_EIGHT_TO_NINE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 10
	// Conversion lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_NINE_TO_TEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 11
	// Conversion lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TEN_TO_ELEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 12
	// Conversion lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_ELEVEN_TO_TWELVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 13
	// Conversion lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWELVE_TO_THIRTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 14
	// Conversion lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 15
	// Conversion lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 16
	// Conversion lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 17
	// Conversion lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 18
	// Conversion lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 19
	// Conversion lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SIXTY_TO_NINETY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 20
	// Conversion adjustment lag bucket from 0 to 1 day. 0 day is included,
	// 1 day is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_LESS_THAN_ONE_DAY ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 21
	// Conversion adjustment lag bucket from 1 to 2 days. 1 day is included,
	// 2 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_ONE_TO_TWO_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 22
	// Conversion adjustment lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWO_TO_THREE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 23
	// Conversion adjustment lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THREE_TO_FOUR_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 24
	// Conversion adjustment lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FOUR_TO_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 25
	// Conversion adjustment lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FIVE_TO_SIX_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 26
	// Conversion adjustment lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SIX_TO_SEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 27
	// Conversion adjustment lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SEVEN_TO_EIGHT_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 28
	// Conversion adjustment lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_EIGHT_TO_NINE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 29
	// Conversion adjustment lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_NINE_TO_TEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 30
	// Conversion adjustment lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TEN_TO_ELEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 31
	// Conversion adjustment lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 32
	// Conversion adjustment lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 33
	// Conversion adjustment lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 34
	// Conversion adjustment lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 35
	// Conversion adjustment lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 36
	// Conversion adjustment lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 37
	// Conversion adjustment lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 38
	// Conversion adjustment lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SIXTY_TO_NINETY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 39
	// Conversion adjustment lag bucket from 90 to 145 days. 90 days is
	// included, 145 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 40
	// Conversion lag bucket UNKNOWN. This is for dates before conversion lag
	// bucket was available in Google Ads.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 41
	// Conversion adjustment lag bucket UNKNOWN. This is for dates before
	// conversion adjustment lag bucket was available in Google Ads.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 42
)

var ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "CONVERSION_LESS_THAN_ONE_DAY",
	3:  "CONVERSION_ONE_TO_TWO_DAYS",
	4:  "CONVERSION_TWO_TO_THREE_DAYS",
	5:  "CONVERSION_THREE_TO_FOUR_DAYS",
	6:  "CONVERSION_FOUR_TO_FIVE_DAYS",
	7:  "CONVERSION_FIVE_TO_SIX_DAYS",
	8:  "CONVERSION_SIX_TO_SEVEN_DAYS",
	9:  "CONVERSION_SEVEN_TO_EIGHT_DAYS",
	10: "CONVERSION_EIGHT_TO_NINE_DAYS",
	11: "CONVERSION_NINE_TO_TEN_DAYS",
	12: "CONVERSION_TEN_TO_ELEVEN_DAYS",
	13: "CONVERSION_ELEVEN_TO_TWELVE_DAYS",
	14: "CONVERSION_TWELVE_TO_THIRTEEN_DAYS",
	15: "CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS",
	16: "CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS",
	17: "CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS",
	18: "CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS",
	19: "CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS",
	20: "CONVERSION_SIXTY_TO_NINETY_DAYS",
	21: "ADJUSTMENT_LESS_THAN_ONE_DAY",
	22: "ADJUSTMENT_ONE_TO_TWO_DAYS",
	23: "ADJUSTMENT_TWO_TO_THREE_DAYS",
	24: "ADJUSTMENT_THREE_TO_FOUR_DAYS",
	25: "ADJUSTMENT_FOUR_TO_FIVE_DAYS",
	26: "ADJUSTMENT_FIVE_TO_SIX_DAYS",
	27: "ADJUSTMENT_SIX_TO_SEVEN_DAYS",
	28: "ADJUSTMENT_SEVEN_TO_EIGHT_DAYS",
	29: "ADJUSTMENT_EIGHT_TO_NINE_DAYS",
	30: "ADJUSTMENT_NINE_TO_TEN_DAYS",
	31: "ADJUSTMENT_TEN_TO_ELEVEN_DAYS",
	32: "ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS",
	33: "ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS",
	34: "ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS",
	35: "ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS",
	36: "ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS",
	37: "ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS",
	38: "ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS",
	39: "ADJUSTMENT_SIXTY_TO_NINETY_DAYS",
	40: "ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS",
	41: "CONVERSION_UNKNOWN",
	42: "ADJUSTMENT_UNKNOWN",
}
var ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_value = map[string]int32{
	"UNSPECIFIED":                                          0,
	"UNKNOWN":                                              1,
	"CONVERSION_LESS_THAN_ONE_DAY":                         2,
	"CONVERSION_ONE_TO_TWO_DAYS":                           3,
	"CONVERSION_TWO_TO_THREE_DAYS":                         4,
	"CONVERSION_THREE_TO_FOUR_DAYS":                        5,
	"CONVERSION_FOUR_TO_FIVE_DAYS":                         6,
	"CONVERSION_FIVE_TO_SIX_DAYS":                          7,
	"CONVERSION_SIX_TO_SEVEN_DAYS":                         8,
	"CONVERSION_SEVEN_TO_EIGHT_DAYS":                       9,
	"CONVERSION_EIGHT_TO_NINE_DAYS":                        10,
	"CONVERSION_NINE_TO_TEN_DAYS":                          11,
	"CONVERSION_TEN_TO_ELEVEN_DAYS":                        12,
	"CONVERSION_ELEVEN_TO_TWELVE_DAYS":                     13,
	"CONVERSION_TWELVE_TO_THIRTEEN_DAYS":                   14,
	"CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS":                 15,
	"CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS":               16,
	"CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS":                 17,
	"CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS":                 18,
	"CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS":                  19,
	"CONVERSION_SIXTY_TO_NINETY_DAYS":                      20,
	"ADJUSTMENT_LESS_THAN_ONE_DAY":                         21,
	"ADJUSTMENT_ONE_TO_TWO_DAYS":                           22,
	"ADJUSTMENT_TWO_TO_THREE_DAYS":                         23,
	"ADJUSTMENT_THREE_TO_FOUR_DAYS":                        24,
	"ADJUSTMENT_FOUR_TO_FIVE_DAYS":                         25,
	"ADJUSTMENT_FIVE_TO_SIX_DAYS":                          26,
	"ADJUSTMENT_SIX_TO_SEVEN_DAYS":                         27,
	"ADJUSTMENT_SEVEN_TO_EIGHT_DAYS":                       28,
	"ADJUSTMENT_EIGHT_TO_NINE_DAYS":                        29,
	"ADJUSTMENT_NINE_TO_TEN_DAYS":                          30,
	"ADJUSTMENT_TEN_TO_ELEVEN_DAYS":                        31,
	"ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS":                     32,
	"ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS":                   33,
	"ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS":                 34,
	"ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS":               35,
	"ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS":                 36,
	"ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS":                 37,
	"ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS":                  38,
	"ADJUSTMENT_SIXTY_TO_NINETY_DAYS":                      39,
	"ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS": 40,
	"CONVERSION_UNKNOWN":                                   41,
	"ADJUSTMENT_UNKNOWN":                                   42,
}

func (x ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) String() string {
	return proto.EnumName(ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_name, int32(x))
}
func (ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_or_adjustment_lag_bucket_ea9c79e631bbd90c, []int{0, 0}
}

// Container for enum representing the number of days between the impression and
// the conversion or between the impression and adjustments to the conversion.
type ConversionOrAdjustmentLagBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionOrAdjustmentLagBucketEnum) Reset()         { *m = ConversionOrAdjustmentLagBucketEnum{} }
func (m *ConversionOrAdjustmentLagBucketEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionOrAdjustmentLagBucketEnum) ProtoMessage()    {}
func (*ConversionOrAdjustmentLagBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_or_adjustment_lag_bucket_ea9c79e631bbd90c, []int{0}
}
func (m *ConversionOrAdjustmentLagBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum.Unmarshal(m, b)
}
func (m *ConversionOrAdjustmentLagBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionOrAdjustmentLagBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum.Merge(dst, src)
}
func (m *ConversionOrAdjustmentLagBucketEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum.Size(m)
}
func (m *ConversionOrAdjustmentLagBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionOrAdjustmentLagBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionOrAdjustmentLagBucketEnum)(nil), "google.ads.googleads.v1.enums.ConversionOrAdjustmentLagBucketEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket", ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_name, ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/conversion_or_adjustment_lag_bucket.proto", fileDescriptor_conversion_or_adjustment_lag_bucket_ea9c79e631bbd90c)
}

var fileDescriptor_conversion_or_adjustment_lag_bucket_ea9c79e631bbd90c = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xdf, 0x72, 0xd2, 0x40,
	0x14, 0xc6, 0x2d, 0xd5, 0x56, 0xb7, 0xd5, 0xe2, 0xaa, 0x55, 0xdb, 0x52, 0xca, 0x1f, 0xdb, 0xda,
	0x8b, 0x30, 0x8c, 0x5e, 0x38, 0xf1, 0x2a, 0x94, 0x2d, 0x44, 0x71, 0xd3, 0x21, 0x01, 0xc4, 0x61,
	0x66, 0x27, 0x2d, 0x4c, 0x06, 0x2d, 0x49, 0x87, 0x40, 0x1f, 0x48, 0xef, 0x7c, 0x14, 0x7d, 0x13,
	0xdf, 0xc0, 0x3b, 0xe7, 0x64, 0x59, 0x48, 0x36, 0xc1, 0xde, 0x74, 0x76, 0xf2, 0x7d, 0xf9, 0xce,
	0xc9, 0x39, 0xbf, 0xb2, 0xa8, 0xe6, 0x78, 0x9e, 0x73, 0x35, 0x28, 0xd9, 0x7d, 0xbf, 0xc4, 0x8f,
	0x70, 0xba, 0x29, 0x97, 0x06, 0xee, 0x74, 0xe4, 0x97, 0x2e, 0x3d, 0xf7, 0x66, 0x30, 0xf6, 0x87,
	0x9e, 0xcb, 0xbc, 0x31, 0xb3, 0xfb, 0x5f, 0xa7, 0xfe, 0x64, 0x34, 0x70, 0x27, 0xec, 0xca, 0x76,
	0xd8, 0xc5, 0xf4, 0xf2, 0xdb, 0x60, 0xa2, 0x5c, 0x8f, 0xbd, 0x89, 0x87, 0x33, 0xfc, 0x6d, 0xc5,
	0xee, 0xfb, 0xca, 0x3c, 0x48, 0xb9, 0x29, 0x2b, 0x41, 0xd0, 0xce, 0x9e, 0xa8, 0x73, 0x3d, 0x2c,
	0xd9, 0xae, 0xeb, 0x4d, 0xec, 0xc9, 0xd0, 0x73, 0x7d, 0xfe, 0x72, 0xfe, 0xf7, 0x26, 0x2a, 0x9c,
	0xce, 0x4b, 0x19, 0x63, 0x6d, 0x5e, 0xa8, 0x61, 0x3b, 0x95, 0xa0, 0x0c, 0x71, 0xa7, 0xa3, 0xfc,
	0x8f, 0x4d, 0x94, 0xbd, 0xc5, 0x87, 0xb7, 0xd0, 0x46, 0x8b, 0x9a, 0xe7, 0xe4, 0x54, 0x3f, 0xd3,
	0x49, 0x35, 0x7d, 0x07, 0x6f, 0xa0, 0xf5, 0x16, 0xfd, 0x48, 0x8d, 0x0e, 0x4d, 0xaf, 0xe0, 0x03,
	0xb4, 0x77, 0x6a, 0xd0, 0x36, 0x69, 0x9a, 0xba, 0x41, 0x59, 0x83, 0x98, 0x26, 0xb3, 0xea, 0x1a,
	0x65, 0x06, 0x25, 0xac, 0xaa, 0x75, 0xd3, 0x29, 0xbc, 0x8f, 0x76, 0x42, 0x0e, 0x78, 0x6e, 0x19,
	0xcc, 0xea, 0x18, 0x20, 0x9b, 0xe9, 0x55, 0x29, 0x01, 0x04, 0xd0, 0xeb, 0x4d, 0x42, 0xb8, 0xe3,
	0x2e, 0xce, 0xa1, 0x4c, 0xd8, 0x11, 0x48, 0x96, 0xc1, 0xce, 0x8c, 0x56, 0x93, 0x5b, 0xee, 0x49,
	0x21, 0x81, 0x02, 0x0e, 0xbd, 0x3d, 0x0b, 0x59, 0xc3, 0x59, 0xb4, 0x1b, 0x76, 0x80, 0x62, 0x19,
	0xcc, 0xd4, 0x3f, 0x73, 0xc3, 0xba, 0x14, 0x01, 0x02, 0xe8, 0xa4, 0x4d, 0x28, 0x77, 0xdc, 0xc7,
	0x79, 0xb4, 0x1f, 0x76, 0x04, 0x92, 0x65, 0x30, 0xa2, 0xd7, 0xea, 0x16, 0xf7, 0x3c, 0x90, 0x7a,
	0xe5, 0x92, 0x65, 0x30, 0xaa, 0xd3, 0x59, 0x27, 0x48, 0xea, 0x24, 0x50, 0xe0, 0x8b, 0x45, 0x9d,
	0x0d, 0xf9, 0x7b, 0x67, 0x55, 0x1a, 0x8b, 0x56, 0x36, 0x71, 0x11, 0x1d, 0x84, 0xcb, 0x34, 0x44,
	0x2f, 0x56, 0x87, 0x34, 0xc4, 0x37, 0x3f, 0xc4, 0x87, 0x28, 0x1f, 0x19, 0x6d, 0xa0, 0x05, 0xd3,
	0xd5, 0x9b, 0x16, 0x11, 0x69, 0x8f, 0xf0, 0x31, 0x2a, 0x46, 0x06, 0x3c, 0x53, 0x67, 0x33, 0x5e,
	0x38, 0xb7, 0xf0, 0x09, 0x3a, 0x94, 0xe6, 0x2c, 0x9c, 0x56, 0x87, 0x50, 0xab, 0x2b, 0xf6, 0x6e,
	0xa6, 0xd3, 0x72, 0xea, 0x42, 0x17, 0x1d, 0x74, 0xb9, 0xf3, 0x71, 0x52, 0xfd, 0x2e, 0xaf, 0x0e,
	0x87, 0xc5, 0x16, 0x31, 0x3e, 0x42, 0x85, 0x48, 0xfd, 0xb9, 0xce, 0x77, 0x29, 0x22, 0x9f, 0xe0,
	0x02, 0xca, 0x46, 0xb7, 0xc9, 0x13, 0x61, 0xda, 0xc2, 0xf4, 0x14, 0x56, 0xae, 0x55, 0x3f, 0xb4,
	0x4c, 0xeb, 0x13, 0xa1, 0x56, 0x02, 0xbc, 0xcf, 0x00, 0xde, 0x90, 0x43, 0x86, 0x77, 0x5b, 0x4a,
	0x88, 0xc3, 0xfb, 0x1c, 0x96, 0x19, 0x76, 0xc4, 0xe1, 0x7d, 0x21, 0x85, 0xc4, 0xe1, 0x7d, 0x09,
	0xc8, 0x84, 0x1d, 0x32, 0xbc, 0x3b, 0x52, 0x44, 0x1c, 0xde, 0x5d, 0x80, 0x37, 0xec, 0x48, 0x80,
	0x77, 0x4f, 0xea, 0x35, 0x01, 0xde, 0x8c, 0xd4, 0x49, 0x0c, 0xde, 0x7d, 0xf9, 0x7b, 0xe3, 0xf0,
	0x66, 0x01, 0xde, 0x70, 0x99, 0x44, 0x78, 0x0f, 0x00, 0xde, 0xc8, 0x68, 0x93, 0xe1, 0xcd, 0x01,
	0x3c, 0x91, 0x01, 0x2f, 0x83, 0x37, 0x0f, 0xf0, 0x4a, 0x73, 0x5e, 0x06, 0x6f, 0x41, 0x4e, 0x5d,
	0x0a, 0x6f, 0x31, 0xa9, 0x7e, 0x22, 0xbc, 0xaf, 0x00, 0xde, 0x48, 0xfd, 0x25, 0xf0, 0x1e, 0x02,
	0xbc, 0xd1, 0x6d, 0xc6, 0xe1, 0x3d, 0xc2, 0xef, 0xd0, 0x5b, 0x69, 0x13, 0xdc, 0x05, 0x4d, 0xd6,
	0x5b, 0xb4, 0xda, 0x24, 0x55, 0xa6, 0xd1, 0x6a, 0xac, 0x8f, 0x63, 0xbc, 0x8d, 0x70, 0xe8, 0x7f,
	0x43, 0xfc, 0x96, 0xbf, 0x86, 0xe7, 0xa1, 0x44, 0xf1, 0xfc, 0xa4, 0xf2, 0x77, 0x05, 0xe5, 0x2e,
	0xbd, 0x91, 0xf2, 0xdf, 0x1b, 0xa9, 0x52, 0xbc, 0xe5, 0x22, 0x39, 0x87, 0x9b, 0xe9, 0x7c, 0xe5,
	0x4b, 0x65, 0x16, 0xe3, 0x78, 0x57, 0xb6, 0xeb, 0x28, 0xde, 0xd8, 0x29, 0x39, 0x03, 0x37, 0xb8,
	0xb7, 0xc4, 0x8d, 0x79, 0x3d, 0xf4, 0x97, 0x5c, 0xa0, 0xef, 0x83, 0xbf, 0xdf, 0x53, 0xab, 0x35,
	0x4d, 0xfb, 0x99, 0xca, 0xd4, 0x78, 0x94, 0xd6, 0xf7, 0x15, 0x7e, 0x84, 0x53, 0xbb, 0xac, 0xc0,
	0xe5, 0xe6, 0xff, 0x12, 0x7a, 0x4f, 0xeb, 0xfb, 0xbd, 0xb9, 0xde, 0x6b, 0x97, 0x7b, 0x81, 0xfe,
	0x27, 0x95, 0xe3, 0x0f, 0x55, 0x55, 0xeb, 0xfb, 0xaa, 0x3a, 0x77, 0xa8, 0x6a, 0xbb, 0xac, 0xaa,
	0x81, 0xe7, 0x62, 0x2d, 0x68, 0xec, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xcd, 0x9e,
	0xae, 0xd8, 0x07, 0x00, 0x00,
}
