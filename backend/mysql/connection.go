package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"takoyaki/bolt"
	"takoyaki/defs"
)

var connPool sync.Map

// getConn if the connetion does not exist, create new connection.
func getConn(source string) (*gorm.DB, error) {
	if value, ok := connPool.Load(source); ok {
		return value.(*gorm.DB), nil
	}
	// query data source config
	dbConf, err := bolt.GetDataSource(source)
	if err != nil {
		return nil, err
	}
	conn, err := newDbConn(dbConf)
	if err != nil {
		return nil, err
	}
	connPool.Store(source, conn)
	return conn, nil
}

// newDbConn create new db connection
func newDbConn(dbconf *defs.DataSourceForm) (*gorm.DB, error) {
	dsn := getDsn(dbconf)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("call gorm.Open failed, err:%v", err)
	}
	return db, err
}

func getDsn(dbconf *defs.DataSourceForm) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.Username, dbconf.Password, dbconf.Host, dbconf.Database)
}
