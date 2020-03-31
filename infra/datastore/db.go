package datastore

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func GormConn() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	if !db.HasTable(&domain.User{}) {
		db.CreateTable(&domain.User{})
	}
	if !db.HasTable(&domain.Artist{}) {
		db.CreateTable(&domain.Artist{})
	}
	if !db.HasTable(&domain.UserArtistTag{}) {
		db.CreateTable(&domain.UserArtistTag{})
	}
	if !db.HasTable(&domain.UserToken{}) {
		db.CreateTable(&domain.UserToken{})
	}
	return db, nil
}
