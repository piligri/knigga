package main

import (
	"flag"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// Ваши структуры (оставлены без изменений)
type GlobalConfig struct {
	OneCconf `yaml:",inline"` // ,inline "поднимает" поля OneCconf на уровень выше в YAML
	Rcenter  `yaml:",inline"`
}

type OneCconf struct {
	IPOneC string `yaml:"iponec"`
	APIkey string `yaml:"apikey"`
}

type Rcenter struct {
	Name      string `yaml:"name"`
	ID        int    `yaml:"id"`
	RcNetConf `yaml:",inline"`
}

type RcNetConf struct {
	IPaddr string `yaml:"ipaddr"`
	UID    string `yaml:"uid"`
}

func main() {
	// 1. Параметры запуска
	// Использование: ./app -env test_lab
	envPtr := flag.String("env", "prod_office", "название блока конфигурации")
	pathPtr := flag.String("path", "cfg.yaml", "путь к файлу конфигурации")
	flag.Parse()

	// 2. Читаем файл
	data, err := os.ReadFile(*pathPtr)
	if err != nil {
		slog.Error("Не удалось прочитать файл", "path", *pathPtr, "err", err)
		os.Exit(1)
	}

	// 3. Парсим в карту (Map), где ключ - строка, а значение - ваша структура
	var allConfigs map[string]GlobalConfig
	if err := yaml.Unmarshal(data, &allConfigs); err != nil {
		slog.Error("Ошибка парсинга YAML", "err", err)
		os.Exit(1)
	}

	// 4. Выбираем нужный блок
	cfg, ok := allConfigs[*envPtr]
	if !ok {
		slog.Error("Блок конфигурации не найден", "env", *envPtr)
		os.Exit(1)
	}

	// Теперь cfg — это строго типизированная структура GlobalConfig
	slog.Info("Конфигурация выбрана успешно",
		"target", *envPtr,
		"office_name", cfg.Name,
		"ip", cfg.IPaddr,
	)

	// Далее используем cfg.IPaddr, cfg.APIkey и т.д.
}
