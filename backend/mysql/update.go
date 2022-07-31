package mysql

import (
	"log"
	"takoyaki/bolt"
	"takoyaki/defs"
)

// Update 更新记录
func Update(form *defs.UpdateForm) error {
	page, err := bolt.GetPage(form.PageName)
	if err != nil {
		return err
	}
	tableName := page.TableName
	db, err := getConn(page.DataSource)
	if err != nil {
		return err
	}
	conditions := map[string]interface{}{
		page.PrimaryKey: form.Record[page.PrimaryKey],
	}
	delete(form.Record, page.PrimaryKey)
	if err := db.Table(tableName).Where(conditions).Updates(form.Record).Error; err != nil {
		log.Printf("call table update failed, err: %v", err)
	}
	return err
}
