// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/pricefeed/v1beta1/pricefeed.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Market defines an asset in the pricefeed.
type Market struct {
	MarketID   string   `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	BaseAsset  string   `protobuf:"bytes,2,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string   `protobuf:"bytes,3,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
	Oracles    []string `protobuf:"bytes,4,rep,name=oracles,proto3" json:"oracles,omitempty"`
	Active     bool     `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Market) Reset()      { *m = Market{} }
func (*Market) ProtoMessage() {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_188809b7a903b86f, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *Market) GetBaseAsset() string {
	if m != nil {
		return m.BaseAsset
	}
	return ""
}

func (m *Market) GetQuoteAsset() string {
	if m != nil {
		return m.QuoteAsset
	}
	return ""
}

func (m *Market) GetOracles() []string {
	if m != nil {
		return m.Oracles
	}
	return nil
}

func (m *Market) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

// PostedPrice defines a price for market posted by a specific oracle.
type PostedPrice struct {
	MarketID      string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	OracleAddress string                                 `protobuf:"bytes,2,opt,name=oracle_address,json=oracleAddress,proto3" json:"oracle_address,omitempty" yaml:"oracle_address"`
	Price         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
	Expiry        time.Time                              `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry" yaml:"expiry"`
}

func (m *PostedPrice) Reset()      { *m = PostedPrice{} }
func (*PostedPrice) ProtoMessage() {}
func (*PostedPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_188809b7a903b86f, []int{1}
}
func (m *PostedPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostedPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostedPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostedPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostedPrice.Merge(m, src)
}
func (m *PostedPrice) XXX_Size() int {
	return m.Size()
}
func (m *PostedPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_PostedPrice.DiscardUnknown(m)
}

var xxx_messageInfo_PostedPrice proto.InternalMessageInfo

func (m *PostedPrice) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *PostedPrice) GetOracleAddress() string {
	if m != nil {
		return m.OracleAddress
	}
	return ""
}

func (m *PostedPrice) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

// CurrentPrice defines a current price for a particular market in the pricefeed
// module.
type CurrentPrice struct {
	MarketID string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
}

