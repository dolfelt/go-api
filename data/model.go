package data

import "time"

// BaseModel is the standard columns
type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
