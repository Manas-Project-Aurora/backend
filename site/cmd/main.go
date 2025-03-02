package main

import (
	"log"

	"github.com/Manas-Project-Aurora/gavna/site/cmd/api"
	"github.com/Manas-Project-Aurora/gavna/site/config"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	api.RegisterRoutes(router, db)
	log.Println("Server running on 8080")
	router.Run(":8080")
}
