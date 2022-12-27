package services

import (
	"elma_hw/internal/actions"
	"elma_hw/pkg/responses"
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
		arr, question, err := actions.GetDataCyclicRotation("%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F")
		check, err := actions.SendData(actions.CrResponse(arr, question))
		responses.Make(w, check, err)
	})
	r.Get("/task/Чудные вхождения в массив", func(w http.ResponseWriter, r *http.Request) {
		arr, question, err := actions.GetDataWonderfulOccurrences("%D0%A7%D1%83%D0%B4%D0%BD%D1%8B%D0%B5%20%D0%B2%D1%85%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%B8%D1%8F%20%D0%B2%20%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2")
		check, err := actions.SendData(actions.WoResponse(arr, question))
		responses.Make(w, check, err)
	})
	r.Get("/task/Проверка последовательности", func(w http.ResponseWriter, r *http.Request) {
		arr, question, err := actions.GetDataCheckingSequence("%D0%9F%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%BA%D0%B0%20%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8")
		check, err := actions.SendData(actions.CsResponse(arr, question))
		responses.Make(w, check, err)
	})
	r.Get("/task/Поиск отсутствующего элемента", func(w http.ResponseWriter, r *http.Request) {
		arr, question, err := actions.GetDataSearchingMissing("%D0%9F%D0%BE%D0%B8%D1%81%D0%BA%20%D0%BE%D1%82%D1%81%D1%83%D1%82%D1%81%D1%82%D0%B2%D1%83%D1%8E%D1%89%D0%B5%D0%B3%D0%BE%20%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%B0")
		check, err := actions.SendData(actions.SmResponse(arr, question))
		responses.Make(w, check, err)
	})
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		check := actions.GetAll()
		responses.Make(w, check, nil)
	})
	return r
}
