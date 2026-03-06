package main

import (
	"flag"
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
)

var Config *GlobalConfig

func init() {
	cfName := flag.String("cf", "line92", "Подгрука конфигурации из файла config.yaml")
	cfPath := flag.String("ph", "config", "Путь и имя файла конфигурации (по умолчанию ./config.yaml)")
	flag.Parse()

	Config = NewConfig(*cfPath, *cfName)
}

func validateData(s any) error {
	v := validator.New()
	if err := v.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			slog.Error("Ошибка валидации", "Ошибка", err.Error(), "Содержимое поля", err.Value())
		}
		return err
	}
	return nil
}

func main() {
	slog.Info("Стартуем!", "Конфигурация", Config)

	//Для теста записи 1с
	val := &RequestData{
		FioUID:  "56565656-5656-5656-5656-123456789011",
		EtapUID: "56565656-5656-5656-5656-123456789011",
		RCUID:   "16565656-5656-5656-5656-123456789011",
		Metrazh: "1",
	}
	res, err := write1c(*val)
	if err != nil {
		slog.Error("Ошибка записи в 1с", "", err)
	}
	spew.Dump(res)
}
