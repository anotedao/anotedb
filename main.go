package main

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db = initDb()

	var miners []*Miner
	db.Find(&miners)

	counter := 0

	for range miners {
		counter++
	}

	log.Println(counter)
}
