package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "notes.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createNewTable(db)

	message := flag.String("add", "", "New todo to add")
	listAll := flag.Int("list", 0, "choose items to list, 0 - inComplete, 1 - complete")
	flag.Parse()

	if flag.Lookup("add") != nil && *message != "" {
		addNewItem(*message, db)
	} else if flag.Parsed() {
		listItems(db, *listAll)
	} else {
		flag.Usage()
	}
}

// Create a new notes table if one doesn't exist
//TODO maybe return bool if table created 
func createNewTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, note TEXT, dateAdded TEXT, completed INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
}

func listItems(db *sql.DB, listAll int) {
	rows, err := db.Query("SELECT note FROM notes WHERE completed=?", listAll)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var note string
		rows.Scan(&note)
		fmt.Println(note)
	}
}

func addNewItem(itemName string, db *sql.DB) {
	insertStmt, err := db.Prepare("INSERT INTO notes(note, dateAdded, completed) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer insertStmt.Close()
	//get the current time and format it
	now := time.Now().UTC().Format("2006-01-02 15:04:05") 

	_, err = insertStmt.Exec(itemName, now, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Note added succesfully")
}
