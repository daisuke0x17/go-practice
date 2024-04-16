package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Tweet struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	UserID    int `gorm:"foreignKey:UserID;references:ID"`
	User      User
	Text      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=sample_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Tweet{})
	fmt.Println("migrated")

	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		db.Create(&User{Name: "user01"})
		db.Create(&User{Name: "user02"})
		db.Create(&User{Name: "user03"})
	}

	var user User
	db.First(&user)
	db.Model(&Tweet{}).Count(&count)
	if count == 0 {
		db.Create(&Tweet{UserID: user.ID, Text: "tweet01"})
		db.Create(&Tweet{UserID: user.ID, Text: "tweet02"})
		db.Create(&Tweet{UserID: user.ID, Text: "tweet03"})
	}

	fmt.Println("===================================================")

	// Preload
	var tweet []Tweet
	db.Preload("User").Where("user_id = ?", user.ID).Find(&tweet)
	jsonPreload, err := json.MarshalIndent(tweet, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonPreload))

	fmt.Println("===================================================")

	// Left Join
	var tweetJoin Tweet
	db.Joins("User").First(&tweetJoin)
	jsonJoin, err := json.MarshalIndent(tweetJoin, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonJoin))
}
