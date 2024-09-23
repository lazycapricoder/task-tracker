package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func main() {
	if len(os.Args) < 2 {
		showCommandList()
		return
	}

	cmd := os.Args[1]
	param := ""
	if len(os.Args) > 2 {
		param = os.Args[2]
	}

	switch cmd {
	case "add":
		add(param)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> <new_name>")
			return
		}
		updateName(param, os.Args[3])
	case "delete":
		remove(param)
	case "mark-in-progress":
		markStatus(param, StatusInProgress)
	case "mark-done":
		markStatus(param, StatusDone)
	case "list":
		list(param)
	default:
		fmt.Println("Command not found, use ./main to see the list of available commands")
	}
}

func showCommandList() {
	fmt.Println("Commands:")
	fmt.Println("  add <name>              : add task")
	fmt.Println("  update <id> <name>      : update task name by id")
	fmt.Println("  delete <id>             : delete task by id")
	fmt.Println("  mark-in-progress <id>   : mark task as in progress")
	fmt.Println("  mark-done <id>          : mark task as done")
	fmt.Println("  list <status>           : list tasks by status")
}

func list(status string) {
	if !checkIfStatusExists(status) {
		fmt.Println("Status doest not exists")
		return
	}
	tasks := readTasks()
	if tasks == nil {
		return
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if status == "" || task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}

	printTasks(filteredTasks)
}

// check if the given status is valid
func checkIfStatusExists(status string) bool {
	switch status {
	case StatusTodo, StatusInProgress, StatusDone, "":
		return true
	default:
		return false
	}
}

func printTasks(tasks []Task) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func markStatus(id, status string) {
	update(id, "status", status)
}

func updateName(id string, param string) {
	update(id, "name", param)
}

func remove(id string) {
	if id == "" {
		fmt.Println("id is required")
		return
	}

	tasks := readTasks()
	if tasks == nil {
		return
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if strconv.Itoa(task.Id) != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	writeTasks(updatedTasks)
}

func update(id, field, value string) {
	if id == "" {
		fmt.Println("id is required")
		return
	}

	tasks := readTasks()
	if tasks == nil {
		return
	}

	for i, task := range tasks {
		if strconv.Itoa(task.Id) == id {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			tasks[i].UpdatedAt = currentTime
			if field == "name" {
				tasks[i].Name = value
			} else if field == "status" {
				tasks[i].Status = value
			}
			writeTasks(tasks)
			return
		}
	}

	fmt.Printf("Task with id %s not found\n", id)
}

func add(param string) {
	tasks := readTasks()
	if tasks == nil {
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	task := Task{
		Id:        len(tasks) + 1,
		Name:      param,
		Status:    StatusTodo,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	tasks = append(tasks, task)
	writeTasks(tasks)
}

// read tasks from a file
func readTasks() []Task {
	filename := "task_tracker.json"
	if !fileExists(filename) {
		return []Task{}
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return nil
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error unmarshalling tasks:", err)
		return nil
	}
	return tasks
}

// write the task to a file
func writeTasks(tasks []Task) {
	filename := "task_tracker.json"
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
		return
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		fmt.Println("Error writing tasks to file:", err)
	}
}

// Check if the file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
