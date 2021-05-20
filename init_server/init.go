package init_server

import (
	"opencodes/api/server/cron"
	"opencodes/api/server/worker"
	"opencodes/pkg/config"
	"opencodes/pkg/database"
	"opencodes/pkg/logger"
	"opencodes/pkg/redis"
	"opencodes/pkg/web_server"
	"opencodes/pkg/ws"
)

func StartAllSever() {
	LogServer()
	DbSever()
	RedisServer()
	worker.WorkSetUp()
	// 判断是否启动定时器任务
	if config.ApplicationConfig.IsSync == 1 {
		CronService()
	}

	WebServer()

}
func DbSever() {
	database.MysqlSetup()
}

//基础服务
func BaseServer() {

}

//日志服务
func LogServer() {
	logger.Setup(config.LogConfig.Dir)
}

func WebServer() {
	web_server.InitWebServer()
}

// redis 服务
func RedisServer() {
	redis.InitRedis()
}

func WebsocketServer() {
	ws.InitWebsocket()
}
func CronService() {
	go cron.InitCron()
}
