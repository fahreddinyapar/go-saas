package biz

import (
	"context"
	"github.com/goxiaoy/go-saas-kit/pkg/data"
	"time"
)

type UserTenantStatus int32

const (
	Active   UserTenantStatus = 0
	Inactive UserTenantStatus = 1
)

func (p UserTenantStatus) String() string {
	switch p {
	case Active:
		return "ACTIVE"
	case Inactive:
		return "INACTIVE"
	default:
		return "UNKNOWN"
	}
}

type UserTenant struct {
	UserId   string           `gorm:"type:char(36);primaryKey" json:"user_id"`
	TenantId string           `json:"tenant_id" gorm:"type:char(36);primaryKey"`
	JoinTime time.Time        `json:"joint_time"`
	Status   UserTenantStatus `json:"status"`
	Extra    data.JSONMap
}

type UserTenantRepo interface {
	JoinTenant(ctx context.Context, userId string, tenantId string) (*UserTenant, error)
	RemoveFromTenant(ctx context.Context, userId string, tenantId string) error
	Get(ctx context.Context, userId string, tenantId string) (*UserTenant, error)
	IsIn(ctx context.Context, userId string, tenantId string) (bool, error)
	Update(ctx context.Context, userTenant *UserTenant) error
}