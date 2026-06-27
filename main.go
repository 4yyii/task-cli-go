package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask(os.Args[2:])
	case "list":
		listTasks(os.Args[2:])
	case "update":
		updateTask(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	case "mark-done":
		markStatus(os.Args[2:], "done")
	case "mark-in-progress":
		markStatus(os.Args[2:], "in-progress")
	default:
		fmt.Println("Unknown command")
	}
}