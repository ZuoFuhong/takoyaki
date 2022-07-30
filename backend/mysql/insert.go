package mysql

import (
	"log"
	"takoyaki/bolt"
	"takoyaki/defs"
)

// Insert 新增记录
func Insert(form *defs.InsertForm) error {
	page, err := bolt.GetPage(form.PageName)
	if err != nil {
		return err
	}
	tableName := page.TableName
	db, err := getConn(page.DataSource)
	if err != nil {
		return err
	}
	if err := db.Table(tableName).Create(form.Record).Error; err != nil {
		log.Printf("call insert table failed, err:%v", err)
	}
	return err
}
