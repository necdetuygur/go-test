package main

import (
	"fmt"
)

type User struct {
	ID       string
	Username string
}

var Users = []User{}

func main() {

	var u1 = User{}
	u1.ID = "001"
	u1.Username = "user1"

	var u2 = User{}
	u2.ID = "002"
	u2.Username = "user2"

	var u3 = User{}
	u3.ID = "003"
	u3.Username = "user3"

	Users = append(Users, u1)
	Users = append(Users, u2)
	Users = append(Users, u3)

	for _, item := range Users {
		fmt.Println(item.Username)
	}

	Users = Remove(Users, u1)

	for _, item := range Users {
		fmt.Println(item.Username)
	}
}

func Remove(s []User, item User) []User {
	index := Find(s, item)
	return append(s[:index], s[index+1:]...)
}

func Find(a []User, x User) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func P(a string){
	fmt.Println(a)
}