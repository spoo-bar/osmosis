// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/module_route.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// PoolType is an enumeration of all supported pool types.
type PoolType int32

const (
	// Balancer is the standard xy=k curve. Its pool model is defined in x/gamm.
	Balancer PoolType = 0
	// Stableswap is the Solidly cfmm stable swap curve. Its pool model is defined
	// in x/gamm.
	Stableswap PoolType = 1
	// Concentrated is the pool model specific to concentrated liquidity. It is
	// defined in x/concentrated-liquidity.
	Concentrated PoolType = 2
	// CosmWasm is the pool model specific to CosmWasm. It is defined in
	// x/cosmwasmpool.
	CosmWasm PoolType = 3
)

var PoolType_name = map[int32]string{
	0: "Balancer",
	1: "Stableswap",
	2: "Concentrated",
	3: "CosmWasm",
}

var PoolType_value = map[string]int32{
	"Balancer":     0,
	"Stableswap":   1,
	"Concentrated": 2,
	"CosmWasm":     3,
}

func (x PoolType) String() string {
	return proto.EnumName(PoolType_name, int32(x))
}

func (PoolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96bfcc7b6d387cee, []int{0}
}

// ModuleRouter defines a route encapsulating pool type.
// It is used as the value of a mapping from pool id to the pool type,
// allowing the pool manager to know which module to route swaps to given the
// pool id.
type ModuleRoute struct {
	// pool_type specifies the type of the pool
	PoolType PoolType `protobuf:"varint,1,opt,name=pool_type,json=poolType,proto3,enum=osmosis.poolmanager.v1beta1.PoolType" json:"pool_type,omitempty"`
	PoolId   uint64   `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
}

func (m *ModuleRoute) Reset()         { *m = ModuleRoute{} }
func (m *ModuleRoute) String() string { return proto.CompactTextString(m) }
func (*ModuleRoute) ProtoMessage()    {}
func (*ModuleRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_96bfcc7b6d387cee, []int{0}
}
func (m *ModuleRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleRoute.Merge(m, src)
}
func (m *ModuleRoute) XXX_Size() int {
	return m.Size()
}
func (m *ModuleRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleRoute proto.InternalMessageInfo

func (m *ModuleRoute) GetPoolType() PoolType {
	if m != nil {
		return m.PoolType
	}
	return Balancer
}

func (m *ModuleRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.poolmanager.v1beta1.PoolType", PoolType_name, PoolType_value)
	proto.RegisterType((*ModuleRoute)(nil), "osmosis.poolmanager.v1beta1.ModuleRoute")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/module_route.proto", fileDescriptor_96bfcc7b6d387cee)
}

var fileDescriptor_96bfcc7b6d387cee = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x2f, 0x7d, 0x4b, 0xdf, 0x1a, 0x4b, 0x39, 0x0e, 0x91, 0x52, 0x21, 0x2d, 0x05, 0xa1,
	0x08, 0x26, 0x54, 0x07, 0xc5, 0xf1, 0x3a, 0x39, 0x28, 0x5a, 0x05, 0xc1, 0xa5, 0xe4, 0x7a, 0xe1,
	0x2c, 0x24, 0xf7, 0x0f, 0x97, 0xb4, 0xda, 0xd5, 0xc9, 0xd1, 0xef, 0xe0, 0x97, 0xe9, 0xd8, 0xd1,
	0xa9, 0x48, 0xfb, 0x0d, 0xfc, 0x04, 0x92, 0xf3, 0x0a, 0xba, 0xb8, 0x3d, 0x21, 0xbf, 0xe7, 0x07,
	0xff, 0x07, 0x53, 0x30, 0x0a, 0xcc, 0xd8, 0x30, 0x0d, 0x20, 0x15, 0x4f, 0x79, 0x22, 0x32, 0x36,
	0xed, 0x45, 0xc2, 0xf2, 0x1e, 0x53, 0x10, 0x4f, 0xa4, 0x18, 0x66, 0x30, 0xb1, 0x82, 0xea, 0x0c,
	0x2c, 0x04, 0x7b, 0x05, 0x4f, 0x7f, 0xf0, 0xb4, 0xe0, 0x9b, 0x3b, 0x09, 0x24, 0x90, 0x73, 0xcc,
	0xa5, 0xef, 0x4a, 0xe7, 0x19, 0xe1, 0xed, 0x8b, 0xdc, 0x34, 0x70, 0xa2, 0x20, 0xc4, 0x5b, 0xae,
	0x3c, 0xb4, 0x33, 0x2d, 0x1a, 0xa8, 0x8d, 0xba, 0xf5, 0xa3, 0x7d, 0xfa, 0x87, 0x96, 0x5e, 0x01,
	0xc8, 0xdb, 0x99, 0x16, 0x83, 0xaa, 0x2e, 0x52, 0xc0, 0xf0, 0xff, 0xdc, 0x31, 0x8e, 0x1b, 0xa5,
	0x36, 0xea, 0x96, 0xc3, 0xdd, 0xf9, 0xb2, 0x85, 0x3e, 0x97, 0xad, 0xfa, 0x8c, 0x2b, 0x79, 0xd6,
	0x29, 0x3e, 0x3b, 0x83, 0x8a, 0x4b, 0xe7, 0xf1, 0xc1, 0x25, 0xae, 0x6e, 0x34, 0x41, 0x0d, 0x57,
	0x43, 0x2e, 0x79, 0x3a, 0x12, 0x99, 0xef, 0x05, 0x75, 0x8c, 0x6f, 0x2c, 0x8f, 0xa4, 0x30, 0x8f,
	0x5c, 0xfb, 0x28, 0xf0, 0x71, 0xad, 0x0f, 0xe9, 0x48, 0xa4, 0x36, 0xe3, 0x56, 0xc4, 0x7e, 0xc9,
	0xf1, 0x7d, 0x30, 0xea, 0x8e, 0x1b, 0xe5, 0xff, 0x6b, 0x96, 0x5f, 0xde, 0x88, 0x17, 0x5e, 0xcf,
	0x57, 0x04, 0x2d, 0x56, 0x04, 0x7d, 0xac, 0x08, 0x7a, 0x5d, 0x13, 0x6f, 0xb1, 0x26, 0xde, 0xfb,
	0x9a, 0x78, 0xf7, 0x27, 0xc9, 0xd8, 0x3e, 0x4c, 0x22, 0x3a, 0x02, 0xc5, 0x8a, 0xab, 0x0e, 0x25,
	0x8f, 0xcc, 0xe6, 0xc1, 0xa6, 0xbd, 0x53, 0xf6, 0xf4, 0x6b, 0x6f, 0xb7, 0x84, 0x89, 0x2a, 0xf9,
	0x5c, 0xc7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x59, 0xce, 0x06, 0x93, 0x01, 0x00, 0x00,
}

func (m *ModuleRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintModuleRoute(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolType != 0 {
		i = encodeVarintModuleRoute(dAtA, i, uint64(m.PoolType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModuleRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovModuleRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ModuleRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolType != 0 {
		n += 1 + sovModuleRoute(uint64(m.PoolType))
	}
	if m.PoolId != 0 {
		n += 1 + sovModuleRoute(uint64(m.PoolId))
	}
	return n
}

func sovModuleRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModuleRoute(x uint64) (n int) {
	return sovModuleRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ModuleRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleRoute
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
			return fmt.Errorf("proto: ModuleRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolType", wireType)
			}
			m.PoolType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolType |= PoolType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModuleRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleRoute
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
func skipModuleRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModuleRoute
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
					return 0, ErrIntOverflowModuleRoute
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
					return 0, ErrIntOverflowModuleRoute
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
				return 0, ErrInvalidLengthModuleRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModuleRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModuleRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModuleRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModuleRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModuleRoute = fmt.Errorf("proto: unexpected end of group")
)
