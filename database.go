package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"bytes"
	"text/template"
)

type connInfo struct {
	Host     string
	DbName   string
	User     string
	Password string
}

// DB is the main Database interface, all model methods and DB
// initialization should be attached to it
type DB struct {
	DB *gorm.DB
}

func (db *DB) init() {
	tmpl, err := template.New("database").Parse("host={{.Host}} user={{.User}} dbname={{.DbName}} password={{.Password}} sslmode=disable")
	if err != nil {
		panic(err)
	}

	var conn bytes.Buffer
	connection := connInfo{
		Host:     "localhost",
		DbName:   "covman",
		User:     "covman",
		Password: "covman",
	}

	err = tmpl.Execute(&conn, connection)
	if err != nil {
		panic(err)
	}

	db.DB, err = gorm.Open("postgres", conn.String())
	if err != nil {
		panic(err)
	}

	db.DB.LogMode(true)
}

func (db *DB) initSchema() {
	db.DB.AutoMigrate(&Organization{})
}
