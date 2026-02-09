package server

import (
	"Yp_Sprint6/internal/handlers"
	"log"
	"net/http"
	"time"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.HandleIndexHTMLFile)

	router.HandleFunc("/upload", handlers.HandleUpload)

	return router
}

func NewServer(logger *log.Logger) *Server {
	// 1. Создаем роутер
	router := createRouter()

	// 2. Создаем и настраиваем http.Server
	httpServer := &http.Server{
		Addr:         ":8080",          // Порт 8080
		Handler:      router,           // Наш роутер
		ErrorLog:     logger,           // Логгер для ошибок
		ReadTimeout:  5 * time.Second,  // Таймаут чтения
		WriteTimeout: 10 * time.Second, // Таймаут записи
		IdleTimeout:  15 * time.Second, // Таймаут ожидания
	}

	// 3. Создаем и возвращаем наш сервер
	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запускается на %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
