package singleton

/*
	- Database singleton for connection to db
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Username string
	Password string
	DbName   string
}

func NewDatabaseInitialize(username string, password string, dbName string) Database {
	return Database{
		Username: username,
		Password: password,
		DbName:   dbName,
	}
}

// Create connection to db
func (db *Database) GetConnection() (*gorm.DB, error) {
	address := fmt.Sprintf("%s:%s@(localhost)/%s?charset=utf8&parseTime=True&loc=Local", db.Username, db.Password, db.DbName)
	database, err := gorm.Open("mysql", address)

	if err == nil {
		fmt.Println("connected to database...")
	}

	defer database.Close()

	return database, err
}
