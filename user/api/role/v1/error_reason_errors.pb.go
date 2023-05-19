// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	context "context"
	errors "github.com/go-kratos/kratos/v2/errors"
	i18n "github.com/go-saas/go-i18n/v2/i18n"
	localize "github.com/go-saas/kit/pkg/localize"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsRolePreserved(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ROLE_PRESERVED.String() && e.Code == 403
}

func ErrorRolePreservedLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(403, ErrorReason_ROLE_PRESERVED.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "RolePreserved",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(403, ErrorReason_ROLE_PRESERVED.String(), msg)
	} else {
		return errors.New(403, ErrorReason_ROLE_PRESERVED.String(), "")
	}
}

func IsRoleNameDuplicate(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ROLE_NAME_DUPLICATE.String() && e.Code == 400
}

func ErrorRoleNameDuplicateLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_ROLE_NAME_DUPLICATE.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "RoleNameDuplicate",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_ROLE_NAME_DUPLICATE.String(), msg)
	} else {
		return errors.New(400, ErrorReason_ROLE_NAME_DUPLICATE.String(), "")
	}
}
