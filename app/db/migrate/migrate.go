package main

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/db"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/entity"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Todo{})
	if err != nil {
		return
	}
}

func main() {
	dbCon := db.Init()
	//	dbを閉じる
	defer db.CloseDB(dbCon)

	migrate(dbCon)
}
