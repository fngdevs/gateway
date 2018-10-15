// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: member.proto

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

type Member struct {
	UserId   string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GuildId  string      `protobuf:"bytes,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	JoinedAt *BytesValue `protobuf:"bytes,3,opt,name=JoinedAt" json:"JoinedAt,omitempty"`
	Nick     *BytesValue `protobuf:"bytes,4,opt,name=nick" json:"nick,omitempty"`
	Deaf     *BoolValue  `protobuf:"bytes,5,opt,name=deaf" json:"deaf,omitempty"`
	Mute     *BoolValue  `protobuf:"bytes,6,opt,name=mute" json:"mute,omitempty"`
	Roles    []string    `protobuf:"bytes,7,rep,name=roles" json:"roles,omitempty"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{0}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(dst, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Member) GetGuildId() string {
	if m != nil {
		return m.GuildId
	}
	return ""
}

func (m *Member) GetJoinedAt() *BytesValue {
	if m != nil {
		return m.JoinedAt
	}
	return nil
}

func (m *Member) GetNick() *BytesValue {
	if m != nil {
		return m.Nick
	}
	return nil
}

func (m *Member) GetDeaf() *BoolValue {
	if m != nil {
		return m.Deaf
	}
	return nil
}

func (m *Member) GetMute() *BoolValue {
	if m != nil {
		return m.Mute
	}
	return nil
}

func (m *Member) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type GetMemberRequest struct {
	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GuildId string `protobuf:"bytes,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
}

func (m *GetMemberRequest) Reset()         { *m = GetMemberRequest{} }
func (m *GetMemberRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemberRequest) ProtoMessage()    {}
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{1}
}
func (m *GetMemberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMemberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberRequest.Merge(dst, src)
}
func (m *GetMemberRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberRequest proto.InternalMessageInfo

func (m *GetMemberRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetMemberRequest) GetGuildId() string {
	if m != nil {
		return m.GuildId
	}
	return ""
}

type GetMemberResponse struct {
	Member *Member `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
}

func (m *GetMemberResponse) Reset()         { *m = GetMemberResponse{} }
func (m *GetMemberResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemberResponse) ProtoMessage()    {}
func (*GetMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{2}
}
func (m *GetMemberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMemberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberResponse.Merge(dst, src)
}
func (m *GetMemberResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberResponse proto.InternalMessageInfo

func (m *GetMemberResponse) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type SetMemberRequest struct {
	Member *Member `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
}

func (m *SetMemberRequest) Reset()         { *m = SetMemberRequest{} }
func (m *SetMemberRequest) String() string { return proto.CompactTextString(m) }
func (*SetMemberRequest) ProtoMessage()    {}
func (*SetMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{3}
}
func (m *SetMemberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMemberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMemberRequest.Merge(dst, src)
}
func (m *SetMemberRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetMemberRequest proto.InternalMessageInfo

func (m *SetMemberRequest) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type SetMemberResponse struct {
}

func (m *SetMemberResponse) Reset()         { *m = SetMemberResponse{} }
func (m *SetMemberResponse) String() string { return proto.CompactTextString(m) }
func (*SetMemberResponse) ProtoMessage()    {}
func (*SetMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{4}
}
func (m *SetMemberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMemberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMemberResponse.Merge(dst, src)
}
func (m *SetMemberResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetMemberResponse proto.InternalMessageInfo

type UpdateMemberRequest struct {
	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Member *Member `protobuf:"bytes,2,opt,name=member" json:"member,omitempty"`
}

func (m *UpdateMemberRequest) Reset()         { *m = UpdateMemberRequest{} }
func (m *UpdateMemberRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMemberRequest) ProtoMessage()    {}
func (*UpdateMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{5}
}
func (m *UpdateMemberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMemberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UpdateMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMemberRequest.Merge(dst, src)
}
func (m *UpdateMemberRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMemberRequest proto.InternalMessageInfo

func (m *UpdateMemberRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateMemberRequest) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type UpdateMemberResponse struct {
}

func (m *UpdateMemberResponse) Reset()         { *m = UpdateMemberResponse{} }
func (m *UpdateMemberResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateMemberResponse) ProtoMessage()    {}
func (*UpdateMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_a712ac81594a29c3, []int{6}
}
func (m *UpdateMemberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMemberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UpdateMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMemberResponse.Merge(dst, src)
}
func (m *UpdateMemberResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMemberResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Member)(nil), "state.Member")
	proto.RegisterType((*GetMemberRequest)(nil), "state.GetMemberRequest")
	proto.RegisterType((*GetMemberResponse)(nil), "state.GetMemberResponse")
	proto.RegisterType((*SetMemberRequest)(nil), "state.SetMemberRequest")
	proto.RegisterType((*SetMemberResponse)(nil), "state.SetMemberResponse")
	proto.RegisterType((*UpdateMemberRequest)(nil), "state.UpdateMemberRequest")
	proto.RegisterType((*UpdateMemberResponse)(nil), "state.UpdateMemberResponse")
}
func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMember(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.GuildId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMember(dAtA, i, uint64(len(m.GuildId)))
		i += copy(dAtA[i:], m.GuildId)
	}
	if m.JoinedAt != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.JoinedAt.Size()))
		n1, err := m.JoinedAt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Nick != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Nick.Size()))
		n2, err := m.Nick.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Deaf != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Deaf.Size()))
		n3, err := m.Deaf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Mute != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Mute.Size()))
		n4, err := m.Mute.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x3a
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
	return i, nil
}

func (m *GetMemberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMemberRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMember(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.GuildId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMember(dAtA, i, uint64(len(m.GuildId)))
		i += copy(dAtA[i:], m.GuildId)
	}
	return i, nil
}

func (m *GetMemberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMemberResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Member != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Member.Size()))
		n5, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *SetMemberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMemberRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Member != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Member.Size()))
		n6, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *SetMemberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMemberResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UpdateMemberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMemberRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMember(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if m.Member != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMember(dAtA, i, uint64(m.Member.Size()))
		n7, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *UpdateMemberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMemberResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintMember(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.GuildId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	if m.JoinedAt != nil {
		l = m.JoinedAt.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Nick != nil {
		l = m.Nick.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Deaf != nil {
		l = m.Deaf.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Mute != nil {
		l = m.Mute.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovMember(uint64(l))
		}
	}
	return n
}

func (m *GetMemberRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.GuildId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	return n
}

func (m *GetMemberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	return n
}

func (m *SetMemberRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	return n
}

func (m *SetMemberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UpdateMemberRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovMember(uint64(l))
	}
	return n
}

func (m *UpdateMemberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMember(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMember(x uint64) (n int) {
	return sovMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuildId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuildId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.JoinedAt == nil {
				m.JoinedAt = &BytesValue{}
			}
			if err := m.JoinedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nick", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nick == nil {
				m.Nick = &BytesValue{}
			}
			if err := m.Nick.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deaf == nil {
				m.Deaf = &BoolValue{}
			}
			if err := m.Deaf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mute == nil {
				m.Mute = &BoolValue{}
			}
			if err := m.Mute.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *GetMemberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: GetMemberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMemberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuildId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuildId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *GetMemberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: GetMemberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMemberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *SetMemberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: SetMemberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMemberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *SetMemberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: SetMemberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMemberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *UpdateMemberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: UpdateMemberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMemberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func (m *UpdateMemberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: UpdateMemberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMemberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMember
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
func skipMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
				return 0, ErrInvalidLengthMember
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMember
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
				next, err := skipMember(dAtA[start:])
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
	ErrInvalidLengthMember = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMember   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("member.proto", fileDescriptor_member_a712ac81594a29c3) }

var fileDescriptor_member_a712ac81594a29c3 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0xcd, 0x4d,
	0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x49, 0x95,
	0xe2, 0x2b, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x2a, 0x86, 0x08, 0x2b, 0x7d, 0x61, 0xe4, 0x62,
	0xf3, 0x05, 0xab, 0x13, 0x12, 0xe7, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0x24, 0xb9, 0x38, 0xd2, 0x4b,
	0x33, 0x73, 0x52, 0x40, 0x32, 0x4c, 0x60, 0x19, 0x76, 0x30, 0xdf, 0x33, 0x45, 0x48, 0x97, 0x8b,
	0xc3, 0x2b, 0x3f, 0x33, 0x2f, 0x35, 0xc5, 0xb1, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x50, 0x0f, 0x6c, 0x91, 0x9e, 0x53, 0x65, 0x49, 0x6a, 0x71, 0x58, 0x62, 0x4e, 0x69, 0x6a, 0x10,
	0x5c, 0x89, 0x90, 0x2a, 0x17, 0x4b, 0x5e, 0x66, 0x72, 0xb6, 0x04, 0x0b, 0x2e, 0xa5, 0x60, 0x69,
	0x21, 0x15, 0x2e, 0x96, 0x94, 0xd4, 0xc4, 0x34, 0x09, 0x56, 0xb0, 0x32, 0x01, 0x98, 0xb2, 0xfc,
	0xfc, 0x1c, 0xa8, 0x2a, 0x90, 0x2c, 0x48, 0x55, 0x6e, 0x69, 0x49, 0xaa, 0x04, 0x1b, 0x2e, 0x55,
	0x20, 0x59, 0x21, 0x11, 0x2e, 0xd6, 0xa2, 0xfc, 0x9c, 0xd4, 0x62, 0x09, 0x76, 0x05, 0x66, 0x0d,
	0xce, 0x20, 0x08, 0x47, 0xc9, 0x8d, 0x4b, 0xc0, 0x3d, 0xb5, 0x04, 0xe2, 0xf1, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x72, 0xfc, 0xaf, 0x64, 0xc5, 0x25, 0x88, 0x64, 0x4e, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x2a, 0x17, 0x1b, 0x24, 0xe8, 0xc1, 0xe6, 0x70, 0x1b, 0xf1, 0x42, 0x9d,
	0x06, 0x55, 0x06, 0x95, 0x54, 0xb2, 0xe4, 0x12, 0x08, 0x46, 0x77, 0x03, 0x91, 0x5a, 0x85, 0xb9,
	0x04, 0x83, 0xd1, 0xad, 0x55, 0x0a, 0xe5, 0x12, 0x0e, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x25, 0xd2,
	0x5b, 0x08, 0xbb, 0x98, 0xf0, 0xd9, 0x25, 0xc6, 0x25, 0x82, 0x6a, 0x2c, 0xc4, 0x3a, 0x27, 0x99,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03,
	0x27, 0x2f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0xfd, 0xfb, 0xf2, 0x85, 0x02, 0x00,
	0x00,
}
