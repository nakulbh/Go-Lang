package main

import (
	"fmt"
	"time"
)

type TodoItems struct {
	Description string
	DueDate     time.Time
	Priority    int
}

var todoList []TodoItems

func addItems(description string, dueDate time.Time, priority int) {
	todo := TodoItems{
		Description: description,
		DueDate:     dueDate,
		Priority:    priority,
	}

	todoList := append(todoList, todo)
	fmt.Println("Todo item added : ", todo.Description)
}

func listTodo() {
	if len(todoList) == 0 {
		fmt.Println("No todo Items present.")
		return
	}
	fmt.Println("Todo List")
	for i, todo := range todoList {
		fmt.Printf("%d. %s - Due %s, Priority : %d\n", i+1, todo.Description, todo.DueDate.Format("2006-01-02"), todo.Priority)
	}
}

func main() {
	addcmd := flag

}
