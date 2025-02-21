package main

import "fmt"

type User struct {
	name           string
	user_id        int
	email          string
	contact_number int
	IsActive       bool
}
type Task struct {
	taskname    string
	title       string
	description string
	completed   bool
}

func main() {
	user := User{
		name:           "Sri Harish",
		user_id:        1,
		email:          "XYZ@gmail.com",
		contact_number: 123445556,
		IsActive:       true,
	}
	task := Task{
		taskname: "Golang", title: "arrays", description: "learning arrays.", completed: true}
	fmt.Println(user)
	fmt.Println(task)

}
