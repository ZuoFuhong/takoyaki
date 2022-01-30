package mysql

import (
	"fmt"
	"testing"
	"time"
)

func Test_Select(t *testing.T) {
	db, err := GetDb()
	if err != nil {
		t.Fatalf("call GetDb failed, err: %v", err)
	}
	records := make([]map[string]interface{}, 0)
	if err := db.Debug().Table("t_book2").Offset(0).Limit(1).Find(&records).Error; err != nil {
		t.Fatalf("call db.Find failed, err: %v", err)
	}
	for k, v := range records[0] {
		switch v.(type) {
		case string:
			fmt.Printf("string key: %s\n", k)
		case int32:
			fmt.Printf("int32 key: %s\n", k)
		case time.Time:
			fmt.Printf("date key: %s\n", k)
		}
	}
}

func Test_Create(t *testing.T) {
	db, err := GetDb()
	if err != nil {
		t.Fatalf("call GetDb failed, err: %v", err)
	}
	record := make(map[string]interface{})
	record["isbn"] = "978-7-5063-6869-2"
	record["name"] = "《天幕红尘》"
	record["price"] = 4800
	record["author"] = "豆豆"
	record["edition"] = "2013 年 6 月第 1 版"
	record["press"] = "作家出版社"
	record["address"] = "北京农展馆南里 10 号"
	if err := db.Debug().Table("t_book2").Create(record).Error; err != nil {
		t.Fatalf("call db.Create failed, err:%v", err)
	}
}

func Test_Update(t *testing.T) {
	db, err := GetDb()
	if err != nil {
		t.Fatalf("call GetDb failed, err: %v", err)
	}

	conditions := make(map[string]interface{})
	conditions["isbn"] = "978-7-5063-6869-1"

	record := make(map[string]interface{})
	record["isbn"] = "978-7-5063-6869-1"
	record["author"] = "豆豆"
	record["price"] = 4800

	if err := db.Debug().Table("t_book2").Where(conditions).Updates(record).Error; err != nil {
		t.Fatalf("call db.update failed, err: %v", err)
	}
}

func Test_Delete(t *testing.T) {
	db, err := GetDb()
	if err != nil {
		t.Fatalf("call GetDb failed, err: %v", err)
	}

	conditions := make(map[string]interface{})
	conditions["isbn"] = "978-7-5063-6869-2"
	conditions["price"] = "4800"

	records := make([]map[string]interface{}, 0)
	if err := db.Debug().Table("t_book2").Where(conditions).Delete(&records).Error; err != nil {
		t.Fatalf("call db.delete failed, err: %v", err)
	}
	fmt.Printf("records: %v", records)
}
