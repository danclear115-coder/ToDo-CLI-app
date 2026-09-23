package taskFunc

import (
	"fmt"
	database "servInTerm/createDb"
)

func CompleteTask() {
	
	width := BoxTop("ПЕРЕКЛЮЧЕНИЕ СТАТУСА", Green)

	taskId, ok := AskID(width, Green)
	if !ok {
		return
	}

	if err := database.CompleteTask(taskId); err != nil {
		BoxLine(width, Green, Red, fmt.Sprintf("Не удалось изменить статус: %v", err))
		BoxBottom(width, Green)
		return
	}

	BoxLine(width, Green, Green, fmt.Sprintf("Статус задачи %d переключён.", taskId))
	BoxBottom(width, Green)
}
