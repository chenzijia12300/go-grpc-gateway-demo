package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc-demo/global"
	model "grpc-demo/model/system"
	"log"
)

func InitDB(initTable bool) {
	var databaseConfig = global.Configs["database"].(map[string]interface{})
	dns := fillDns(databaseConfig)
	log.Printf("dns: %s", dns)
	if db, err := gorm.Open(mysql.Open(dns)); err != nil {
		panic("Failed to open database")
	} else {
		global.DB = db
	}
	if initTable {
		registerTable(global.DB)
	}
}

func registerTable(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Goods{},
	)
	if err != nil {
		panic(err)
	}
	log.Printf("register table success")
}

func fillDns(mysqlConfig map[string]interface{}) (dns string) {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
	return dns
}
