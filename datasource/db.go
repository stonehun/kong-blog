package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"kong_blog/configs"
	"kong_blog/models"
	"strings"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	dbConfig := configs.GetConfig().Db
	path := strings.Join([]string{dbConfig.Name, ":", dbConfig.Password, "@(", dbConfig.Host,
		":", dbConfig.Port, ")/", dbConfig.Database, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	GetDB().AutoMigrate(
		&models.User{},
	)

}
