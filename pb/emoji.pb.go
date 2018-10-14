// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emoji.proto

package pb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Emoji struct {
	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Roles         []string `protobuf:"bytes,3,rep,name=roles" json:"roles,omitempty"`
	Managed       bool     `protobuf:"varint,4,opt,name=managed,proto3" json:"managed,omitempty"`
	RequireColons bool     `protobuf:"varint,5,opt,name=require_colons,json=requireColons,proto3" json:"require_colons,omitempty"`
	Animated      bool     `protobuf:"varint,6,opt,name=animated,proto3" json:"animated,omitempty"`
}

func (m *Emoji) Reset()         { *m = Emoji{} }
func (m *Emoji) String() string { return proto.CompactTextString(m) }
func (*Emoji) ProtoMessage()    {}
func (*Emoji) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_ec1be6588912c796, []int{0}
}
func (m *Emoji) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Emoji) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Emoji.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Emoji) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Emoji.Merge(dst, src)
}
func (m *Emoji) XXX_Size() int {
	return m.Size()
}
func (m *Emoji) XXX_DiscardUnknown() {
	xxx_messageInfo_Emoji.DiscardUnknown(m)
}

var xxx_messageInfo_Emoji proto.InternalMessageInfo

func (m *Emoji) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Emoji) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Emoji) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Emoji) GetManaged() bool {
	if m != nil {
		return m.Managed
	}
	return false
}

func (m *Emoji) GetRequireColons() bool {
	if m != nil {
		return m.RequireColons
	}
	return false
}

func (m *Emoji) GetAnimated() bool {
	if m != nil {
		return m.Animated
	}
	return false
}

type SetEmojiRequest struct {
	Emoji *Emoji `protobuf:"bytes,1,opt,name=emoji" json:"emoji,omitempty"`
}

