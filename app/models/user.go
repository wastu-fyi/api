package models

import (
	"encoding/json"
	"time"

	"github.com/goravel/framework/database/orm"
)

type UserRole uint8

const (
	UserRoleDeveloper UserRole = 0
	UserRoleSuper     UserRole = 1
	UserRoleMember    UserRole = 2
	UserRoleCommunity UserRole = 3
)

func (r UserRole) Label() string {
	switch r {
	case UserRoleDeveloper:
		return "Pengembang"
	case UserRoleSuper:
		return "Super Admin"
	case UserRoleMember:
		return "Member"
	case UserRoleCommunity:
		return "Komunitas"
	default:
		return "Unknown"
	}
}

type User struct {
	orm.Model
	StudentID                      *uint           `gorm:"column:student_id" json:"student_id,omitempty"`
	Username                       *string         `gorm:"column:username" json:"username,omitempty"`
	Email                          string          `gorm:"column:email" json:"email"`
	Password                       *string         `gorm:"column:password" json:"-"`
	Name                           string          `gorm:"column:name" json:"name"`
	Avatar                         *string         `gorm:"column:avatar" json:"avatar,omitempty"`
	Role                           UserRole        `gorm:"column:role" json:"role"`
	Settings                       json.RawMessage `gorm:"column:settings;type:json;default:'{}'" json:"settings"`
	IsActive                       bool            `gorm:"column:is_active" json:"is_active"`
	RememberToken                  *string         `gorm:"column:remember_token" json:"-"`
	OauthProvider                  *string         `gorm:"column:oauth_provider" json:"oauth_provider,omitempty"`
	OauthID                        *string         `gorm:"column:oauth_id" json:"oauth_id,omitempty"`
	AppAuthenticationRecoveryCodes *string         `gorm:"column:app_authentication_recovery_codes" json:"-"`
	AppAuthenticationSecret        *string         `gorm:"column:app_authentication_secret" json:"-"`
	EmailVerifiedAt                *time.Time      `gorm:"column:email_verified_at" json:"email_verified_at,omitempty"`

	Student *Student `gorm:"foreignKey:StudentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"student,omitempty"`
}

func (User) TableName() string {
	return "users"
}
