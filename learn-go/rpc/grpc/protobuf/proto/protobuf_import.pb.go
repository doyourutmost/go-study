// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: protobuf_import.proto

package proto

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

type Sex int32

const (
	Sex_WOMAN Sex = 0
	Sex_MAN   Sex = 1
	Sex_OTHER Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "WOMAN",
		1: "MAN",
		2: "OTHER",
	}
	Sex_value = map[string]int32{
		"WOMAN": 0,
		"MAN":   1,
		"OTHER": 2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_import_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_protobuf_import_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_import_proto_rawDescGZIP(), []int{0}
}

var File_protobuf_import_proto protoreflect.FileDescriptor

var file_protobuf_import_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x24, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x09,
	0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_import_proto_rawDescOnce sync.Once
	file_protobuf_import_proto_rawDescData = file_protobuf_import_proto_rawDesc
)

func file_protobuf_import_proto_rawDescGZIP() []byte {
	file_protobuf_import_proto_rawDescOnce.Do(func() {
		file_protobuf_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_import_proto_rawDescData)
	})
	return file_protobuf_import_proto_rawDescData
}

var file_protobuf_import_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_import_proto_goTypes = []interface{}{
	(Sex)(0), // 0: Sex
}
var file_protobuf_import_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_import_proto_init() }
func file_protobuf_import_proto_init() {
	if File_protobuf_import_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_import_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_import_proto_goTypes,
		DependencyIndexes: file_protobuf_import_proto_depIdxs,
		EnumInfos:         file_protobuf_import_proto_enumTypes,
	}.Build()
	File_protobuf_import_proto = out.File
	file_protobuf_import_proto_rawDesc = nil
	file_protobuf_import_proto_goTypes = nil
	file_protobuf_import_proto_depIdxs = nil
}
