package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	email string
}

func NewUser(id int, name, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

func (user *User) getName() string {
	return user.name
}

func (user *User) setName(name string) {
	user.name = name
}

func main() {

	user := NewUser(505, "Aurelio", "abfm")
	fmt.Println(*user)
	user.setName(user.getName() + " Buarque")
	fmt.Println(*user)

}
