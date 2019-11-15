package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)


var db *gorm.DB
var err error

func CreateConnection(v *viper.Viper)  {
	username := v.GetString("nameuser")
	password := v.GetString("password")
	dbName := v.GetString("dbName")
	dbHost := v.GetString("dbHost")
	dbPort := v.GetString("dbPort")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, username, dbName, password)
	fmt.Println(dbUri)
	db, err = gorm.Open("postgres",dbUri)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection established")

}
/*
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)



var Db *gorm.DB

var err error

func BuildConnection() {
	v := viper.New()
	v.SetConfigName(".env")
	v.AddConfigPath("$GOPATH/src/webapp/connection/")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	User := v.GetString("nameuser")
	Password := v.GetString("password")
	DB := v.GetString("dbName")
	Host := v.GetString("dbHost")
	Port := v.GetString("dbPort")

	URI:=fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",Host,Port,User,DB,Password)

	fmt.Println(URI)
	Db,err=gorm.Open("postgres",URI)
	fmt.Println(Db)
	if err!=nil{
		log.Println("\nconnection failed")
	}else{
	log.Println("\nconnection established")}

}
func GetDB() *gorm.DB {
	return Db
}
*/