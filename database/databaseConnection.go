package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/piendop/mysql_analytic/config"
)

var (
	once sync.Once
	db   *gorm.DB
)

//get connection to db, return pointer to struct gorm db
func GetConnectionDb() *gorm.DB {
	once.Do(func() {
		connString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetInst().DbUsername, config.GetInst().DbPassword, config.GetInst().DbName)
		gormDb, err := gorm.Open("mysql", connString)
		if err != nil {
			log.Fatal("Connect to Db failed: " + err.Error())
		}
		db = gormDb
	})
	return db
}
