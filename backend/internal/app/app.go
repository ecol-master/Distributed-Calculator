package app

import (
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	server *http.Server
}

func New() (*App, error) {
	const (
		defaultHTTPServerWriteTimeout = time.Second * 15
		defaultHTTPServerReadTimeout  = time.Second * 15
	)

	router := mux.NewRouter()
	router.HandleFunc("/new_expression", HandlerNewExpression).Methods("POST")
	router.HandleFunc("/new_user", HandlerNewUser).Methods("POST")
	router.HandleFunc("/list_of_expressions", HandlerSelectUserExpressions).Methods("GET")
	router.HandleFunc("/get_expression", HandlerSelectExpression).Methods("GET")
	//	router.HandleFunc("/get_config", handler.HandlerGetConfig).Methods("GET")
	//	router.HandleFunc("/set_config", handler.HandlePostConfig).Methods("POST")

	err := http.ListenAndServe(config.ServerAddress,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))

	return &App{
		server: &http.Server{
			Handler:      router,
			Addr:         config.ServerAddress,
			WriteTimeout: defaultHTTPServerWriteTimeout,
			ReadTimeout:  defaultHTTPServerReadTimeout,
		},
	}, err
}

func (a *App) Run() error {
	logger.Info("starting http server")

	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server was stop with err: %w", err)
	}

	logger.Info("server was stop")
	return nil
}
