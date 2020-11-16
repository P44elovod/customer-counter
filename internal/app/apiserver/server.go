package apiserver

import (
	"customer-counter/util"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *util.Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *util.Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("start api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleCounter())

}

func (s *Server) handleCounter() http.HandlerFunc {
	var mutex = &sync.Mutex{}
	var counter = 0

	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()

		counter++
		fmt.Fprintf(w, strconv.Itoa(counter))
		s.logger.Info(counter)
		mutex.Unlock()
	}
}
