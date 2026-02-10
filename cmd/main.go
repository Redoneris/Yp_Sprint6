package main

import (
	"log"
	"os"

	"Yp_Sprint6/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	srv := server.NewServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
