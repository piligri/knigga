package main

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// Структура глобальной конфигурации приложения
type GlobalConfig struct {
	LineCfg `yaml:",inline"` //Конфигурация линии
	OnecCfg `yaml:",inline"` //Конфигурация 1С
}

// Структура описания конфигурации подключения к 1с
type OnecCfg struct {
	URL      string `yaml:"URL" validate:"http_url"` //URL подключения к API 1С
	WriteAPI string `yaml:"WriteAPI"`                //Адрес функции записи этапа
	GetApi   string `yaml:"GetAPI"`                  //Адрес функции чтения данных
	Token    string `yaml:"Token"`                   //Токен авторизации 1с
}

// Структура описания конфигурации линии
type LineCfg struct {
	LineIP   string `yaml:"LineIP" validate:"ip"`    //Сетевой адрес линии
	LineName string `yaml:"LineName"`                //Название линии
	LineUID  string `yaml:"LineUID" validate:"uuid"` //UID линии в 1с
}

// Создание базовой конфигурации приложения
func NewConfig(path, name string) *GlobalConfig {
	a := &GlobalConfig{}
	a.loadConfig(path, name)
	return a
}

func (s GlobalConfig) validateReq() error {
	if err := validateData(s); err != nil {
		return err
	}

	return nil
}

// Загрузка данных из конфигурационного файла YAML
func (s *GlobalConfig) loadConfig(path string, linename string) {
	cfFile, err := os.ReadFile(path + ".yaml")
	if err != nil {
		slog.Error("Ошибка чтения файла!", "Ошибка", err)
		os.Exit(1)
	}

	var data struct {
		Onec OnecCfg            `yaml:"onec"`
		Line map[string]LineCfg `yaml:",inline"`
	}

	err = yaml.Unmarshal(cfFile, &data)
	if err != nil {
		slog.Error("Ошибка файла конфигурации!", "Ошибка", err)
		os.Exit(1)
	}

	s.OnecCfg = data.Onec
	s.LineCfg = data.Line[linename]

	if err := s.validateReq(); err != nil {
		os.Exit(1)
	}
}
