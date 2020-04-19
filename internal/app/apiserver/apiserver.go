package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/isratmir/restapi/internal/app/store"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

//NewAPIServer constructor
func NewAPIServer(c *Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start server method
func (s *APIServer) Start() error  {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")

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

func (s *APIServer) configureRouter()  {
	s.router.HandleFunc("/", s.handleMain())
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.NewStore(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleMain() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Main")
	}
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		   io.WriteString(w, "Hello")
	}
}