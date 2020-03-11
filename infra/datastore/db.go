package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func GormConn() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	//defer db.Close()
	return db, nil
}
