package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lzientek/octopush-middleware/config"
	"github.com/lzientek/octopush-middleware/models"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	c := config.GetConfig()

	db, err := gorm.Open("postgres", "host="+c.GetString("db.host")+"user="+c.GetString("db.user")+" dbname="+c.GetString("db.dbname")+" sslmode=disable password="+c.GetString("db.password"))
	if err != nil {
		panic("failed to connect database : " + err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.SendTemplate{}, &models.SmsTemplate{})

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
