package model

import (
	"time"
)

type User struct {
	UserID    string    `json:"-" db:"user_id, omitempty"`
	Email     string    `json:"email,omitempty" db:"email, omitempty"`
	Phone     string    `json:"phone,omitempty" db:"phone, omitempty"`
	Password  string    `json:"-" db:"password, omitempty"`
	Address   string    `json:"address,omitempty" db:"address, omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name, omitempty"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar, omitempty"`
	Role      string    `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at, omitempty"`
	UpdateAt  time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"token,omitempty"`
}
