// Code generated by protoc-gen-gogo.
// source: gogovanity.proto
// DO NOT EDIT!

/*
	Package vanity is a generated protocol buffer package.

	It is generated from these files:
		gogovanity.proto

	It has these top-level messages:
		B
*/
package vanity

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type B struct {
	String_ *string `protobuf:"bytes,1,opt,name=String,json=string" json:"String,omitempty"`
	Int64   int64   `protobuf:"varint,2,opt,name=Int64,json=int64" json:"Int64"`
	Int32   *int32  `protobuf:"varint,3,opt,name=Int32,json=int32,def=1234" json:"Int32,omitempty"`
}

func (m *B) Reset()                    { *m = B{} }
func (m *B) String() string            { return proto.CompactTextString(m) }
func (*B) ProtoMessage()               {}
func (*B) Descriptor() ([]byte, []int) { return fileDescriptorGogovanity, []int{0} }

const Default_B_Int32 int32 = 1234

func (m *B) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func (m *B) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *B) GetInt32() int32 {
	if m != nil && m.Int32 != nil {
		return *m.Int32
	}
	return Default_B_Int32
}

func init() {
	proto.RegisterType((*B)(nil), "vanity.B")
}
func (m *B) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *B) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.String_ != nil {
		data[i] = 0xa
		i++
		i = encodeVarintGogovanity(data, i, uint64(len(*m.String_)))
		i += copy(data[i:], *m.String_)
	}
	data[i] = 0x10
	i++
	i = encodeVarintGogovanity(data, i, uint64(m.Int64))
	if m.Int32 != nil {
		data[i] = 0x18
		i++
		i = encodeVarintGogovanity(data, i, uint64(*m.Int32))
	}
	return i, nil
}

func encodeFixed64Gogovanity(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Gogovanity(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGogovanity(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *B) Size() (n int) {
	var l int
	_ = l
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovGogovanity(uint64(l))
	}
	n += 1 + sovGogovanity(uint64(m.Int64))
	if m.Int32 != nil {
		n += 1 + sovGogovanity(uint64(*m.Int32))
	}
	return n
}

func sovGogovanity(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGogovanity(x uint64) (n int) {
	return sovGogovanity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *B) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogovanity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGogovanity
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Int64 |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int32 = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGogovanity(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogovanity
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
func skipGogovanity(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGogovanity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGogovanity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGogovanity
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGogovanity(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGogovanity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGogovanity   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("gogovanity.proto", fileDescriptorGogovanity) }

var fileDescriptorGogovanity = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcf, 0x4f, 0xcf,
	0x2f, 0x4b, 0xcc, 0xcb, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0,
	0xa4, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0x8a, 0xf4,
	0xc1, 0xd2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xb4, 0x29, 0x45, 0x72, 0x31,
	0x3a, 0x09, 0xc9, 0x70, 0xb1, 0x05, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x3a, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x18, 0xc4, 0x56, 0x0c, 0x16, 0x13, 0x92, 0xe2, 0x62,
	0xf5, 0xcc, 0x2b, 0x31, 0x33, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x06, 0x4b, 0x32, 0x04, 0xb1,
	0x66, 0x82, 0x84, 0xa0, 0x72, 0xc6, 0x46, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x56, 0x2c, 0x86,
	0x46, 0xc6, 0x26, 0x60, 0x39, 0x63, 0x23, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x06, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0xdf, 0x35, 0x69, 0xb6, 0x00, 0x00, 0x00,
}
