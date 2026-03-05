package main

import (
	"flag"
	"log/slog"
)

var Config GlobalConfig

func init() {
	cfName := flag.String("cf", "line92", "Подгрука конфигурации из файла config.yaml")
	cfPath := flag.String("ph", "config", "Путь и имя файла конфигурации (по умолчанию ./config.yaml)")
	flag.Parse()

	Config = NewConfig(*cfPath, *cfName)
}

func main() {
	slog.Info("Стартуем!", "Конфигурация", Config)
}
