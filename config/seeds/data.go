package seeds

import (
	"math/rand"
	"museum/app/models"
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

func TypeEventsData() [][]interface{} {
	return [][]interface{}{
		{"Экскурсии", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Мастер-классы", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Спектакли", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Выставки", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Интерактивные занятия", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Концерты", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Мероприятия по генеалогии", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Лекции", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Творческие встречи", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Aртист-токи", RandStringBytes(24), true, time.Now(), time.Now()},
		{"Кинопоказы", RandStringBytes(24), true, time.Now(), time.Now()},
	}
}

func TagsData() [][]interface{} {
	groups := &models.Tag{}
	return [][]interface{}{
		{"18+", groups.Groups()["ByAge"], RandStringBytes(24), time.Now(), time.Now()},
		{"12+", groups.Groups()["ByAge"], RandStringBytes(24), time.Now(), time.Now()},
		{"6+", groups.Groups()["ByAge"], RandStringBytes(24), time.Now(), time.Now()},
		{"0+", groups.Groups()["ByAge"], RandStringBytes(24), time.Now(), time.Now()},
		{"Новые", groups.Groups()["ByAge"], RandStringBytes(24), time.Now(), time.Now()},
		{"Только летом", groups.Groups()["Else"], RandStringBytes(24), time.Now(), time.Now()},
		{"Выбор редакции", groups.Groups()["Else"], RandStringBytes(24), time.Now(), time.Now()},
	}
}

func AreasData() [][]interface{} {
	return [][]interface{}{
		{"Дом Качки", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Водонапорная башня", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Дом Метенкова", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Креативный кластер Л52", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Дом Маклецкого", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Мемориал", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
		{"Пешеходная", RandStringBytes(24), true, RandStringBytes(24), time.Now(), time.Now()},
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
