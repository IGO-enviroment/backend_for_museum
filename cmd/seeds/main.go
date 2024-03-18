package main

import (
	"log"
	"museum/config"
	"museum/config/seeds"
)

func main() {
	// Загрузка настроек
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки настроек: %s", err)
	}

	seeds.NewSeeds(conf).Run()
}