func (m *CurrentPrice) Reset()      { *m = CurrentPrice{} }
func (*CurrentPrice) ProtoMessage() {}
func (*CurrentPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_188809b7a903b86f, []int{2}
}
func (m *CurrentPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentPrice.Merge(m, src)
}
func (m *CurrentPrice) XXX_Size() int {
	return m.Size()
}
func (m *CurrentPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentPrice.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentPrice proto.InternalMessageInfo

func (m *CurrentPrice) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func init() {
	proto.RegisterType((*Market)(nil), "kava.pricefeed.v1beta1.Market")
	proto.RegisterType((*PostedPrice)(nil), "kava.pricefeed.v1beta1.PostedPrice")
	proto.RegisterType((*CurrentPrice)(nil), "kava.pricefeed.v1beta1.CurrentPrice")
}

func init() {
	proto.RegisterFile("kava/pricefeed/v1beta1/pricefeed.proto", fileDescriptor_188809b7a903b86f)
}

var fileDescriptor_188809b7a903b86f = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0xa6, 0x69, 0x48, 0x36, 0x29, 0x42, 0x16, 0x54, 0x6e, 0x24, 0xbc, 0x96, 0x0f, 0x55,
	0x10, 0xaa, 0xad, 0xc2, 0xad, 0x12, 0x88, 0x9a, 0x72, 0xe8, 0xa1, 0x52, 0x65, 0xf5, 0xc4, 0x25,
	0x5a, 0xdb, 0x5b, 0x63, 0x25, 0x66, 0x8d, 0x77, 0x13, 0x35, 0x7f, 0xd1, 0x0b, 0x12, 0x47, 0xee,
	0x7c, 0x01, 0x7f, 0xd0, 0x63, 0x8f, 0x88, 0x83, 0x41, 0xc9, 0x1f, 0xe4, 0x0b, 0xd0, 0xee, 0x38,
	0x4d, 0x10, 0x27, 0xa4, 0x9e, 0x32, 0xf3, 0xe6, 0xed, 0xcc, 0xbc, 0x79, 0x31, 0xde, 0x1f, 0xd1,
	0x29, 0xf5, 0x8b, 0x32, 0x8b, 0xd9, 0x25, 0x63, 0x89, 0x3f, 0x3d, 0x8c, 0x98, 0xa4, 0x87, 0x6b,
	0xc4, 0x2b, 0x4a, 0x2e, 0xb9, 0xb9, 0xab, 0x78, 0xde, 0x1a, 0xad, 0x79, 0xfd, 0xc7, 0x29, 0x4f,
	0xb9, 0xa6, 0xf8, 0x2a, 0x02, 0x76, 0x9f, 0xa4, 0x9c, 0xa7, 0x63, 0xe6, 0xeb, 0x2c, 0x9a, 0x5c,
	0xfa, 0x32, 0xcb, 0x99, 0x90, 0x34, 0x2f, 0x80, 0xe0, 0x7e, 0x43, 0xb8, 0x75, 0x46, 0xcb, 0x11,
	0x93, 0xe6, 0x33, 0xdc, 0xc9, 0x75, 0x34, 0xcc, 0x12, 0x0b, 0x39, 0x68, 0xd0, 0x09, 0x7a, 0xf3,
	0x8a, 0xb4, 0xa1, 0x7c, 0x7a, 0x12, 0xb6, 0xa1, 0x7c, 0x9a, 0x98, 0x4f, 0x31, 0x8e, 0xa8, 0x60,
	0x43, 0x2a, 0x04, 0x93, 0x56, 0x43, 0x71, 0xc3, 0x8e, 0x42, 0x8e, 0x15, 0x60, 0x12, 0xdc, 0xfd,
	0x34, 0xe1, 0x72, 0x55, 0xdf, 0xd2, 0x75, 0xac, 0x21, 0x20, 0x58, 0xf8, 0x01, 0x2f, 0x69, 0x3c,
	0x66, 0xc2, 0x6a, 0x3a, 0x5b, 0x83, 0x4e, 0xb8, 0x4a, 0xcd, 0x5d, 0xdc, 0xa2, 0xb1, 0xcc, 0xa6,
	0xcc, 0xda, 0x76, 0xd0, 0xa0, 0x1d, 0xd6, 0xd9, 0x51, 0xf3, 0xcb, 0x57, 0x62, 0xb8, 0xdf, 0x1b,
	0xb8, 0x7b, 0xce, 0x85, 0x64, 0xc9, 0xb9, 0x3a, 0x80, 0xf9, 0xea, 0xdf, 0x95, 0x9d, 0xcd, 0x95,
	0x97, 0x15, 0x79, 0x34, 0xa3, 0xf9, 0xf8, 0xc8, 0xbd, 0xa3, 0xb9, 0x1b, 0x32, 0xde, 0xe0, 0x87,
	0x30, 0x77, 0x48, 0x93, 0xa4, 0x64, 0x42, 0x80, 0x94, 0x60, 0x6f, 0x59, 0x91, 0x27, 0xf0, 0xee,
	0xef, 0xba, 0x1b, 0xee, 0x00, 0x70, 0x0c, 0xb9, 0x79, 0x81, 0xb7, 0xb5, 0x15, 0xa0, 0x31, 0x78,
	0x7d, 0x53, 0x11, 0xe3, 0x67, 0x45, 0xf6, 0xd3, 0x4c, 0x7e, 0x98, 0x44, 0x5e, 0xcc, 0x73, 0x3f,
	0xe6, 0x22, 0xe7, 0xa2, 0xfe, 0x39, 0x10, 0xc9, 0xc8, 0x97, 0xb3, 0x82, 0x09, 0xef, 0x84, 0xc5,
	0xcb, 0x8a, 0xf4, 0x60, 0x8c, 0x6e, 0xe2, 0x86, 0xd0, 0xcc, 0x3c, 0xc3, 0x2d, 0x76, 0x55, 0x64,
	0xe5, 0xcc, 0x6a, 0x3a, 0x68, 0xd0, 0x7d, 0xd1, 0xf7, 0xc0, 0x46, 0x6f, 0x65, 0xa3, 0x77, 0xb1,
	0xb2, 0x31, 0xd8, 0x53, 0x23, 0x97, 0x15, 0xd9, 0x81, 0x46, 0xf0, 0xce, 0xbd, 0xfe, 0x45, 0x50,
	0x58, 0x37, 0xa9, 0x6f, 0xf7, 0x19, 0xe1, 0xde, 0xdb, 0x49, 0x59, 0xb2, 0x8f, 0x12, 0x8e, 0xf7,
	0x1f, 0x7e, 0xdf, 0xc9, 0x6c, 0xdc, 0xa3, 0x4c, 0xd8, 0x2b, 0x78, 0x77, 0x33, 0xb7, 0xd1, 0xed,
	0xdc, 0x46, 0xbf, 0xe7, 0x36, 0xba, 0x5e, 0xd8, 0xc6, 0xed, 0xc2, 0x36, 0x7e, 0x2c, 0x6c, 0xe3,
	0xfd, 0xf3, 0x8d, 0xf6, 0xea, 0x5f, 0x7f, 0x30, 0xa6, 0x91, 0xd0, 0x91, 0x7f, 0xb5, 0xf1, 0xa5,
	0xe8, 0x39, 0x51, 0x4b, 0xdf, 0xe6, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x3c, 0x38,
	0xe2, 0x48, 0x03, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Oracles) > 0 {
		for iNdEx := len(m.Oracles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Oracles[iNdEx])
			copy(dAtA[i:], m.Oracles[iNdEx])
			i = encodeVarintPricefeed(dAtA, i, uint64(len(m.Oracles[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.QuoteAsset) > 0 {
		i -= len(m.QuoteAsset)
		copy(dAtA[i:], m.QuoteAsset)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.QuoteAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAsset) > 0 {
		i -= len(m.BaseAsset)
		copy(dAtA[i:], m.BaseAsset)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.BaseAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostedPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostedPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostedPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPricefeed(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPricefeed(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.OracleAddress) > 0 {
		i -= len(m.OracleAddress)
		copy(dAtA[i:], m.OracleAddress)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.OracleAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrentPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPricefeed(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPricefeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovPricefeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = len(m.BaseAsset)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = len(m.QuoteAsset)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	if len(m.Oracles) > 0 {
		for _, s := range m.Oracles {
			l = len(s)
			n += 1 + l + sovPricefeed(uint64(l))
		}
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *PostedPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = len(m.OracleAddress)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovPricefeed(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovPricefeed(uint64(l))
	return n
}

func (m *CurrentPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovPricefeed(uint64(l))
	return n
}

func sovPricefeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPricefeed(x uint64) (n int) {
	return sovPricefeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracles = append(m.Oracles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PostedPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PostedPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostedPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CurrentPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CurrentPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPricefeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPricefeed
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPricefeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPricefeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPricefeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPricefeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPricefeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPricefeed = fmt.Errorf("proto: unexpected end of group")
)
