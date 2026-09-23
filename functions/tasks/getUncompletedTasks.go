package taskFunc

import (
	"fmt"
	DB "servInTerm/createDb"
)

func GetUncompletedTasks() {

	tasks, err := DB.GetUncompletedTasks()
	if err != nil {
		fmt.Printf("%s[!] Ошибка БД: %v%s\n\n", Red, err, Reset)
		return
	}
	
	renderTaskList(tasks, "АКТИВНЫЕ ЗАДАЧИ", "Незавершенных задач нет")
}
