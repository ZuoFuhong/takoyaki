package mysql

import (
	"log"
	"takoyaki/bolt"
	"takoyaki/defs"
)

// Delete 更新记录
func Delete(form *defs.DeleteForm) error {
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
	empty := make(map[string]interface{})
	if err := db.Debug().Table(tableName).Where(conditions).Delete(empty).Limit(1).Error; err != nil {
		log.Printf("call table delete failed, err: %v", err)
	}
	return err
}
