// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stellarstation/planservice/api/mymessage.proto

package api // import "example.com/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MyMessage_MyEnum int32

const (
	MyMessage_CASE0 MyMessage_MyEnum = 0
	MyMessage_CASE1 MyMessage_MyEnum = 1
	MyMessage_CASE2 MyMessage_MyEnum = 2
	MyMessage_CASE3 MyMessage_MyEnum = 3
	MyMessage_CASE4 MyMessage_MyEnum = 4
	MyMessage_CASE5 MyMessage_MyEnum = 5
	MyMessage_CASE6 MyMessage_MyEnum = 6
)

var MyMessage_MyEnum_name = map[int32]string{
	0: "CASE0",
	1: "CASE1",
	2: "CASE2",
	3: "CASE3",
	4: "CASE4",
	5: "CASE5",
	6: "CASE6",
}
var MyMessage_MyEnum_value = map[string]int32{
	"CASE0": 0,
	"CASE1": 1,
	"CASE2": 2,
	"CASE3": 3,
	"CASE4": 4,
	"CASE5": 5,
	"CASE6": 6,
}

func (x MyMessage_MyEnum) String() string {
	return proto.EnumName(MyMessage_MyEnum_name, int32(x))
}
func (MyMessage_MyEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mymessage_dfcee293f34cb772, []int{0, 0}
}

type MyMessage struct {
	E                    MyMessage_MyEnum `protobuf:"varint,15,opt,name=e,proto3,enum=example.api.MyMessage_MyEnum" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_mymessage_dfcee293f34cb772, []int{0}
}
func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(dst, src)
}
func (m *MyMessage) XXX_Size() int {
	return m.Size()
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetE() MyMessage_MyEnum {
	if m != nil {
		return m.E
	}
	return MyMessage_CASE0
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "example.api.MyMessage")
	proto.RegisterEnum("example.api.MyMessage_MyEnum", MyMessage_MyEnum_name, MyMessage_MyEnum_value)
}
func (m *MyMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MyMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.E != 0 {
		dAtA[i] = 0x78
		i++
		i = encodeVarintMymessage(dAtA, i, uint64(m.E))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMymessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MyMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.E != 0 {
		n += 1 + sovMymessage(uint64(m.E))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMymessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMymessage(x uint64) (n int) {
	return sovMymessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MyMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MyMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MyMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			m.E = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.E |= (MyMessage_MyEnum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMymessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMymessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMymessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMymessage
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
					return 0, ErrIntOverflowMymessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMymessage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMymessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMymessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipMymessage(dAtA[start:])
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
	ErrInvalidLengthMymessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMymessage   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("stellarstation/planservice/api/mymessage.proto", fileDescriptor_mymessage_dfcee293f34cb772)
}

var fileDescriptor_mymessage_dfcee293f34cb772 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x2e, 0x49, 0xcd,
	0xc9, 0x49, 0x2c, 0x2a, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0xc8, 0x49, 0xcc, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0xad, 0xcc, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0x54, 0xea, 0x67, 0xe4, 0xe2, 0xf4, 0xad, 0xf4, 0x85,
	0x28, 0x10, 0xd2, 0xe6, 0x62, 0x4c, 0x95, 0xe0, 0x57, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd5, 0x43,
	0x52, 0xa6, 0x07, 0x57, 0xa2, 0xe7, 0x5b, 0xe9, 0x9a, 0x57, 0x9a, 0x1b, 0xc4, 0x98, 0xaa, 0x14,
	0xca, 0xc5, 0x06, 0xe1, 0x08, 0x71, 0x72, 0xb1, 0x3a, 0x3b, 0x06, 0xbb, 0x1a, 0x08, 0x30, 0xc0,
	0x98, 0x86, 0x02, 0x8c, 0x30, 0xa6, 0x91, 0x00, 0x13, 0x8c, 0x69, 0x2c, 0xc0, 0x0c, 0x63, 0x9a,
	0x08, 0xb0, 0xc0, 0x98, 0xa6, 0x02, 0xac, 0x30, 0xa6, 0x99, 0x00, 0x9b, 0x93, 0xe2, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x14,
	0x3f, 0xcc, 0x25, 0xc9, 0xf9, 0xb9, 0x20, 0x1f, 0x25, 0xb1, 0x81, 0x3d, 0x62, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0xaa, 0x1a, 0x3f, 0xfa, 0x00, 0x00, 0x00,
}
