package mysql

import (
	"fmt"
	"log"
	"sync"
	"takoyaki/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

// GetDb 获取 Db 连接
func GetDb() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := getDsn()
		if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			log.Printf("call gorm.Open failed, err:%v", err)
		}
	})
	return db, err
}

func getDsn() string {
	dbconf := conf.GetConfig().Db
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.User, dbconf.Pass, dbconf.Addr, dbconf.Dbname)
}
