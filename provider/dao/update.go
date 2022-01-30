package dao

import (
	"log"
	"takoyaki/conf"
	"takoyaki/pkg/mysql"
)

// Update 更新记录
func Update(conditions, record map[string]interface{}) error {
	tableName := conf.GetConfig().TableName
	db, err := mysql.GetDb()
	if err != nil {
		return err
	}

	if err := db.Debug().Table(tableName).Where(conditions).Updates(record).Error; err != nil {
		log.Printf("call table update failed, err: %v", err)
		return err
	}
	return nil
}
