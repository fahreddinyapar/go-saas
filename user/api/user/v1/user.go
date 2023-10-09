package v1

import "google.golang.org/protobuf/proto"

func (x *CreateUserRequest) StringWithMask(mask string) string {
	ret := proto.Clone(x).(*CreateUserRequest)
	ret.Password = mask
	return ret.String()
}

func (x *AdminUpdateUserRequest) StringWithMask(mask string) string {
	ret := proto.Clone(x).(*AdminUpdateUserRequest)
	if ret.User != nil {
		ret.User.Password = mask
	}
	return ret.String()
}
