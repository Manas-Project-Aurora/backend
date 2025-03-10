package main

import (
	"github.com/Manas-Project-Aurora/gavna/site/internal/server"
)

func main() {
	server.NewServer().SetPort(8080).SetDBConfig("dbconfig.yaml").Run()
}
