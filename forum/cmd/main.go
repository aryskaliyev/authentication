package main

import (
	"lincoln.boris/forum"
	"lincoln.boris/forum/internal/handler"
	"lincoln.boris/forum/internal/repository"
	"lincoln.boris/forum/internal/service"
	"lincoln.boris/forum/pkg/logger"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log := logger.NewLogger()

	db, err := repository.NewSQLiteDB(repository.Config{
		DSN: "./database/mydb.db",
	})

	if err != nil {
		log.ErrorLog.Println(err)
	}

	log.InfoLog.Println("Connnected to DB")

	repos := repository.NewRepository(db)
	log.InfoLog.Println("Connnected to Repository")

	services := service.NewService(repos)
	log.InfoLog.Println("Connnected to Service")

	handlers := handler.NewHandler(services)
	log.InfoLog.Println("Connnected to Handler")

	srv := new(forum.Server)

	if err = srv.Run(handlers.InitRoutes()); err != nil {
		log.ErrorLog.Println(err)
	}
}
