package data

import "github.com/jinzhu/gorm"

// DB stores all database info
type DB struct {
	*gorm.DB
}
