package service

import (
	"opencodes/model/db"
	config2 "opencodes/pkg/config"
	"opencodes/pkg/database"
	"opencodes/pkg/logger"
	"testing"
	"time"
)

func init() {
	config2.Setup("../../config/settings.dev.yml")
	logger.Setup()
	database.MysqlSetup()
}

func TestCreateSysOperationRecord(t *testing.T) {
	settings := db.SysSetting{
		Group:     "sys",
		Key:       "name",
		Content:   "www.opencodes.top",
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}

	err, affected := CreateSettings(settings)

	t.Log(err)
	t.Log(affected)
}
