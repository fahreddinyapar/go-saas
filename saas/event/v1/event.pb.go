// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: saas/event/v1/event.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TenantCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Region        string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	SeparateDb    bool   `protobuf:"varint,6,opt,name=separate_db,json=separateDb,proto3" json:"separate_db,omitempty"`
	AdminEmail    string `protobuf:"bytes,7,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
	AdminUsername string `protobuf:"bytes,8,opt,name=admin_username,json=adminUsername,proto3" json:"admin_username,omitempty"`
	AdminPassword string `protobuf:"bytes,9,opt,name=admin_password,json=adminPassword,proto3" json:"admin_password,omitempty"`
}

func (x *TenantCreatedEvent) Reset() {
	*x = TenantCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantCreatedEvent) ProtoMessage() {}

func (x *TenantCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_saas_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantCreatedEvent.ProtoReflect.Descriptor instead.
func (*TenantCreatedEvent) Descriptor() ([]byte, []int) {
	return file_saas_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *TenantCreatedEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantCreatedEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TenantCreatedEvent) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *TenantCreatedEvent) GetSeparateDb() bool {
	if x != nil {
		return x.SeparateDb
	}
	return false
}

func (x *TenantCreatedEvent) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

func (x *TenantCreatedEvent) GetAdminUsername() string {
	if x != nil {
		return x.AdminUsername
	}
	return ""
}

func (x *TenantCreatedEvent) GetAdminPassword() string {
	if x != nil {
		return x.AdminPassword
	}
	return ""
}

type TenantReadyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *TenantReadyEvent) Reset() {
	*x = TenantReadyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantReadyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantReadyEvent) ProtoMessage() {}

func (x *TenantReadyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_saas_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantReadyEvent.ProtoReflect.Descriptor instead.
func (*TenantReadyEvent) Descriptor() ([]byte, []int) {
	return file_saas_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *TenantReadyEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantReadyEvent) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

var File_saas_event_v1_event_proto protoreflect.FileDescriptor

var file_saas_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73,
	0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x44, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x45, 0x0a, 0x10, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x78, 0x69, 0x61, 0x6f, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2d, 0x6b, 0x69,
	0x74, 0x2f, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saas_event_v1_event_proto_rawDescOnce sync.Once
	file_saas_event_v1_event_proto_rawDescData = file_saas_event_v1_event_proto_rawDesc
)

func file_saas_event_v1_event_proto_rawDescGZIP() []byte {
	file_saas_event_v1_event_proto_rawDescOnce.Do(func() {
		file_saas_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_saas_event_v1_event_proto_rawDescData)
	})
	return file_saas_event_v1_event_proto_rawDescData
}

var file_saas_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_saas_event_v1_event_proto_goTypes = []interface{}{
	(*TenantCreatedEvent)(nil), // 0: saas.event.v1.TenantCreatedEvent
	(*TenantReadyEvent)(nil),   // 1: saas.event.v1.TenantReadyEvent
}
var file_saas_event_v1_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_saas_event_v1_event_proto_init() }
func file_saas_event_v1_event_proto_init() {
	if File_saas_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saas_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantCreatedEvent); i {
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
		file_saas_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantReadyEvent); i {
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
			RawDescriptor: file_saas_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saas_event_v1_event_proto_goTypes,
		DependencyIndexes: file_saas_event_v1_event_proto_depIdxs,
		MessageInfos:      file_saas_event_v1_event_proto_msgTypes,
	}.Build()
	File_saas_event_v1_event_proto = out.File
	file_saas_event_v1_event_proto_rawDesc = nil
	file_saas_event_v1_event_proto_goTypes = nil
	file_saas_event_v1_event_proto_depIdxs = nil
}
