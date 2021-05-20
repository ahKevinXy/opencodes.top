package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgDatabase *viper.Viper //数据库配置

var cfgApplication *viper.Viper //应用配置

var cfgLog *viper.Viper //日志配置

var cfgRedis *viper.Viper //redis配置

var cfgUpload *viper.Viper //上传配置
var cfgJwt *viper.Viper    //jwt 配置
//载入配置文件
func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// 初始化starPool 配置
	cfgDatabase = viper.Sub("database")
	if cfgDatabase == nil {
		panic("config not found database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	//初始化应用配置
	cfgApplication = viper.Sub("application")
	if cfgApplication == nil {
		panic("config not found application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	//初始化日志配置
	cfgLog = viper.Sub("log")
	if cfgLog == nil {
		panic("config not found log")
	}
	LogConfig = InitLog(cfgLog)

	//初始化Redis配置
	cfgRedis = viper.Sub("redis")
	if cfgRedis == nil {
		panic("config not found redis")
	}
	RedisConfig = InitRedis(cfgRedis)
	//初始化上传配置
	cfgUpload = viper.Sub("upload")
	if cfgUpload == nil {
		panic("config not found redis")
	}
	UploadConfig = InitUpload(cfgUpload)

	cfgJwt = viper.Sub("jwt")
	if cfgJwt == nil {
		panic("config not found jwt")
	}
	JwtConfig = InitJwt(cfgJwt)
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	viper.WriteConfig()
}
