// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: todo/todo.proto

package todo

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Card) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type AllCards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *AllCards) Reset() {
	*x = AllCards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCards) ProtoMessage() {}

func (x *AllCards) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCards.ProtoReflect.Descriptor instead.
func (*AllCards) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{1}
}

func (x *AllCards) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_todo_todo_proto protoreflect.FileDescriptor

var file_todo_todo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x4c, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x2c, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x20, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x32, 0xb3, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0a,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a,
	0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00, 0x12, 0x2a, 0x0a,
	0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x41, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x0a, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x44, 0x46, 0x6f, 0x72, 0x67, 0x65,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_todo_proto_rawDescOnce sync.Once
	file_todo_todo_proto_rawDescData = file_todo_todo_proto_rawDesc
)

func file_todo_todo_proto_rawDescGZIP() []byte {
	file_todo_todo_proto_rawDescOnce.Do(func() {
		file_todo_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_todo_proto_rawDescData)
	})
	return file_todo_todo_proto_rawDescData
}

var file_todo_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_todo_todo_proto_goTypes = []interface{}{
	(*Card)(nil),     // 0: todo.Card
	(*AllCards)(nil), // 1: todo.AllCards
}
var file_todo_todo_proto_depIdxs = []int32{
	0, // 0: todo.AllCards.cards:type_name -> todo.Card
	0, // 1: todo.TodoService.NewCard:input_type -> todo.Card
	0, // 2: todo.TodoService.GetAllCards:input_type -> todo.Card
	0, // 3: todo.TodoService.RemoveCard:input_type -> todo.Card
	0, // 4: todo.TodoService.MarkCardAsDone:input_type -> todo.Card
	0, // 5: todo.TodoService.NewCard:output_type -> todo.Card
	1, // 6: todo.TodoService.GetAllCards:output_type -> todo.AllCards
	0, // 7: todo.TodoService.RemoveCard:output_type -> todo.Card
	0, // 8: todo.TodoService.MarkCardAsDone:output_type -> todo.Card
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_todo_proto_init() }
func file_todo_todo_proto_init() {
	if File_todo_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_todo_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCards); i {
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
			RawDescriptor: file_todo_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_todo_proto_goTypes,
		DependencyIndexes: file_todo_todo_proto_depIdxs,
		MessageInfos:      file_todo_todo_proto_msgTypes,
	}.Build()
	File_todo_todo_proto = out.File
	file_todo_todo_proto_rawDesc = nil
	file_todo_todo_proto_goTypes = nil
	file_todo_todo_proto_depIdxs = nil
}