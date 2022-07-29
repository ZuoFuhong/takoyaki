package dao

import (
	"log"
	"takoyaki/conf"
	"takoyaki/pkg/mysql"
)

// Delete 更新记录
func Delete(conditions map[string]interface{}) error {
	tableName := conf.GetConfig().TableName
	db, err := mysql.GetDb()
	if err != nil {
		return err
	}
	empty := make(map[string]interface{})
	if err := db.Debug().Table(tableName).Where(conditions).Delete(empty).Limit(1).Error; err != nil {
		log.Printf("call table delete failed, err: %v", err)
		return err
	}
	return nil
}
