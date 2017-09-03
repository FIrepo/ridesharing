package model

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var Secretkey = "indonesiatanahairbeta"
var Db *sql.DB

func DbConnect() {
	// Mysql Db Connection
	var err error
	Db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ridesharing")
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func IsLoginCredValid(passwordHash []byte, passwordString []byte) bool {
	err := bcrypt.CompareHashAndPassword(passwordHash, passwordString)
	if err == nil {
		return true
	} else {
		return false
	}
}
