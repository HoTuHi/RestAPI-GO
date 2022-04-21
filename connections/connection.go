package connections

import (
	"golangEx/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=Daucham@ dbname=goLangEx port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	DB = conn
	conn.AutoMigrate(models.User{})
}
