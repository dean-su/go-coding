// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: model/book.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Category int32

const (
	Category_Novel     Category = 0
	Category_SciFi     Category = 1
	Category_Fantasy   Category = 2
	Category_Spiritual Category = 3
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "Novel",
		1: "SciFi",
		2: "Fantasy",
		3: "Spiritual",
	}
	Category_value = map[string]int32{
		"Novel":     0,
		"SciFi":     1,
		"Fantasy":   2,
		"Spiritual": 3,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_model_book_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_model_book_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title    string    `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Authors  []*Author `protobuf:"bytes,3,rep,name=Authors,proto3" json:"Authors,omitempty"`
	Category Category  `protobuf:"varint,4,opt,name=Category,proto3,enum=Category" json:"Category,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_model_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Book) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_Novel
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_model_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_model_book_proto protoreflect.FileDescriptor

var file_model_book_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x76, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x3c, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x63, 0x69, 0x46, 0x69, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61,
	0x6e, 0x74, 0x61, 0x73, 0x79, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x75, 0x61, 0x6c, 0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x30, 0x32, 0x36, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_book_proto_rawDescOnce sync.Once
	file_model_book_proto_rawDescData = file_model_book_proto_rawDesc
)

func file_model_book_proto_rawDescGZIP() []byte {
	file_model_book_proto_rawDescOnce.Do(func() {
		file_model_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_book_proto_rawDescData)
	})
	return file_model_book_proto_rawDescData
}

var file_model_book_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_book_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_book_proto_goTypes = []interface{}{
	(Category)(0),  // 0: Category
	(*Book)(nil),   // 1: Book
	(*Author)(nil), // 2: Author
}
var file_model_book_proto_depIdxs = []int32{
	2, // 0: Book.Authors:type_name -> Author
	0, // 1: Book.Category:type_name -> Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_book_proto_init() }
func file_model_book_proto_init() {
	if File_model_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_book_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_book_proto_goTypes,
		DependencyIndexes: file_model_book_proto_depIdxs,
		EnumInfos:         file_model_book_proto_enumTypes,
		MessageInfos:      file_model_book_proto_msgTypes,
	}.Build()
	File_model_book_proto = out.File
	file_model_book_proto_rawDesc = nil
	file_model_book_proto_goTypes = nil
	file_model_book_proto_depIdxs = nil
}
