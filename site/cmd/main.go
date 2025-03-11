package main

import (
	"github.com/Manas-Project-Aurora/gavna/site/internal/config"
	"github.com/Manas-Project-Aurora/gavna/site/internal/server"
)

func main() {
	cfg := config.ParseFlags()
	server.NewServer().SetPort(cfg.Port).SetDBConfig(cfg.DBConfigPath).Run()
}
