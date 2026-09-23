package taskFunc

import (
	"fmt"
	database "servInTerm/createDb"
)

func GetUncompletedTasks() {
	tasks, err := database.GetUncompletedTasks()
	if err != nil {
		fmt.Printf("%s[!] Ошибка БД: %v%s\n\n", Red, err, Reset)
		return
	}
	renderTaskList(tasks, "АКТИВНЫЕ ЗАДАЧИ", "Незавершенных задач нет")
}
