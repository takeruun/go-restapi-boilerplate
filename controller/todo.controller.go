package controller

import (
	"app/usecase"
	"encoding/json"
	"net/http"
)

type TodoController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	todoU usecase.TodoUsecase
}

func NewTodoController(todoU usecase.TodoUsecase) TodoController {
	return &todoController{
		todoU: todoU,
	}
}

func (todoCon *todoController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := todoCon.todoU.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var data []byte
	data, err = json.MarshalIndent(&todos, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(data)
}

func (todoCon *todoController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func (todoCon *todoController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func (todoCon *todoController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func (todoCon *todoController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
