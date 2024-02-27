package main

import (
	"log"
	"museum/config"
	v1 "museum/config/web/v1"
)

func main() {
	// Загрузка настроек
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки настроек: %s", err)
	}

	v1.RunServer(conf)
}
