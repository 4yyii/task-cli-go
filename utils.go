package main

import "fmt"

func printTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}
}