package main

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// Структура глобальной конфигурации приложения
type GlobalConfig struct {
	LineCfg  `yaml:",inline"` //Конфигурация линии
	OnecCfg  `yaml:",inline"` //Конфигурация 1С
	QRdata   `yaml:",inline"` //Конфигурация тегов данных QR
	QRstatus `yaml:",inline"` //Конфигурация тегов статуса QR
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

// Структура описания тегов данных QR
type QRdata struct {
	Fio      string `yaml:"Fio"`      //ФИО исполнителя
	Etap     string `yaml:"Etap"`     //Номер Этапа
	Operacia string `yaml:"Operacia"` //Строка операции
	Rc       string `yaml:"Rc"`       //Наименование рабочего центра
	Uchastok string `yaml:"Uchastok"` //Участок исполнителя
	PlanM    string `yaml:"PlanM"`    //Метраж из заявки на этап
	FaktM    string `yaml:"FaktM"`    //Фактическая выработка из счетчика линии
	QRFiao   string `yaml:"QRFiao"`   //Считанный и обработанный QR Исполнителя
	QREtap   string `yaml:"QREtap"`   //Считанный и обработанный QR Этапа
}

// Структура описания тегов статусов QR
type QRstatus struct {
	InitQRWindow        string `yaml:"InitQRWindow"`        //Окно считывания QR кода открыто
	FioReadComplite     string `yaml:"FioReadComplite"`     //Код QR ФИО считан
	EtapReadComplite    string `yaml:"EtapReadComplite"`    //Код QR Этапа считан
	EtapProcess         string `yaml:"EtapProcess"`         //Этап в работе - TRUE, этап не начат FALSE
	EtapEnd             string `yaml:"EtapEnd"`             //Этап завершен
	AllDataComplite     string `yaml:"AllDataComplite"`     //Все данные считаны и готовы к записи в 1с
	AllDataReadytoWrite string `yaml:"AllDataReadytoWrite"` //?????
	NextQRAccept        string `yaml:"NextQRAccept"`        //Разрешение считывания QR (инверсное поле, FALSE - разрешено чтение)
	ScanerQREnable      string `yaml:"ScanerQREnable"`      //Разрешение работы сканера
	QR                  string `yaml:"QR"`                  //Считанный QR
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
		Onec          OnecCfg            `yaml:"onec"`
		Line          map[string]LineCfg `yaml:",inline"`
		OpcDataTags   QRdata             `yaml:"OpcDataTags"`
		OpcStatusTags QRstatus           `yaml:"OpcStatusTags"`
	}

	err = yaml.Unmarshal(cfFile, &data)
	if err != nil {
		slog.Error("Ошибка файла конфигурации!", "Ошибка", err)
		os.Exit(1)
	}

	s.OnecCfg = data.Onec
	s.LineCfg = data.Line[linename]
	s.QRstatus = data.OpcStatusTags
	s.QRdata = data.OpcDataTags

	if err := s.validateReq(); err != nil {
		os.Exit(1)
	}
}
