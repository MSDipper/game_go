package utils

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func Init(router *mux.Router) *mux.Router {
	router.HandleFunc("/", HomeHandler)
	// обработчик статических файлов

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))
	return router
}

// HomeHandler обрабатывает запросы на главную страницу
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("../../templates/review.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.ExecuteTemplate(w, "review", nil)
}
