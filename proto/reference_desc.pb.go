// Code generated by protoc-gen-gogo.
// source: reference_desc.proto
// DO NOT EDIT!

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReferenceDescriptor_ReferenceType int32

const (
	ReferenceDescriptor_Tombstone ReferenceDescriptor_ReferenceType = 0
	ReferenceDescriptor_Future    ReferenceDescriptor_ReferenceType = 1
	ReferenceDescriptor_Concrete  ReferenceDescriptor_ReferenceType = 2
	ReferenceDescriptor_Stream    ReferenceDescriptor_ReferenceType = 3
	ReferenceDescriptor_Value     ReferenceDescriptor_ReferenceType = 4
	ReferenceDescriptor_Error     ReferenceDescriptor_ReferenceType = 5
)

var ReferenceDescriptor_ReferenceType_name = map[int32]string{
	0: "Tombstone",
	1: "Future",
	2: "Concrete",
	3: "Stream",
	4: "Value",
	5: "Error",
}
var ReferenceDescriptor_ReferenceType_value = map[string]int32{
	"Tombstone": 0,
	"Future":    1,
	"Concrete":  2,
	"Stream":    3,
	"Value":     4,
	"Error":     5,
}

func (x ReferenceDescriptor_ReferenceType) String() string {
	return proto.EnumName(ReferenceDescriptor_ReferenceType_name, int32(x))
}
func (ReferenceDescriptor_ReferenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorReferenceDesc, []int{0, 0}
}

type ReferenceDescriptor_ReferenceScope int32

const (
	// TODO(malte): really encode like this?
	ReferenceDescriptor_Public  ReferenceDescriptor_ReferenceScope = 0
	ReferenceDescriptor_Private ReferenceDescriptor_ReferenceScope = 1
)

var ReferenceDescriptor_ReferenceScope_name = map[int32]string{
	0: "Public",
	1: "Private",
}
var ReferenceDescriptor_ReferenceScope_value = map[string]int32{
	"Public":  0,
	"Private": 1,
}

func (x ReferenceDescriptor_ReferenceScope) String() string {
	return proto.EnumName(ReferenceDescriptor_ReferenceScope_name, int32(x))
}
func (ReferenceDescriptor_ReferenceScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorReferenceDesc, []int{0, 1}
}

type ReferenceDescriptor struct {
	Id               []byte                             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type             ReferenceDescriptor_ReferenceType  `protobuf:"varint,2,opt,name=type,proto3,enum=firmament.ReferenceDescriptor_ReferenceType" json:"type,omitempty"`
	Scope            ReferenceDescriptor_ReferenceScope `protobuf:"varint,3,opt,name=scope,proto3,enum=firmament.ReferenceDescriptor_ReferenceScope" json:"scope,omitempty"`
	NonDeterministic bool                               `protobuf:"varint,4,opt,name=non_deterministic,json=nonDeterministic,proto3" json:"non_deterministic,omitempty"`
	Size_            uint64                             `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Location         string                             `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	InlineData       []byte                             `protobuf:"bytes,7,opt,name=inline_data,json=inlineData,proto3" json:"inline_data,omitempty"`
	ProducingTask    uint64                             `protobuf:"varint,8,opt,name=producing_task,json=producingTask,proto3" json:"producing_task,omitempty"`
	TimeToCompute    uint64                             `protobuf:"varint,9,opt,name=time_to_compute,json=timeToCompute,proto3" json:"time_to_compute,omitempty"`
	Version          uint64                             `protobuf:"varint,10,opt,name=version,proto3" json:"version,omitempty"`
	// Modifications for storage engine
	IsModified bool `protobuf:"varint,11,opt,name=is_modified,json=isModified,proto3" json:"is_modified,omitempty"`
}

func (m *ReferenceDescriptor) Reset()                    { *m = ReferenceDescriptor{} }
func (m *ReferenceDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ReferenceDescriptor) ProtoMessage()               {}
func (*ReferenceDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorReferenceDesc, []int{0} }

