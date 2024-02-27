package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QinniuAccessKey string
	QinniuSecretKey string
	Valine          bool
	ValineAppid     string
	ValineAppKey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "Eugene Go 博客"
	Cfg.System.Version = 0.1
	Cfg.System.CurrentDir, _ = os.Getwd()
	if _, err := toml.DecodeFile("config/config.toml", &Cfg); err != nil {
		panic(err)
	}
}
