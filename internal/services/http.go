package services

import (
	"elma_hw/internal/actions"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func New() ServerImplementation {
	return ServerImplementation{}
}

type ServerImplementation struct {
}

func (s ServerImplementation) BuildRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/task/Циклическая ротация", func(w http.ResponseWriter, r *http.Request) {
		arr, err := actions.GetDataCyclicRotation()
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		fmt.Fprint(w, arr)
		//result, data, err := actions.GetLogin()
		//responses.DrawPage(w, result, data, err)
	})
	return r
}
