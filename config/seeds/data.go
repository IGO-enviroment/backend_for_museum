package seeds

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

const (
	lettersRus = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	letterEng  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Генарция случайных строк, определенной длины.
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterEng[rand.Intn(len(letterEng))]
	}

	return string(b)
}

func TypeEventsNames() []string {
	return []string{
		"Экскурсии",
		"Мастер-классы",
		"Спектакли",
		"Выставки",
		"Интерактивные занятия",
		"Концерты",
		"Мероприятия по генеалогии",
		"Лекции",
		"Творческие встречи",
		"Фестивали",
		"Артист-токи",
		"Кинопоказы",
	}
}

func TagsNames() []string {
	return []string{
		"18+",
		"12+",
		"6+",
		"Регулярные",
	}
}

func AreasNames() []string {
	return []string{
		"Дом Качки",
		"Водонапорная башня",
		"Дом Метенкова",
		"Креативный кластер Л52",
		"Дом Маклецкого",
		"Мемориал",
		"Пешеходная",
	}
}

type EventData struct {
	Name        string
	TicketCount int
	Publish     bool
	Price       decimal.Decimal
	Duration    int
	StartAt     time.Time
}

func EventsData() []EventData {
	return []EventData{
		{
			Name:        "Концерт",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now(),
		},
		{
			Name:        "Ешее",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now().Add(time.Hour + (time.Duration(rand.Intn(250) + 14))),
		},
		{
			Name:        "выаыва",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now().Add(time.Hour + (time.Duration(rand.Intn(250) + 14))),
		},
		{
			Name:        "ваыва",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now().Add(time.Hour + (time.Duration(rand.Intn(250) + 14))),
		},
		{
			Name:        "Пешеходная",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now().Add(time.Hour + (time.Duration(rand.Intn(250) + 14))),
		},
		{
			Name:        "Мемориал",
			TicketCount: rand.Intn(90) + 10,
			Publish:     true,
			Price:       decimal.NewFromInt(int64(rand.Intn(900) + 100)),
			Duration:    rand.Intn(6200) + 1000,
			StartAt:     time.Now().Add(time.Hour + (time.Duration(rand.Intn(250) + 14))),
		},
	}
}

func RolesNames() []string {
	return []string{
		"Директор",
		"Просмотр мероприятий",
		"Создание мероприятий",
	}
}
