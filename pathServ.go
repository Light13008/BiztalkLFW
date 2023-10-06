package main

// Paths DB Driver for BLFW
// Path: A path is a network or local URI used for Send Ports / Receive Locations that uses FILE Adapter
// Description: Driver for interacting SQLite DB for saving, viewing and modifying paths.
// Usage : import functions of pathServ from external go modules
// SCHEMA :
// BLFW.DB -> TABLES:
//				-> PATHS
// 					-> COLUMNS:
// 						-> PATH_ID: UNIQUE: UUID (PRIMARY KEY)
// 						-> PORT_NAME: STRING
// 						-> PORT_TYPE: STRING
// 						-> PATH_URI: STRING

// TODO:
// 1. Create SQLITE DB "blfw" if not exists, connect to SQLite database (initialise) [v]
// 2. Create CRUD ops for querying to DB and export functions [WIP]

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
	ID        uuid.UUID `gorm:"primaryKey"`
	Port_Name string
	Port_Type string
	Path_Uri  string
}

// defining DB object

var db *gorm.DB

func initialiseDB() {
	var connCheck bool = true
	// Step 1: Check if DB Exists, if not create DB file
	_, fileCreationError := os.Stat("blfw.db")
	if fileCreationError != nil {
		fmt.Println("BLFW DB is not Present in the current directory, proceeding to create the DB")
		os.Create("./blfw.db")
	} else {
		if fileCreationError == nil {
			fmt.Println("BLFW DB Found")
		} else {
			connCheck = false
			fmt.Println("An Error Occured, creating the Database, exception is below:")
			fmt.Println(fileCreationError)
		}
	}
	if connCheck {
		// Step 2: Connect to BLFW DB
		var err error
		db, err = gorm.Open(sqlite.Open("blfw.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to Open Database Connection, exception is printed below:")
			fmt.Println(err)
		} else {
			fmt.Println("Database Connection Successful...")
		}
		// Step 3: Migrate Schema
		db.AutoMigrate(&Paths{})
	} else {
		fmt.Println("Could not connect to database...Exiting")
	}
}

func insert(port_name string, port_type string, path_uri string) {
	// Generate Random UUID and insert row to sqlite db
	u, uuidErr := uuid.NewRandom()
	if uuidErr == nil {
		db.Create(&Paths{ID: u, Port_Name: port_name, Port_Type: port_type, Path_Uri: path_uri})
	} else {
		fmt.Println("An error occured generating UUID")
		fmt.Println(uuidErr)
	}
}

func delete(path_id string) {
	// Convert path_id string to uuid
	uuidPathId, err := uuid.Parse(path_id)
	if err == nil {
		db.Delete(&Paths{}, uuidPathId)
		fmt.Println("Path has been successfully deleted")
	} else {
		fmt.Println("Error Occured Deleting the Path")
		fmt.Println(err)
	}
}

func update(path_id string, path_uri string) {
	// Only designed to update the Path URI
	uuidPathId, err := uuid.Parse(path_id)
	if err == nil {
		result := db.Save(&Paths{ID: uuidPathId, Path_Uri: path_uri})
		fmt.Println(result.RowsAffected)
	}

}

func retrieveAllPaths() {
	var paths []Paths
	db.Find(&paths)
	for _, path := range paths {
		fmt.Println(path)
	}
}

func main() {
	// use main for testing only
}
