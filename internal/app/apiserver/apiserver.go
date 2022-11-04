package apiserver

import (
	"io"
	"net/http"

	generatorService "github.com/DrDezmond/memeGen/generatorService/generator"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config    *Config
	logger    *logrus.Logger
	router    *mux.Router
	generator *generatorService.GeneratorData
}

func New(config *Config) *APIserver {
	return &APIserver{
		config:    config,
		logger:    logrus.New(),
		router:    mux.NewRouter(),
		generator: generatorService.New(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting api server on port ", s.config.BindAddr)
	s.configureRouter()
	s.logger.Info("Setup api routing")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
}

func (s *APIserver) configureRouter() {
	s.router.Handle("/upload-images", s.HandleImagesUpload()).Methods("POST", "OPTIONS")
	s.router.Handle("/upload-generator-data", s.HandleGeneratorDataUpload()).Methods("POST", "OPTIONS")
	s.router.Handle("/", s.helloHandler()).Methods("GET", "OPTIONS")
}
