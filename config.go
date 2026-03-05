package main

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	LineCfg `yaml:",inline"`
}

type LineCfg struct {
	LineIP   string `yaml:"LineIP"`
	LineName string `yaml:"LineName"`
	LineUID  string `yaml:"LineUID"`
}

func NewConfig(path, name string) GlobalConfig {
	var a GlobalConfig
	a.loadConfig(path, name)
	return a
}

func (s *GlobalConfig) loadConfig(path string, linename string) {
	cfFile, err := os.ReadFile(path + ".yaml")
	if err != nil {
		slog.Error("Ошибка чтения файла!", "%v", err)
	}
	var yam map[string]GlobalConfig
	err = yaml.Unmarshal(cfFile, &yam)
	if err != nil {
		slog.Error("Ошибка файла конфигурации!", "%v", err)

	}
	*s = yam[linename]
}
