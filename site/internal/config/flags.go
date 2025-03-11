package config

import (
	"github.com/spf13/pflag"
)

type CLIConfig struct {
	Port         uint
	DBConfigPath string
}

func ParseFlags() CLIConfig {
	port := pflag.UintP("port", "p", 8080, "Port for the server")
	path := pflag.StringP("dbpath", "d", "dbconfig.yaml", "Enable debug mode")
	pflag.Parse()

	return CLIConfig{
		Port:         *port,
		DBConfigPath: *path,
	}
}