func (m *SetEmojiRequest) Reset()         { *m = SetEmojiRequest{} }
func (m *SetEmojiRequest) String() string { return proto.CompactTextString(m) }
func (*SetEmojiRequest) ProtoMessage()    {}
func (*SetEmojiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_ec1be6588912c796, []int{1}
}
func (m *SetEmojiRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetEmojiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetEmojiRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetEmojiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetEmojiRequest.Merge(dst, src)
}
func (m *SetEmojiRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetEmojiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetEmojiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetEmojiRequest proto.InternalMessageInfo

func (m *SetEmojiRequest) GetEmoji() *Emoji {
	if m != nil {
		return m.Emoji
	}
	return nil
}

type SetEmojiResponse struct {
}

func (m *SetEmojiResponse) Reset()         { *m = SetEmojiResponse{} }
func (m *SetEmojiResponse) String() string { return proto.CompactTextString(m) }
func (*SetEmojiResponse) ProtoMessage()    {}
func (*SetEmojiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_ec1be6588912c796, []int{2}
}
func (m *SetEmojiResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetEmojiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetEmojiResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetEmojiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetEmojiResponse.Merge(dst, src)
}
func (m *SetEmojiResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetEmojiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetEmojiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetEmojiResponse proto.InternalMessageInfo

type GetEmojiRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetEmojiRequest) Reset()         { *m = GetEmojiRequest{} }
func (m *GetEmojiRequest) String() string { return proto.CompactTextString(m) }
func (*GetEmojiRequest) ProtoMessage()    {}
func (*GetEmojiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_ec1be6588912c796, []int{3}
}
func (m *GetEmojiRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetEmojiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetEmojiRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetEmojiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEmojiRequest.Merge(dst, src)
}
func (m *GetEmojiRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetEmojiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEmojiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEmojiRequest proto.InternalMessageInfo

func (m *GetEmojiRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEmojiResponse struct {
	Emoji *Emoji `protobuf:"bytes,1,opt,name=emoji" json:"emoji,omitempty"`
}

func (m *GetEmojiResponse) Reset()         { *m = GetEmojiResponse{} }
func (m *GetEmojiResponse) String() string { return proto.CompactTextString(m) }
func (*GetEmojiResponse) ProtoMessage()    {}
func (*GetEmojiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_ec1be6588912c796, []int{4}
}
func (m *GetEmojiResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetEmojiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetEmojiResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetEmojiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEmojiResponse.Merge(dst, src)
}
func (m *GetEmojiResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetEmojiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEmojiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEmojiResponse proto.InternalMessageInfo

func (m *GetEmojiResponse) GetEmoji() *Emoji {
	if m != nil {
		return m.Emoji
	}
	return nil
}

func init() {
	proto.RegisterType((*Emoji)(nil), "state.Emoji")
	proto.RegisterType((*SetEmojiRequest)(nil), "state.SetEmojiRequest")
	proto.RegisterType((*SetEmojiResponse)(nil), "state.SetEmojiResponse")
	proto.RegisterType((*GetEmojiRequest)(nil), "state.GetEmojiRequest")
	proto.RegisterType((*GetEmojiResponse)(nil), "state.GetEmojiResponse")
}
func (m *Emoji) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Emoji) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmoji(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEmoji(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Managed {
		dAtA[i] = 0x20
		i++
		if m.Managed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RequireColons {
		dAtA[i] = 0x28
		i++
		if m.RequireColons {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Animated {
		dAtA[i] = 0x30
		i++
		if m.Animated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *SetEmojiRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetEmojiRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Emoji != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmoji(dAtA, i, uint64(m.Emoji.Size()))
		n1, err := m.Emoji.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SetEmojiResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetEmojiResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetEmojiRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEmojiRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmoji(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *GetEmojiResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEmojiResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Emoji != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmoji(dAtA, i, uint64(m.Emoji.Size()))
		n2, err := m.Emoji.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintEmoji(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Emoji) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEmoji(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEmoji(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovEmoji(uint64(l))
		}
	}
	if m.Managed {
		n += 2
	}
	if m.RequireColons {
		n += 2
	}
	if m.Animated {
		n += 2
	}
	return n
}

func (m *SetEmojiRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Emoji != nil {
		l = m.Emoji.Size()
		n += 1 + l + sovEmoji(uint64(l))
	}
	return n
}

func (m *SetEmojiResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetEmojiRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEmoji(uint64(l))
	}
	return n
}

func (m *GetEmojiResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Emoji != nil {
		l = m.Emoji.Size()
		n += 1 + l + sovEmoji(uint64(l))
	}
	return n
}

func sovEmoji(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEmoji(x uint64) (n int) {
	return sovEmoji(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Emoji) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmoji
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
			return fmt.Errorf("proto: Emoji: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Emoji: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Managed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Managed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequireColons", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RequireColons = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Animated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Animated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEmoji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmoji
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
func (m *SetEmojiRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmoji
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
			return fmt.Errorf("proto: SetEmojiRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetEmojiRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emoji", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Emoji == nil {
				m.Emoji = &Emoji{}
			}
			if err := m.Emoji.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmoji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmoji
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
func (m *SetEmojiResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmoji
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
			return fmt.Errorf("proto: SetEmojiResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetEmojiResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEmoji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmoji
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
func (m *GetEmojiRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmoji
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
			return fmt.Errorf("proto: GetEmojiRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEmojiRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmoji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmoji
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
func (m *GetEmojiResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmoji
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
			return fmt.Errorf("proto: GetEmojiResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEmojiResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emoji", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmoji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEmoji
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Emoji == nil {
				m.Emoji = &Emoji{}
			}
			if err := m.Emoji.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmoji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmoji
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
func skipEmoji(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEmoji
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
					return 0, ErrIntOverflowEmoji
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
					return 0, ErrIntOverflowEmoji
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
				return 0, ErrInvalidLengthEmoji
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEmoji
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
				next, err := skipEmoji(dAtA[start:])
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
	ErrInvalidLengthEmoji = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEmoji   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("emoji.proto", fileDescriptor_emoji_ec1be6588912c796) }

var fileDescriptor_emoji_ec1be6588912c796 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xb4, 0x29, 0xed, 0x2d, 0xb4, 0x95, 0xc5, 0x60, 0x21, 0x64, 0x05, 0x4b, 0x48,
	0x99, 0x32, 0x80, 0xe0, 0x01, 0x40, 0xa8, 0x7b, 0xd8, 0x58, 0x90, 0x4b, 0xae, 0x50, 0x50, 0x63,
	0xa7, 0xb1, 0xfb, 0x1e, 0x3c, 0x01, 0xcf, 0xc3, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0xb8, 0xe1,
	0x47, 0xca, 0xd4, 0xcd, 0xe7, 0xf3, 0xb9, 0xf6, 0x39, 0x17, 0xa6, 0x58, 0xda, 0x97, 0x22, 0xad,
	0x6a, 0xeb, 0x2d, 0x8f, 0x9c, 0xd7, 0x1e, 0xd5, 0x1b, 0x83, 0xe8, 0xee, 0x1b, 0xf3, 0x19, 0x84,
	0x45, 0x2e, 0x58, 0xcc, 0x92, 0x49, 0x16, 0x16, 0x39, 0xe7, 0x30, 0x34, 0xba, 0x44, 0x11, 0x12,
	0xa1, 0x33, 0x3f, 0x86, 0xa8, 0xb6, 0x6b, 0x74, 0x62, 0x10, 0x0f, 0x92, 0x49, 0xd6, 0x09, 0x2e,
	0xe0, 0xa0, 0xd4, 0x46, 0x3f, 0x63, 0x2e, 0x86, 0x31, 0x4b, 0xc6, 0xd9, 0xaf, 0xe4, 0xe7, 0x30,
	0xab, 0x71, 0xb3, 0x2d, 0x6a, 0x7c, 0x7c, 0xb2, 0x6b, 0x6b, 0x9c, 0x88, 0xc8, 0x70, 0xf4, 0x43,
	0x6f, 0x09, 0xf2, 0x13, 0x18, 0x6b, 0x53, 0x94, 0xda, 0x63, 0x2e, 0x46, 0x64, 0xf8, 0xd3, 0xea,
	0x0a, 0xe6, 0xf7, 0xe8, 0x29, 0x62, 0x86, 0x9b, 0x2d, 0x3a, 0xcf, 0x15, 0x44, 0xd4, 0x84, 0xc2,
	0x4e, 0x2f, 0x0e, 0x53, 0xaa, 0x92, 0x76, 0x9e, 0xee, 0x4a, 0x71, 0x58, 0xfc, 0x8f, 0xb9, 0xca,
	0x1a, 0x87, 0xea, 0x0c, 0xe6, 0xcb, 0xde, 0x53, 0xbd, 0xd2, 0xea, 0x1a, 0x16, 0xcb, 0xde, 0xd8,
	0x3e, 0xdf, 0xdd, 0x9c, 0xbe, 0x37, 0x92, 0xed, 0x1a, 0xc9, 0x3e, 0x1b, 0xc9, 0x5e, 0x5b, 0x19,
	0xec, 0x5a, 0x19, 0x7c, 0xb4, 0x32, 0x78, 0x08, 0xab, 0xd5, 0x6a, 0x44, 0x2b, 0xbf, 0xfc, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0x18, 0x5b, 0xcf, 0x81, 0x01, 0x00, 0x00,
}
