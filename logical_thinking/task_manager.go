package main

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Learn Go", Description: "Study the Go programming language", Status: "In Progress"},
		{ID: 2, Title: "Build API", Description: "Develop a RESTful API", Status: "Completed"},
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, task.Status)
	}
}
