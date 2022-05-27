package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName string `gorm:"type:varchar(800)"`
	Price     uint64 `gorm:"default:0"`
}

func (g *Goods) TableName() string {
	return "good"
}
