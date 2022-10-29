// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: basic.proto

package __

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

type Gender int32

const (
	Gender_NONE   Gender = 0
	Gender_MALE   Gender = 1
	Gender_FEMALE Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "NONE",
		1: "MALE",
		2: "FEMALE",
	}
	Gender_value = map[string]int32{
		"NONE":   0,
		"MALE":   1,
		"FEMALE": 2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_basic_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_basic_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{0}
}

type PersonItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string            `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName   *string           `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3,oneof" json:"second_name,omitempty"`
	ThirdNames   []string          `protobuf:"bytes,3,rep,name=third_names,json=thirdNames,proto3" json:"third_names,omitempty"`
	ForthNameMap map[string]string `protobuf:"bytes,4,rep,name=forth_name_map,json=forthNameMap,proto3" json:"forth_name_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Gender       Gender            `protobuf:"varint,5,opt,name=gender,proto3,enum=lib.Gender" json:"gender,omitempty"`
}

func (x *PersonItem) Reset() {
	*x = PersonItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonItem) ProtoMessage() {}

func (x *PersonItem) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonItem.ProtoReflect.Descriptor instead.
func (*PersonItem) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{0}
}

func (x *PersonItem) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PersonItem) GetSecondName() string {
	if x != nil && x.SecondName != nil {
		return *x.SecondName
	}
	return ""
}

func (x *PersonItem) GetThirdNames() []string {
	if x != nil {
		return x.ThirdNames
	}
	return nil
}

func (x *PersonItem) GetForthNameMap() map[string]string {
	if x != nil {
		return x.ForthNameMap
	}
	return nil
}

func (x *PersonItem) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_NONE
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P *PersonItem `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	L *Lesson     `protobuf:"bytes,2,opt,name=l,proto3" json:"l,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{1}
}

func (x *Student) GetP() *PersonItem {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *Student) GetL() *Lesson {
	if x != nil {
		return x.L
	}
	return nil
}

var File_basic_proto protoreflect.FileDescriptor

var file_basic_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6c,
	0x69, 0x62, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1,
	0x02, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x74, 0x68, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69,
	0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x46, 0x6f, 0x72,
	0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x66, 0x6f, 0x72, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x23, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x6c,
	0x69, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x43, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x01, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x01, 0x70, 0x12, 0x19, 0x0a, 0x01,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x52, 0x01, 0x6c, 0x2a, 0x28, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x02, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_proto_rawDescOnce sync.Once
	file_basic_proto_rawDescData = file_basic_proto_rawDesc
)

func file_basic_proto_rawDescGZIP() []byte {
	file_basic_proto_rawDescOnce.Do(func() {
		file_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_proto_rawDescData)
	})
	return file_basic_proto_rawDescData
}

var file_basic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_basic_proto_goTypes = []interface{}{
	(Gender)(0),        // 0: lib.Gender
	(*PersonItem)(nil), // 1: lib.PersonItem
	(*Student)(nil),    // 2: lib.Student
	nil,                // 3: lib.PersonItem.ForthNameMapEntry
	(*Lesson)(nil),     // 4: lib.Lesson
}
var file_basic_proto_depIdxs = []int32{
	3, // 0: lib.PersonItem.forth_name_map:type_name -> lib.PersonItem.ForthNameMapEntry
	0, // 1: lib.PersonItem.gender:type_name -> lib.Gender
	1, // 2: lib.Student.p:type_name -> lib.PersonItem
	4, // 3: lib.Student.l:type_name -> lib.Lesson
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_basic_proto_init() }
func file_basic_proto_init() {
	if File_basic_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonItem); i {
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
		file_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
	file_basic_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basic_proto_goTypes,
		DependencyIndexes: file_basic_proto_depIdxs,
		EnumInfos:         file_basic_proto_enumTypes,
		MessageInfos:      file_basic_proto_msgTypes,
	}.Build()
	File_basic_proto = out.File
	file_basic_proto_rawDesc = nil
	file_basic_proto_goTypes = nil
	file_basic_proto_depIdxs = nil
}
