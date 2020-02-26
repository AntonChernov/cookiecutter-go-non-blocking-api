package utils

{%if cookiecutter.use_postgres !="n" and cookiecutter.use_GORM != "n" -%}

import (
	"log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB // Global variable handle connection to db
)

//GORMDBConnection Connection to Postgress DB use GORM lib
func GORMDBConnection() (*gorm.DB, error) {
	connStr := "{{cookiecutter.use_postgres}}"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println("Database connection not created!")
		log.Fatal(err)
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		log.Panic(err)
	}

	return db, err

}

func init() {
	DB, _ = GORMDBConnection()
}

{%elif cookiecutter.use_postgres !="n" and cookiecutter.use_GORM == "n" -%}
import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


var (
	DB *sql.DB // Global variable handle connection to db
)



//PostgressDBConnection Connection to Postgress DB
func PostgressDBConnection() (*sql.DB, error) {
	connStr := "{{cookiecutter.use_postgres}}"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Database connection not created!")
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db, err

}


func init() {
	DB, _ = PostgressDBConnection()
}
{%endif -%}