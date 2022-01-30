package dao

import (
	"log"
	"takoyaki/conf"
	"takoyaki/pkg/mysql"
)

// Insert 新增记录
func Insert(record map[string]interface{}) error {
	tableName := conf.GetConfig().TableName
	db, err := mysql.GetDb()
	if err != nil {
		return err
	}
	if err := db.Debug().Table(tableName).Create(record).Error; err != nil {
		log.Printf("call insert table failed, err:%v", err)
		return err
	}
	return nil
}
