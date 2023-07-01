// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/bet_odds.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BetOdds is the type to store odds of a market.
type BetOdds struct {
	// uid is universal unique identifier of odds.
	// Required | Unique | uuid-v4 or code
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// market_uid is the parent, used for fast retrieving.
	// Required | NonUnique | -
	MarketUID string `protobuf:"bytes,2,opt,name=market_uid,proto3" json:"market_uid"`
	// value of the odds in corresponding odds type proposed in bet placement
	// message. Required | NonUnique | "1.286" or "2/7" or "+500"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// max_loss_multiplier is the factor for calculating max loss for given odds
	MaxLossMultiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_loss_multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_loss_multiplier"`
}

func (m *BetOdds) Reset()         { *m = BetOdds{} }
func (m *BetOdds) String() string { return proto.CompactTextString(m) }
func (*BetOdds) ProtoMessage()    {}
func (*BetOdds) Descriptor() ([]byte, []int) {
	return fileDescriptor_2629a03d0a23fb04, []int{0}
}

func (m *BetOdds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *BetOdds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetOdds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *BetOdds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetOdds.Merge(m, src)
}

func (m *BetOdds) XXX_Size() int {
	return m.Size()
}

func (m *BetOdds) XXX_DiscardUnknown() {
	xxx_messageInfo_BetOdds.DiscardUnknown(m)
}

var xxx_messageInfo_BetOdds proto.InternalMessageInfo

func (m *BetOdds) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *BetOdds) GetMarketUID() string {
	if m != nil {
		return m.MarketUID
	}
	return ""
}

func (m *BetOdds) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*BetOdds)(nil), "sgenetwork.sge.bet.BetOdds")
}

func init() { proto.RegisterFile("sge/bet/bet_odds.proto", fileDescriptor_2629a03d0a23fb04) }

var fileDescriptor_2629a03d0a23fb04 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4b, 0x3b, 0x31,
	0x18, 0xc6, 0xef, 0xda, 0xff, 0x5f, 0x69, 0x06, 0xc1, 0xb3, 0x48, 0x51, 0x48, 0x8a, 0x83, 0xb8,
	0xf4, 0x6e, 0x70, 0x74, 0x91, 0xd2, 0x45, 0xb0, 0x08, 0x07, 0x2e, 0x5d, 0x8e, 0xbb, 0xe6, 0x25,
	0x1e, 0xbd, 0x33, 0xe5, 0xde, 0x9c, 0xd6, 0xef, 0xa0, 0xe0, 0xc7, 0xea, 0xd8, 0x51, 0x1c, 0x82,
	0xa4, 0x5b, 0x3f, 0x85, 0x24, 0x2d, 0x72, 0x43, 0x87, 0x24, 0x4f, 0x9e, 0xfc, 0xf2, 0xe6, 0xe5,
	0x09, 0x39, 0x45, 0x01, 0x51, 0x06, 0xca, 0x8e, 0x44, 0x72, 0x8e, 0xe1, 0xbc, 0x92, 0x4a, 0x06,
	0x01, 0x0a, 0x78, 0x06, 0xf5, 0x2a, 0xab, 0x59, 0x88, 0x02, 0xc2, 0x0c, 0xd4, 0x59, 0x57, 0x48,
	0x21, 0xdd, 0x71, 0x64, 0xd5, 0x96, 0xbc, 0xf8, 0x68, 0x91, 0xc3, 0x21, 0xa8, 0x07, 0xce, 0x31,
	0xe8, 0x93, 0x76, 0x9d, 0xf3, 0x9e, 0xdf, 0xf7, 0xaf, 0x3a, 0xc3, 0x23, 0xa3, 0x59, 0xfb, 0xf1,
	0x6e, 0xb4, 0xd1, 0xcc, 0xba, 0xb1, 0x9d, 0x82, 0x1b, 0x42, 0xca, 0xb4, 0x9a, 0x81, 0x4a, 0x2c,
	0xd8, 0x72, 0xe0, 0xb9, 0xd1, 0xac, 0x33, 0x76, 0xee, 0x16, 0x6f, 0x20, 0x71, 0x43, 0x07, 0x5d,
	0xf2, 0xff, 0x25, 0x2d, 0x6a, 0xe8, 0xb5, 0xed, 0xbd, 0x78, 0xbb, 0x09, 0xde, 0x7d, 0x72, 0x52,
	0xa6, 0x8b, 0xa4, 0x90, 0x88, 0x49, 0x59, 0x17, 0x2a, 0x9f, 0x17, 0x39, 0x54, 0xbd, 0x7f, 0xae,
	0xf8, 0x64, 0xa9, 0x99, 0xf7, 0xad, 0xd9, 0xa5, 0xc8, 0xd5, 0x53, 0x9d, 0x85, 0x53, 0x59, 0x46,
	0x53, 0x89, 0xa5, 0xc4, 0xdd, 0x32, 0x40, 0x3e, 0x8b, 0xd4, 0xdb, 0x1c, 0x30, 0x1c, 0xc1, 0xd4,
	0x68, 0x76, 0x3c, 0x4e, 0x17, 0xf7, 0x12, 0x71, 0xfc, 0x57, 0x6a, 0xa3, 0xd9, 0xbe, 0x17, 0xe2,
	0x7d, 0xe6, 0xf0, 0x76, 0x69, 0xa8, 0xbf, 0x32, 0xd4, 0xff, 0x31, 0xd4, 0xff, 0x5c, 0x53, 0x6f,
	0xb5, 0xa6, 0xde, 0xd7, 0x9a, 0x7a, 0x93, 0x66, 0x0b, 0x28, 0x60, 0xb0, 0xcb, 0xd7, 0xea, 0x68,
	0xe1, 0x3e, 0xc1, 0xb5, 0x91, 0x1d, 0xb8, 0x60, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x87, 0x45, 0xa6, 0x9c, 0x01, 0x00, 0x00,
}

func (m *BetOdds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetOdds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetOdds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLossMultiplier.Size()
		i -= size
		if _, err := m.MaxLossMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBetOdds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MarketUID) > 0 {
		i -= len(m.MarketUID)
		copy(dAtA[i:], m.MarketUID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.MarketUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBetOdds(dAtA []byte, offset int, v uint64) int {
	offset -= sovBetOdds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *BetOdds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.MarketUID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = m.MaxLossMultiplier.Size()
	n += 1 + l + sovBetOdds(uint64(l))
	return n
}

func sovBetOdds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozBetOdds(x uint64) (n int) {
	return sovBetOdds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *BetOdds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBetOdds
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
			return fmt.Errorf("proto: BetOdds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetOdds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLossMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLossMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBetOdds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBetOdds
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

func skipBetOdds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
				return 0, ErrInvalidLengthBetOdds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBetOdds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBetOdds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBetOdds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBetOdds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBetOdds = fmt.Errorf("proto: unexpected end of group")
)
