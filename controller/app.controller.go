package controller

import (
	"net/http"
)

type AppController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type appController struct{}

func NewAppController() AppController {
	return &appController{}
}

func (appCon *appController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "ok"}`))
}
