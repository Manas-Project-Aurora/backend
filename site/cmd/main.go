package main

import (
	"github.com/Manas-Project-Aurora/gavna/site/internal/server"
)

func main() {
	server.NewServer(8080).Run()
}
