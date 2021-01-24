package db

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

type M map[string]interface{}

func Connect(dbname string) *sql.DB{
	connStr := fmt.Sprintf("dbname=%s sslmode=disable", dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Query(db *sql.DB, query string) *sql.Rows{
	rows, err := db.Query(query)
	if (err != nil){
		panic(err)
	}
	return rows
}

func Exec(db *sql.DB, query string, args ...interface{}) bool{
	_, execErr := db.Exec(query, args...)
	if (execErr != nil){
		log.Fatal(execErr)
		return false
	}
	return true
}