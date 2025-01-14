// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/orderbook/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the orderbook module's genesis state.
type GenesisState struct {
	// params defines all the parameters related to order book.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// order_book_list defines the order books available at genesis.
	OrderBookList []OrderBook `protobuf:"bytes,2,rep,name=order_book_list,json=orderBookList,proto3" json:"order_book_list"`
	// order_book_participation_list defines the book participations available at
	// genesis.
	OrderBookParticipationList []OrderBookParticipation `protobuf:"bytes,3,rep,name=order_book_participation_list,json=orderBookParticipationList,proto3" json:"order_book_participation_list"`
	// order_book_exposure_list defines the order book exposures available at
	// genesis.
	OrderBookExposureList []OrderBookOddsExposure `protobuf:"bytes,4,rep,name=order_book_exposure_list,json=orderBookExposureList,proto3" json:"order_book_exposure_list"`
	// participation_exposure_list defines the participation exposures available
	// at genesis.
	ParticipationExposureList []ParticipationExposure `protobuf:"bytes,5,rep,name=participation_exposure_list,json=participationExposureList,proto3" json:"participation_exposure_list"`
	// participation_exposure_by_index_list defines the participation exposures by
	// the indices available at genesis.
	ParticipationExposureByIndexList []ParticipationExposure `protobuf:"bytes,6,rep,name=participation_exposure_by_index_list,json=participationExposureByIndexList,proto3" json:"participation_exposure_by_index_list"`
	// historical_participation_exposure_list defines the historical participation
	// exposures available at genesis.
	HistoricalParticipationExposureList []ParticipationExposure `protobuf:"bytes,7,rep,name=historical_participation_exposure_list,json=historicalParticipationExposureList,proto3" json:"historical_participation_exposure_list"`
	// historical_participation_exposure_list defines the participation bet pair
	// exposures available at genesis.
	ParticipationBetPairExposureList []ParticipationBetPair `protobuf:"bytes,8,rep,name=participation_bet_pair_exposure_list,json=participationBetPairExposureList,proto3" json:"participation_bet_pair_exposure_list"`
	// stats is the statistics of the order book.
	Stats OrderBookStats `protobuf:"bytes,9,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54e9379cfb7d94d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOrderBookList() []OrderBook {
	if m != nil {
		return m.OrderBookList
	}
	return nil
}

func (m *GenesisState) GetOrderBookParticipationList() []OrderBookParticipation {
	if m != nil {
		return m.OrderBookParticipationList
	}
	return nil
}

func (m *GenesisState) GetOrderBookExposureList() []OrderBookOddsExposure {
	if m != nil {
		return m.OrderBookExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureByIndexList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureByIndexList
	}
	return nil
}

func (m *GenesisState) GetHistoricalParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.HistoricalParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationBetPairExposureList() []ParticipationBetPair {
	if m != nil {
		return m.ParticipationBetPairExposureList
	}
	return nil
}

func (m *GenesisState) GetStats() OrderBookStats {
	if m != nil {
		return m.Stats
	}
	return OrderBookStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.orderbook.GenesisState")
}

func init() { proto.RegisterFile("sge/orderbook/genesis.proto", fileDescriptor_b54e9379cfb7d94d) }

var fileDescriptor_b54e9379cfb7d94d = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xb6, 0xdd, 0xea, 0x54, 0x11, 0x82, 0x42, 0x9a, 0xda, 0x18, 0xad, 0xc8, 0x1e,
	0x34, 0x91, 0x7a, 0xf7, 0x10, 0xfc, 0x81, 0x20, 0x34, 0xea, 0xcd, 0x4b, 0x98, 0x6c, 0x86, 0x74,
	0xd8, 0x36, 0x2f, 0xcc, 0xcc, 0xe2, 0xe6, 0xae, 0x17, 0x4f, 0xfe, 0x59, 0x3d, 0xf6, 0xe8, 0x49,
	0x64, 0x17, 0xff, 0x0f, 0xc9, 0xcc, 0x64, 0xcd, 0xb4, 0x89, 0x06, 0x7a, 0x9b, 0xdd, 0xf7, 0x7d,
	0xdf, 0xcf, 0xf7, 0x3d, 0x32, 0x83, 0xf6, 0x78, 0x4e, 0x42, 0x60, 0x19, 0x61, 0x29, 0xc0, 0x2c,
	0xcc, 0x49, 0x41, 0x38, 0xe5, 0x41, 0xc9, 0x40, 0x80, 0xed, 0xf0, 0xfa, 0xb7, 0xf8, 0x0c, 0x6c,
	0x16, 0xf0, 0x9c, 0x04, 0x6b, 0x9d, 0x7b, 0x27, 0x87, 0x1c, 0xa4, 0x28, 0xac, 0x4f, 0x4a, 0xef,
	0xba, 0xa6, 0x59, 0x89, 0x19, 0x3e, 0xd5, 0x5e, 0xee, 0xbe, 0x59, 0x5b, 0x9f, 0x74, 0xf9, 0xc1,
	0xa5, 0x56, 0x41, 0xa7, 0xb4, 0xc4, 0x82, 0x42, 0xa1, 0x25, 0xbb, 0xa6, 0x84, 0x0b, 0x2c, 0x1a,
	0xf3, 0x7b, 0x66, 0x89, 0x2c, 0x4a, 0xe0, 0x73, 0x46, 0x54, 0xf5, 0xe1, 0xef, 0x6d, 0x74, 0xf3,
	0x8d, 0x1a, 0xec, 0xa3, 0xc0, 0x82, 0xd8, 0x2f, 0xd0, 0x58, 0x65, 0x73, 0x2c, 0xdf, 0x9a, 0xec,
	0x1c, 0xfa, 0x41, 0xdf, 0xa0, 0x41, 0x2c, 0x75, 0xd1, 0xe6, 0xd9, 0xcf, 0xfb, 0xa3, 0x0f, 0xba,
	0xcb, 0x7e, 0x8f, 0x6e, 0x4b, 0x45, 0x52, 0x4b, 0x92, 0x13, 0xca, 0x85, 0x73, 0xcd, 0xdf, 0x98,
	0xec, 0x1c, 0x1e, 0xf4, 0x1b, 0x1d, 0xd5, 0xa7, 0x08, 0x60, 0xa6, 0xbd, 0x6e, 0x41, 0xf3, 0xc7,
	0x3b, 0xca, 0x85, 0x5d, 0xa1, 0xfd, 0x96, 0xa5, 0x31, 0xbe, 0x02, 0x6c, 0x48, 0xc0, 0xb3, 0x01,
	0x80, 0xb8, 0xdd, 0xac, 0x69, 0x2e, 0x74, 0x56, 0x25, 0xba, 0x40, 0x4e, 0x0b, 0xdd, 0xec, 0x4e,
	0x51, 0x37, 0x25, 0x35, 0x1c, 0x40, 0x3d, 0xca, 0x32, 0xfe, 0x4a, 0xf7, 0x6a, 0xe8, 0xdd, 0x35,
	0xb4, 0x29, 0x48, 0xde, 0x1c, 0xed, 0x99, 0xf3, 0x99, 0xc8, 0xad, 0xff, 0x21, 0x8d, 0x09, 0x2e,
	0x20, 0x77, 0xcb, 0xae, 0xa2, 0xc4, 0x7e, 0xb5, 0xd0, 0xa3, 0x1e, 0x6e, 0x5a, 0x25, 0xb4, 0xc8,
	0xc8, 0x42, 0x05, 0x18, 0x5f, 0x25, 0x80, 0xdf, 0x19, 0x20, 0xaa, 0xde, 0xd6, 0xfe, 0x32, 0xc7,
	0x37, 0x0b, 0x3d, 0x3e, 0xa6, 0x5c, 0x00, 0xa3, 0x53, 0x7c, 0x92, 0xfc, 0x6b, 0x15, 0xdb, 0x57,
	0x49, 0x72, 0xf0, 0x17, 0x12, 0xf7, 0x2e, 0xe5, 0xcb, 0xa5, 0xa5, 0xa4, 0x44, 0x24, 0x25, 0xa6,
	0xec, 0x42, 0x94, 0xeb, 0x32, 0x4a, 0x30, 0x30, 0x4a, 0x44, 0x44, 0x8c, 0x29, 0xeb, 0xdc, 0x89,
	0xae, 0x19, 0x31, 0x5e, 0xa2, 0x2d, 0x79, 0x9d, 0x9d, 0x1b, 0xf2, 0x3e, 0x4e, 0x06, 0x7c, 0x6f,
	0xf5, 0x4d, 0x6e, 0xee, 0xa5, 0x6a, 0x8e, 0x5e, 0x9f, 0x2d, 0x3d, 0xeb, 0x7c, 0xe9, 0x59, 0xbf,
	0x96, 0x9e, 0xf5, 0x7d, 0xe5, 0x8d, 0xce, 0x57, 0xde, 0xe8, 0xc7, 0xca, 0x1b, 0x7d, 0x7a, 0x92,
	0x53, 0x71, 0x3c, 0x4f, 0x83, 0x29, 0x9c, 0x86, 0x3c, 0x27, 0x4f, 0xb5, 0x77, 0x7d, 0x0e, 0x17,
	0xad, 0x87, 0x43, 0x54, 0x25, 0xe1, 0xe9, 0x58, 0x3e, 0x1b, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x80, 0x96, 0xed, 0x1c, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.ParticipationBetPairExposureList) > 0 {
		for iNdEx := len(m.ParticipationBetPairExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationBetPairExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for iNdEx := len(m.HistoricalParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricalParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for iNdEx := len(m.ParticipationExposureByIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureByIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ParticipationExposureList) > 0 {
		for iNdEx := len(m.ParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OrderBookExposureList) > 0 {
		for iNdEx := len(m.OrderBookExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.OrderBookParticipationList) > 0 {
		for iNdEx := len(m.OrderBookParticipationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookParticipationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OrderBookList) > 0 {
		for iNdEx := len(m.OrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.OrderBookList) > 0 {
		for _, e := range m.OrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrderBookParticipationList) > 0 {
		for _, e := range m.OrderBookParticipationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrderBookExposureList) > 0 {
		for _, e := range m.OrderBookExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureList) > 0 {
		for _, e := range m.ParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for _, e := range m.ParticipationExposureByIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for _, e := range m.HistoricalParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationBetPairExposureList) > 0 {
		for _, e := range m.ParticipationBetPairExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Stats.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookList = append(m.OrderBookList, OrderBook{})
			if err := m.OrderBookList[len(m.OrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookParticipationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookParticipationList = append(m.OrderBookParticipationList, OrderBookParticipation{})
			if err := m.OrderBookParticipationList[len(m.OrderBookParticipationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookExposureList = append(m.OrderBookExposureList, OrderBookOddsExposure{})
			if err := m.OrderBookExposureList[len(m.OrderBookExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationExposureList = append(m.ParticipationExposureList, ParticipationExposure{})
			if err := m.ParticipationExposureList[len(m.ParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureByIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationExposureByIndexList = append(m.ParticipationExposureByIndexList, ParticipationExposure{})
			if err := m.ParticipationExposureByIndexList[len(m.ParticipationExposureByIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalParticipationExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HistoricalParticipationExposureList = append(m.HistoricalParticipationExposureList, ParticipationExposure{})
			if err := m.HistoricalParticipationExposureList[len(m.HistoricalParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationBetPairExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationBetPairExposureList = append(m.ParticipationBetPairExposureList, ParticipationBetPair{})
			if err := m.ParticipationBetPairExposureList[len(m.ParticipationBetPairExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
