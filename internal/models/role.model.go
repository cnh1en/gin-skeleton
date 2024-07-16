package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model
	ID       int64  `gorm:"primary_key;auto_increment"`
	RoleName string `gorm:"type:varchar(255)"`
	RoleNote string `gorm:"type:text"`
}

func (u *Role) TableName() string {
	return "roles"
}
