package config

import "github.com/spf13/viper"

type Log struct {
	Dir          string
	Level        string
	Format       string
	LogInConsole bool
}

func InitLog(cfg *viper.Viper) *Log {
	return &Log{
		Dir:          cfg.GetString("dir"),
		Level:        cfg.GetString("level"),
		Format:       cfg.GetString("format"),
		LogInConsole: cfg.GetBool("log-in-console"),
	}
}

var LogConfig = new(Log)
