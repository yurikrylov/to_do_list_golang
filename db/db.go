package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type SQLiteRepository struct {
	db *sql.DB
}

func createDB(pathToDB string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(ProjectTabelDefinition)
	db.Exec(TaskTabelDefinition)
	return db
}

func NewSQLiteRepository() *SQLiteRepository {
	var db *sql.DB

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		db = createDB(dbName)
		fmt.Println("DB isn't exist")
	} else {
		db, err = sql.Open("sqlite3", dbName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB already exists")
	}

	return &SQLiteRepository{
		db: db,
	}
}

func (s *SQLiteRepository) Close() {
	s.db.Close()
}
