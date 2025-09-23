package model

import "time"

type UserModel struct {
	Model
	Name            string     `gorm:"not null" json:"name"`
	Email           string     `gorm:"not null;unique" json:"email"`
	Phone           *string    `gorm:"unique" json:"phone"`
	Password        string     `gorm:"not null" json:"-"`
	IsActive        bool       `gorm:"not null;default:1" json:"is_active"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
}

func (UserModel) TableName() string {
	return "users"
}
