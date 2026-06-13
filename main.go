package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const filename = "task.json"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func saveTaskToFile(task []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(task)

}

func appendTaskToFile(newTask Task, filename string) error {
	var tasks []Task

	data, err := os.ReadFile(filename)
	if err == nil {

		if err := json.Unmarshal(data, &tasks); err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newTask.ID = maxID + 1

	tasks = append(tasks, newTask)

	return saveTaskToFile(tasks, filename)
}

func listAllTasks(filename string) error {
	var allTasks []Task

	task, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks found. File dosen't exist yet.")
			return nil
		}
		return err
	}

	if err := json.Unmarshal(task, &allTasks); err != nil {
		return err
	}

	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	fmt.Println("\nTodo List: \n----------")
	for _, t := range allTasks {
		status := " "
		if t.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Title)
	}

	return nil
}

func deleteTask(taskID int, filname string) error {
	var alltasks []Task

	task, err := os.ReadFile(filname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File dosen't exist yet.")
			return nil
		}
		return err
	}

	if err := json.Unmarshal(task, &alltasks); err != nil {
		return err
	}

	if len(alltasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	found := false
	for i, t := range alltasks {
		if t.ID == taskID {
			alltasks = append(alltasks[:i], alltasks[i+1:]...)
			found = true
			break
		}

	}

	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	return saveTaskToFile(alltasks, filname)
}

func markAsDone(taskID int, filename string) error {
	var allTasks []Task

	task, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exists.")
			return nil
		}
		return err
	}

	if err := json.Unmarshal(task, &allTasks); err != nil {
		return err
	}

	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	found := false
	for i := range allTasks {
		if allTasks[i].ID == taskID {
			allTasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	return saveTaskToFile(allTasks, filename)

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task title.")
			return
		}
		task := Task{
			Title:     os.Args[2],
			Completed: false,
		}

		fmt.Printf("Added task: %s\n", task.Title)
		appendTaskToFile(task, "task.json")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID. Type \"todo list\" to see all tasks.")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID. Please provide a number.")
			return
		}

		if err := deleteTask(id, filename); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task %d deleted\n", id)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID.")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}

		if err := markAsDone(id, filename); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Task %d marked as done.\n", id)
	case "list":
		if err := listAllTasks(filename); err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
		}

	case "help":
		fmt.Println("Usage: ./todo \n add [task] \n list \n done [task id] \n delete [task id] \n help")

	default:
		fmt.Println("Invalid command. Try \"todo help\" to see available commands.")
	}

}
