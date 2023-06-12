package main

import (
	"context"
	"fmt"
	"github.com/Asqar95/crud-app/internal/config"
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"github.com/Asqar95/crud-app/internal/server"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/Asqar95/crud-app/internal/transport/rest"
	"github.com/Asqar95/crud-app/pkg/database"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(new(log.JSONFormatter))
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v\n", cfg)
	//init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	//init deps
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// init & run server
	srv := server.NewServer(&http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        handlers.Init(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	go func() {
		log.Fatal(srv.Run())
	}()

	log.Println("Starting server on port 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to shutdown server: ", err)
	}

	if err = db.Close(); err != nil {
		log.Fatal("Failed to close database: ", err)
	}

	fmt.Println("Server stopped")
}
