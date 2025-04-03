package server

import (
	"context"
	"fmt"
	"github.com/Manas-Project-Aurora/gavna/auth/internal/config"
	"github.com/Manas-Project-Aurora/gavna/auth/internal/db"
	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Server struct {
	port       uint
	dbYamlPath string
	basePath   string
	Domain     string
}

func NewServer(c config.CLIConfig) *Server {
	return &Server{port: c.Port, dbYamlPath: c.DBConfigPath, basePath: c.BasePath, Domain: c.Domain}
}

func (s *Server) SetPort(port uint) *Server {
	s.port = port
	return s
}

func (s *Server) SetDBConfig(path string) *Server {
	s.dbYamlPath = path
	return s
}
func (s *Server) SetBasePath(basePath string) *Server {
	s.basePath = basePath
	return s
}

func (s *Server) SetSiteDomain(domain string) *Server {
	s.Domain = domain
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
	router := SetupRoutes(db, s.basePath)
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return strings.HasSuffix(origin, "."+s.Domain) || origin == "https://"+s.Domain
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
