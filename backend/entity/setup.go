package entity

// import (
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func DB() *gorm.DB {
// 	return db
// }

// func SetupDatabase() {

// 	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

// 	if err != nil {

// 		panic("failed to connect database")

// 	}

// 	// Migrate the schema

// 	database.AutoMigrate(
// 		&User{},
// 		&Activity{},
// 		&Student{},
// 		&Ac_his{},
// 	)

// 	db = database

// }

import (
	// "fmt"
	// "time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Video{}, &User{}, &Playlist{}, &Resolution{}, WatchVideo{},
	)

	db = database
}
