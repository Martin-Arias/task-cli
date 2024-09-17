package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Martin-Arias/task-cli/pkg/models"
)

type IRepository interface {
	ReadJSON() []models.Task
	WriteJSON(desc string) error
}

type repository struct{}

func New() IRepository {
	return repository{}
}

func (r repository) ReadJSON() []models.Task {
	data := []models.Task{}
	file, _ := os.OpenFile("tasks.json", os.O_CREATE, os.ModePerm)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	return data
}

func (r repository) WriteJSON(desc string) error {
	tasks := r.ReadJSON()
	newId := len(tasks) + 1
	data := models.Task{
		Id:          newId,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now().Local().String(),
		UpdatedAt:   time.Now().Local().String(),
	}
	tasks = append(tasks, data)
	parsedData, _ := json.Marshal(tasks)
	_ = os.WriteFile("tasks.json", parsedData, os.ModePerm)
	return nil
}
