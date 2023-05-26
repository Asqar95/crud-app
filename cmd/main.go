package main

import (
	"fmt"
	"github.com/Asqar95/crud-app/internal/config"
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/Asqar95/crud-app/internal/transport/rest"
	"github.com/Asqar95/crud-app/pkg/database"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	ConfigDir  = "configs"
	ConfigFile = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(ConfigDir, ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("config: %+v\n", cfg)

	//init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: "crudapp",
		DBName:   "crudapp",
		SSLMode:  cfg.DB.SSLMode,
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
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
