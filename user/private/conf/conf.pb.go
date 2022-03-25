// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: private/conf/conf.proto

package conf

import (
	blob "github.com/goxiaoy/go-saas-kit/pkg/blob"
	conf "github.com/goxiaoy/go-saas-kit/pkg/conf"
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

	Data     *Data          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Security *conf.Security `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	Services *conf.Services `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
	User     *UserConf      `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Logging  *conf.Logging  `protobuf:"bytes,6,opt,name=logging,proto3" json:"logging,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_private_conf_conf_proto_msgTypes[0]
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
	return file_private_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetData() *Data {
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints *conf.Endpoints             `protobuf:"bytes,1,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	Blobs     map[string]*blob.BlobConfig `protobuf:"bytes,2,rep,name=blobs,proto3" json:"blobs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_private_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_private_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetEndpoints() *conf.Endpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Data) GetBlobs() map[string]*blob.BlobConfig {
	if x != nil {
		return x.Blobs
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
		mi := &file_private_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_private_conf_conf_proto_msgTypes[2]
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
	return file_private_conf_conf_proto_rawDescGZIP(), []int{2}
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
	PasswordScoreMin   int32                `protobuf:"varint,1,opt,name=password_score_min,json=passwordScoreMin,proto3" json:"password_score_min,omitempty"`
	Admin              *Admin               `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	Auth               *Auth                `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	RootUrl            string               `protobuf:"bytes,4,opt,name=root_url,json=rootUrl,proto3" json:"root_url,omitempty"`
	EmailConfirmExpiry *durationpb.Duration `protobuf:"bytes,5,opt,name=email_confirm_expiry,json=emailConfirmExpiry,proto3" json:"email_confirm_expiry,omitempty"`
	PhoneConfirmExpiry *durationpb.Duration `protobuf:"bytes,6,opt,name=phone_confirm_expiry,json=phoneConfirmExpiry,proto3" json:"phone_confirm_expiry,omitempty"`
	EmailRecoverExpiry *durationpb.Duration `protobuf:"bytes,7,opt,name=email_recover_expiry,json=emailRecoverExpiry,proto3" json:"email_recover_expiry,omitempty"`
	PhoneRecoverExpiry *durationpb.Duration `protobuf:"bytes,8,opt,name=phone_recover_expiry,json=phoneRecoverExpiry,proto3" json:"phone_recover_expiry,omitempty"`
}

func (x *UserConf) Reset() {
	*x = UserConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConf) ProtoMessage() {}

func (x *UserConf) ProtoReflect() protoreflect.Message {
	mi := &file_private_conf_conf_proto_msgTypes[3]
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
	return file_private_conf_conf_proto_rawDescGZIP(), []int{3}
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

func (x *UserConf) GetAuth() *Auth {
	if x != nil {
		return x.Auth
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

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie         *conf.Cookie            `protobuf:"bytes,4,opt,name=cookie,proto3" json:"cookie,omitempty"`
	SessionCookie  *conf.Cookie            `protobuf:"bytes,6,opt,name=session_cookie,json=sessionCookie,proto3" json:"session_cookie,omitempty"`
	Totp_2FaIssuer *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=totp_2fa_issuer,json=totp2faIssuer,proto3" json:"totp_2fa_issuer,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_private_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_private_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Auth) GetCookie() *conf.Cookie {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *Auth) GetSessionCookie() *conf.Cookie {
	if x != nil {
		return x.SessionCookie
	}
	return nil
}

func (x *Auth) GetTotp_2FaIssuer() *wrapperspb.StringValue {
	if x != nil {
		return x.Totp_2FaIssuer
	}
	return nil
}

var File_private_conf_conf_proto protoreflect.FileDescriptor

var file_private_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x62, 0x6c, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x22, 0xb4, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2d, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x31,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x42, 0x6c, 0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x62,
	0x73, 0x1a, 0x4a, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a,
	0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xd6,
	0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2c, 0x0a, 0x12, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x4b, 0x0a, 0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x12, 0x4b, 0x0a, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4b, 0x0a,
	0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x14, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0xa7, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x06,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x0d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x74,
	0x6f, 0x74, 0x70, 0x5f, 0x32, 0x66, 0x61, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x70, 0x32, 0x66, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x78, 0x69, 0x61, 0x6f, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_private_conf_conf_proto_rawDescOnce sync.Once
	file_private_conf_conf_proto_rawDescData = file_private_conf_conf_proto_rawDesc
)

func file_private_conf_conf_proto_rawDescGZIP() []byte {
	file_private_conf_conf_proto_rawDescOnce.Do(func() {
		file_private_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_conf_conf_proto_rawDescData)
	})
	return file_private_conf_conf_proto_rawDescData
}

var file_private_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_private_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),              // 0: kratos.api.Bootstrap
	(*Data)(nil),                   // 1: kratos.api.Data
	(*Admin)(nil),                  // 2: kratos.api.Admin
	(*UserConf)(nil),               // 3: kratos.api.UserConf
	(*Auth)(nil),                   // 4: kratos.api.Auth
	nil,                            // 5: kratos.api.Data.BlobsEntry
	(*conf.Security)(nil),          // 6: conf.Security
	(*conf.Services)(nil),          // 7: conf.Services
	(*conf.Logging)(nil),           // 8: conf.Logging
	(*conf.Endpoints)(nil),         // 9: conf.Endpoints
	(*durationpb.Duration)(nil),    // 10: google.protobuf.Duration
	(*conf.Cookie)(nil),            // 11: conf.Cookie
	(*wrapperspb.StringValue)(nil), // 12: google.protobuf.StringValue
	(*blob.BlobConfig)(nil),        // 13: blob.BlobConfig
}
var file_private_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	6,  // 1: kratos.api.Bootstrap.security:type_name -> conf.Security
	7,  // 2: kratos.api.Bootstrap.services:type_name -> conf.Services
	3,  // 3: kratos.api.Bootstrap.user:type_name -> kratos.api.UserConf
	8,  // 4: kratos.api.Bootstrap.logging:type_name -> conf.Logging
	9,  // 5: kratos.api.Data.endpoints:type_name -> conf.Endpoints
	5,  // 6: kratos.api.Data.blobs:type_name -> kratos.api.Data.BlobsEntry
	2,  // 7: kratos.api.UserConf.admin:type_name -> kratos.api.Admin
	4,  // 8: kratos.api.UserConf.auth:type_name -> kratos.api.Auth
	10, // 9: kratos.api.UserConf.email_confirm_expiry:type_name -> google.protobuf.Duration
	10, // 10: kratos.api.UserConf.phone_confirm_expiry:type_name -> google.protobuf.Duration
	10, // 11: kratos.api.UserConf.email_recover_expiry:type_name -> google.protobuf.Duration
	10, // 12: kratos.api.UserConf.phone_recover_expiry:type_name -> google.protobuf.Duration
	11, // 13: kratos.api.Auth.cookie:type_name -> conf.Cookie
	11, // 14: kratos.api.Auth.session_cookie:type_name -> conf.Cookie
	12, // 15: kratos.api.Auth.totp_2fa_issuer:type_name -> google.protobuf.StringValue
	13, // 16: kratos.api.Data.BlobsEntry.value:type_name -> blob.BlobConfig
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_private_conf_conf_proto_init() }
func file_private_conf_conf_proto_init() {
	if File_private_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_private_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_private_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_private_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_private_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
			RawDescriptor: file_private_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_conf_conf_proto_goTypes,
		DependencyIndexes: file_private_conf_conf_proto_depIdxs,
		MessageInfos:      file_private_conf_conf_proto_msgTypes,
	}.Build()
	File_private_conf_conf_proto = out.File
	file_private_conf_conf_proto_rawDesc = nil
	file_private_conf_conf_proto_goTypes = nil
	file_private_conf_conf_proto_depIdxs = nil
}
