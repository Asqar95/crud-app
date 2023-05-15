package main

import (
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/Asqar95/crud-app/internal/transport/rest"
	"github.com/Asqar95/crud-app/pkg/database"
	"log"
	"net/http"
	"time"
)

func main() {
	//init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "crudapp",
		DBName:   "crudapp",
		SSLMode:  "disable",
		Password: "crudapp",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//init deps
	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8000",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
