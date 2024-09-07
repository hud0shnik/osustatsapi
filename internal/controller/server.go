package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hud0shnik/osustatsapi/internal/handler"
)

// Server - структура http сервера
type Server struct {
	basePath       string
	requestTimeout time.Duration
	router         http.Handler
	Server         *http.Server
}

// NewServer создаёт новый http сервер
func NewHttpServer(config *Config) *Server {

	s := &Server{
		basePath:       config.BasePath,
		requestTimeout: config.RequestTimeout,
	}

	s.NewRouter()

	s.Server = &http.Server{
		Addr:              config.ServerPort,
		Handler:           s.router,
		ReadTimeout:       config.RequestTimeout,
		ReadHeaderTimeout: config.RequestTimeout,
	}

	return s
}

// NewRouter создаёт новый роутер
func (s *Server) NewRouter() {

	// Роутер
	router := chi.NewRouter()

	// Маршруты
	router.Get(s.basePath+"/user", handler.User)
	router.Get(s.basePath+"/online", handler.Online)
	router.Get(s.basePath+"/map", handler.Map)
	router.Get(s.basePath+"/historical", handler.Historical)

	s.router = router

}
