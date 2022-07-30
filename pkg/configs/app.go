package configs

// docker run --name dockerdb -e MYSQL_ROOT_PASSWORD=qwer1234 -d -p 3306:3306 mysql:latest
// docker exec -it dockerdb /bin/bash

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )
// _ "github.com/jinzhu/gorm/dialects/postgres"
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	// dsn := "admin:qwer1234@tcp(127.0.0.1:3306)/Local instance 3306?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// d, err := gorm.Open(mysql.New(mysql.Config{DSN: "admin:qwer1234@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"}), &gorm.Config{})
	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, err := gorm.Open("mysql", "root:qwer1234@tcp(127.0.0.1:3306)/IHAR_practice?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("postgres", "testuser:qwer1234@tcp(127.0.0.1:5432)/testdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
