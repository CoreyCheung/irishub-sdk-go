// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: record/record.proto

package record

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

// Content defines the detailed information for a record.
type Content struct {
	Digest     string `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	DigestAlgo string `protobuf:"bytes,2,opt,name=digest_algo,json=digestAlgo,proto3" json:"digest_algo,omitempty" yaml:"digest_algo"`
	URI        string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Meta       string `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_197cabccbeb2a7b7, []int{0}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return m.Size()
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

type Record struct {
	TxHash   string    `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
	Contents []Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents"`
	Creator  string    `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_197cabccbeb2a7b7, []int{1}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Content)(nil), "irismod.record.Content")
	proto.RegisterType((*Record)(nil), "irismod.record.Record")
}

func init() { proto.RegisterFile("record/record.proto", fileDescriptor_197cabccbeb2a7b7) }

var fileDescriptor_197cabccbeb2a7b7 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x8d, 0x9b, 0x28, 0x01, 0x57, 0xea, 0x60, 0x50, 0x09, 0x0c, 0x49, 0x95, 0xa9, 0x12, 0x6a,
	0x82, 0x60, 0x40, 0x74, 0x23, 0x2c, 0xb0, 0x30, 0x58, 0x62, 0x61, 0xa9, 0xd2, 0x24, 0x72, 0x22,
	0x92, 0x1a, 0xd9, 0x8e, 0x54, 0xbe, 0x82, 0x7e, 0x02, 0x9f, 0xd3, 0xb1, 0x23, 0x53, 0x04, 0xe9,
	0xc2, 0xdc, 0x2f, 0x40, 0xb1, 0x03, 0x82, 0xe9, 0xde, 0xbb, 0x77, 0xd6, 0x7b, 0x77, 0x86, 0x07,
	0x2c, 0x8d, 0x29, 0x4b, 0x02, 0x55, 0xfc, 0x67, 0x46, 0x05, 0x45, 0x83, 0x9c, 0xe5, 0xbc, 0xa4,
	0x89, 0xaf, 0xba, 0x27, 0x87, 0x84, 0x12, 0x2a, 0xa5, 0xa0, 0x45, 0x6a, 0xca, 0x7b, 0x05, 0xd0,
	0xba, 0xa1, 0x0b, 0x91, 0x2e, 0x04, 0x1a, 0x42, 0x33, 0xc9, 0x49, 0xca, 0x85, 0x0d, 0x46, 0x60,
	0xbc, 0x8f, 0x3b, 0x86, 0x2e, 0x61, 0x5f, 0xa1, 0x59, 0x54, 0x10, 0x6a, 0xf7, 0x5a, 0x31, 0x1c,
	0xee, 0x6a, 0x17, 0xbd, 0x44, 0x65, 0x31, 0xf5, 0xfe, 0x88, 0x1e, 0x86, 0x8a, 0x5d, 0x17, 0x84,
	0xa2, 0x63, 0xa8, 0x57, 0x2c, 0xb7, 0x75, 0xf9, 0xc0, 0x6a, 0x6a, 0x57, 0x7f, 0xc0, 0x77, 0xb8,
	0xed, 0x21, 0x04, 0x8d, 0x32, 0x15, 0x91, 0x6d, 0x48, 0x27, 0x89, 0xa7, 0xc6, 0xd7, 0x9b, 0x0b,
	0xbc, 0x15, 0x80, 0x26, 0x96, 0x91, 0xd1, 0x29, 0xb4, 0xc4, 0x72, 0x96, 0x45, 0x3c, 0x53, 0x89,
	0x42, 0xb4, 0xab, 0xdd, 0x81, 0x32, 0xed, 0x04, 0x0f, 0x9b, 0x62, 0x79, 0x1b, 0xf1, 0x0c, 0x5d,
	0xc1, 0xbd, 0x58, 0x2d, 0xc2, 0xed, 0xde, 0x48, 0x1f, 0xf7, 0xcf, 0x8f, 0xfc, 0xff, 0x27, 0xf0,
	0xbb, 0x45, 0x43, 0x63, 0x5d, 0xbb, 0x1a, 0xfe, 0x1d, 0x47, 0x36, 0xb4, 0x62, 0x96, 0x46, 0x82,
	0x32, 0x95, 0x15, 0xff, 0x50, 0x15, 0x29, 0xbc, 0x5f, 0x7f, 0x3a, 0xda, 0xba, 0x71, 0xc0, 0xa6,
	0x71, 0xc0, 0x47, 0xe3, 0x80, 0xd5, 0xd6, 0xd1, 0x36, 0x5b, 0x47, 0x7b, 0xdf, 0x3a, 0xda, 0xe3,
	0x19, 0xc9, 0x45, 0x56, 0xcd, 0xfd, 0x98, 0x96, 0x41, 0x6b, 0xb8, 0x48, 0x85, 0xac, 0x59, 0x35,
	0x9f, 0xf0, 0xe4, 0x69, 0x42, 0x68, 0x50, 0xd2, 0xa4, 0x2a, 0x52, 0xde, 0x7d, 0xd0, 0xdc, 0x94,
	0xb7, 0xbf, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x4e, 0x10, 0xde, 0xb8, 0x01, 0x00, 0x00,
}

func (this *Content) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Content)
	if !ok {
		that2, ok := that.(Content)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Digest != that1.Digest {
		return false
	}
	if this.DigestAlgo != that1.DigestAlgo {
		return false
	}
	if this.URI != that1.URI {
		return false
	}
	if this.Meta != that1.Meta {
		return false
	}
	return true
}
func (this *Record) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Record)
	if !ok {
		that2, ok := that.(Record)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TxHash != that1.TxHash {
		return false
	}
	if len(this.Contents) != len(that1.Contents) {
		return false
	}
	for i := range this.Contents {
		if !this.Contents[i].Equal(&that1.Contents[i]) {
			return false
		}
	}
	if this.Creator != that1.Creator {
		return false
	}
	return true
}
func (m *Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DigestAlgo) > 0 {
		i -= len(m.DigestAlgo)
		copy(dAtA[i:], m.DigestAlgo)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.DigestAlgo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contents) > 0 {
		for iNdEx := len(m.Contents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.DigestAlgo)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovRecord(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DigestAlgo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, Content{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
