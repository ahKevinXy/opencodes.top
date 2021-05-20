package service

import (
	"opencodes/model/db"
	orm "opencodes/pkg/database"

	logs "github.com/cihub/seelog"
	"gorm.io/gorm"
)

func CreateSettings(settings db.SysSetting) (err error, affected int64) {
	tb := orm.Eloquent.Table("tb_sys_settings").Save(&settings)
	if tb.Error != nil && tb.Error != gorm.ErrRecordNotFound {
		logs.Errorf("%+v", tb.Error)
	}
	return tb.Error, tb.RowsAffected
}

func GetSysSettings(g, k string) {
	//tb := orm.Eloquent.Table("tb_sys_settings").Select("content,type").Where("group=?",g).Where("key=?",k).Scan(&)
}
