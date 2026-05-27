package models

import (
	"encoding/json"

	"github.com/goravel/framework/database/orm"
)

type Setting struct {
	orm.Model
	Group   string          `gorm:"column:group" json:"group"`
	Name    string          `gorm:"column:name" json:"name"`
	Locked  bool            `gorm:"column:locked" json:"locked"`
	Payload json.RawMessage `gorm:"column:payload" json:"payload"`
}

func (Setting) TableName() string {
	return "settings"
}
