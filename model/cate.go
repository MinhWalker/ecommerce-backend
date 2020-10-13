package model

import "time"

type Cate struct {
	CateID    string    `json:"-" db:"cate_id, omitempty"`
	CateName    string    `json:"cateName,omitempty" db:"cate_name, omitempty"`
	CateImage    string    `json:"cateImage,omitempty" db:"cate_image, omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at, omitempty"`
	UpdatedAt  time.Time `json:"-" db:"updated_at, omitempty"`
}
