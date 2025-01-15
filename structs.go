package main

import (
	"fmt"
	"time"
)

type User struct {
	userName  string
	userAge   string
	createdAt time.Time
}

func ProvideStructDetails() {
	name := getUserData("Enter name : ")
	age := getUserData("Enter age : ")

	//declaration
	var appUser User

	//initialization
	appUser = User{
		userName:  name,
		userAge:   age,
		createdAt: time.Now(),
	}
	fmt.Println(appUser.userName)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
