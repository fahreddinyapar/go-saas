// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: localize/i18n.proto

package localize

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_localize_i18n_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1208,
		Name:          "localize.msg_key",
		Tag:           "bytes,1208,opt,name=msg_key",
		Filename:      "localize/i18n.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1209,
		Name:          "localize.skip_locale",
		Tag:           "varint,1209,opt,name=skip_locale",
		Filename:      "localize/i18n.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string msg_key = 1208;
	E_MsgKey = &file_localize_i18n_proto_extTypes[0]
	// optional bool skip_locale = 1209;
	E_SkipLocale = &file_localize_i18n_proto_extTypes[1]
)

var File_localize_i18n_proto protoreflect.FileDescriptor

var file_localize_i18n_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x69, 0x31, 0x38, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3a, 0x3b, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xb8, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x3a, 0x43,
	0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_localize_i18n_proto_goTypes = []interface{}{
	(*descriptorpb.EnumValueOptions)(nil), // 0: google.protobuf.EnumValueOptions
}
var file_localize_i18n_proto_depIdxs = []int32{
	0, // 0: localize.msg_key:extendee -> google.protobuf.EnumValueOptions
	0, // 1: localize.skip_locale:extendee -> google.protobuf.EnumValueOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_localize_i18n_proto_init() }
func file_localize_i18n_proto_init() {
	if File_localize_i18n_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_localize_i18n_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_localize_i18n_proto_goTypes,
		DependencyIndexes: file_localize_i18n_proto_depIdxs,
		ExtensionInfos:    file_localize_i18n_proto_extTypes,
	}.Build()
	File_localize_i18n_proto = out.File
	file_localize_i18n_proto_rawDesc = nil
	file_localize_i18n_proto_goTypes = nil
	file_localize_i18n_proto_depIdxs = nil
}
