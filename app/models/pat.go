package models

import "time"

type PersonalAccessToken struct {
	ID            uint64     `gorm:"column:id;primaryKey"`
	TokenableType string     `gorm:"column:tokenable_type"`
	TokenableID   uint64     `gorm:"column:tokenable_id"`
	Name          string     `gorm:"column:name;type:text"`
	Token         string     `gorm:"column:token;size:64"`
	Abilities     *string    `gorm:"column:abilities;type:text"`
	LastUsedAt    *time.Time `gorm:"column:last_used_at"`
	ExpiresAt     *time.Time `gorm:"column:expires_at"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
}

func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens"
}
