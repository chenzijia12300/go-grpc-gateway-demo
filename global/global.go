package global

import (
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Configs map[string]interface{}
)
