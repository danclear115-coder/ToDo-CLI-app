package taskFunc

import (
	"fmt"
	"servInTerm/auth"
)

const maxLoginAttempts = 3

func RequireAuth() bool {

	if !auth.IsConfigured() {
		return setupPassword()
	}

	return login()

}

func setupPassword() bool {

	width := BoxTop("ПЕРВЫЙ ЗАПУСК — ЗАДАЙТЕ ПАРОЛЬ", Green)
	BoxLine(width, Green, Gray, "Этот пароль будет запрашиваться при каждом входе в приложение.")

	for {

		pass := ReadPassword(BoxPrompt(Green, "Новый пароль: "))
		if pass == "" {
			BoxLine(width, Green, Red, "Пароль не может быть пустым.")
			continue
		}

		confirm := ReadPassword(BoxPrompt(Green, "Повторите пароль: "))
		if pass != confirm {
			BoxLine(width, Green, Red, "Пароли не совпадают, попробуйте ещё раз.")
			continue
		}

		if err := auth.SetPassword(pass); err != nil {
			BoxLine(width, Green, Red, fmt.Sprintf("Не удалось сохранить пароль: %v", err))
			BoxBottom(width, Green)
			return false
		}

		BoxLine(width, Green, Green, "Пароль сохранён. Запомните его — восстановить будет нельзя.")
		BoxBottom(width, Green)
		return true

	}
}

func login() bool {

	width := BoxTop("АВТОРИЗАЦИЯ", Cyan)

	for attempt := 1; attempt <= maxLoginAttempts; attempt++ {

		pass := ReadPassword(BoxPrompt(Cyan, "Пароль: "))

		ok, err := auth.Check(pass)
		if err != nil {
			BoxLine(width, Cyan, Red, fmt.Sprintf("Ошибка проверки пароля: %v", err))
			BoxBottom(width, Cyan)
			return false
		}
		if ok {
			BoxBottom(width, Cyan)
			return true
		}

		left := maxLoginAttempts - attempt
		if left > 0 {
			BoxLine(width, Cyan, Red, fmt.Sprintf("Неверный пароль. Осталось попыток: %d.", left))
		}

	}

	BoxLine(width, Cyan, Red, "Слишком много неверных попыток. Доступ закрыт.")
	BoxBottom(width, Cyan)
	return false
	
}