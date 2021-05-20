package config

import "github.com/spf13/viper"

//数据库配置
type Database struct {
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

// 初始化数据库配置
func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		DbType:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
		Name:     cfg.GetString("name"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}

var DatabaseConfig = new(Database)

var IpfsUnionConfig = new(Database)
