// Code generated by protoc-gen-go. DO NOT EDIT.
// source: library/book_service.proto

// Package library exposes a list of books
// over a gRPC API.

package library

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// BookType describes the different types
// a book in the library can be.
type BookType int32

const (
	// Hardcover is a book with a hard back.
	BookType_HARDCOVER BookType = 0
	// Paperback is a book with a soft back.
	BookType_PAPERBACK BookType = 1
	// Audiobook is an audio recording of the book.
	BookType_AUDIOBOOK BookType = 2
)

var BookType_name = map[int32]string{
	0: "HARDCOVER",
	1: "PAPERBACK",
	2: "AUDIOBOOK",
}

var BookType_value = map[string]int32{
	"HARDCOVER": 0,
	"PAPERBACK": 1,
	"AUDIOBOOK": 2,
}

func (x BookType) String() string {
	return proto.EnumName(BookType_name, int32(x))
}

func (BookType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b32465fb641f4481, []int{0}
}

// Publisher describes a Book Publisher.
type Publisher struct {
	// Name is the name of the Publisher.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Publisher) Reset()         { *m = Publisher{} }
func (m *Publisher) String() string { return proto.CompactTextString(m) }
func (*Publisher) ProtoMessage()    {}
func (*Publisher) Descriptor() ([]byte, []int) {
	return fileDescriptor_b32465fb641f4481, []int{0}
}

func (m *Publisher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Publisher.Unmarshal(m, b)
}
func (m *Publisher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Publisher.Marshal(b, m, deterministic)
}
func (m *Publisher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publisher.Merge(m, src)
}
func (m *Publisher) XXX_Size() int {
	return xxx_messageInfo_Publisher.Size(m)
}
func (m *Publisher) XXX_DiscardUnknown() {
	xxx_messageInfo_Publisher.DiscardUnknown(m)
}

var xxx_messageInfo_Publisher proto.InternalMessageInfo

func (m *Publisher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Book represents a book in the library.
type Book struct {
	// Isbn is the ISBN number of the book.
	Isbn int64 `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	// Title is the title of the book.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Author is the author of the book.
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	// BookType is the type of the book.
	BookType BookType `protobuf:"varint,4,opt,name=book_type,json=bookType,proto3,enum=library.BookType" json:"book_type,omitempty"`
	// PublishingMethod is the publishing method
	// used for this Book.
	//
	// Types that are valid to be assigned to PublishingMethod:
	//	*Book_SelfPublished
	//	*Book_Publisher
	PublishingMethod isBook_PublishingMethod `protobuf_oneof:"publishing_method"`
	// PublicationDate is the time of publication of the book.
	PublicationDate      *timestamp.Timestamp `protobuf:"bytes,7,opt,name=publication_date,json=publicationDate,proto3" json:"publication_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_b32465fb641f4481, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetBookType() BookType {
	if m != nil {
		return m.BookType
	}
	return BookType_HARDCOVER
}

type isBook_PublishingMethod interface {
	isBook_PublishingMethod()
}

type Book_SelfPublished struct {
	SelfPublished bool `protobuf:"varint,5,opt,name=self_published,json=selfPublished,proto3,oneof"`
}

type Book_Publisher struct {
	Publisher *Publisher `protobuf:"bytes,6,opt,name=publisher,proto3,oneof"`
}

func (*Book_SelfPublished) isBook_PublishingMethod() {}

func (*Book_Publisher) isBook_PublishingMethod() {}

func (m *Book) GetPublishingMethod() isBook_PublishingMethod {
	if m != nil {
		return m.PublishingMethod
	}
	return nil
}

func (m *Book) GetSelfPublished() bool {
	if x, ok := m.GetPublishingMethod().(*Book_SelfPublished); ok {
		return x.SelfPublished
	}
	return false
}

func (m *Book) GetPublisher() *Publisher {
	if x, ok := m.GetPublishingMethod().(*Book_Publisher); ok {
		return x.Publisher
	}
	return nil
}

func (m *Book) GetPublicationDate() *timestamp.Timestamp {
	if m != nil {
		return m.PublicationDate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Book) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Book_SelfPublished)(nil),
		(*Book_Publisher)(nil),
	}
}

