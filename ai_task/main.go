package main

import (
	"fmt"
)

// **Описание**: Создайте функцию для сравнения двух хешированных паролей
//
// **Входные данные**: Две строки - hash1 и hash2 (хешированные пароли в формате bcrypt)
//
// **Выходные данные**: Булево значение true, если хеши одинаковые, иначе false
//
// **Ограничения**:
// - Хеши могут быть любой длины
// - Пустые строки считаются разными хешами
// - Сравнение должно быть точным (без декодирования)
// - Используйте простое строковое сравнение
//
// **Примеры**:
// Input: "$2a$10$abcdef123456", "$2a$10$abcdef123456"
// Output: true
//
// Input: "$2a$10$abcdef123456", "$2a$10$xyz789012345"
// Output: false

func CompareHashes(hash1, hash2 string) bool {
	return hash1 == hash2
}

func main() {
	// Тестовые данные
	hash1 := "$2a$10$abcdef123456"
	hash2 := "$2a$10$abcdef123456"
	hash3 := "$2a$10$xyz789012345"

	fmt.Println(CompareHashes(hash1, hash2))
	fmt.Println(CompareHashes(hash1, hash3))
}
