// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aizel/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// period is the amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// epoch_identifier for inflation
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// epochs_per_period is the number of epochs after which inflation is recalculated
	EpochsPerPeriod int64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
	// skipped_epochs is the number of epochs that have passed while inflation is disabled
	SkippedEpochs uint64 `protobuf:"varint,5,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_22d7da4a03e630cd, []int{0}
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

func (m *GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *GenesisState) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *GenesisState) GetEpochsPerPeriod() int64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func (m *GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

// Params holds parameters for the inflation module.
type Params struct {
	// mint_denom specifies the type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// exponential_calculation takes in the variables to calculate exponential inflation
	ExponentialCalculation ExponentialCalculation `protobuf:"bytes,2,opt,name=exponential_calculation,json=exponentialCalculation,proto3" json:"exponential_calculation"`
	// inflation_distribution of the minted denom
	InflationDistribution InflationDistribution `protobuf:"bytes,3,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// enable_inflation is the parameter that enables inflation and halts increasing the skipped_epochs
	EnableInflation bool `protobuf:"varint,4,opt,name=enable_inflation,json=enableInflation,proto3" json:"enable_inflation,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_22d7da4a03e630cd, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetExponentialCalculation() ExponentialCalculation {
	if m != nil {
		return m.ExponentialCalculation
	}
	return ExponentialCalculation{}
}

func (m *Params) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func (m *Params) GetEnableInflation() bool {
	if m != nil {
		return m.EnableInflation
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "aizel.inflation.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "aizel.inflation.v1.Params")
}

func init() { proto.RegisterFile("aizel/inflation/v1/genesis.proto", fileDescriptor_22d7da4a03e630cd) }

var fileDescriptor_22d7da4a03e630cd = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x75, 0x54, 0xd4, 0x03, 0xc6, 0x2c, 0x28, 0x55, 0x25, 0x42, 0x54, 0x09, 0xa9,
	0xeb, 0x21, 0x66, 0x45, 0xe2, 0xc6, 0x81, 0xb1, 0x09, 0xed, 0x82, 0xaa, 0x70, 0xe3, 0x12, 0xb9,
	0xcd, 0x5b, 0x67, 0x35, 0xb1, 0xad, 0xd8, 0x2d, 0x83, 0x4f, 0xc1, 0xc7, 0xe0, 0x06, 0x1f, 0x63,
	0xc7, 0x1e, 0x39, 0x21, 0xd4, 0x1e, 0xf8, 0x1a, 0x28, 0x2f, 0x21, 0x2d, 0x5a, 0x2e, 0x91, 0xfd,
	0xf3, 0xff, 0xfd, 0xff, 0xcf, 0x2f, 0xa6, 0xbe, 0x90, 0x5f, 0x20, 0xe1, 0x52, 0x5d, 0x26, 0xc2,
	0x49, 0xad, 0xf8, 0xf2, 0x84, 0xcf, 0x40, 0x81, 0x95, 0x36, 0x30, 0x99, 0x76, 0x9a, 0x31, 0x54,
	0x04, 0x95, 0x22, 0x58, 0x9e, 0xf4, 0x8e, 0x44, 0x2a, 0x95, 0xe6, 0xf8, 0x2d, 0x64, 0xbd, 0x47,
	0x33, 0x3d, 0xd3, 0xb8, 0xe4, 0xf9, 0xaa, 0xa4, 0xfd, 0x1a, 0xfb, 0xad, 0x13, 0x6a, 0xfa, 0x1b,
	0x42, 0xef, 0xbd, 0x2b, 0x22, 0x3f, 0x38, 0xe1, 0x80, 0xbd, 0xa6, 0x2d, 0x23, 0x32, 0x91, 0xda,
	0x2e, 0xf1, 0xc9, 0xe0, 0x60, 0xd4, 0x0b, 0x6e, 0xb7, 0x10, 0x8c, 0x51, 0x71, 0xda, 0xbe, 0xf9,
	0xf5, 0xac, 0xf1, 0xed, 0xcf, 0x8f, 0x21, 0x09, 0xcb, 0x22, 0xd6, 0xa1, 0x2d, 0x03, 0x99, 0xd4,
	0x71, 0x77, 0xcf, 0x27, 0x83, 0xfd, 0xb0, 0xdc, 0xb1, 0x63, 0xfa, 0x10, 0x8c, 0x9e, 0x5e, 0x45,
	0x32, 0x06, 0xe5, 0xe4, 0xa5, 0x84, 0xac, 0xdb, 0xf4, 0xc9, 0xa0, 0x1d, 0x1e, 0x22, 0xbf, 0xa8,
	0x30, 0x1b, 0xd2, 0x23, 0x44, 0x36, 0x32, 0x90, 0x45, 0xa5, 0xdb, 0xbe, 0x4f, 0x06, 0xcd, 0x52,
	0x6b, 0xc7, 0x90, 0x8d, 0x0b, 0xdb, 0xe7, 0xf4, 0x81, 0x9d, 0x4b, 0x63, 0x20, 0x8e, 0x8a, 0xa3,
	0xee, 0x1d, 0x8c, 0xbd, 0x5f, 0xd2, 0x73, 0x84, 0xfd, 0xef, 0x7b, 0xb4, 0x55, 0xf4, 0xcc, 0x9e,
	0x52, 0x9a, 0x4a, 0xe5, 0xa2, 0x18, 0x94, 0x4e, 0xf1, 0x8e, 0xed, 0xb0, 0x9d, 0x93, 0xb3, 0x1c,
	0x30, 0x45, 0x9f, 0xc0, 0xb5, 0xd1, 0x2a, 0xef, 0x46, 0x24, 0xd1, 0x54, 0x24, 0xd3, 0x45, 0x71,
	0x6f, 0xbc, 0xd0, 0xc1, 0x68, 0x58, 0x37, 0x8f, 0xf3, 0x6d, 0xc9, 0xdb, 0x6d, 0xc5, 0xee, 0x7c,
	0x3a, 0x50, 0x2b, 0x61, 0x73, 0xda, 0xa9, 0x9c, 0xa2, 0x58, 0x5a, 0x97, 0xc9, 0xc9, 0x02, 0xe3,
	0x9a, 0x18, 0x77, 0x5c, 0x17, 0x77, 0xf1, 0x6f, 0x73, 0xb6, 0x53, 0xb0, 0x9b, 0xf6, 0x58, 0xd6,
	0x29, 0xf0, 0x27, 0x28, 0x31, 0x49, 0x20, 0xaa, 0xce, 0x71, 0xb0, 0x77, 0xc3, 0xc3, 0x82, 0x57,
	0xc6, 0xa7, 0xe3, 0x9b, 0xb5, 0x47, 0x56, 0x6b, 0x8f, 0xfc, 0x5e, 0x7b, 0xe4, 0xeb, 0xc6, 0x6b,
	0xac, 0x36, 0x5e, 0xe3, 0xe7, 0xc6, 0x6b, 0x7c, 0x7c, 0x35, 0x93, 0xee, 0x6a, 0x31, 0x09, 0xa6,
	0x3a, 0xe5, 0x6f, 0xf2, 0xde, 0xde, 0x83, 0xfb, 0xa4, 0xb3, 0x39, 0x87, 0x65, 0xaa, 0x2d, 0x5f,
	0x8e, 0x5e, 0xf0, 0xeb, 0xff, 0xdf, 0x9c, 0xfb, 0x6c, 0xc0, 0x4e, 0x5a, 0xf8, 0xe0, 0x5e, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xde, 0x53, 0xd0, 0x1c, 0xf5, 0x02, 0x00, 0x00,
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
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableInflation {
		i--
		if m.EnableInflation {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ExponentialCalculation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
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
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ExponentialCalculation.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EnableInflation {
		n += 2
	}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsPerPeriod", wireType)
			}
			m.EpochsPerPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsPerPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedEpochs", wireType)
			}
			m.SkippedEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentialCalculation", wireType)
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
			if err := m.ExponentialCalculation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
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
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableInflation", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableInflation = bool(v != 0)
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
