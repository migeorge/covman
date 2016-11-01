package main

import "github.com/jinzhu/gorm"

// Handlers provides a common interface to attach route handlers
type Handlers struct {
	DB *gorm.DB
}

// All common route handlers that do not relate to a specific type
// should live here
