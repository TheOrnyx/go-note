//TODO
/// Sanitize inputs lmao

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
	listAll := flag.Int("list", 3, "choose items to list, 0 - inComplete, 1 - complete")
	removeItem := flag.Bool("remove", false, "Whether to open the item remove")
	flag.Parse()

	//flag handling
	if flag.Lookup("add") != nil && *message != "" {
		addNewItem(*message, db)
	} else if *listAll != 3 {
		listItems(db, *listAll, "id,note")
	} else if *removeItem {
		removeFromTable(db)
	} else {
		flag.Usage()
	}
}

// TODO - maybe add a like thing where if person
//        puts an int then it deletes that otherwise chooser
func removeFromTable(db *sql.DB) {
	var id int
	listItems(db, 0, "id,note")
	fmt.Println("---------------------")
	fmt.Println("Enter id number")
	fmt.Printf("> ")
	fmt.Scanln(&id)

	_, err := db.Exec("UPDATE notes SET completed=1 WHERE id=?", id)
	if err !=nil {
		log.Fatal(err)
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

func listItems(db *sql.DB, listAll int, cols string) {
	query := fmt.Sprintf("SELECT %s FROM notes WHERE completed=?", cols)
	rows, err := db.Query(query, listAll)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var note string
		var id string
		rows.Scan(&id,&note)
		fmt.Println(id,note)
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
