package database

import "gorm.io/gorm"

var tb *gorm.DB

func GetDb() *gorm.DB {
	if tb == nil {
		tb = MysqlSetup()
	}
	return tb
}
