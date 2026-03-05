package main

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	LineCfg `yaml:",inline"`
	OnecCfg `yaml:",inline"`
}

type OnecCfg struct {
	URL      string `yaml:"URL"`
	WriteAPI string `yaml:"WriteAPI"`
	GetApi   string `yaml:"GetAPI"`
	Token    string `yaml:"Token"`
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

	var data struct {
		Onec OnecCfg            `yaml:"onec"`
		Line map[string]LineCfg `yaml:",inline"`
	}

	err = yaml.Unmarshal(cfFile, &data)
	if err != nil {
		slog.Error("Ошибка файла конфигурации!", "%v", err)
	}

	s.OnecCfg = data.Onec
	s.LineCfg = data.Line[linename]
}
