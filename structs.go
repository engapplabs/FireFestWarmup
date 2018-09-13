package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	email string
}

// NewUser ...
func NewUser(id int, name, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

func main() {

	var user = User{id: 0, name: "Aurelio", email: "abmf"}
	fmt.Println(user)

	var user1 = NewUser(0, "aurelio", "abmf")
	fmt.Println(*user1)

}
