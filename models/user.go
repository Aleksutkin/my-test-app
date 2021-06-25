package models

//не создавать дубли, поискать
type User struct {
	//tableName struct{} `tableName:"users"` //системный тег
	Id    int    `json:"id" sql:"id"`
	Name  string `json:"name" sql:"name" form:"name"`
	Email string `json:"email" sql:"email" form:"email"`
}

type Users []User
