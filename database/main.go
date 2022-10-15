package database

import "gorm.io/gorm"

var Handler *BaseHandler

type BaseHandler struct {
	Db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		Db: db,
	}
}
