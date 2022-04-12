// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: user/api/auth/v1/web.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogoutChallenge string `protobuf:"bytes,1,opt,name=logout_challenge,json=logoutChallenge,proto3" json:"logout_challenge,omitempty"`
}

func (x *GetLogoutRequest) Reset() {
	*x = GetLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogoutRequest) ProtoMessage() {}

func (x *GetLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogoutRequest.ProtoReflect.Descriptor instead.
func (*GetLogoutRequest) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{0}
}

func (x *GetLogoutRequest) GetLogoutChallenge() string {
	if x != nil {
		return x.LogoutChallenge
	}
	return ""
}

type GetLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge string `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *GetLogoutResponse) Reset() {
	*x = GetLogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogoutResponse) ProtoMessage() {}

func (x *GetLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogoutResponse.ProtoReflect.Descriptor instead.
func (*GetLogoutResponse) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogoutResponse) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

type GetConsentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsentChallenge string `protobuf:"bytes,1,opt,name=consent_challenge,json=consentChallenge,proto3" json:"consent_challenge,omitempty"`
}

func (x *GetConsentRequest) Reset() {
	*x = GetConsentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentRequest) ProtoMessage() {}

func (x *GetConsentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentRequest.ProtoReflect.Descriptor instead.
func (*GetConsentRequest) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{2}
}

func (x *GetConsentRequest) GetConsentChallenge() string {
	if x != nil {
		return x.ConsentChallenge
	}
	return ""
}

type GetConsentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge      string       `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	RequestedScope []string     `protobuf:"bytes,2,rep,name=requested_scope,json=requestedScope,proto3" json:"requested_scope,omitempty"`
	UserId         string       `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Client         *OAuthClient `protobuf:"bytes,4,opt,name=client,proto3" json:"client,omitempty"`
	Redirect       string       `protobuf:"bytes,10,opt,name=redirect,proto3" json:"redirect,omitempty"`
}

func (x *GetConsentResponse) Reset() {
	*x = GetConsentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentResponse) ProtoMessage() {}

func (x *GetConsentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentResponse.ProtoReflect.Descriptor instead.
func (*GetConsentResponse) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{3}
}

func (x *GetConsentResponse) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *GetConsentResponse) GetRequestedScope() []string {
	if x != nil {
		return x.RequestedScope
	}
	return nil
}

func (x *GetConsentResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetConsentResponse) GetClient() *OAuthClient {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *GetConsentResponse) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

type GrantConsentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge string `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	//user granted scope
	GrantScope []string `protobuf:"bytes,2,rep,name=grant_scope,json=grantScope,proto3" json:"grant_scope,omitempty"`
	Reject     bool     `protobuf:"varint,5,opt,name=reject,proto3" json:"reject,omitempty"`
}

func (x *GrantConsentRequest) Reset() {
	*x = GrantConsentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantConsentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantConsentRequest) ProtoMessage() {}

func (x *GrantConsentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantConsentRequest.ProtoReflect.Descriptor instead.
func (*GrantConsentRequest) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{4}
}

func (x *GrantConsentRequest) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *GrantConsentRequest) GetGrantScope() []string {
	if x != nil {
		return x.GrantScope
	}
	return nil
}

func (x *GrantConsentRequest) GetReject() bool {
	if x != nil {
		return x.Reject
	}
	return false
}

type GrantConsentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Redirect string `protobuf:"bytes,10,opt,name=redirect,proto3" json:"redirect,omitempty"`
}

func (x *GrantConsentResponse) Reset() {
	*x = GrantConsentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantConsentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantConsentResponse) ProtoMessage() {}

func (x *GrantConsentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantConsentResponse.ProtoReflect.Descriptor instead.
func (*GrantConsentResponse) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{5}
}

func (x *GrantConsentResponse) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

type OAuthClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogoUrl  string           `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Metadata *structpb.Struct `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *OAuthClient) Reset() {
	*x = OAuthClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_auth_v1_web_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthClient) ProtoMessage() {}

func (x *OAuthClient) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_auth_v1_web_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthClient.ProtoReflect.Descriptor instead.
func (*OAuthClient) Descriptor() ([]byte, []int) {
	return file_user_api_auth_v1_web_proto_rawDescGZIP(), []int{6}
}

func (x *OAuthClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OAuthClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OAuthClient) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *OAuthClient) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_user_api_auth_v1_web_proto protoreflect.FileDescriptor

var file_user_api_auth_v1_web_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x22, 0x6c, 0x0a, 0x13,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x22, 0x81,
	0x01, 0x0a, 0x0b, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x33, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x32, 0xcf, 0x05, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x57, 0x65, 0x62, 0x12, 0x70,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x75, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x65, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x62, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x6e, 0x0a,
	0x09, 0x57, 0x65, 0x62, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x77, 0x65, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x7e, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x3a, 0x01, 0x2a, 0x42, 0x40, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x78, 0x69, 0x61, 0x6f, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_api_auth_v1_web_proto_rawDescOnce sync.Once
	file_user_api_auth_v1_web_proto_rawDescData = file_user_api_auth_v1_web_proto_rawDesc
)

