package main

import (
	"flag"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	cfName := flag.String("cf", "line92", "Подгрука конфигурации из файла config.yaml")
	cfPath := flag.String("ph", "config", "Путь и имя файла конфигурации (по умолчанию ./config.yaml)")
	flag.Parse()

	a := NewConfig(*cfPath, *cfName)

	spew.Dump(a)
}
