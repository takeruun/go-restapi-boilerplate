package controller

import (
	"app/controller/dto"
	"app/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type TodoController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request, todoId int)
	Edit(w http.ResponseWriter, r *http.Request, todoId int)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request, todoId int)
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
	todos, err := todoCon.todoU.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	var data []byte
	data, err = json.MarshalIndent(&todos, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (todoCon *todoController) Show(w http.ResponseWriter, r *http.Request, todoId int) {
	todo, err := todoCon.todoU.Show(r.Context(), todoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	var data []byte
	data, err = json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (todoCon *todoController) Edit(w http.ResponseWriter, r *http.Request, todoId int) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var updateParams dto.TodoUpdateRequestDto
	json.Unmarshal(body, &updateParams)

	todo, err := todoCon.todoU.Edit(todoId, &updateParams)

	var data []byte
	data, err = json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (todoCon *todoController) Create(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var createParams dto.TodoCreateRequestDto
	json.Unmarshal(body, &createParams)

	todo, err := todoCon.todoU.Create(r.Context(), &createParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	var data []byte
	data, err = json.MarshalIndent(&todo, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (todoCon *todoController) Delete(w http.ResponseWriter, r *http.Request, todoId int) {
	if err := todoCon.todoU.Delete(r.Context(), todoId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write([]byte(`{"message": "ok"}`))
}
