package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter your choice: ")
		switch option {
		case "1":
			showTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			markTaskAsCompleted(&tasks)
		case "4":
			saveTasksToFile(tasks)
		case "5":
			fmt.Println("Exiting the todo application")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Save Tasks to File")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("You have no tasks.")
		return
	}
	fmt.Println("Your tasks:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func addTask(tasks *[]Task) {
	text := getUserInput("Enter task description: ")
	*tasks = append(*tasks, Task{Text: text})
	fmt.Println("Task added successfully.")
}

func markTaskAsCompleted(tasks *[]Task) {
	showTasks(*tasks)
	taskIndex := getUserInput("Enter task number to mark as completed: ")
	index, err := strconv.Atoi(taskIndex)
	if err != nil || index < 1 || index > len(*tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}
	(*tasks)[index-1].Completed = true
	fmt.Println("Task marked as completed.")
}

func saveTasksToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
	}
	fmt.Println("Tasks saved to file.")
}
