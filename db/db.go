package db

import (
	"database/sql"
	"fmt"

	"github.com/02amanag/p-02/config"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var dbOpen *sql.DB
var db *gorp.DbMap

func Init() {

	conf := config.NewConfig(".env")

	// SECRET_KEY is one of the key whose value is fetch and used
	dbhost, err := conf.GetConfig("DB_HOST")
	if err != nil {
		panic(err)
	}

	dbuser, err := conf.GetConfig("DB_USER")
	if err != nil {
		panic(err)
	}

	dbpass, err := conf.GetConfig("DB_PASS")
	if err != nil {
		panic(err)
	}

	dbname, err := conf.GetConfig("DB_NAME")
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbuser, dbpass, dbname)

	dbOpen, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = dbOpen.Ping()
	if err != nil {
		panic(err)
	}

	db = &gorp.DbMap{Db: dbOpen, Dialect: gorp.PostgresDialect{}}
}

func GetDB() *gorp.DbMap {
	Init()
	return db
}

func CloseDB() {
	db.Db.Close()
}
