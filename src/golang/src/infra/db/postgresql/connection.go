package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=myuser dbname=db password=secret")

	if err != nil {
		panic(err)
	}
	return db
}

func CloseConn() {
	db.Close()
}
