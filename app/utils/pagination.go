package utils

import "math"

// Общее количество страниц.
func TotalPages(perPage, total int) int {
	return int(math.Ceil(float64(total) / float64(perPage)))
}

// Текущая страница пагинации.
func CurrentPageCalc(requestPage, countPage int) int {
	if requestPage > countPage {
		return countPage
	}

	if requestPage < 0 {
		return 0
	}

	return requestPage
}

// Расчет числа offset для корректной выдачи на страницу.
func PageOffsetCalc(currentPage, perPage int) int {
	if currentPage <= 0 {
		return currentPage
	}

	return (currentPage - 1) * perPage
}
