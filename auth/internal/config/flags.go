package config

import (
	"github.com/spf13/pflag"
	"strings"
)

type CLIConfig struct {
	Port         uint
	DBConfigPath string
	BasePath     string
	Domain       string
}

func ParseFlags() CLIConfig {
	port := pflag.UintP("port", "p", 8080, "Port for the server")
	path := pflag.StringP("dbpath", "d", "dbconfig.yaml", "Set dbConfig.yaml path")
	base := pflag.StringP("basepath", "b", "", "Set base path")
	domain := pflag.StringP("url", "u", "avtomoykabot.store", "Set site's url")

	pflag.Parse()

	*base = strings.TrimSuffix(*base, "/")
	*domain = strings.TrimPrefix(*domain, "https://")
	*domain = strings.TrimPrefix(*domain, "http://")

	return CLIConfig{
		Port:         *port,
		DBConfigPath: *path,
		BasePath:     *base,
		Domain:       *domain,
	}
}
