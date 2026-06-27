package main

import (
	"fmt"
	"strconv"
	"time"
)

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func addTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-cli add \"description\"")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTask := Task{
		ID:          getNextID(tasks),
		Description: args[0],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Task added (ID: %d)\n", newTask.ID)
}

func listTasks(args []string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(args) == 0 {
		printTasks(tasks)
		return
	}

	filter := args[0]

	var filtered []Task
	for _, t := range tasks {
		if t.Status == filter {
			filtered = append(filtered, t)
		}
	}

	printTasks(filtered)
}

func updateTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: task-cli update <id> \"desc\"")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	newDesc := args[1]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Println("Error saving:", err)
				return
			}

			fmt.Println("Task updated")
			return
		}
	}

	fmt.Println("Task not found")
}

func deleteTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			if err := saveTasks(tasks); err != nil {
				fmt.Println("Error saving:", err)
				return
			}

			fmt.Println("Task deleted")
			return
		}
	}

	fmt.Println("Task not found")
}

func markStatus(args []string, status string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-cli mark-<status> <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Println("Error saving:", err)
				return
			}

			fmt.Println("Status updated")
			return
		}
	}

	fmt.Println("Task not found")
}