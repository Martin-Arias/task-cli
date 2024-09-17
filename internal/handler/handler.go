package handler

import (
	"fmt"
	"os"
	"slices"

	"github.com/Martin-Arias/task-cli/internal/repository"
)

type IHandler interface {
	Handle()
}

type handler struct {
	r repository.IRepository
}

func New(r repository.IRepository) IHandler {
	return handler{
		r: r,
	}
}

func (h handler) Handle() {
	// First arg posible values: add, update, delete, mark-in-progress, mark-done, list
	posibleArgs := []string{"add", "update", "delete", "mark-in-progress", "mark-done", "list"}
	if !slices.Contains(posibleArgs, os.Args[1]) {
		fmt.Printf("invalid arg: %s \n", os.Args[1])
		return
	}
	switch os.Args[1] {
	case "add":
		if os.Args[2] == "" {
			fmt.Println("description can not be null")
		}
		_ = h.r.WriteJSON(os.Args[2])
	case "update":
		fmt.Println(os.Args[1])
	case "delete":
		fmt.Println(os.Args[1])
	case "mark-in-progress":
		fmt.Println(os.Args[1])
	case "mark-done":
		fmt.Println(os.Args[1])
	case "list":
		data := h.r.ReadJSON()
		fmt.Println("<-------List--------->")
		fmt.Println("<-------------------->")
		for _, v := range data {
			fmt.Printf("Id:%v - Description: %s - Status: %s \n", v.Id, v.Description, v.Status)
		}
		fmt.Println("<-------------------->")
		fmt.Println("<-------------------->")
	}
}