// GetBookRequest is the input to the GetBook method.
type GetBookRequest struct {
	// Isbn is the ISBN with which
	// to match against the ISBN of a book in the library.
	Isbn                 int64    `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b32465fb641f4481, []int{2}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

// QueryBooksRequest is the input to the QueryBooks method.
type QueryBooksRequest struct {
	// AuthorPrefix is the prefix with which
	// to match against the author of a book in the library.
	AuthorPrefix         string   `protobuf:"bytes,1,opt,name=author_prefix,json=authorPrefix,proto3" json:"author_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryBooksRequest) Reset()         { *m = QueryBooksRequest{} }
func (m *QueryBooksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBooksRequest) ProtoMessage()    {}
func (*QueryBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b32465fb641f4481, []int{3}
}

func (m *QueryBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryBooksRequest.Unmarshal(m, b)
}
func (m *QueryBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryBooksRequest.Marshal(b, m, deterministic)
}
func (m *QueryBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBooksRequest.Merge(m, src)
}
func (m *QueryBooksRequest) XXX_Size() int {
	return xxx_messageInfo_QueryBooksRequest.Size(m)
}
func (m *QueryBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBooksRequest proto.InternalMessageInfo

func (m *QueryBooksRequest) GetAuthorPrefix() string {
	if m != nil {
		return m.AuthorPrefix
	}
	return ""
}

func init() {
	proto.RegisterEnum("library.BookType", BookType_name, BookType_value)
	proto.RegisterType((*Publisher)(nil), "library.Publisher")
	proto.RegisterType((*Book)(nil), "library.Book")
	proto.RegisterType((*GetBookRequest)(nil), "library.GetBookRequest")
	proto.RegisterType((*QueryBooksRequest)(nil), "library.QueryBooksRequest")
}

func init() { proto.RegisterFile("library/book_service.proto", fileDescriptor_b32465fb641f4481) }

var fileDescriptor_b32465fb641f4481 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xdb, 0x34, 0x1f, 0x53, 0x12, 0x92, 0x05, 0x81, 0x95, 0x4b, 0xa3, 0x80, 0x44, 0x84,
	0x14, 0x1b, 0xa5, 0x07, 0x90, 0x38, 0x25, 0x4d, 0xd4, 0x54, 0x3d, 0x24, 0x98, 0xc2, 0x81, 0x4b,
	0xb4, 0x9b, 0x4c, 0xec, 0xa5, 0xb6, 0xd7, 0xec, 0xae, 0x29, 0x39, 0xf1, 0x9b, 0xf8, 0x87, 0xc8,
	0x6b, 0x3b, 0xa5, 0x2a, 0xb7, 0x99, 0x37, 0xef, 0x8d, 0xdf, 0x3c, 0x2f, 0xf4, 0x42, 0xce, 0x24,
	0x95, 0x7b, 0x97, 0x09, 0x71, 0xbb, 0x56, 0x28, 0x7f, 0xf2, 0x0d, 0x3a, 0x89, 0x14, 0x5a, 0x90,
	0x7a, 0x31, 0xeb, 0x9d, 0xf9, 0x42, 0xf8, 0x21, 0xba, 0x06, 0x66, 0xe9, 0xce, 0xd5, 0x3c, 0x42,
	0xa5, 0x69, 0x94, 0xe4, 0xcc, 0xc1, 0x19, 0x34, 0x57, 0x29, 0x0b, 0xb9, 0x0a, 0x50, 0x12, 0x02,
	0xd5, 0x98, 0x46, 0x68, 0x5b, 0x7d, 0x6b, 0xd8, 0xf4, 0x4c, 0x3d, 0xf8, 0x73, 0x04, 0xd5, 0xa9,
	0x10, 0xb7, 0xd9, 0x90, 0x2b, 0x16, 0x9b, 0xe1, 0xb1, 0x67, 0x6a, 0xf2, 0x1c, 0x4e, 0x34, 0xd7,
	0x21, 0xda, 0x47, 0x46, 0x91, 0x37, 0xe4, 0x05, 0xd4, 0x68, 0xaa, 0x03, 0x21, 0xed, 0x63, 0x03,
	0x17, 0x1d, 0x71, 0xa0, 0x69, 0xbc, 0xea, 0x7d, 0x82, 0x76, 0xb5, 0x6f, 0x0d, 0xdb, 0xe3, 0xae,
	0x53, 0x38, 0x75, 0xb2, 0x6f, 0xdc, 0xec, 0x13, 0xf4, 0x1a, 0xac, 0xa8, 0xc8, 0x1b, 0x68, 0x2b,
	0x0c, 0x77, 0xeb, 0xa4, 0x30, 0xb8, 0xb5, 0x4f, 0xfa, 0xd6, 0xb0, 0xb1, 0xa8, 0x78, 0xad, 0x0c,
	0x2f, 0x7d, 0x6f, 0xc9, 0x18, 0x9a, 0x25, 0x47, 0xda, 0xb5, 0xbe, 0x35, 0x3c, 0x1d, 0x93, 0xc3,
	0xe2, 0xc3, 0x79, 0x8b, 0x8a, 0x77, 0x4f, 0x23, 0x73, 0xe8, 0x98, 0x66, 0x43, 0x35, 0x17, 0xf1,
	0x7a, 0x4b, 0x35, 0xda, 0x75, 0x23, 0xed, 0x39, 0x79, 0x68, 0x4e, 0x19, 0x9a, 0x73, 0x53, 0x86,
	0xe6, 0x3d, 0xfd, 0x47, 0x33, 0xa3, 0x1a, 0xa7, 0xcf, 0xa0, 0x5b, 0xec, 0xe4, 0xb1, 0xbf, 0x8e,
	0x50, 0x07, 0x62, 0x3b, 0x78, 0x0d, 0xed, 0x4b, 0xd4, 0xd9, 0x45, 0x1e, 0xfe, 0x48, 0x51, 0xe9,
	0xff, 0x85, 0x37, 0xf8, 0x00, 0xdd, 0x4f, 0x29, 0xca, 0x7d, 0xc6, 0x53, 0x25, 0xf1, 0x15, 0xb4,
	0xf2, 0xb4, 0xd6, 0x89, 0xc4, 0x1d, 0xff, 0x55, 0xfc, 0x8b, 0x27, 0x39, 0xb8, 0x32, 0xd8, 0xdb,
	0xf7, 0xd0, 0x28, 0xe3, 0x22, 0x2d, 0x68, 0x2e, 0x26, 0xde, 0xec, 0x62, 0xf9, 0x75, 0xee, 0x75,
	0x2a, 0x59, 0xbb, 0x9a, 0xac, 0xe6, 0xde, 0x74, 0x72, 0x71, 0xdd, 0xb1, 0xb2, 0x76, 0xf2, 0x65,
	0x76, 0xb5, 0x9c, 0x2e, 0x97, 0xd7, 0x9d, 0xa3, 0xf1, 0x6f, 0x38, 0xcd, 0x84, 0x9f, 0xf3, 0xc7,
	0x42, 0xce, 0xa1, 0x5e, 0xf8, 0x24, 0x2f, 0x0f, 0x79, 0x3d, 0x74, 0xde, 0x6b, 0x3d, 0xf8, 0x43,
	0x83, 0x0a, 0xf9, 0x08, 0x70, 0x6f, 0x9b, 0xf4, 0x0e, 0xe3, 0x47, 0xb7, 0x3c, 0x92, 0xbe, 0xb3,
	0xa6, 0x57, 0xdf, 0x2e, 0x7d, 0xae, 0x83, 0x94, 0x39, 0x1b, 0x11, 0xb9, 0xdf, 0x45, 0x40, 0x63,
	0x26, 0x69, 0xbc, 0x0d, 0x84, 0x54, 0xda, 0x95, 0xa9, 0xd2, 0x23, 0x5f, 0x26, 0x9b, 0xd1, 0x1d,
	0xb2, 0xd1, 0x1d, 0x55, 0xd1, 0x48, 0xa3, 0xd2, 0xf9, 0xeb, 0x75, 0x7d, 0x8c, 0x5d, 0x5f, 0xb8,
	0xc5, 0x4e, 0x56, 0x33, 0xe8, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x58, 0x80, 0x53,
	0x08, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/library.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BookService_serviceDesc.Streams[0], "/library.BookService/QueryBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceQueryBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceQueryBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(*QueryBooksRequest, BookService_QueryBooksServer) error
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) GetBook(ctx context.Context, req *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookServiceServer) QueryBooks(req *QueryBooksRequest, srv BookService_QueryBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryBooks not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_QueryBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).QueryBooks(m, &bookServiceQueryBooksServer{stream})
}

type BookService_QueryBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceQueryBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceQueryBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "library.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryBooks",
			Handler:       _BookService_QueryBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "library/book_service.proto",
}
