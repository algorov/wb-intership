package service

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"io"
	"l0Service/internal/app/store"
	"l0Service/internal/util/jsonutil"
	"net/http"
)

// Service ...
type Service struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
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

	if err := s.configureStore(); err != nil {
		return err
	}

	err := s.configureAndSubscribeBroker()

	if err != nil {
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

func (s *Service) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *Service) configureAndSubscribeBroker() error {
	sn, err := stan.Connect(s.config.NatsStreaming.NatsClusterId, s.config.NatsStreaming.NatsClientId)
	if err != nil {
		return err
	}

	if _, err := sn.Subscribe(s.config.NatsStreaming.NatsTopic, func(m *stan.Msg) {
		data := string(m.Data)

		s.logger.Info("New data from NATS-STREAMING\n" + data)

		if valid := jsonutil.ValidateJsonData(data); valid {
			order, err := jsonutil.GetUnmarshallingJsonData(data)
			if err != nil {
				s.logger.Info(err)
				return
			}

			result, err := s.store.AddOrder(order)
			if err != nil {
				s.logger.Info(err)
				return
			}

			if result {
				s.logger.Info("Data was saved")
			}
		} else {
			s.logger.Info("Invalid data")
		}
	}, stan.DurableName(s.config.NatsStreaming.DurableName)); err != nil {
		return err
	}

	return nil
}

func (s *Service) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "GET: INDEX")
	}
}
