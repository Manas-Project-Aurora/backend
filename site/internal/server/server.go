package server

import (
	"context"
	"fmt"
	"github.com/Manas-Project-Aurora/gavna/site/internal/config"
	"github.com/Manas-Project-Aurora/gavna/site/internal/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	port       uint
	dbYamlPath string
	BasePath   string
}

func NewServer(c config.CLIConfig) *Server {
	return &Server{port: c.Port, dbYamlPath: c.DBConfigPath, BasePath: c.BasePath}
}

func (s *Server) SetPort(port uint) *Server {
	s.port = port
	return s
}

func (s *Server) SetDBConfig(path string) *Server {
	s.dbYamlPath = path
	return s
}

func (s *Server) Run() {
	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	db, err := db.ConnectToDB(s.dbYamlPath)
	if err != nil {
		errChan <- fmt.Errorf("Базе пизда: %v", err)
	}
	router := gin.Default()
	RegisterRoutes(router, db, s.BasePath)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router,
	}

	log.Printf("Server running on %d\n", s.port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("Server error: %v", err)
		}
	}()

	select {
	case sig := <-sigChan:
		log.Printf("Signal recieved: %s. Shutting down...", sig)
	case err := <-errChan:
		log.Printf("Error occured: %v\nShutting down...", err)
	}
	shutdown(db, srv)
}

func shutdown(db *gorm.DB, s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	closeDB(db)

	log.Println("Server shut down")
}

func closeDB(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting *sqlDB: %v\n", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing DB connection: %v\n", err)
	}
}
