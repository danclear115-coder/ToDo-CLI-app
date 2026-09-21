package taskFunc

import (
	"fmt"
	database "servInTerm/createDb"
)

func GetTasks() {
	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Printf("%s[!] Ошибка БД: %v%s\n\n", Red, err, Reset)
		return
	}
	renderTaskList(tasks, "СПИСОК ПУСТ", "В базе пока нет ни одной задачи.")
}
