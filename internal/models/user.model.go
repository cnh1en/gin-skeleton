package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	UUID     uuid.UUID `gorm:"type:varchar(255);index:idx_uuid;unique"`
	Username string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "users"
}
