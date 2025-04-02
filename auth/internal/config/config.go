package config

import (
	"github.com/spf13/pflag"
)

type CLIConfig struct {
	Port         uint
	DBConfigPath string
	BasePath     string
}

func ParseFlags() CLIConfig {
	port := pflag.UintP("port", "p", 8080, "Port for the server")
	path := pflag.StringP("dbpath", "d", "dbconfig.yaml", "Set dbConfig.yaml path")
	base := pflag.StringP("basepath", "b", "", "Set base path")
	pflag.Parse()

	return CLIConfig{
		Port:         *port,
		DBConfigPath: *path,
		BasePath:     *base,
	}
}
