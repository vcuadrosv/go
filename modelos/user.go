package modelos

import (
	"time"
)

type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	Status   bool
}

func (this *User) AddUser(id int, name string, createat time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.CreateAt = createat
	this.Status = status
}
