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

	db, err = gorm.Open("postgres", "host="+c.GetString("db.host")+" user="+c.GetString("db.user")+" dbname="+c.GetString("db.dbname")+" sslmode=disable password="+c.GetString("db.password"))
	if err != nil {
		panic("failed to connect database : " + err.Error())
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(&models.SendTemplate{}, &models.SmsTemplate{}, &models.User{})

	db.Model(&models.SmsTemplate{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.SendTemplate{}).AddForeignKey("sms_template_id", "sms_templates(id)", "RESTRICT", "RESTRICT")
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
