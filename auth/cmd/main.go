package main

import (
	"github.com/Manas-Project-Aurora/backend/auth/internal/config"
	"github.com/Manas-Project-Aurora/backend/auth/internal/server"
)

func main() {
	cfg := config.ParseFlags()
	server.NewServer(cfg).Run()
}
