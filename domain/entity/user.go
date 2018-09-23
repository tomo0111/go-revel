package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type User struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	Username  string    `gorm:"type:varchar(128);not null"validate:"required"json:"username"`
	Email     string    `gorm:"type:varchar(128);not null;index:email"validate:"required,email"json:"email"`
	Password  string    `gorm:"type:varchar(128);not null"validate:"required"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m User) TableName() string {
	return "users"
}

func (m User) UserUuid() uuid.UUID {
	return m.Uuid
}