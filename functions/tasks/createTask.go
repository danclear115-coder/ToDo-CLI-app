package taskFunc

import (
	"fmt"
	DB "servInTerm/createDb"
)

func CreateTask() {
	
	width := BoxTop("НОВАЯ ЗАДАЧА", Green)

	title := ReadLine(BoxPrompt(Green, "Заголовок: "))
	content := ReadLine(BoxPrompt(Green, "Описание: "))
	priority := ReadLine(BoxPrompt(Green, "Приоритет (High/Medium/Low): "))

	if err := DB.CreateTask(title, content, priority); err != nil {
		BoxLine(width, Green, Red, fmt.Sprintf("Не удалось создать: %v", err))
		BoxBottom(width, Green)
		return
	}

	BoxLine(width, Green, Green, "Задача создана.")
	BoxBottom(width, Green)
}
