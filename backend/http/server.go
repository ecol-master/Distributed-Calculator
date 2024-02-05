package http

import (
	"github.com/ecol-master/distributed_calculator/http/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/new_expression", handler.HandlerNewExpression).Methods("GET")

	http.ListenAndServe(":8000",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))

}
