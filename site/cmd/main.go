package main

import (
	"github.com/Manas-Project-Aurora/backend/site/internal/config"
	"github.com/Manas-Project-Aurora/backend/site/internal/server"
)

func main() {
	cfg := config.ParseFlags()
	server.NewServer(cfg).Run()
}
