package config

import (
	"juejinCollections/tool"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Debug  bool   `default:"false"`
	Host   string `default:"localhost"`
	Port   int    `default:"8014"`
	DbFile string `default:"./main.db"`
}{}

func init() {
	err := configor.Load(&Config, "config.yml")
	tool.PanicErr(err)

	if Config.Host == "" {
		Config.Host = "localhost"
	}
}
