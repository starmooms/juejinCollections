package config

import (
	"juejinCollections/tool"

	"github.com/jinzhu/configor"
)

var Config = struct {
	IsDebug bool   `default:"false" yaml:"isDebug"` // gopkg.in/yaml.v2 只支持小写字母，如果要在yaml文件中使用大写字母，则必须在结构中声明yaml
	Host    string `default:"localhost"`
	Port    int    `default:"8014"`
	DbFile  string `default:"./main.db" yaml:"dbFile"`
}{}

func init() {
	err := configor.Load(&Config, "config.yml")
	tool.PanicErr(err)

	if Config.Host == "" {
		Config.Host = "localhost"
	}
}
