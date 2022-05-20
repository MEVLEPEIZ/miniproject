/*package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gorm-many2many/models"
)

var database *gorm.DB

func Init() {
	database, connectionError := gorm.Open("postgres", "host=localhost port=5432 user=zoharyzgeav dbname=many2many sslmode=disable")
	if connectionError != nil {
		panic(connectionError)
	}

	database.AutoMigrate(&models.Feature{})
	database.AutoMigrate(&models.Photo{})
	database.AutoMigrate(&models.Place{})
}

// Get TODO: save database instance in context to eliminate the Get() method
func Get() *gorm.DB {
	database, _ := gorm.Open("postgres", "host=localhost port=5432 user=zoharyzgeav dbname=many2many sslmode=disable")
	return database
}*/
package config
import (
	"fmt"
	"github.com/MEVLEPEIZ/Mini-Project/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "*Eren1964",
		DB_Port:     "3360",
		DB_Host:     "127.0.0.1",
		DB_Name:     "db_miniproject",
	}

	connectionString :=	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	//DB.AutoMigrate(&models.Feature{})
	//DB.AutoMigrate(&models.Photo{})
	//DB.AutoMigrate(&models.Place{})
}