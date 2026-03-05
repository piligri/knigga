package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	LineIP string `yaml:"LineIP"`
}

func main() {

	// var config = GlobalConfig{
	// 	Rcenter: Rcenter{
	// 		ID:   1,
	// 		Name: "Главный офис",
	// 		RcNetConf: RcNetConf{
	// 			IPaddr: "192.168.1.10",
	// 		},
	// 	},
	// 	OneCconf: OneCconf{
	// 		IPOneC: "10.0.0.1",
	// 	},
	// }

	// cfName := flag.String("cf", "line92", "Подгрука конфигурации из файла config.yaml")
	cfPath := flag.String("ph", "config", "Путь и имя файла конфигурации (по умолчанию ./config.yaml)")
	flag.Parse()

	cfFile, err := os.ReadFile(*cfPath + ".yaml")
	if err != nil {
		slog.Error("Ошибка чтения файла!", "%v", err)
	}
	var yam GlobalConfig
	err = yaml.Unmarshal(cfFile, &yam)
	spew.Dump(yam, err)
}
