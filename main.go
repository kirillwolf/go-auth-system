package main

import (
	"fmt"
	"strings"
)

func main() {
	var name, pass string
	var age int

	//const admiLog = "admin"
	//const admiPass = "1122"

	isAdmin := false

	for {
		fmt.Println("\n1. User")
		fmt.Println("2. Admin")

		if isAdmin {
			fmt.Println("3. Secret admin panel")
			fmt.Println("4. Exit")
		} else {
			fmt.Println("3. Exit")
		}

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nHello User")

			fmt.Println("What is your name?")
			fmt.Scan(&name)

			if !isValidName(name) {
				fmt.Println("Ошибка: логин должен содержать только буквы.")
				continue
			}

			fmt.Println("How old are you?")
			fmt.Scan(&age)

			if age < 18 {
				fmt.Println("Доступ запрещён: нужен возраст 18+")
				continue
			} //более новый вариант

			fmt.Printf("У вас есть доступ, %s, %d лет\n", name, age)

		case 2:
			// Получение возраста
			age, err := getAge()
			if err != nil {
				fmt.Println("ошибка:", err)
				continue
			}
			if age < 18 {
				fmt.Println("Доступ запрещён: нужен возраст 18+")
				continue
			} //более новый вариант

			/*fmt.Println("How old are you?")
			fmt.Scan(&age)

			if age < 18 {
				fmt.Println("Доступ запрещён: нужен возраст 18+")
				continue
			}*/

			// Ввод логина
			fmt.Println("Введите Логин:")
			fmt.Scan(&name)

			// Проверка на допустимость логина
			if !isValidName(name) {
				fmt.Println("Ошибка: логин должен содержать только буквы.")
				continue
			}

			/*if !isValidName(name) && name != "admin" && name != "Админ" && name != "админ" {
				fmt.Println("Ошибка: логин должен содержать только буквы.")
				continue
			}*/

			// Ввод пароля
			fmt.Println("Введите Пароль:")
			fmt.Scan(&pass)

			// Проверка логина и пароля по карте Users
			savedPass, ok := Users[name]
			if !ok {
				fmt.Println("Пользователь не найден.")
			} else if pass != savedPass {
				fmt.Println("Неверный пароль.")
				continue

			}

			fmt.Println("Добро пожаловать,", name)
			/*else {

				fmt.Println("Добро пожаловать,", name)



				if name == "admin" || name == "Admin" || name == "Админ" || name == "админ" {
					isAdmin = true // сброс перед каждым входом
					fmt.Println("Hi admin")
					continue // ← ДОБАВИТЬ это

				}
			}*/

			// Проверка, админ ли это
			if strings.ToLower(name) == "admin" {
				isAdmin = true
				fmt.Println("Hi admin")
			}

		case 3:
			if isAdmin {
				fmt.Println("У вас есть доступ к секретной панели!")
				// Здесь можно добавить команды администратора
			} else {
				fmt.Println("Выход из программы")
				// здесь ничего не нужно — меню появится снова само
				// continue — чтобы вернуться в начало меню

				return // завершаем main()

			}

		case 4:
			if isAdmin {
				fmt.Println("Выход администратора")
				//return // ← обязательно!(Как правильно:)
			} else {
				fmt.Println("Выход пользователя")
				//return // ← тоже остаётся(Как правильно:)
			}
			return // ← общий для обоих(Или проще)
		default:
			fmt.Println("Не сделан выбор")
		}
	}

}

/*(Старый вариант без map(string))if name == admiLog && pass == admiPass {
	isAdmin = true
	fmt.Println("Hi admin")
} else {
	fmt.Println("Отказано в доступе")
	continue
}*/
