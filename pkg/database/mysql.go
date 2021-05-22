package database

import (
	"bytes"
	"opencodes/pkg/config"
	"strconv"

	"gorm.io/driver/mysql" //加载mysql
	"gorm.io/gorm"
)

var Eloquent *gorm.DB

func MysqlSetup() *gorm.DB {
	databaseConfig := config.DatabaseConfig
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: GetMysqlConnect(databaseConfig),
	}), &gorm.Config{})
	if err != nil {
		panic("服务器错误")
	}
	Eloquent = db

	sqlDb, _ := db.DB()
	//设置空闲连接池中连接的最大数
	sqlDb.SetConnMaxLifetime(500)
	//设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(5000)

	defer sqlDb.Close()
	if Eloquent.Error != nil {
		panic(Eloquent.Error)
	}
	return db

}

func GetMysqlConnect(conf *config.Database) string {
	var conn bytes.Buffer
	conn.WriteString(conf.Username)
	conn.WriteString(":")
	conn.WriteString(conf.Password)
	conn.WriteString("@tcp(")
	conn.WriteString(conf.Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(conf.Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(conf.Name)
	conn.WriteString("?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}
type Mysql struct {
}
