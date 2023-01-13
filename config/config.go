package config

import (
	"encoding/json"
	"fmt"
	"juejinCollections/tool"

	"github.com/jinzhu/configor"
)

var Config = struct {
	IsDebug       bool   `default:"false" yaml:"isDebug"`       // gopkg.in/yaml.v2 只支持小写字母，如果要在yaml文件中使用大写字母，则必须在结构中声明yaml
	IsDevelopment bool   `default:"false" yaml:"isDevelopment"` // 是否开发模式 statikFs 读取的是真实的文件
	UseMock       bool   `default:"false" yaml:"useMock"`
	Host          string `default:"localhost"`
	Port          int    `default:"0"`
	DbFile        string `default:"./main.db" yaml:"dbFile"`
}{}

func init() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		tool.ShowErr(err)
	}

	if Config.Host == "" {
		Config.Host = "localhost"
	}

	if Config.Port == 0 {
		Config.Port = tool.GetEnablePort()
	}

	configData, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		tool.ShowErr(err)
		return
	}

	fmt.Println(string(configData))
}
