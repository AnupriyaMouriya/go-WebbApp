package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"webapp/model"
)

type Control struct {
	ds model.AllFunc
}
func NewController(ds model.AllFunc) *Control{
	return &Control{ds:ds}
}

func (c Control) AddUser(w http.ResponseWriter , r *http.Request) {
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

	if match_email && match_name {
		b:=c.ds.RetrievePeopleByEmail(user.Email)
		if !b{
			c.ds.AddPeople(user)
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
func(c Control) AllArticle(w http.ResponseWriter, r *http.Request){
	article := c.ds.RetrieveAllArticle()
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&article)
}

//right
func (c Control)AllUser(w http.ResponseWriter, r *http.Request){
	people:=c.ds.RetrieveAllPeople()
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&people)
}

//right
func(c Control) UserArticle(w http.ResponseWriter, r *http.Request){

	params:=mux.Vars(r)
	b:=c.ds.RetrievePeopleById(params["user_id"])
	//err:=connection.Db.Debug().Where("id=?",params["user_id"]).Find(&user).Error
	if !b {
		w.WriteHeader(http.StatusBadRequest)
		x, _ := json.Marshal("No user found")
		_, _ = w.Write(x)
	}else {
		find_article:=c.ds.RetrieveSubscriptionTable(params["user_id"])
		//connection.Db.Where("people_id=?", params["user_id"]).Find(&find_article)
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
func(c Control) SubscribeArticle(w http.ResponseWriter, r *http.Request){

	params:=mux.Vars(r)
	user:=&model.People{}
	b:=c.ds.RetrievePeopleById(params["user_id"])
	/*err:=connection.Db.Debug().Where("id=?",params["user_id"]).Find(&user).Error

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		x,_:=json.Marshal("No user found")
		_, _ = w.Write(x)

	}*/
	//err:=connection.Db.Debug().Where("id=?",params["user_id"]).Find(&user).Error
	if !b {
		w.WriteHeader(http.StatusBadRequest)
		x, _ := json.Marshal("No user found")
		_, _ = w.Write(x)
	} else {
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


func (c Control) Initializer() {
	c.ds.CreateTable()
	article1:=model.Article{Model:     gorm.Model{}, Topic:     "topic1", Content:   "content1" }
	article2:=model.Article{Model:     gorm.Model{}, Topic:     "topic2", Content:   "content2" }
	article3:=model.Article{Model:     gorm.Model{}, Topic:     "topic3", Content:   "content3" }
	article4:=model.Article{Model:     gorm.Model{}, Topic:     "topic4", Content:   "content4" }

	c.ds.AddArticle(article1)
	c.ds.AddArticle(article2)
	c.ds.AddArticle(article3)
	c.ds.AddArticle(article4)


}