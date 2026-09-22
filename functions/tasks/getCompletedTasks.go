package taskFunc

import (
	"fmt"
	Db "servInTerm/createDb"
)

func GetCompletedTasks() {

	tasks, err := Db.GetCompletedTasks()

	if err != nil {
		fmt.Printf("%s[!] Ошибка БД: %v%s\n\n", Red, err, Reset)
		return
	}

	renderTaskList(tasks, "СПИСОК ПУСТ", "В базе пока нет ни одной выполненной задачи.")

}
