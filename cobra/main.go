package main

import (
	"log"

	"fuk/fukkk"
	"fuk/serv"
)

func main() {
	db, err := fukkk.Dbinit()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = fukkk.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	serv.Server(db)
}
