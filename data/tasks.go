package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

var (
	Green   = color.New(color.FgGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Reset   = color.New(color.Reset).SprintFunc()
)

type Task struct {
	ID          int    `json:"id"`
	Priority    string `json:"priority"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var configDir string

func init() {
	if runtime.GOOS == "windows" {
		appData, _ := os.UserConfigDir()
		configDir = filepath.Join(appData, "task-cli", "todos")
	} else {
		homeDir, _ := os.UserHomeDir()
		configDir = filepath.Join(homeDir, ".config", "todos")
	}
	_ = os.MkdirAll(configDir, 0755)
}

func AddTask(priority string, description string) (int, error) {
	id := generateID()
	task := Task{ID: id, Priority: priority, Description: description, Status: "todo"}
	filePath := filepath.Join(configDir, fmt.Sprintf("%d.json", id))
	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(task)
	return id, nil
}

func UpdateTask(id int, newPriority string, newDescription string) error {
	task, err := getTask(id)
	if err != nil {
		return err
	}
	task.Priority = newPriority
	task.Description = newDescription
	return saveTask(task)
}

func DeleteTask(id int) error {
	filePath := filepath.Join(configDir, fmt.Sprintf("%d.json", id))
	return os.Remove(filePath)
}

func MarkTask(id int, status string) error {
	task, err := getTask(id)
	if err != nil {
		return err
	}
	task.Status = status
	return saveTask(task)
}

func ListTasks(status string) ([]Task, error) {
	files, err := os.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, file := range files {
		var task Task
		data, _ := os.ReadFile(filepath.Join(configDir, file.Name()))
		json.Unmarshal(data, &task)
		if status == "" || task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func getTask(id int) (*Task, error) {
	filePath := filepath.Join(configDir, fmt.Sprintf("%d.json", id))
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("Tarefa não encontrada")
	}

	var task Task
	json.Unmarshal(data, &task)
	return &task, nil
}

func saveTask(task *Task) error {
	filePath := filepath.Join(configDir, fmt.Sprintf("%d.json", task.ID))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(task)
}

func generateID() int {
	files, _ := os.ReadDir(configDir)
	return len(files) + 1
}
