package tasksHandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"` // 1 - это todo, 2 - in-progress, 3 - done
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var (
	filePath = "tasks.json"

	IDnotFoundErr = errors.New("Указанное ID не найдено")
)

func AppendTask(newTask task) error {

	// Чтение содержимого файла
	fileData, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file:%w", err)
	}

	// Инициализация массива задач
	var tasks []task

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	// Добавление нового задания
	tasks = append(tasks, newTask)

	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marschal JSON: %w", err)
	}

	// Запись JSON обратно в файл
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("Task added successfully!!!")
	return nil
}

func ReadTasks() ([]task, error) {
	// Чтение содержимого файла
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Если файл не существует, возвращаем пустой массив
			return []task{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Инициализация массива задач
	var tasks []task

	// Парсинг JSON
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return tasks, nil
}

func TaskConstructor(description string) task {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println(err)
	}
	nextID := generateNextID(tasks)

	fmt.Printf("Task added successfully (ID: %d)\n", nextID)
	return task{
		ID:          nextID,
		Description: description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

}

func UpdateTaskStatus(newStatus int, targetID int) error {
	if newStatus > 3 || newStatus < 1 {
		return fmt.Errorf("Статус должен быть от одного до трёх!")
	}

	// Чтение содержимого файла
	fileData, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file:%w", err)
	}

	// Инициализация массива задач
	var tasks []task

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	// update task
	found := false
	for i := range tasks {
		if tasks[i].ID == targetID {
			found = true
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			break
		}
	}
	if !found {
		return IDnotFoundErr
	}

	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marschal JSON: %w", err)
	}

	// Запись JSON обратно в файл
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func UpdateTaskDescription(newDescription string, targetID int) error {
	// Чтение содержимого файла
	fileData, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file:%w", err)
	}

	// Инициализация массива задач
	var tasks []task

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	// update task description
	found := false
	for i := range tasks {
		if tasks[i].ID == targetID {
			found = true
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			break
		}
	}
	if !found {
		return IDnotFoundErr
	}

	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marschal JSON: %w", err)
	}

	// Запись JSON обратно в файл
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func DeleteTask(targetID int) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file:%w", err)
	}

	// Инициализация массива задач
	var tasks []task

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	// update task
	found := false
	for i := range tasks {
		if tasks[i].ID == targetID {
			found = true
			tasks = slices.Delete(tasks, i, i+1)
			break
		}
	}
	if !found {
		return IDnotFoundErr
	}

	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marschal JSON: %w", err)
	}

	// Запись JSON обратно в файл
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func generateNextID(tasks []task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
