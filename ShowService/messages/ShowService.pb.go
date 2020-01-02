// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ShowService.proto

package ShowService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KillShowsMovieMessage struct {
	MovieID              int32    `protobuf:"varint,1,opt,name=movieID,proto3" json:"movieID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KillShowsMovieMessage) Reset()         { *m = KillShowsMovieMessage{} }
func (m *KillShowsMovieMessage) String() string { return proto.CompactTextString(m) }
func (*KillShowsMovieMessage) ProtoMessage()    {}
func (*KillShowsMovieMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{0}
}

func (m *KillShowsMovieMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KillShowsMovieMessage.Unmarshal(m, b)
}
func (m *KillShowsMovieMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KillShowsMovieMessage.Marshal(b, m, deterministic)
}
func (m *KillShowsMovieMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KillShowsMovieMessage.Merge(m, src)
}
func (m *KillShowsMovieMessage) XXX_Size() int {
	return xxx_messageInfo_KillShowsMovieMessage.Size(m)
}
func (m *KillShowsMovieMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KillShowsMovieMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KillShowsMovieMessage proto.InternalMessageInfo

func (m *KillShowsMovieMessage) GetMovieID() int32 {
	if m != nil {
		return m.MovieID
	}
	return 0
}

type KillShowsMovieResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KillShowsMovieResponse) Reset()         { *m = KillShowsMovieResponse{} }
func (m *KillShowsMovieResponse) String() string { return proto.CompactTextString(m) }
func (*KillShowsMovieResponse) ProtoMessage()    {}
func (*KillShowsMovieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{1}
}

func (m *KillShowsMovieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KillShowsMovieResponse.Unmarshal(m, b)
}
func (m *KillShowsMovieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KillShowsMovieResponse.Marshal(b, m, deterministic)
}
func (m *KillShowsMovieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KillShowsMovieResponse.Merge(m, src)
}
func (m *KillShowsMovieResponse) XXX_Size() int {
	return xxx_messageInfo_KillShowsMovieResponse.Size(m)
}
func (m *KillShowsMovieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KillShowsMovieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KillShowsMovieResponse proto.InternalMessageInfo

func (m *KillShowsMovieResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type KillShowsHallMessage struct {
	HallID               int32    `protobuf:"varint,1,opt,name=hallID,proto3" json:"hallID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KillShowsHallMessage) Reset()         { *m = KillShowsHallMessage{} }
func (m *KillShowsHallMessage) String() string { return proto.CompactTextString(m) }
func (*KillShowsHallMessage) ProtoMessage()    {}
func (*KillShowsHallMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{2}
}

func (m *KillShowsHallMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KillShowsHallMessage.Unmarshal(m, b)
}
func (m *KillShowsHallMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KillShowsHallMessage.Marshal(b, m, deterministic)
}
func (m *KillShowsHallMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KillShowsHallMessage.Merge(m, src)
}
func (m *KillShowsHallMessage) XXX_Size() int {
	return xxx_messageInfo_KillShowsHallMessage.Size(m)
}
func (m *KillShowsHallMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KillShowsHallMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KillShowsHallMessage proto.InternalMessageInfo

func (m *KillShowsHallMessage) GetHallID() int32 {
	if m != nil {
		return m.HallID
	}
	return 0
}

type KillShowsHallResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KillShowsHallResponse) Reset()         { *m = KillShowsHallResponse{} }
func (m *KillShowsHallResponse) String() string { return proto.CompactTextString(m) }
func (*KillShowsHallResponse) ProtoMessage()    {}
func (*KillShowsHallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{3}
}

func (m *KillShowsHallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KillShowsHallResponse.Unmarshal(m, b)
}
func (m *KillShowsHallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KillShowsHallResponse.Marshal(b, m, deterministic)
}
func (m *KillShowsHallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KillShowsHallResponse.Merge(m, src)
}
func (m *KillShowsHallResponse) XXX_Size() int {
	return xxx_messageInfo_KillShowsHallResponse.Size(m)
}
func (m *KillShowsHallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KillShowsHallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KillShowsHallResponse proto.InternalMessageInfo

func (m *KillShowsHallResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CreateShowMessage struct {
	MovieID              int32    `protobuf:"varint,1,opt,name=movieID,proto3" json:"movieID,omitempty"`
	HallID               int32    `protobuf:"varint,2,opt,name=hallID,proto3" json:"hallID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShowMessage) Reset()         { *m = CreateShowMessage{} }
func (m *CreateShowMessage) String() string { return proto.CompactTextString(m) }
func (*CreateShowMessage) ProtoMessage()    {}
func (*CreateShowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{4}
}

func (m *CreateShowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShowMessage.Unmarshal(m, b)
}
func (m *CreateShowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShowMessage.Marshal(b, m, deterministic)
}
func (m *CreateShowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShowMessage.Merge(m, src)
}
func (m *CreateShowMessage) XXX_Size() int {
	return xxx_messageInfo_CreateShowMessage.Size(m)
}
func (m *CreateShowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShowMessage proto.InternalMessageInfo

func (m *CreateShowMessage) GetMovieID() int32 {
	if m != nil {
		return m.MovieID
	}
	return 0
}

func (m *CreateShowMessage) GetHallID() int32 {
	if m != nil {
		return m.HallID
	}
	return 0
}

type CreateShowResponse struct {
	ShowID               int32    `protobuf:"varint,1,opt,name=showID,proto3" json:"showID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShowResponse) Reset()         { *m = CreateShowResponse{} }
func (m *CreateShowResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShowResponse) ProtoMessage()    {}
func (*CreateShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{5}
}

func (m *CreateShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShowResponse.Unmarshal(m, b)
}
func (m *CreateShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShowResponse.Marshal(b, m, deterministic)
}
func (m *CreateShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShowResponse.Merge(m, src)
}
func (m *CreateShowResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShowResponse.Size(m)
}
func (m *CreateShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShowResponse proto.InternalMessageInfo

func (m *CreateShowResponse) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

type DeleteShowMessage struct {
	ShowID               int32    `protobuf:"varint,1,opt,name=showID,proto3" json:"showID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowMessage) Reset()         { *m = DeleteShowMessage{} }
func (m *DeleteShowMessage) String() string { return proto.CompactTextString(m) }
func (*DeleteShowMessage) ProtoMessage()    {}
func (*DeleteShowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{6}
}

func (m *DeleteShowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowMessage.Unmarshal(m, b)
}
func (m *DeleteShowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowMessage.Marshal(b, m, deterministic)
}
func (m *DeleteShowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowMessage.Merge(m, src)
}
func (m *DeleteShowMessage) XXX_Size() int {
	return xxx_messageInfo_DeleteShowMessage.Size(m)
}
func (m *DeleteShowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowMessage proto.InternalMessageInfo

func (m *DeleteShowMessage) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

type DeleteShowResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShowResponse) Reset()         { *m = DeleteShowResponse{} }
func (m *DeleteShowResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteShowResponse) ProtoMessage()    {}
func (*DeleteShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{7}
}

func (m *DeleteShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShowResponse.Unmarshal(m, b)
}
func (m *DeleteShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShowResponse.Marshal(b, m, deterministic)
}
func (m *DeleteShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShowResponse.Merge(m, src)
}
func (m *DeleteShowResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteShowResponse.Size(m)
}
func (m *DeleteShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShowResponse proto.InternalMessageInfo

func (m *DeleteShowResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type BlockSeatMessage struct {
	BookingID            int32    `protobuf:"varint,3,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	ShowID               int32    `protobuf:"varint,1,opt,name=showID,proto3" json:"showID,omitempty"`
	SeatID               []int32  `protobuf:"varint,2,rep,packed,name=seatID,proto3" json:"seatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockSeatMessage) Reset()         { *m = BlockSeatMessage{} }
func (m *BlockSeatMessage) String() string { return proto.CompactTextString(m) }
func (*BlockSeatMessage) ProtoMessage()    {}
func (*BlockSeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{8}
}

func (m *BlockSeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockSeatMessage.Unmarshal(m, b)
}
func (m *BlockSeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockSeatMessage.Marshal(b, m, deterministic)
}
func (m *BlockSeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockSeatMessage.Merge(m, src)
}
func (m *BlockSeatMessage) XXX_Size() int {
	return xxx_messageInfo_BlockSeatMessage.Size(m)
}
func (m *BlockSeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockSeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockSeatMessage proto.InternalMessageInfo

func (m *BlockSeatMessage) GetBookingID() int32 {
	if m != nil {
		return m.BookingID
	}
	return 0
}

func (m *BlockSeatMessage) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

func (m *BlockSeatMessage) GetSeatID() []int32 {
	if m != nil {
		return m.SeatID
	}
	return nil
}

type BlockSeatResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	BookingID            int32    `protobuf:"varint,2,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockSeatResponse) Reset()         { *m = BlockSeatResponse{} }
func (m *BlockSeatResponse) String() string { return proto.CompactTextString(m) }
func (*BlockSeatResponse) ProtoMessage()    {}
func (*BlockSeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{9}
}

func (m *BlockSeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockSeatResponse.Unmarshal(m, b)
}
func (m *BlockSeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockSeatResponse.Marshal(b, m, deterministic)
}
func (m *BlockSeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockSeatResponse.Merge(m, src)
}
func (m *BlockSeatResponse) XXX_Size() int {
	return xxx_messageInfo_BlockSeatResponse.Size(m)
}
func (m *BlockSeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockSeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockSeatResponse proto.InternalMessageInfo

func (m *BlockSeatResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BlockSeatResponse) GetBookingID() int32 {
	if m != nil {
		return m.BookingID
	}
	return 0
}

type LockSeatMessage struct {
	ShowID               int32    `protobuf:"varint,1,opt,name=showID,proto3" json:"showID,omitempty"`
	BookingID            int32    `protobuf:"varint,2,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockSeatMessage) Reset()         { *m = LockSeatMessage{} }
func (m *LockSeatMessage) String() string { return proto.CompactTextString(m) }
func (*LockSeatMessage) ProtoMessage()    {}
func (*LockSeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{10}
}

func (m *LockSeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockSeatMessage.Unmarshal(m, b)
}
func (m *LockSeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockSeatMessage.Marshal(b, m, deterministic)
}
func (m *LockSeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockSeatMessage.Merge(m, src)
}
func (m *LockSeatMessage) XXX_Size() int {
	return xxx_messageInfo_LockSeatMessage.Size(m)
}
func (m *LockSeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LockSeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LockSeatMessage proto.InternalMessageInfo

func (m *LockSeatMessage) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

func (m *LockSeatMessage) GetBookingID() int32 {
	if m != nil {
		return m.BookingID
	}
	return 0
}

type LockSeatResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	BookingID            int32    `protobuf:"varint,2,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockSeatResponse) Reset()         { *m = LockSeatResponse{} }
func (m *LockSeatResponse) String() string { return proto.CompactTextString(m) }
func (*LockSeatResponse) ProtoMessage()    {}
func (*LockSeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{11}
}

func (m *LockSeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockSeatResponse.Unmarshal(m, b)
}
func (m *LockSeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockSeatResponse.Marshal(b, m, deterministic)
}
func (m *LockSeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockSeatResponse.Merge(m, src)
}
func (m *LockSeatResponse) XXX_Size() int {
	return xxx_messageInfo_LockSeatResponse.Size(m)
}
func (m *LockSeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockSeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockSeatResponse proto.InternalMessageInfo

func (m *LockSeatResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LockSeatResponse) GetBookingID() int32 {
	if m != nil {
		return m.BookingID
	}
	return 0
}

type FreeSeatMessage struct {
	ShowID               int32    `protobuf:"varint,1,opt,name=showID,proto3" json:"showID,omitempty"`
	BookingID            int32    `protobuf:"varint,2,opt,name=bookingID,proto3" json:"bookingID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreeSeatMessage) Reset()         { *m = FreeSeatMessage{} }
func (m *FreeSeatMessage) String() string { return proto.CompactTextString(m) }
func (*FreeSeatMessage) ProtoMessage()    {}
func (*FreeSeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{12}
}

func (m *FreeSeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreeSeatMessage.Unmarshal(m, b)
}
func (m *FreeSeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreeSeatMessage.Marshal(b, m, deterministic)
}
func (m *FreeSeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreeSeatMessage.Merge(m, src)
}
func (m *FreeSeatMessage) XXX_Size() int {
	return xxx_messageInfo_FreeSeatMessage.Size(m)
}
func (m *FreeSeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FreeSeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FreeSeatMessage proto.InternalMessageInfo

func (m *FreeSeatMessage) GetShowID() int32 {
	if m != nil {
		return m.ShowID
	}
	return 0
}

func (m *FreeSeatMessage) GetBookingID() int32 {
	if m != nil {
		return m.BookingID
	}
	return 0
}

type FreeSeatResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreeSeatResponse) Reset()         { *m = FreeSeatResponse{} }
func (m *FreeSeatResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSeatResponse) ProtoMessage()    {}
func (*FreeSeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{13}
}

func (m *FreeSeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreeSeatResponse.Unmarshal(m, b)
}
func (m *FreeSeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreeSeatResponse.Marshal(b, m, deterministic)
}
func (m *FreeSeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreeSeatResponse.Merge(m, src)
}
func (m *FreeSeatResponse) XXX_Size() int {
	return xxx_messageInfo_FreeSeatResponse.Size(m)
}
func (m *FreeSeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FreeSeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FreeSeatResponse proto.InternalMessageInfo

func (m *FreeSeatResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type VerifySeatMessage struct {
	SeatID               []int32  `protobuf:"varint,1,rep,packed,name=seatID,proto3" json:"seatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifySeatMessage) Reset()         { *m = VerifySeatMessage{} }
func (m *VerifySeatMessage) String() string { return proto.CompactTextString(m) }
func (*VerifySeatMessage) ProtoMessage()    {}
func (*VerifySeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{14}
}

func (m *VerifySeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifySeatMessage.Unmarshal(m, b)
}
func (m *VerifySeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifySeatMessage.Marshal(b, m, deterministic)
}
func (m *VerifySeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifySeatMessage.Merge(m, src)
}
func (m *VerifySeatMessage) XXX_Size() int {
	return xxx_messageInfo_VerifySeatMessage.Size(m)
}
func (m *VerifySeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifySeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_VerifySeatMessage proto.InternalMessageInfo

func (m *VerifySeatMessage) GetSeatID() []int32 {
	if m != nil {
		return m.SeatID
	}
	return nil
}

type VerifySeatResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifySeatResponse) Reset()         { *m = VerifySeatResponse{} }
func (m *VerifySeatResponse) String() string { return proto.CompactTextString(m) }
func (*VerifySeatResponse) ProtoMessage()    {}
func (*VerifySeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_932e6a77e85c512c, []int{15}
}

func (m *VerifySeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifySeatResponse.Unmarshal(m, b)
}
func (m *VerifySeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifySeatResponse.Marshal(b, m, deterministic)
}
func (m *VerifySeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifySeatResponse.Merge(m, src)
}
func (m *VerifySeatResponse) XXX_Size() int {
	return xxx_messageInfo_VerifySeatResponse.Size(m)
}
func (m *VerifySeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifySeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifySeatResponse proto.InternalMessageInfo

func (m *VerifySeatResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*KillShowsMovieMessage)(nil), "KillShowsMovieMessage")
	proto.RegisterType((*KillShowsMovieResponse)(nil), "KillShowsMovieResponse")
	proto.RegisterType((*KillShowsHallMessage)(nil), "KillShowsHallMessage")
	proto.RegisterType((*KillShowsHallResponse)(nil), "KillShowsHallResponse")
	proto.RegisterType((*CreateShowMessage)(nil), "CreateShowMessage")
	proto.RegisterType((*CreateShowResponse)(nil), "CreateShowResponse")
	proto.RegisterType((*DeleteShowMessage)(nil), "DeleteShowMessage")
	proto.RegisterType((*DeleteShowResponse)(nil), "DeleteShowResponse")
	proto.RegisterType((*BlockSeatMessage)(nil), "BlockSeatMessage")
	proto.RegisterType((*BlockSeatResponse)(nil), "BlockSeatResponse")
	proto.RegisterType((*LockSeatMessage)(nil), "LockSeatMessage")
	proto.RegisterType((*LockSeatResponse)(nil), "LockSeatResponse")
	proto.RegisterType((*FreeSeatMessage)(nil), "FreeSeatMessage")
	proto.RegisterType((*FreeSeatResponse)(nil), "FreeSeatResponse")
	proto.RegisterType((*VerifySeatMessage)(nil), "VerifySeatMessage")
	proto.RegisterType((*VerifySeatResponse)(nil), "VerifySeatResponse")
}

func init() { proto.RegisterFile("ShowService.proto", fileDescriptor_932e6a77e85c512c) }

var fileDescriptor_932e6a77e85c512c = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xe5, 0x43, 0xa5, 0x65, 0xaa, 0xb6, 0xf6, 0xb6, 0x50, 0x64, 0xf5, 0x50, 0xed, 0xa9, 0x52,
	0xd1, 0x4a, 0x21, 0x8a, 0x72, 0x8d, 0x02, 0xf9, 0xc4, 0x5c, 0x40, 0xca, 0x39, 0x60, 0x6d, 0xc0,
	0xc2, 0x61, 0x91, 0x97, 0x10, 0xe5, 0xcf, 0xe6, 0xb7, 0x64, 0xbc, 0xc4, 0x5e, 0xaf, 0x6d, 0x02,
	0x52, 0x72, 0x9c, 0xf5, 0xbc, 0xf7, 0xe6, 0xe3, 0x79, 0xc0, 0x1e, 0xcd, 0xc4, 0xe3, 0x88, 0x87,
	0x6b, 0xdf, 0xe3, 0x6c, 0x19, 0x8a, 0x95, 0xa0, 0x07, 0xd0, 0xe8, 0xfb, 0x41, 0x10, 0x7d, 0x90,
	0x03, 0xb1, 0xf6, 0xf9, 0x80, 0x4b, 0x39, 0x9e, 0x72, 0xd2, 0x82, 0xcf, 0xf7, 0x51, 0x7c, 0xd5,
	0x6b, 0x95, 0xff, 0x96, 0xff, 0x7d, 0x1a, 0xc6, 0x21, 0xed, 0x40, 0xd3, 0x84, 0x0c, 0xb9, 0x5c,
	0x8a, 0x85, 0x54, 0x18, 0xf9, 0xe0, 0x79, 0xc8, 0xa0, 0x30, 0x5f, 0x86, 0x71, 0x48, 0x19, 0xfc,
	0x4a, 0x30, 0x97, 0xe3, 0x20, 0x88, 0x55, 0x9a, 0x50, 0x9b, 0x61, 0x98, 0x88, 0xbc, 0x46, 0x46,
	0x59, 0x51, 0xfe, 0x1e, 0x12, 0x67, 0x60, 0x77, 0x43, 0x3e, 0x5e, 0xf1, 0x08, 0xb4, 0xb3, 0x8b,
	0x94, 0x72, 0xc5, 0x50, 0x6e, 0x03, 0xd1, 0x34, 0x89, 0x2c, 0x66, 0x4b, 0x8c, 0x75, 0x9d, 0x9b,
	0x88, 0xfe, 0x07, 0xbb, 0xc7, 0x03, 0x6e, 0x8a, 0x6e, 0x4b, 0x66, 0x40, 0x74, 0xf2, 0x1e, 0x1d,
	0xdd, 0x82, 0x75, 0x1a, 0x08, 0x6f, 0x3e, 0xc2, 0x72, 0x62, 0xee, 0x3f, 0x50, 0x9f, 0x08, 0x31,
	0xf7, 0x17, 0x53, 0xa4, 0xaf, 0x2a, 0x7a, 0xfd, 0xb0, 0x4d, 0x59, 0xbd, 0x23, 0x89, 0x6a, 0xb6,
	0xaa, 0xde, 0x55, 0x44, 0xfb, 0x60, 0x27, 0x0a, 0xbb, 0x0b, 0x32, 0xc5, 0x2b, 0x19, 0x71, 0x7a,
	0x01, 0x3f, 0xdc, 0x4c, 0xb5, 0xdb, 0xea, 0x79, 0x9b, 0xe8, 0x1a, 0x2c, 0xf7, 0x03, 0x8b, 0x3a,
	0x0f, 0x39, 0x7f, 0x7f, 0x51, 0x6d, 0xb0, 0x62, 0xa2, 0x3d, 0x56, 0x87, 0xbe, 0xb8, 0xe1, 0xa1,
	0x7f, 0xf7, 0x94, 0x15, 0xde, 0x6c, 0xa1, 0x6c, 0x6c, 0x01, 0x7d, 0xa1, 0x93, 0x77, 0x93, 0x77,
	0x9e, 0xab, 0xf0, 0x35, 0xf5, 0x27, 0x93, 0x63, 0x00, 0x6d, 0x59, 0x42, 0x58, 0xee, 0x37, 0x70,
	0x7e, 0xb2, 0xbc, 0xa7, 0x69, 0x29, 0x02, 0x6a, 0x43, 0x22, 0x30, 0x67, 0x65, 0x04, 0xe6, 0x1d,
	0x8b, 0xc0, 0x23, 0x80, 0xc4, 0x37, 0x92, 0xd8, 0x2c, 0x6b, 0x53, 0x87, 0xb0, 0x9c, 0xaf, 0x10,
	0xd6, 0x81, 0xba, 0x9b, 0xa0, 0x2c, 0x96, 0x71, 0x8b, 0x63, 0x33, 0xb7, 0x10, 0x13, 0xcf, 0x3d,
	0xc2, 0x64, 0x96, 0x89, 0x98, 0xec, 0x56, 0x36, 0x7d, 0xe9, 0x81, 0x62, 0x5f, 0xb9, 0x55, 0x60,
	0x5f, 0xf9, 0x89, 0x23, 0xf0, 0x04, 0xbe, 0x19, 0x67, 0x87, 0x34, 0x58, 0xd1, 0xd9, 0x72, 0x9a,
	0xac, 0xf0, 0x3a, 0x21, 0x43, 0x17, 0xbe, 0x9b, 0xc7, 0x91, 0xa4, 0x72, 0xd3, 0x07, 0xd6, 0xf9,
	0xcd, 0x8a, 0xaf, 0x28, 0x2d, 0x4d, 0x6a, 0xea, 0x36, 0x1f, 0xbe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x7d, 0x37, 0xcc, 0xb0, 0x05, 0x00, 0x00,
}