func file_user_api_auth_v1_web_proto_rawDescGZIP() []byte {
	file_user_api_auth_v1_web_proto_rawDescOnce.Do(func() {
		file_user_api_auth_v1_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_api_auth_v1_web_proto_rawDescData)
	})
	return file_user_api_auth_v1_web_proto_rawDescData
}

var file_user_api_auth_v1_web_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_api_auth_v1_web_proto_goTypes = []interface{}{
	(*GetLogoutRequest)(nil),     // 0: user.api.auth.v1.GetLogoutRequest
	(*GetLogoutResponse)(nil),    // 1: user.api.auth.v1.GetLogoutResponse
	(*GetConsentRequest)(nil),    // 2: user.api.auth.v1.GetConsentRequest
	(*GetConsentResponse)(nil),   // 3: user.api.auth.v1.GetConsentResponse
	(*GrantConsentRequest)(nil),  // 4: user.api.auth.v1.GrantConsentRequest
	(*GrantConsentResponse)(nil), // 5: user.api.auth.v1.GrantConsentResponse
	(*OAuthClient)(nil),          // 6: user.api.auth.v1.OAuthClient
	(*structpb.Struct)(nil),      // 7: google.protobuf.Struct
	(*GetLoginRequest)(nil),      // 8: user.api.auth.v1.GetLoginRequest
	(*WebLoginAuthRequest)(nil),  // 9: user.api.auth.v1.WebLoginAuthRequest
	(*LogoutRequest)(nil),        // 10: user.api.auth.v1.LogoutRequest
	(*GetLoginResponse)(nil),     // 11: user.api.auth.v1.GetLoginResponse
	(*WebLoginAuthReply)(nil),    // 12: user.api.auth.v1.WebLoginAuthReply
	(*LogoutResponse)(nil),       // 13: user.api.auth.v1.LogoutResponse
}
var file_user_api_auth_v1_web_proto_depIdxs = []int32{
	6,  // 0: user.api.auth.v1.GetConsentResponse.client:type_name -> user.api.auth.v1.OAuthClient
	7,  // 1: user.api.auth.v1.OAuthClient.metadata:type_name -> google.protobuf.Struct
	8,  // 2: user.api.auth.v1.AuthWeb.GetWebLogin:input_type -> user.api.auth.v1.GetLoginRequest
	9,  // 3: user.api.auth.v1.AuthWeb.WebLogin:input_type -> user.api.auth.v1.WebLoginAuthRequest
	0,  // 4: user.api.auth.v1.AuthWeb.GetWebLogout:input_type -> user.api.auth.v1.GetLogoutRequest
	10, // 5: user.api.auth.v1.AuthWeb.WebLogout:input_type -> user.api.auth.v1.LogoutRequest
	2,  // 6: user.api.auth.v1.AuthWeb.GetConsent:input_type -> user.api.auth.v1.GetConsentRequest
	4,  // 7: user.api.auth.v1.AuthWeb.GrantConsent:input_type -> user.api.auth.v1.GrantConsentRequest
	11, // 8: user.api.auth.v1.AuthWeb.GetWebLogin:output_type -> user.api.auth.v1.GetLoginResponse
	12, // 9: user.api.auth.v1.AuthWeb.WebLogin:output_type -> user.api.auth.v1.WebLoginAuthReply
	1,  // 10: user.api.auth.v1.AuthWeb.GetWebLogout:output_type -> user.api.auth.v1.GetLogoutResponse
	13, // 11: user.api.auth.v1.AuthWeb.WebLogout:output_type -> user.api.auth.v1.LogoutResponse
	3,  // 12: user.api.auth.v1.AuthWeb.GetConsent:output_type -> user.api.auth.v1.GetConsentResponse
	5,  // 13: user.api.auth.v1.AuthWeb.GrantConsent:output_type -> user.api.auth.v1.GrantConsentResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_api_auth_v1_web_proto_init() }
func file_user_api_auth_v1_web_proto_init() {
	if File_user_api_auth_v1_web_proto != nil {
		return
	}
	file_user_api_auth_v1_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_api_auth_v1_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogoutRequest); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogoutResponse); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentRequest); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentResponse); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantConsentRequest); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantConsentResponse); i {
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
		file_user_api_auth_v1_web_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthClient); i {
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
			RawDescriptor: file_user_api_auth_v1_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_api_auth_v1_web_proto_goTypes,
		DependencyIndexes: file_user_api_auth_v1_web_proto_depIdxs,
		MessageInfos:      file_user_api_auth_v1_web_proto_msgTypes,
	}.Build()
	File_user_api_auth_v1_web_proto = out.File
	file_user_api_auth_v1_web_proto_rawDesc = nil
	file_user_api_auth_v1_web_proto_goTypes = nil
	file_user_api_auth_v1_web_proto_depIdxs = nil
}
