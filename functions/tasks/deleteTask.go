package taskFunc

import (
	"fmt"
	DB "servInTerm/createDb"
	"strings"
)

func DeleteTask() {
	
	width := BoxTop("УДАЛЕНИЕ ЗАДАЧИ", Red)
	BoxLine(width, Red, Yellow, "Внимание: действие необратимо.")

	taskId, ok := AskID(width, Red)
	if !ok {
		return
	}

	confirm := ReadLine(BoxPrompt(Red, fmt.Sprintf("Удалить задачу %d? (y/n): ", taskId)))
	if strings.ToLower(confirm) != "y" {
		BoxLine(width, Red, Gray, "Отменено, ничего не удалено.")
		BoxBottom(width, Red)
		return
	}

	if err := DB.DeleteTask(taskId); err != nil {
		BoxLine(width, Red, Red, fmt.Sprintf("Удаление отклонено: %v", err))
		BoxBottom(width, Red)
		return
	}

	BoxLine(width, Red, Green, fmt.Sprintf("Задача %d удалена.", taskId))
	BoxBottom(width, Red)
}
