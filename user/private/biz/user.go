package biz

import (
	"context"
	"github.com/goxiaoy/go-saas-kit/pkg/data"
	"github.com/goxiaoy/go-saas-kit/pkg/gorm"
	v1 "github.com/goxiaoy/go-saas-kit/user/api/user/v1"
	gorm3 "github.com/goxiaoy/go-saas/gorm"
	concurrency "github.com/goxiaoy/gorm-concurrency"
	gorm2 "gorm.io/gorm"
	"time"
)

type User struct {
	gorm.UIDBase        `json:",squash"`
	concurrency.Version `gorm:"type:char(36)"`
	gorm.AuditedModel
	gorm3.MultiTenancy

	DeletedAt gorm2.DeletedAt `gorm:"index"`

	Name      *string `json:"name"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`

	Username *string `json:"username" gorm:"index"`
	// NormalizedUsername uppercase normalized userName
	NormalizedUsername *string `json:"normalized_username" gorm:"index"`

	// Phone
	Phone          *string `json:"phone" gorm:"index"`
	PhoneConfirmed bool    `json:"phone_confirmed"`

	// Email
	Email *string `json:"email" gorm:"index"`
	// NormalizedEmail uppercase normalized email
	NormalizedEmail *string `json:"normalized_email" gorm:"index"`
	EmailConfirmed  bool    `json:"email_confirmed"`

	// Password hashed
	Password *string `json:"password"`

	//Security

	AccessFailedCount int        `json:"accessFailedCount"`
	LastLoginAttempt  *time.Time `json:"lastLoginAttempt"`
	LockoutEndDateUtc *time.Time `json:"lockoutEndDateUtc"`

	// Recover
	RecoverSelector    string
	RecoverVerifier    string
	RecoverTokenExpiry *time.Time

	//2FA
	TwoFactorEnabled bool `json:"two_factor_enabled"`

	Roles []Role `gorm:"many2many:user_roles"`

	Location *string `json:"location"`
	Tags     *string `json:"tags"`

	// Avatar could be an id of asset or simple url
	Avatar   *string    `json:"avatar"`
	Birthday *time.Time `json:"birthday"`
	Gender   *string    `json:"gender"`

	Tenants []UserTenant `json:"tenants"`

	Extra data.JSONMap
}

type UserRepo interface {
	List(ctx context.Context, query *v1.ListUsersRequest) ([]*User, error)
	Count(ctx context.Context, query *v1.UserFilter) (total int64, filtered int64, err error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByName(ctx context.Context, name string) (*User, error)
	FindByPhone(ctx context.Context, phone string) (*User, error)
	FindByRecoverSelector(ctx context.Context, r string) (*User, error)
	FindByConfirmSelector(ctx context.Context, c string) (*User, error)
	AddLogin(ctx context.Context, user *User, userLogin *UserLogin) error
	RemoveLogin(ctx context.Context, user *User, loginProvider string, providerKey string) error
	ListLogin(ctx context.Context, user *User) ([]*UserLogin, error)
	FindByLogin(ctx context.Context, loginProvider string, providerKey string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	SetToken(ctx context.Context, user *User, loginProvider string, name string, value string) error
	RemoveToken(ctx context.Context, user *User, loginProvider string, name string) error
	GetToken(ctx context.Context, user *User, loginProvider string, name string) (*string, error)
	GetRoles(ctx context.Context, user *User) ([]*Role, error)
	UpdateRoles(ctx context.Context, user *User, roles []*Role) error
	AddToRole(ctx context.Context, user *User, role *Role) error
	RemoveFromRole(ctx context.Context, user *User, role *Role) error
}