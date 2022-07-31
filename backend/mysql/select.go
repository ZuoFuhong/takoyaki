package mysql

import (
	"log"
	"takoyaki/bolt"
	"takoyaki/defs"
)

// Select 查询数据
func Select(form *defs.SelectForm) ([]map[string]interface{}, int64, error) {
	page, err := bolt.GetPage(form.PageName)
	if err != nil {
		return nil, 0, err
	}
	tableName := page.TableName
	db, err := getConn(page.DataSource)
	if err != nil {
		return nil, 0, err
	}
	// filter select condtions
	conditions := form.Conditions
	for k, v := range conditions {
		if v == "" {
			delete(conditions, k)
		}
	}
	// query record from db
	records := make([]map[string]interface{}, 0)
	if err := db.Table(tableName).Where(conditions).Offset((form.Page - 1) * form.PageSize).Limit(form.PageSize).Find(&records).Error; err != nil {
		log.Printf("call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	var total int64
	if err := db.Table(tableName).Where(conditions).Count(&total).Error; err != nil {
		log.Printf("call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return records, total, nil
}
