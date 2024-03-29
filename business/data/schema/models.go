package schema

import (
	"time"
)

type NewUser struct {
	FullName    string
	Username    string
	Email       string
	Hash        string
	Preferences string //JSON
}

type User struct {
	Id          int    `gorm:"primaryKey;column:id"`
	FullName    string `gorm:"column:full_name"`
	Username    string `gorm:"column:username"`
	Hash        string `gorm:"column:hash"`
	Email       string `gorm:"column:email"`
	Preferences string `gorm:"column:preferences"`

	Roles []Role `gorm:"many2many:auth_users_roles;ForeignKey:id;References:id"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "auth_users"
}

type Role struct {
	Id          int    `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return "auth_roles"
}
