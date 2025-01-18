package db

import "gorm.io/gorm"

type Database interface {
	Connect() *gorm.DB
}
