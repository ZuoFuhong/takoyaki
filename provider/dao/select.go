package dao

import (
	"log"
	"takoyaki/conf"
	"takoyaki/pkg/mysql"
)

// Select 查询数据
func Select(conditions map[string]interface{}, offset, limit int) ([]map[string]interface{}, error) {
	tableName := conf.GetConfig().TableName
	db, err := mysql.GetDb()
	if err != nil {
		return nil, err
	}

	records := make([]map[string]interface{}, 0)
	if err := db.Debug().Table(tableName).Where(conditions).Offset(offset).Limit(limit).Find(&records).Error; err != nil {
		log.Printf("call db.Find failed, err: %v", err)
		return nil, err
	}
	return records, nil
}
