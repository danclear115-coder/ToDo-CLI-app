package taskFunc

import (
	"fmt"
	database "servInTerm/createDb"
)

func ChangeTask() {
	
	width := BoxTop("ИЗМЕНЕНИЕ ЗАДАЧИ", Cyan)

	taskId, ok := AskID(width, Cyan)
	if !ok {
		return
	}

	title := ReadLine(BoxPrompt(Cyan, "Новый заголовок: "))
	content := ReadLine(BoxPrompt(Cyan, "Новое описание: "))
	priority := ReadLine(BoxPrompt(Cyan, "Новый приоритет: "))

	if err := database.ChangeTask(title, content, priority, taskId); err != nil {
		BoxLine(width, Cyan, Red, fmt.Sprintf("Обновление отклонено: %v", err))
		BoxBottom(width, Cyan)
		return
	}

	BoxLine(width, Cyan, Green, fmt.Sprintf("Задача %d обновлена.", taskId))
	BoxBottom(width, Cyan)
}
