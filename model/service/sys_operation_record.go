package service

import (
	"opencodes/model/db"
	orm "opencodes/pkg/database"
)

//创建操作记录
func CreateSysOperationRecord(record db.SysOperationRecord) error {
	err := orm.Eloquent.Table("sys_operation_record").Create(&record).Error
	return err
}
