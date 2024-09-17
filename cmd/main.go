package main

import (
	"github.com/Martin-Arias/task-cli/internal/handler"
	"github.com/Martin-Arias/task-cli/internal/repository"
)

func main() {
	r := repository.New()
	h := handler.New(r)
	h.Handle()
}
