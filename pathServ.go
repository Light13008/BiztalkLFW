package main

// Paths Driver for BLFW
// Path: A path is a
// Description: Driver for interacting SQLite DB for saving, viewing and modifying paths.
// Usage : import functions of pathServ from external go modules
// SCHEMA :
// BLFW.DB -> TABLES:
//				-> PATHS
// 					-> COLUMNS:
// 						-> PATH_ID: UNIQUE: UINT
// 						-> PORT_NAME: STRING
// 						-> PORT_TYPE: VARCHAR(10)
// 						-> PATH_URI: STRING

// TODO:
// 1. Create SQLITE DB "blfw" if not exists, connect to SQLite database (initialise) [v]
// 2. Create CRUD ops for querying to DB and export functions

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// defining types
type Paths struct {
	gorm.Model
	Path_Id   uuid.UUID
	Port_Name string
	Port_Type string
	Path_Uri  string
}

// defining DB object

var db *gorm.DB

func initialiseDB() {
	var connCheck bool = true
	// Step 1: Check if DB Exists, if not create DB file
	_, fileCreationerror := os.Stat("blfw.db")
	if fileCreationerror != nil {
		fmt.Println("BLFW DB is not Present in the current directory, proceeding to create the DB")
		os.Create("./blfw.db")
	} else {
		if fileCreationerror == nil {
			fmt.Println("BLFW DB Found")
		} else {
			connCheck = false
			fmt.Println("An Error Occured, creating the Database, exception is below:")
			fmt.Println(fileCreationerror)
		}
	}
	if connCheck {
		// Step 3: Connect to BLFW DB
		var err error
		db, err = gorm.Open(sqlite.Open("blfw.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to Open Database Connection, exception is printed below:")
			fmt.Println(err)
		} else {
			fmt.Println("Database Connection Successful...")
		}
		// Step 4: Migrate Schema
		db.AutoMigrate(&Paths{})
	} else {
		fmt.Println("Could not connect to database...Exiting")
	}
}

func insertToPaths(port_name string, port_type string, path_uri string) {
	u, uuidErr := uuid.NewRandom()
	if uuidErr == nil {
		db.Create(&Paths{Path_Id: u, Port_Name: port_name, Port_Type: port_type, Path_Uri: path_uri})
	} else {
		fmt.Println("An error occured generating UUID")
		fmt.Println(uuidErr)
	}
}

func deletePath() {

}

func updatePath() {

}

func pathLists() {
	var paths []Paths
	db.Find(&paths)
	for _, path := range paths {
		fmt.Println(path)
	}
}

func main() {
	// use main for testing only
	// initialiseDB()
	// insertToPaths("Send To WMS-MPI From ROSS-BPO", "SEND", "https://www.sample.edu/?aunt=pail&men=bit#slave")
	// pathLists()
}
