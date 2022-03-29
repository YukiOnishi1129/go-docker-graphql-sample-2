package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	err error
)

func Init() *gorm.DB {
	if err = godotenv.Load(); err != nil {
		fmt.Println(err)
		return nil
	}
	// MySQLへの接続情報を定義
	dsn := os.Getenv(("MYSQL_USER")) + ":" + os.Getenv(("MYSQL_PASSWORD")) + "@tcp(" + os.Getenv(("MYSQL_HOST")) + ":" + os.Getenv(("MYSQL_PORT")) + ")/" + os.Getenv(("MYSQL_DATABASE")) + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, dbErr := db.DB()
	if dbErr != nil {
		panic(dbErr)
	}
	if err = sqlDB.Close(); err != nil {
		panic(err)
	}
}
