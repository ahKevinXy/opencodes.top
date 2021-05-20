package config

import "github.com/spf13/viper"

type Upload struct {
	Dir string
}

func InitUpload(cfg *viper.Viper) *Upload {
	return &Upload{
		Dir: cfg.GetString("dir"),
	}
}

var UploadConfig = new(Upload)
