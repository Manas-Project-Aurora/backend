package server

import (
	"fmt"
	"log"

	"github.com/Manas-Project-Aurora/gavna/site/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port       uint
	dbYamlPath string
}

func NewServer() *Server {
	return &Server{port: 0, dbYamlPath: "dbconfig.yaml"}
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
	db, err := db.ConnectToDB(s.dbYamlPath)
	if err != nil {
		log.Fatalf("Базе пизда: %v", err)
	}
	router := gin.Default()
	RegisterRoutes(router, db)
	log.Println(fmt.Sprintf("Server running on %d", s.port))
	router.Run(fmt.Sprintf(":%d", s.port))
}
