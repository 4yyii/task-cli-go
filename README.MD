# Task CLI (Go)

Simple CLI task tracker built with Go (no external libraries).

## ✨ Features

- ✅ Add task
- ✏️ Update task
- ❌ Delete task
- 🔄 Mark task status:
  - todo
  - in-progress
  - done
- 📋 List all tasks
- 🔍 Filter tasks by status
- 💾 Persistent storage (JSON file)

---

## 📦 Tech Stack

- Go (Golang)
- Standard Library only (no external dependencies)

---

## 📁 Project Structure

task-cli/
├── main.go # Entry point
├── commands.go # CLI commands (add, list, update, dll)
├── storage.go # Load & save JSON
├── task.go # Struct model
├── utils.go # Helper (printTasks)
├── tasks.json # Data storage (ignored by git)
├── .gitignore
└── go.mod


---

## Usage
```bash
go run . add "Belajar Go" # add task
go run . list # list all task
go run . list todo # list by status 
go run . list in-progress # list by status
go run . list done # list by status
go run . update 1 "Belajar Go lebih dalam" # updated task
go run . delete 1 # delete task
go run . mark-in-progress 1 # mark task status as in-progress
go run . mark-done 1 # mark task status as done
