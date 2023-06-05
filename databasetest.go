package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main2() {

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name FROM Tag")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}
}
