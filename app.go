package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"webapp/controller"
	"webapp/model"

)


func main() {
	v := loadConfig()
	model.CreateConnection(v)
	ds := model.NewDbgorm()
	ct := controller.NewController(ds)
	ct.Initializer()
	router := mux.NewRouter()
	router.HandleFunc("/user/add", controller.AddUser).Methods("POST")
	router.HandleFunc("/user", controller.AllUser).Methods("GET")
	router.HandleFunc("/article", controller.AllArticle).Methods("GET")
	router.HandleFunc("/article/{user_id}", controller.UserArticle).Methods("GET")
	router.HandleFunc("/article/{user_id}/subscribe/{article_id}", controller.SubscribeArticle).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}


func loadConfig() *viper.Viper{
	v :=  viper.New()
	v.AutomaticEnv()
	v.SetConfigName(".env")
	v.AddConfigPath("$GOPATH/src/webapp/model/")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return v
}


/*
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
	"webapp/connection"
	"webapp/model"
)


func ValidateName(name string)(bool){
	matched,_:=regexp.MatchString(`^[a-zA-Z ]+$`,name)
	return matched
}


func ValidateEmail(email string)(bool){
	match,_:=regexp.MatchString(`^.+@[a-zA-Z0-9-.]+.([a-zA-Z]{2,3}|[0-9]{1,3}$)`,email)
	return match
}


//right
func AddUser(w http.ResponseWriter , r *http.Request) {
	user:=model.People{}
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		x,_:=json.Marshal("Incorrect Json Format")
		_, _ = w.Write(x)
		return

	}
	match_name:=ValidateName(user.Name)
	match_email:=ValidateEmail(user.Email)

	if match_email && match_name{

		ss:=connection.Db.Where("user_email = ?", user.Email).Find(&user).Error
		if ss!=nil{
			connection.Db.Create(&user)
		}else{

			w.WriteHeader(http.StatusConflict)
			x,_:=json.Marshal("Email already present")
			_, _ = w.Write(x)
			return
		}
	}else {
		w.WriteHeader(http.StatusBadRequest)
		if match_email {
			x, _ := json.Marshal("Incorrect name Format")
			_, _ = w.Write(x)
		} else {
			x, _ := json.Marshal("Incorrect Email format")
			_, _ = w.Write(x)
		}
	}
}


//right
func AllArticle(w http.ResponseWriter, r *http.Request){
	var article []model.Article
	connection.Db.Find(&article)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&article)
}

//right
func AllUser(w http.ResponseWriter, r *http.Request){
	var people []model.People
	connection.Db.Find(&people)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&people)
}

//right
func UserArticle(w http.ResponseWriter, r *http.Request){

	params:=mux.Vars(r)
	user:=&model.People{}
	err:=connection.Db.Debug().Where("id=?",params["user_id"]).Find(&user).Error

	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		x, _ := json.Marshal("No user found")
		_, _ = w.Write(x)
	}else {
		var find_article []model.SubscriptionTable
		connection.Db.Where("people_id=?", params["user_id"]).Find(&find_article)
		list := []int{}
		for x, _ := range find_article {
			list = append(list, (find_article[x].ArticleID))
		}
		w.WriteHeader(http.StatusOK)
		fmt.Println(list)
		_ = json.NewEncoder(w).Encode(&list)
	}
}

//right
func SubscribeArticle(w http.ResponseWriter, r *http.Request){

	params:=mux.Vars(r)
	user:=&model.People{}
	err:=connection.Db.Debug().Where("id=?",params["user_id"]).Find(&user).Error

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		x,_:=json.Marshal("No user found")
		_, _ = w.Write(x)

	}else {
		connection.Db.Debug().Model(&user).Association("Subscription").Find(&user.Subscription)

	article:=&model.Article{}
	err=connection.Db.Debug().Where("id=?",params["article_id"]).Find(&article).Error

	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		x,_:=json.Marshal("No article found")
		_, _ = w.Write(x)
	}else {
		user.Subscription = append(user.Subscription, *article)
		connection.Db.Debug().Model(&model.People{}).Where("id=?", params["user_id"]).Update("Subscription", user.Subscription)
		connection.Db.Save(&user)
	}
	}
}
//right
func main() {

	connection.BuildConnection()
	model.CreateTable()

	router := mux.NewRouter()

	router.HandleFunc("/user/add",AddUser).Methods("POST")
	router.HandleFunc("/user",AllUser).Methods("GET")
	router.HandleFunc("/article",AllArticle).Methods("GET")
	router.HandleFunc("/article/{user_id}",UserArticle).Methods("GET")
	router.HandleFunc("/article/{user_id}/subscribe/{article_id}",SubscribeArticle).Methods("POST")

	db:=connection.GetDB()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
*/

