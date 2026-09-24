package taskFunc

import (
	"fmt"
	DB "servInTerm/createDb"
)

func ChangeTask() {

	width := BoxTop("ИЗМЕНЕНИЕ ЗАДАЧИ", Cyan)

	taskId, ok := AskID(width, Cyan)
	if !ok {
		return
	}

	oldTask, err := DB.GetTasksOldParams(taskId)
	if err != nil {
		BoxLine(width, Cyan, Red, fmt.Sprintf("Задача с ID %d не найдена.", taskId))
		BoxBottom(width, Cyan)
		return
	}

	title := ReadLineWithDefault(BoxPrompt(Cyan, "Новый заголовок: "), oldTask.Title)
	content := ReadLineWithDefault(BoxPrompt(Cyan, "Новое описание: "), oldTask.Content)
	priority := ReadLineWithDefault(BoxPrompt(Cyan, "Новый приоритет: "), oldTask.Priority)

	if err := DB.ChangeTask(title, content, priority, taskId); err != nil {
		BoxLine(width, Cyan, Red, fmt.Sprintf("Обновление отклонено: %v", err))
		BoxBottom(width, Cyan)
		return
	}

	BoxLine(width, Cyan, Green, fmt.Sprintf("Задача %d обновлена.", taskId))
	BoxBottom(width, Cyan)
}