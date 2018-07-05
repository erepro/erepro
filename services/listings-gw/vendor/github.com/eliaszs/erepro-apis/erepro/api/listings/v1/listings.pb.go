// Code generated by protoc-gen-go. DO NOT EDIT.
// source: erepro/api/listings/v1/listings.proto

package listings // import "github.com/eliaszs/erepro-apis/erepro/api/listings/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/longrunning"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Listing_Status int32

const (
	Listing_UNSPECIFIED Listing_Status = 0
	Listing_ACTIVE      Listing_Status = 1
	Listing_SOLD        Listing_Status = 2
	Listing_RENTED      Listing_Status = 3
)

var Listing_Status_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ACTIVE",
	2: "SOLD",
	3: "RENTED",
}
var Listing_Status_value = map[string]int32{
	"UNSPECIFIED": 0,
	"ACTIVE":      1,
	"SOLD":        2,
	"RENTED":      3,
}

func (x Listing_Status) String() string {
	return proto.EnumName(Listing_Status_name, int32(x))
}
func (Listing_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{0, 0}
}

type ListingOperation_OperationType int32

const (
	ListingOperation_OPERATION_TYPE_UNSPECIFIED ListingOperation_OperationType = 0
	ListingOperation_CREATE                     ListingOperation_OperationType = 1
)

var ListingOperation_OperationType_name = map[int32]string{
	0: "OPERATION_TYPE_UNSPECIFIED",
	1: "CREATE",
}
var ListingOperation_OperationType_value = map[string]int32{
	"OPERATION_TYPE_UNSPECIFIED": 0,
	"CREATE":                     1,
}

func (x ListingOperation_OperationType) String() string {
	return proto.EnumName(ListingOperation_OperationType_name, int32(x))
}
func (ListingOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{9, 0}
}

type Listing struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Address              string               `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	NoBeds               uint32               `protobuf:"varint,4,opt,name=no_beds,json=noBeds" json:"no_beds,omitempty"`
	NoBaths              uint32               `protobuf:"varint,5,opt,name=no_baths,json=noBaths" json:"no_baths,omitempty"`
	Price                uint64               `protobuf:"varint,6,opt,name=price" json:"price,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Listing) Reset()         { *m = Listing{} }
func (m *Listing) String() string { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()    {}
func (*Listing) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{0}
}
func (m *Listing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listing.Unmarshal(m, b)
}
func (m *Listing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listing.Marshal(b, m, deterministic)
}
func (dst *Listing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listing.Merge(dst, src)
}
func (m *Listing) XXX_Size() int {
	return xxx_messageInfo_Listing.Size(m)
}
func (m *Listing) XXX_DiscardUnknown() {
	xxx_messageInfo_Listing.DiscardUnknown(m)
}

var xxx_messageInfo_Listing proto.InternalMessageInfo

func (m *Listing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listing) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Listing) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Listing) GetNoBeds() uint32 {
	if m != nil {
		return m.NoBeds
	}
	return 0
}

func (m *Listing) GetNoBaths() uint32 {
	if m != nil {
		return m.NoBaths
	}
	return 0
}

