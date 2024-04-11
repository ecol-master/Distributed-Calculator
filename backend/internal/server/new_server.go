package server

import (
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/server/handler"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/new_expression", handler.HandlerNewExpression).Methods("GET")

	err := http.ListenAndServe(config.ServerAddress,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))
	return err
}
