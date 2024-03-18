package seeds

import "math/rand"

const (
	lettersRus = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	letterEng = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Генарция случайных строк, определенной длины.
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = lettersRus[rand.Intn(len(lettersRus))]
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

func RolesNames() []string {
	return []string{
		"Директор",
		"Просмотр мероприятий",
		"Создание мероприятий",
	}
}
