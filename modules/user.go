package modules

import (
	"gin-starter-template/db"
)

type User struct {
	Base
	Name string `json:"name,omitempty"`
	Age  int    `gorm:"type:int(11);not null;default:0" json:"age,omitempty"`
}

func (u *User) GetById() error {
	db := db.GetDb()
	db.First(&u, u.ID)
	return nil
}

func (u *User) Save() error {
	db := db.GetDb()
	db.Save(&u)
	return nil
}