func init() {
	proto.RegisterType((*ReferenceDescriptor)(nil), "firmament.ReferenceDescriptor")
	proto.RegisterEnum("firmament.ReferenceDescriptor_ReferenceType", ReferenceDescriptor_ReferenceType_name, ReferenceDescriptor_ReferenceType_value)
	proto.RegisterEnum("firmament.ReferenceDescriptor_ReferenceScope", ReferenceDescriptor_ReferenceScope_name, ReferenceDescriptor_ReferenceScope_value)
}
func (m *ReferenceDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReferenceDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	if m.Type != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.Type))
	}
	if m.Scope != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.Scope))
	}
	if m.NonDeterministic {
		data[i] = 0x20
		i++
		if m.NonDeterministic {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Size_ != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.Size_))
	}
	if len(m.Location) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(len(m.Location)))
		i += copy(data[i:], m.Location)
	}
	if len(m.InlineData) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(len(m.InlineData)))
		i += copy(data[i:], m.InlineData)
	}
	if m.ProducingTask != 0 {
		data[i] = 0x40
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.ProducingTask))
	}
	if m.TimeToCompute != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.TimeToCompute))
	}
	if m.Version != 0 {
		data[i] = 0x50
		i++
		i = encodeVarintReferenceDesc(data, i, uint64(m.Version))
	}
	if m.IsModified {
		data[i] = 0x58
		i++
		if m.IsModified {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64ReferenceDesc(data []byte, offset int, v uint64) int {
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
func encodeFixed32ReferenceDesc(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintReferenceDesc(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ReferenceDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovReferenceDesc(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovReferenceDesc(uint64(m.Type))
	}
	if m.Scope != 0 {
		n += 1 + sovReferenceDesc(uint64(m.Scope))
	}
	if m.NonDeterministic {
		n += 2
	}
	if m.Size_ != 0 {
		n += 1 + sovReferenceDesc(uint64(m.Size_))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovReferenceDesc(uint64(l))
	}
	l = len(m.InlineData)
	if l > 0 {
		n += 1 + l + sovReferenceDesc(uint64(l))
	}
	if m.ProducingTask != 0 {
		n += 1 + sovReferenceDesc(uint64(m.ProducingTask))
	}
	if m.TimeToCompute != 0 {
		n += 1 + sovReferenceDesc(uint64(m.TimeToCompute))
	}
	if m.Version != 0 {
		n += 1 + sovReferenceDesc(uint64(m.Version))
	}
	if m.IsModified {
		n += 2
	}
	return n
}

func sovReferenceDesc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReferenceDesc(x uint64) (n int) {
	return sovReferenceDesc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReferenceDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReferenceDesc
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
			return fmt.Errorf("proto: ReferenceDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferenceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReferenceDesc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], data[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (ReferenceDescriptor_ReferenceType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			m.Scope = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Scope |= (ReferenceDescriptor_ReferenceScope(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonDeterministic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NonDeterministic = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Size_ |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
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
				return ErrInvalidLengthReferenceDesc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReferenceDesc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InlineData = append(m.InlineData[:0], data[iNdEx:postIndex]...)
			if m.InlineData == nil {
				m.InlineData = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducingTask", wireType)
			}
			m.ProducingTask = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ProducingTask |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeToCompute", wireType)
			}
			m.TimeToCompute = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TimeToCompute |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsModified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferenceDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsModified = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipReferenceDesc(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReferenceDesc
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
func skipReferenceDesc(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReferenceDesc
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
					return 0, ErrIntOverflowReferenceDesc
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
					return 0, ErrIntOverflowReferenceDesc
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
				return 0, ErrInvalidLengthReferenceDesc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReferenceDesc
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
				next, err := skipReferenceDesc(data[start:])
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
	ErrInvalidLengthReferenceDesc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReferenceDesc   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorReferenceDesc = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0xdd, 0xf4, 0x4f, 0xa6, 0xdb, 0x12, 0x0c, 0x07, 0x6b, 0x0f, 0x0b, 0xaa, 0x04,
	0x02, 0xc1, 0x76, 0x25, 0x78, 0x01, 0x44, 0x0b, 0x37, 0xa4, 0x55, 0xb6, 0xe2, 0xc0, 0x25, 0x4a,
	0x1c, 0xb7, 0x8c, 0x68, 0x3c, 0x95, 0xed, 0xac, 0x04, 0x4f, 0xc2, 0xbb, 0xf0, 0x02, 0x1c, 0x79,
	0x04, 0x04, 0x2f, 0xc2, 0xc4, 0x61, 0x0b, 0x95, 0x38, 0x70, 0xb0, 0x34, 0xf3, 0x1b, 0x7f, 0x5f,
	0x3d, 0x5f, 0x03, 0x77, 0xad, 0x5e, 0x6b, 0xab, 0x8d, 0xd2, 0x79, 0xa5, 0x9d, 0x9a, 0xef, 0x2c,
	0x79, 0x12, 0xc9, 0x1a, 0x6d, 0x5d, 0xd4, 0xda, 0xf8, 0xd3, 0xf3, 0x0d, 0xfa, 0xf7, 0x4d, 0x39,
	0x57, 0x54, 0x5f, 0x6c, 0x68, 0x43, 0x17, 0xe1, 0x46, 0xd9, 0xac, 0x43, 0x17, 0x9a, 0x50, 0x75,
	0xca, 0xd9, 0x97, 0x18, 0xee, 0x64, 0x37, 0x96, 0x4b, 0x76, 0xb4, 0xb8, 0xf3, 0x64, 0xc5, 0x14,
	0x7a, 0x58, 0xc9, 0xe8, 0x7e, 0xf4, 0xe8, 0x24, 0xe3, 0x4a, 0xbc, 0x80, 0xd8, 0x7f, 0xdc, 0x69,
	0xd9, 0x63, 0x32, 0x7d, 0xf6, 0x74, 0xbe, 0xff, 0xc1, 0xf9, 0x3f, 0xd4, 0x7f, 0xd8, 0x8a, 0x35,
	0x59, 0x50, 0x8a, 0x05, 0xf4, 0x9d, 0x22, 0xb6, 0x38, 0x0e, 0x16, 0xe7, 0xff, 0x6b, 0x71, 0xd5,
	0x8a, 0xb2, 0x4e, 0x2b, 0x9e, 0xc0, 0x6d, 0x43, 0x86, 0x57, 0xf7, 0xda, 0xd6, 0x68, 0xd0, 0x79,
	0x54, 0x32, 0x66, 0xc3, 0x51, 0x96, 0xf2, 0x60, 0xf9, 0x37, 0x17, 0x02, 0x62, 0x87, 0x9f, 0xb4,
	0xec, 0xf3, 0x3c, 0xce, 0x42, 0x2d, 0x4e, 0x61, 0xb4, 0x25, 0x55, 0x78, 0x24, 0x23, 0x07, 0xcc,
	0x93, 0x6c, 0xdf, 0x8b, 0x7b, 0x30, 0x46, 0xb3, 0x45, 0xc3, 0xd1, 0x16, 0xbe, 0x90, 0xc3, 0xb0,
	0x3c, 0x74, 0x68, 0xc9, 0x44, 0x3c, 0x80, 0x29, 0xa7, 0x56, 0x35, 0x0a, 0xcd, 0x26, 0xf7, 0x85,
	0xfb, 0x20, 0x47, 0xc1, 0x7a, 0xb2, 0xa7, 0x2b, 0x86, 0xe2, 0x21, 0xdc, 0xf2, 0x58, 0xeb, 0xdc,
	0x53, 0xce, 0xff, 0xc2, 0xae, 0xf1, 0x5a, 0x26, 0xdd, 0xbd, 0x16, 0xaf, 0x68, 0xd1, 0x41, 0x21,
	0x61, 0x78, 0xad, 0xad, 0x6b, 0x9f, 0x02, 0x61, 0x7e, 0xd3, 0x86, 0x97, 0xb8, 0xbc, 0xa6, 0x0a,
	0xd7, 0xa8, 0x2b, 0x39, 0x0e, 0x0b, 0x02, 0xba, 0x37, 0xbf, 0xc9, 0xec, 0x1d, 0x4c, 0x0e, 0x32,
	0x16, 0x13, 0x48, 0x56, 0x54, 0x97, 0xce, 0x93, 0xd1, 0xe9, 0x91, 0x00, 0x18, 0xbc, 0x6e, 0x7c,
	0x63, 0x75, 0x1a, 0x89, 0x13, 0x18, 0x2d, 0xc8, 0x28, 0xcb, 0xd9, 0xa4, 0xbd, 0x76, 0x72, 0xe5,
	0xad, 0x2e, 0xea, 0xf4, 0x58, 0x24, 0xd0, 0x7f, 0x5b, 0x6c, 0x1b, 0x9d, 0xc6, 0x6d, 0xf9, 0xca,
	0x5a, 0xb2, 0x69, 0x7f, 0xf6, 0x18, 0xa6, 0x87, 0xe1, 0xb7, 0x9a, 0xcb, 0xa6, 0xdc, 0xa2, 0x62,
	0xe7, 0x31, 0x0c, 0x2f, 0x2d, 0x5e, 0x17, 0x6c, 0x16, 0xbd, 0x4c, 0xbf, 0xfe, 0x38, 0x8b, 0xbe,
	0xf1, 0xf9, 0xce, 0xe7, 0xf3, 0xcf, 0xb3, 0xa3, 0x72, 0x10, 0x3e, 0xab, 0xe7, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0xbf, 0x19, 0x0a, 0xa8, 0x02, 0x00, 0x00,
}