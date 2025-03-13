package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var filename = "tasks.json"

func loadTasks() {
	file, err := os.ReadFile(filename)

	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(filename, data, 0644)
}

func addTask(name string) {
	task := Task{
		ID:        len(tasks) + 1,
		Name:      name,
		Completed: false,
	}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Task Added Successfully")
}

func listTasks() {
	fmt.Println(("\nYour ToDo List:"))
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}

func completeTask(id int) {
	for i, task := range tasks {
		if i == task.ID {
			tasks[i].Completed = true
			saveTasks()
			fmt.Println("Task mark as completed")
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	loadTasks()
	fmt.Println("To do List App")

	for {
		fmt.Println("\nChoose an option: \n1. Add Task\n2. List Tasks\n3. Complete Task\n4. Exit")
		fmt.Print("> ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Print("Enter Task: ")
			var name string
			fmt.Scanln(&name)
			addTask(name)
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter task ID to complete: ")
			var idStr string
			fmt.Scanln(&idStr)
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err == nil {
				completeTask(id)
			} else {
				fmt.Println("Invalid")
			}
		case "4":
			fmt.Println("good bye")
			return
		default:
			fmt.Println("invalid, try again")
		}
	}
}
