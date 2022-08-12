// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: user/private/conf/conf.proto

package conf

import (
	_ "github.com/go-saas/kit/pkg/authz/authz"
	_ "github.com/go-saas/kit/pkg/blob"
	conf "github.com/go-saas/kit/pkg/conf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *conf.Data      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Security *conf.Security  `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	Services *conf.Services  `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
	User     *UserConf       `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Logging  *conf.Logging   `protobuf:"bytes,6,opt,name=logging,proto3" json:"logging,omitempty"`
	Tracing  *conf.Tracers   `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	App      *conf.AppConfig `protobuf:"bytes,8,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_user_private_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetData() *conf.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetSecurity() *conf.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Bootstrap) GetServices() *conf.Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Bootstrap) GetUser() *UserConf {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Bootstrap) GetLogging() *conf.Logging {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Bootstrap) GetTracing() *conf.Tracers {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Bootstrap) GetApp() *conf.AppConfig {
	if x != nil {
		return x.App
	}
	return nil
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_user_private_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Admin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Admin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum password score. [0-5]
	PasswordScoreMin   int32                   `protobuf:"varint,1,opt,name=password_score_min,json=passwordScoreMin,proto3" json:"password_score_min,omitempty"`
	Admin              *Admin                  `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	RootUrl            string                  `protobuf:"bytes,4,opt,name=root_url,json=rootUrl,proto3" json:"root_url,omitempty"`
	EmailConfirmExpiry *durationpb.Duration    `protobuf:"bytes,5,opt,name=email_confirm_expiry,json=emailConfirmExpiry,proto3" json:"email_confirm_expiry,omitempty"`
	PhoneConfirmExpiry *durationpb.Duration    `protobuf:"bytes,6,opt,name=phone_confirm_expiry,json=phoneConfirmExpiry,proto3" json:"phone_confirm_expiry,omitempty"`
	EmailRecoverExpiry *durationpb.Duration    `protobuf:"bytes,7,opt,name=email_recover_expiry,json=emailRecoverExpiry,proto3" json:"email_recover_expiry,omitempty"`
	PhoneRecoverExpiry *durationpb.Duration    `protobuf:"bytes,8,opt,name=phone_recover_expiry,json=phoneRecoverExpiry,proto3" json:"phone_recover_expiry,omitempty"`
	Totp_2FaIssuer     *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=totp_2fa_issuer,json=totp2faIssuer,proto3,oneof" json:"totp_2fa_issuer,omitempty"`
}

func (x *UserConf) Reset() {
	*x = UserConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConf) ProtoMessage() {}

func (x *UserConf) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConf.ProtoReflect.Descriptor instead.
func (*UserConf) Descriptor() ([]byte, []int) {
	return file_user_private_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *UserConf) GetPasswordScoreMin() int32 {
	if x != nil {
		return x.PasswordScoreMin
	}
	return 0
}

func (x *UserConf) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

func (x *UserConf) GetRootUrl() string {
	if x != nil {
		return x.RootUrl
	}
	return ""
}

func (x *UserConf) GetEmailConfirmExpiry() *durationpb.Duration {
	if x != nil {
		return x.EmailConfirmExpiry
	}
	return nil
}

func (x *UserConf) GetPhoneConfirmExpiry() *durationpb.Duration {
	if x != nil {
		return x.PhoneConfirmExpiry
	}
	return nil
}

func (x *UserConf) GetEmailRecoverExpiry() *durationpb.Duration {
	if x != nil {
		return x.EmailRecoverExpiry
	}
	return nil
}

func (x *UserConf) GetPhoneRecoverExpiry() *durationpb.Duration {
	if x != nil {
		return x.PhoneRecoverExpiry
	}
	return nil
}

func (x *UserConf) GetTotp_2FaIssuer() *wrapperspb.StringValue {
	if x != nil {
		return x.Totp_2FaIssuer
	}
	return nil
}

var File_user_private_conf_conf_proto protoreflect.FileDescriptor

var file_user_private_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x64, 0x65, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x73, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x03, 0x61,
	0x70, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x3f,
	0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x92, 0x04, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x4b, 0x0a, 0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4b,
	0x0a, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x14, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x14, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x70, 0x5f, 0x32, 0x66,
	0x61, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x70, 0x32, 0x66, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x6f, 0x74, 0x70, 0x5f, 0x32, 0x66, 0x61, 0x5f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_private_conf_conf_proto_rawDescOnce sync.Once
	file_user_private_conf_conf_proto_rawDescData = file_user_private_conf_conf_proto_rawDesc
)

func file_user_private_conf_conf_proto_rawDescGZIP() []byte {
	file_user_private_conf_conf_proto_rawDescOnce.Do(func() {
		file_user_private_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_private_conf_conf_proto_rawDescData)
	})
	return file_user_private_conf_conf_proto_rawDescData
}

var file_user_private_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_private_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),              // 0: user.internal.Bootstrap
	(*Admin)(nil),                  // 1: user.internal.Admin
	(*UserConf)(nil),               // 2: user.internal.UserConf
	(*conf.Data)(nil),              // 3: conf.Data
	(*conf.Security)(nil),          // 4: conf.Security
	(*conf.Services)(nil),          // 5: conf.Services
	(*conf.Logging)(nil),           // 6: conf.Logging
	(*conf.Tracers)(nil),           // 7: conf.Tracers
	(*conf.AppConfig)(nil),         // 8: conf.AppConfig
	(*durationpb.Duration)(nil),    // 9: google.protobuf.Duration
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
}
var file_user_private_conf_conf_proto_depIdxs = []int32{
	3,  // 0: user.internal.Bootstrap.data:type_name -> conf.Data
	4,  // 1: user.internal.Bootstrap.security:type_name -> conf.Security
	5,  // 2: user.internal.Bootstrap.services:type_name -> conf.Services
	2,  // 3: user.internal.Bootstrap.user:type_name -> user.internal.UserConf
	6,  // 4: user.internal.Bootstrap.logging:type_name -> conf.Logging
	7,  // 5: user.internal.Bootstrap.tracing:type_name -> conf.Tracers
	8,  // 6: user.internal.Bootstrap.app:type_name -> conf.AppConfig
	1,  // 7: user.internal.UserConf.admin:type_name -> user.internal.Admin
	9,  // 8: user.internal.UserConf.email_confirm_expiry:type_name -> google.protobuf.Duration
	9,  // 9: user.internal.UserConf.phone_confirm_expiry:type_name -> google.protobuf.Duration
	9,  // 10: user.internal.UserConf.email_recover_expiry:type_name -> google.protobuf.Duration
	9,  // 11: user.internal.UserConf.phone_recover_expiry:type_name -> google.protobuf.Duration
	10, // 12: user.internal.UserConf.totp_2fa_issuer:type_name -> google.protobuf.StringValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_user_private_conf_conf_proto_init() }
func file_user_private_conf_conf_proto_init() {
	if File_user_private_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_private_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_user_private_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_user_private_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConf); i {
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
	file_user_private_conf_conf_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_private_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_private_conf_conf_proto_goTypes,
		DependencyIndexes: file_user_private_conf_conf_proto_depIdxs,
		MessageInfos:      file_user_private_conf_conf_proto_msgTypes,
	}.Build()
	File_user_private_conf_conf_proto = out.File
	file_user_private_conf_conf_proto_rawDesc = nil
	file_user_private_conf_conf_proto_goTypes = nil
	file_user_private_conf_conf_proto_depIdxs = nil
}
