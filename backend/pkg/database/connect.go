package database

import (
	"fmt"
	"log"
	"time"

	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	var count int = cfg.DB.Retry_attempts

	for count >= 1 {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(time.Second * 2)
			if count == 1 {
				panic(err)
			}
			count--
			fmt.Printf("retry... count:%v\n", count)
			continue
		}
		DB = db
		return
	}

	log.Fatal("Cannot connect to database")
}
