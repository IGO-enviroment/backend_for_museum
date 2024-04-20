package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"museum/config"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/text/language"
)

type Server struct {
	cfg *config.Config
	app *fiber.App
	db  *postgres.Postgres
	l   *logger.Logger
}

func RunServer(cfg *config.Config) {
	ser := &Server{cfg: cfg}
	ser.Run()
}

func (s *Server) Run() {
	s.dependency()

	// Инициализация сервера
	s.app = fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	// Настройка сервера
	s.settingApp()

	// Настройка роутов
	SetRoutes(s)

	// Запуск сервера
	log.Fatal(s.app.Listen(fmt.Sprintf("%s:%s", s.cfg.HTTP.HOST, s.cfg.HTTP.Port)))
}

// Настройка сервера fiber.
func (s *Server) settingApp() {
	// Перехват паники
	s.app.Use(recover.New())

	// Подключение логера
	s.app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: s.l.Log,
	}))

	// Локализация текста
	s.app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:        "./config/localize",
			AcceptLanguages: []language.Tag{language.Russian, language.English},
			DefaultLanguage: language.Russian,
		}),
	)

	// Подключение сжатия
	s.app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

// Настройка зависимых приложений.
func (s *Server) dependency() {
	var err error
	// Настройка логера
	s.l = logger.New(s.cfg.Log.Level)

	// Настройка базы
	s.db, err = postgres.New(s.cfg.PG.URL)
	if err != nil {
		s.l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Подключение отоложенных задач
	// queueCli, err = queue.New("")
	// if err != nil {
	// 	s.l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	// }
	// defer queueCli.Close()
}
