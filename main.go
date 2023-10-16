package main

import (
	"flag"
	"fmt"
	"time"
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "notes.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	message := flag.String("add", "", "New todo to add")
	flag.Parse()
	
	if flag.Lookup("add") != nil {
		addNewItem(*message, db)
	}
}

func addNewItem (itemName string, db *sql.DB){
	insertStmt, err := db.Prepare("INSERT INTO notes(note, dateAdded, completed) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	
	defer insertStmt.Close()
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = insertStmt.Exec(itemName, now, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Note added succesfully")
}
