package service

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Service ...
type Service struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *Service {
	return &Service{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *Service) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	sn, err := stan.Connect(s.config.NatsClusterId, s.config.NatsClientId)
	if err != nil {
		return err
	}

	if _, err := sn.Subscribe(s.config.NatsTopic, func(m *stan.Msg) {
		s.logger.Info("New data: " + string(m.Data))
	}, stan.DurableName(s.config.DurableName)); err != nil {
		return err
	}

	s.logger.Info("starting service...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Service) configureRouter() {
	s.router.HandleFunc("/index", s.handleIndex())

}

func (s *Service) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "GET: INDEX")
	}
}
