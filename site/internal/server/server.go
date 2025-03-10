package server

import (
	"fmt"
	"log"

	"github.com/Manas-Project-Aurora/gavna/site/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port uint
}

func NewServer(port uint) *Server {
	return &Server{Port: port}
}

func (s *Server) Run() {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Базе пизда: %v", err)
	}
	router := gin.Default()
	RegisterRoutes(router, db)
	log.Println(fmt.Sprintf("Server running on %d", s.Port))
	router.Run(fmt.Sprintf(":%d", s.Port))
}
