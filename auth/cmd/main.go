package main

import (
	"github.com/Manas-Project-Aurora/gavna/auth/internal/config"
	"github.com/Manas-Project-Aurora/gavna/auth/internal/server"
)

func main() {
	cfg := config.ParseFlags()
	server.NewServer(cfg).Run()
}
