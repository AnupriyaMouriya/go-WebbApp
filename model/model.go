package model

import (
	"github.com/jinzhu/gorm"
)


type AllFunc interface {
	AddPeople(People) (bool)
	AddArticle(Article) (bool)
	RetrieveAllPeople() []People
	RetrieveSubscriptionTable(string) []SubscriptionTable
	RetrieveAllArticle() []Article
	RetrievePeopleByEmail(string) (bool)
	RetrievePeopleById(string) (bool)
	CreateTable()
}

type People struct {
	gorm.Model
	Name string `gorm:"Column:user_name;size:255;NOT NULL"`
	Email string `gorm:"Column:user_email;type:varchar(100);UNIQUE;NOT NULL"`
	Subscription []Article `gorm:"many2many:subscription_tables";"foreignkey:id"` //;"association_foreignkey:article_id";"foreignkey:person_id"`
}

type Article struct {
	gorm.Model
	Topic     string `gorm:"Column:topic;NOT NULL"`
	Content   string `gorm:"Column:content;NOT NULL"`
}

type SubscriptionTable struct {
	PeopleID int
	ArticleID int
}

type Dbgorm struct {
	db *gorm.DB
}
func NewDbgorm() *Dbgorm{
	return &Dbgorm{db:db}
}

func(d *Dbgorm) AddPeople(people People) (bool){
	d.db.Create(&people)
return true
}

func(d *Dbgorm) AddArticle(article Article) (bool){
	d.db.Create(&article)
	return true
}

func(d *Dbgorm) RetrieveSubscriptionTable(id string) []SubscriptionTable{
	var find_article []SubscriptionTable
	d.db.Where("people_id=?", id).Find(&find_article)
	return find_article
}

func(d *Dbgorm) RetrieveAllPeople() []People{
	var people []People
	d.db.Find(&people)
	return people
}
func(d *Dbgorm) RetrievePeopleByEmail(email string) bool{
	var people People
	if cc:=d.db.Where("user_email = ?", email).Find(&people) ; cc!=nil{ return true}
	return false
}

func(d *Dbgorm) RetrievePeopleById(id string) bool{
	var people People
	if cc:=d.db.Where("id = ?", id).Find(&people) ; cc!=nil{ return true}
	return false
}

func(d *Dbgorm) RetrieveAllArticle() []Article{
	var article []Article
	d.db.Find(&article)
	return article
}
func(d *Dbgorm) CreateTable()  {
	d.db.SingularTable(true)
	d.db.AutoMigrate(&Article{},&People{},&SubscriptionTable{})
	d.db.Model(SubscriptionTable{}).AddForeignKey("people_id", "people(id)", "CASCADE", "CASCADE")
	d.db.Model(SubscriptionTable{}).AddForeignKey("article_id", "article(id)", "CASCADE", "CASCADE")
}


/*
import (
	"github.com/jinzhu/gorm"
	"webapp/connection"
)

type People struct {
	gorm.Model
	Name string `gorm:"Column:user_name;size:255;NOT NULL"`
	Email string `gorm:"Column:user_email;type:varchar(100);UNIQUE;NOT NULL"`
	Subscription []Article `gorm:"many2many:subscription_tables";"foreignkey:id"`//;"association_foreignkey:article_id";"foreignkey:person_id"`
}

type Article struct {
	gorm.Model
	Topic     string `gorm:"Column:topic;NOT NULL"`
	Content   string `gorm:"Column:content;NOT NULL"`

}

type SubscriptionTable struct {
	PeopleID int
	ArticleID int
}

func CreateTable() {

	Db := connection.GetDB()
	Db.AutoMigrate(&Article{},&People{},&SubscriptionTable{})
	Db.Model(SubscriptionTable{}).AddForeignKey("people_id", "peoples(id)", "CASCADE", "CASCADE")
	Db.Model(SubscriptionTable{}).AddForeignKey("article_id", "articles(id)", "CASCADE", "CASCADE")
    AddArticle()
}

func AddArticle(){
	article1:=&Article{ Model:    gorm.Model{}, Topic:     "topic1", Content:   "content1" }
	article2:=&Article{Model:     gorm.Model{}, Topic:     "topic2", Content:   "content2" }
	article3:=&Article{Model:     gorm.Model{}, Topic:     "topic3", Content:   "content3" }
	article4:=&Article{Model:     gorm.Model{}, Topic:     "topic4", Content:   "content4" }

	connection.GetDB().Create(article1)
	connection.GetDB().Create(article2)
	connection.GetDB().Create(article3)
	connection.GetDB().Create(article4)
}*/