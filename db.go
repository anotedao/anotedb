package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDb() *gorm.DB {
	var db *gorm.DB
	var err error
	dbconf := gorm.Config{}

	dbconf.Logger = logger.Default.LogMode(logger.Error)

	db, err = gorm.Open(sqlite.Open("../../anote-mobile/mobile.db"), &dbconf)

	if err != nil {
		log.Println(err)
	}

	return db
}

func initDbPg() *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	if err := db.AutoMigrate(&Miner{}, &IpAddress{}, &KeyValue{}, &User{}, &Alpha{}); err != nil {
		panic(err.Error())
	}

	return db
}
