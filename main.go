package main

import (
	"log"

	"gorm.io/gorm"
)

var conf *Config

var db *gorm.DB

var dbpg *gorm.DB

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	db = initDb()

	dbpg = initDbPg()

	var miners []*Miner
	db.Find(&miners)

	// counter := 0

	// for _, m := range miners {
	// 	dbpg.Save(m)
	// }

	log.Println("done")
}
