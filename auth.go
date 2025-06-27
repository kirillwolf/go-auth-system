package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Глобальная карта пользователей
var Users = map[string]string{
	"admin":  "1122",
	"Admin":  "1122",
	"админ":  "1122",
	"Админ":  "1122",
	"kirill": "1234",
}

func isValidName(s string) bool {

	// Убираем пробелы в начале и в конце
	s = strings.TrimSpace(s)

	// Проверка: пустая строка
	if len(s) == 0 {
		return false
	}
	// Преобразуем строку в массив рун (символов)
	runes := []rune(s)

	// Первый символ должен быть буквой и заглавной
	if !unicode.IsLetter(runes[0]) || !unicode.IsUpper(runes[0]) {
		return false
	}
	// Остальные символы — только буквы
	for _, r := range runes {
		if !unicode.IsLetter(r) {
			return false

		}
	}
	return true
}

func getAge() (int, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите возврат: ") // ← вот строка, которая спрашивает возраст

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения:%v", err)
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return 0, fmt.Errorf("вы ничего не ввели")
	}

	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("нужно ввести число")
	}
	if age < 0 {
		return 0, fmt.Errorf("возраст не может быть отрицательным")
	}
	return age, nil
}
