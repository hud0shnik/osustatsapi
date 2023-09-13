package controllers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hud0shnik/OsuStatsApi/api"
)

// Server - структура сервера
type Server struct {
	basePath       string
	requestTimeout time.Duration
	router         http.Handler
	Server         *http.Server
}

// NewServer создаёт новый сервер
func NewServer(config *Config) *Server {

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
	router.Get(s.basePath+"/user", api.User)
	router.Get(s.basePath+"/online", api.Online)
	router.Get(s.basePath+"/map", api.Map)
	router.Get(s.basePath+"/historical", api.Historical)

	s.router = router

}
