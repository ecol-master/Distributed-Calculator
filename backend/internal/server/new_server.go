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
	router.HandleFunc("/new_user", handler.HandlerNewUser).Methods("GET")
	// router.HandleFunc("/list_of_expressions", handler.HandlerListExpressions).Methods("GET")
	router.HandleFunc("/get_expression", handler.HandlerSelectExpression).Methods("GET")
	//	router.HandleFunc("/get_config", handler.HandlerGetConfig).Methods("GET")
	//	router.HandleFunc("/set_config", handler.HandlePostConfig).Methods("POST")

	err := http.ListenAndServe(config.ServerAddress,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))
	return err
}
