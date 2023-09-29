package main

import (
	"fmt"
	"time"
	"flag"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	records = flag.Int("records", 100, "Number of records that will be added to database")
)

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", "appuser:app123@/app")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	fmt.Println("Connected to database")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	for i:=0; i < *records; i++ {
		insert, err := db.Query("INSERT INTO jobs(name) VALUES(UUID())")
	    if err !=nil {
	        panic(err.Error())
	    }
	    insert.Close()
	    fmt.Printf("Added record %d\n", i+1)
	}
}