func (m *Listing) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Listing) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Listing) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type ListListingsRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	ShowDelete           bool     `protobuf:"varint,3,opt,name=show_delete,json=showDelete" json:"show_delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListListingsRequest) Reset()         { *m = ListListingsRequest{} }
func (m *ListListingsRequest) String() string { return proto.CompactTextString(m) }
func (*ListListingsRequest) ProtoMessage()    {}
func (*ListListingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{1}
}
func (m *ListListingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListListingsRequest.Unmarshal(m, b)
}
func (m *ListListingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListListingsRequest.Marshal(b, m, deterministic)
}
func (dst *ListListingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListListingsRequest.Merge(dst, src)
}
func (m *ListListingsRequest) XXX_Size() int {
	return xxx_messageInfo_ListListingsRequest.Size(m)
}
func (m *ListListingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListListingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListListingsRequest proto.InternalMessageInfo

func (m *ListListingsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListListingsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListListingsRequest) GetShowDelete() bool {
	if m != nil {
		return m.ShowDelete
	}
	return false
}

type ListListingsResponse struct {
	Listings             []*Listing `protobuf:"bytes,1,rep,name=listings" json:"listings,omitempty"`
	NextPageToken        string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListListingsResponse) Reset()         { *m = ListListingsResponse{} }
func (m *ListListingsResponse) String() string { return proto.CompactTextString(m) }
func (*ListListingsResponse) ProtoMessage()    {}
func (*ListListingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{2}
}
func (m *ListListingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListListingsResponse.Unmarshal(m, b)
}
func (m *ListListingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListListingsResponse.Marshal(b, m, deterministic)
}
func (dst *ListListingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListListingsResponse.Merge(dst, src)
}
func (m *ListListingsResponse) XXX_Size() int {
	return xxx_messageInfo_ListListingsResponse.Size(m)
}
func (m *ListListingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListListingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListListingsResponse proto.InternalMessageInfo

func (m *ListListingsResponse) GetListings() []*Listing {
	if m != nil {
		return m.Listings
	}
	return nil
}

func (m *ListListingsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type SearchListingsRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	Query                string   `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchListingsRequest) Reset()         { *m = SearchListingsRequest{} }
func (m *SearchListingsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchListingsRequest) ProtoMessage()    {}
func (*SearchListingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{3}
}
func (m *SearchListingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchListingsRequest.Unmarshal(m, b)
}
func (m *SearchListingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchListingsRequest.Marshal(b, m, deterministic)
}
func (dst *SearchListingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchListingsRequest.Merge(dst, src)
}
func (m *SearchListingsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchListingsRequest.Size(m)
}
func (m *SearchListingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchListingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchListingsRequest proto.InternalMessageInfo

func (m *SearchListingsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchListingsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchListingsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchListingsResponse struct {
	Listings             []*Listing `protobuf:"bytes,1,rep,name=listings" json:"listings,omitempty"`
	NextPageToken        string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchListingsResponse) Reset()         { *m = SearchListingsResponse{} }
func (m *SearchListingsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchListingsResponse) ProtoMessage()    {}
func (*SearchListingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{4}
}
func (m *SearchListingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchListingsResponse.Unmarshal(m, b)
}
func (m *SearchListingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchListingsResponse.Marshal(b, m, deterministic)
}
func (dst *SearchListingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchListingsResponse.Merge(dst, src)
}
func (m *SearchListingsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchListingsResponse.Size(m)
}
func (m *SearchListingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchListingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchListingsResponse proto.InternalMessageInfo

func (m *SearchListingsResponse) GetListings() []*Listing {
	if m != nil {
		return m.Listings
	}
	return nil
}

func (m *SearchListingsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetListingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListingRequest) Reset()         { *m = GetListingRequest{} }
func (m *GetListingRequest) String() string { return proto.CompactTextString(m) }
func (*GetListingRequest) ProtoMessage()    {}
func (*GetListingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{5}
}
func (m *GetListingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListingRequest.Unmarshal(m, b)
}
func (m *GetListingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListingRequest.Marshal(b, m, deterministic)
}
func (dst *GetListingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListingRequest.Merge(dst, src)
}
func (m *GetListingRequest) XXX_Size() int {
	return xxx_messageInfo_GetListingRequest.Size(m)
}
func (m *GetListingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListingRequest proto.InternalMessageInfo

func (m *GetListingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateListingRequest struct {
	Listing              *Listing `protobuf:"bytes,1,opt,name=listing" json:"listing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateListingRequest) Reset()         { *m = CreateListingRequest{} }
func (m *CreateListingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateListingRequest) ProtoMessage()    {}
func (*CreateListingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{6}
}
func (m *CreateListingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateListingRequest.Unmarshal(m, b)
}
func (m *CreateListingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateListingRequest.Marshal(b, m, deterministic)
}
func (dst *CreateListingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateListingRequest.Merge(dst, src)
}
func (m *CreateListingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateListingRequest.Size(m)
}
func (m *CreateListingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateListingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateListingRequest proto.InternalMessageInfo

func (m *CreateListingRequest) GetListing() *Listing {
	if m != nil {
		return m.Listing
	}
	return nil
}

type UpdateListingRequest struct {
	Listing              *Listing              `protobuf:"bytes,1,opt,name=listing" json:"listing,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateListingRequest) Reset()         { *m = UpdateListingRequest{} }
func (m *UpdateListingRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateListingRequest) ProtoMessage()    {}
func (*UpdateListingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{7}
}
func (m *UpdateListingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateListingRequest.Unmarshal(m, b)
}
func (m *UpdateListingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateListingRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateListingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateListingRequest.Merge(dst, src)
}
func (m *UpdateListingRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateListingRequest.Size(m)
}
func (m *UpdateListingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateListingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateListingRequest proto.InternalMessageInfo

func (m *UpdateListingRequest) GetListing() *Listing {
	if m != nil {
		return m.Listing
	}
	return nil
}

func (m *UpdateListingRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteListingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteListingRequest) Reset()         { *m = DeleteListingRequest{} }
func (m *DeleteListingRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteListingRequest) ProtoMessage()    {}
func (*DeleteListingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{8}
}
func (m *DeleteListingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteListingRequest.Unmarshal(m, b)
}
func (m *DeleteListingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteListingRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteListingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteListingRequest.Merge(dst, src)
}
func (m *DeleteListingRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteListingRequest.Size(m)
}
func (m *DeleteListingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteListingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteListingRequest proto.InternalMessageInfo

func (m *DeleteListingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListingOperation struct {
	Name                 string                         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	OperationType        ListingOperation_OperationType `protobuf:"varint,2,opt,name=operation_type,json=operationType,enum=erepro.api.listings.v1.ListingOperation_OperationType" json:"operation_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListingOperation) Reset()         { *m = ListingOperation{} }
func (m *ListingOperation) String() string { return proto.CompactTextString(m) }
func (*ListingOperation) ProtoMessage()    {}
func (*ListingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_listings_5e5a227cf74f0377, []int{9}
}
func (m *ListingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListingOperation.Unmarshal(m, b)
}
func (m *ListingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListingOperation.Marshal(b, m, deterministic)
}
func (dst *ListingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListingOperation.Merge(dst, src)
}
func (m *ListingOperation) XXX_Size() int {
	return xxx_messageInfo_ListingOperation.Size(m)
}
func (m *ListingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ListingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ListingOperation proto.InternalMessageInfo

func (m *ListingOperation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListingOperation) GetOperationType() ListingOperation_OperationType {
	if m != nil {
		return m.OperationType
	}
	return ListingOperation_OPERATION_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Listing)(nil), "erepro.api.listings.v1.Listing")
	proto.RegisterType((*ListListingsRequest)(nil), "erepro.api.listings.v1.ListListingsRequest")
	proto.RegisterType((*ListListingsResponse)(nil), "erepro.api.listings.v1.ListListingsResponse")
	proto.RegisterType((*SearchListingsRequest)(nil), "erepro.api.listings.v1.SearchListingsRequest")
	proto.RegisterType((*SearchListingsResponse)(nil), "erepro.api.listings.v1.SearchListingsResponse")
	proto.RegisterType((*GetListingRequest)(nil), "erepro.api.listings.v1.GetListingRequest")
	proto.RegisterType((*CreateListingRequest)(nil), "erepro.api.listings.v1.CreateListingRequest")
	proto.RegisterType((*UpdateListingRequest)(nil), "erepro.api.listings.v1.UpdateListingRequest")
	proto.RegisterType((*DeleteListingRequest)(nil), "erepro.api.listings.v1.DeleteListingRequest")
	proto.RegisterType((*ListingOperation)(nil), "erepro.api.listings.v1.ListingOperation")
	proto.RegisterEnum("erepro.api.listings.v1.Listing_Status", Listing_Status_name, Listing_Status_value)
	proto.RegisterEnum("erepro.api.listings.v1.ListingOperation_OperationType", ListingOperation_OperationType_name, ListingOperation_OperationType_value)
}

func init() {
	proto.RegisterFile("erepro/api/listings/v1/listings.proto", fileDescriptor_listings_5e5a227cf74f0377)
}

var fileDescriptor_listings_5e5a227cf74f0377 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x66, 0xf3, 0x63, 0xbb, 0xc7, 0x71, 0x6a, 0xa6, 0x6e, 0xbb, 0x18, 0xd2, 0x98, 0x45, 0x80,
	0x31, 0x74, 0x57, 0x35, 0x12, 0x12, 0x89, 0x40, 0xe4, 0x67, 0x8b, 0x22, 0x95, 0x24, 0xac, 0x5d,
	0x24, 0x90, 0xd0, 0x6a, 0x12, 0x4f, 0x9d, 0x51, 0xec, 0x9d, 0xed, 0xce, 0xb8, 0x25, 0x89, 0x2c,
	0x24, 0x1e, 0xa0, 0x5c, 0xf0, 0x08, 0xbc, 0x09, 0xaf, 0xc0, 0x25, 0xb7, 0x3c, 0x08, 0x9a, 0xd9,
	0x9d, 0x8d, 0xd7, 0x59, 0x63, 0x5f, 0x54, 0xbd, 0x9b, 0x39, 0xe7, 0x3b, 0x73, 0x7e, 0xfc, 0x7d,
	0x67, 0x0d, 0x1f, 0x92, 0x88, 0x84, 0x11, 0x73, 0x70, 0x48, 0x9d, 0x01, 0xe5, 0x82, 0x06, 0x7d,
	0xee, 0xbc, 0x78, 0x94, 0x9e, 0xed, 0x30, 0x62, 0x82, 0xa1, 0x7b, 0x31, 0xcc, 0xc6, 0x21, 0xb5,
	0x53, 0xd7, 0x8b, 0x47, 0xf5, 0xf7, 0xfa, 0x8c, 0xf5, 0x07, 0x44, 0x85, 0xe3, 0x20, 0x60, 0x02,
	0x0b, 0xca, 0x82, 0x24, 0xaa, 0xfe, 0x41, 0xe2, 0x1d, 0xb0, 0xa0, 0x1f, 0x8d, 0x82, 0x80, 0x06,
	0x7d, 0x87, 0x85, 0x24, 0xca, 0x80, 0x36, 0x13, 0x90, 0xba, 0x9d, 0x8c, 0x9e, 0x39, 0x82, 0x0e,
	0x09, 0x17, 0x78, 0x18, 0x26, 0x80, 0xc6, 0x34, 0xe0, 0x19, 0x25, 0x83, 0x9e, 0x3f, 0xc4, 0xfc,
	0x3c, 0x46, 0x58, 0xff, 0x2c, 0x41, 0xf1, 0x49, 0x5c, 0x15, 0x42, 0xb0, 0x12, 0xe0, 0x21, 0x31,
	0x8d, 0x86, 0xd1, 0xbc, 0xe5, 0xa9, 0x33, 0x7a, 0x1f, 0xd6, 0x7a, 0x94, 0x87, 0x03, 0x7c, 0xe1,
	0x2b, 0xdf, 0x92, 0xf2, 0x95, 0x13, 0xdb, 0xa1, 0x84, 0x98, 0x50, 0xc4, 0xbd, 0x5e, 0x44, 0x38,
	0x37, 0x97, 0x95, 0x57, 0x5f, 0xd1, 0x7d, 0x28, 0x06, 0xcc, 0x3f, 0x21, 0x3d, 0x6e, 0xae, 0x34,
	0x8c, 0x66, 0xc5, 0x2b, 0x04, 0x6c, 0x97, 0xf4, 0x38, 0x7a, 0x07, 0x4a, 0xd2, 0x81, 0xc5, 0x19,
	0x37, 0x57, 0x95, 0xa7, 0x18, 0xb0, 0x5d, 0x79, 0x45, 0x35, 0x58, 0x0d, 0x23, 0x7a, 0x4a, 0xcc,
	0x42, 0xc3, 0x68, 0xae, 0x78, 0xf1, 0x05, 0x6d, 0x43, 0xf9, 0x34, 0x22, 0x58, 0x10, 0x5f, 0xb6,
	0x68, 0x16, 0x1b, 0x46, 0xb3, 0xdc, 0xae, 0xdb, 0x71, 0x7b, 0xb6, 0x6e, 0xcf, 0xee, 0xea, 0xfe,
	0x3d, 0x88, 0xe1, 0xd2, 0x20, 0x83, 0x47, 0x61, 0x2f, 0x0d, 0x2e, 0xcd, 0x0f, 0x8e, 0xe1, 0xd2,
	0x60, 0x6d, 0x43, 0xa1, 0x23, 0xb0, 0x18, 0x71, 0x74, 0x1b, 0xca, 0x4f, 0x0f, 0x3b, 0xc7, 0xee,
	0xde, 0xc1, 0xe3, 0x03, 0x77, 0xbf, 0xfa, 0x16, 0x02, 0x28, 0xec, 0xec, 0x75, 0x0f, 0x7e, 0x70,
	0xab, 0x06, 0x2a, 0xc1, 0x4a, 0xe7, 0xe8, 0xc9, 0x7e, 0x75, 0x49, 0x5a, 0x3d, 0xf7, 0xb0, 0xeb,
	0xee, 0x57, 0x97, 0xad, 0x08, 0xee, 0xc8, 0xe1, 0x26, 0x03, 0xe6, 0x1e, 0x79, 0x3e, 0x22, 0x5c,
	0xa0, 0x77, 0xe1, 0x56, 0x88, 0xfb, 0xc4, 0xe7, 0xf4, 0x32, 0x9e, 0xf6, 0xaa, 0x57, 0x92, 0x86,
	0x0e, 0xbd, 0x24, 0x68, 0x03, 0x40, 0x39, 0x05, 0x3b, 0x27, 0x41, 0x32, 0x6f, 0x05, 0xef, 0x4a,
	0x03, 0xda, 0x84, 0x32, 0x3f, 0x63, 0x2f, 0xfd, 0x1e, 0x19, 0x10, 0x41, 0xd4, 0xc4, 0x4b, 0x1e,
	0x48, 0xd3, 0xbe, 0xb2, 0x58, 0x57, 0x50, 0xcb, 0xe6, 0xe4, 0x21, 0x0b, 0xb8, 0x9c, 0x42, 0x49,
	0xd3, 0xcf, 0x34, 0x1a, 0xcb, 0xcd, 0x72, 0x7b, 0xd3, 0xce, 0xa7, 0xa6, 0x9d, 0xc4, 0x7a, 0x69,
	0x00, 0xfa, 0x08, 0x6e, 0x07, 0xe4, 0x17, 0xe1, 0xdf, 0xa8, 0xac, 0x22, 0xcd, 0xc7, 0xba, 0x3a,
	0x8b, 0xc2, 0xdd, 0x0e, 0xc1, 0xd1, 0xe9, 0xd9, 0xeb, 0x6c, 0xb9, 0x06, 0xab, 0xcf, 0x47, 0x24,
	0xba, 0x48, 0xe8, 0x15, 0x5f, 0xac, 0x31, 0xdc, 0x9b, 0x4e, 0xf5, 0x26, 0x3b, 0xfd, 0x18, 0xde,
	0xfe, 0x96, 0xe8, 0x29, 0xeb, 0x2e, 0x73, 0x14, 0x64, 0x7d, 0x0f, 0xb5, 0x3d, 0xc5, 0xc5, 0x29,
	0xec, 0x97, 0x50, 0x4c, 0x92, 0x2a, 0xf8, 0x02, 0x45, 0x6a, 0xbc, 0xf5, 0xca, 0x80, 0xda, 0x53,
	0x45, 0xd1, 0xd7, 0xf6, 0xe6, 0x84, 0x48, 0xe4, 0x76, 0x50, 0x3d, 0xe7, 0x89, 0xe4, 0xb1, 0x5c,
	0x20, 0xdf, 0x61, 0x7e, 0xae, 0x45, 0x22, 0xcf, 0x56, 0x0b, 0x6a, 0x31, 0xfb, 0x16, 0x98, 0xc7,
	0x5f, 0x06, 0x54, 0x13, 0xd8, 0x91, 0x5e, 0x68, 0xb9, 0xab, 0xe7, 0x67, 0x58, 0x4f, 0x37, 0x9e,
	0x2f, 0x2e, 0xc2, 0x78, 0xf9, 0xac, 0xb7, 0xbf, 0x98, 0xd3, 0x53, 0xfa, 0xaa, 0x9d, 0x9e, 0xba,
	0x17, 0x21, 0xf1, 0x2a, 0x6c, 0xf2, 0x6a, 0x6d, 0x43, 0x25, 0xe3, 0x47, 0x0f, 0xa0, 0x7e, 0x74,
	0xec, 0x7a, 0x3b, 0xdd, 0x83, 0xa3, 0x43, 0xbf, 0xfb, 0xe3, 0xb1, 0xeb, 0xdf, 0x90, 0xfb, 0x9e,
	0xe7, 0xee, 0x74, 0xdd, 0xaa, 0xd1, 0xfe, 0xb3, 0x00, 0x25, 0xcd, 0x3b, 0xf4, 0x2b, 0xac, 0x4d,
	0x2a, 0x0e, 0x7d, 0xfa, 0x7f, 0x05, 0x4e, 0x09, 0xa3, 0xfe, 0xd9, 0x62, 0xe0, 0x98, 0xda, 0x56,
	0xed, 0xb7, 0xbf, 0xff, 0xfd, 0x63, 0x69, 0x1d, 0xad, 0x4d, 0x7e, 0x68, 0xd0, 0xef, 0x06, 0xac,
	0x67, 0xb5, 0x80, 0x1e, 0xce, 0x7a, 0x36, 0x57, 0x9e, 0x75, 0x7b, 0x51, 0x78, 0x52, 0xc7, 0x03,
	0x55, 0x87, 0x69, 0xdd, 0x99, 0xac, 0x63, 0x8b, 0x2b, 0xf0, 0x96, 0xd1, 0x42, 0x2f, 0x01, 0xae,
	0xd5, 0x81, 0x3e, 0x99, 0xf5, 0xfa, 0x0d, 0x05, 0xd5, 0xe7, 0x11, 0xd6, 0xda, 0x50, 0x99, 0xef,
	0xa3, 0xbb, 0x32, 0xf3, 0x95, 0x24, 0xca, 0x57, 0x1a, 0x36, 0x76, 0x5a, 0xe8, 0x0a, 0x2a, 0x19,
	0xb5, 0xa1, 0x99, 0xf3, 0xcd, 0x13, 0x65, 0x7d, 0x43, 0x13, 0x7e, 0xe2, 0xbb, 0x7b, 0x4d, 0x23,
	0x9d, 0xdc, 0xca, 0x8c, 0x7f, 0x2b, 0xd5, 0xd0, 0x2b, 0x03, 0x2a, 0x19, 0x5d, 0xce, 0xce, 0x9e,
	0x27, 0xdf, 0xf9, 0xcd, 0xdb, 0x2a, 0x7f, 0xb3, 0xbd, 0xa1, 0x9a, 0x4f, 0x10, 0x76, 0x66, 0x08,
	0x4e, 0x6b, 0x7c, 0x5d, 0xd0, 0x18, 0x2a, 0x19, 0x5d, 0xce, 0xae, 0x27, 0x4f, 0xbe, 0x0b, 0xff,
	0x18, 0xad, 0x9c, 0x1f, 0xc3, 0x69, 0x8d, 0x77, 0xbf, 0xf9, 0xe9, 0xeb, 0x3e, 0x15, 0x67, 0xa3,
	0x13, 0xfb, 0x94, 0x0d, 0x1d, 0x32, 0xa0, 0x98, 0x5f, 0x72, 0x27, 0x7e, 0xf3, 0x21, 0x0e, 0xa9,
	0x3e, 0x4f, 0xff, 0x85, 0xda, 0xd6, 0xe7, 0x93, 0x82, 0x5a, 0x3c, 0x9f, 0xff, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x4f, 0x9e, 0x0f, 0x6c, 0x09, 0x00, 0x00,
}