package main

import (
	"fmt"
	"github.com/Asqar95/crud-app/internal/config"
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"github.com/Asqar95/crud-app/internal/service"
	grpc_client "github.com/Asqar95/crud-app/internal/transport/grpc"
	"github.com/Asqar95/crud-app/internal/transport/rest"
	"github.com/Asqar95/crud-app/pkg/database"
	"github.com/Asqar95/crud-app/pkg/hash"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// @title CRUD App API
// @version 1.0
// @description CRUD Server for Books Application

// @host localhost:8080
// @BasePath /

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	hasher := hash.NewSHA1Hasher("salt")

	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)

	usersRepo := psql.NewUsers(db)
	tokensRepo := psql.NewTokens(db)

	auditClient, err := grpc_client.NewClient(9000)
	usersService := service.NewUsers(usersRepo, tokensRepo, auditClient, hasher, []byte("sample secret"))

	handler := rest.NewHandler(booksService, usersService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
