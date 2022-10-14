package database

import "gorm.io/gorm"

var Handler *BaseHandler

type BaseHandler struct {
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}
