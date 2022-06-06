package users

import (
	"go-oauth/data/models"
	"time"
)

type User struct {
	models.Model
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	FirstName string   `gorm:"column:first_name;not null;" json:"first_name" form:"first_name"`
	LastName  string   `gorm:"column:last_name;not null;" json:"last_name" form:"last_name"`
	Password  string   `gorm:"column:password;not null;" json:"password"`
	Role      UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
