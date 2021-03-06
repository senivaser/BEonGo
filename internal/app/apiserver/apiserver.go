package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senivaser/BEonGo/internal/app/model"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *model.Store
}

func New(config *Config, store *model.Store) *APIServer {

	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
}

func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting API server...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
