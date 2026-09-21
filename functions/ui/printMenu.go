package uiFunc

import (
	"fmt"
	taskFunc "servInTerm/functions/tasks"
	"strings"
)

func menuItem(width int, key, text string) {
	content := fmt.Sprintf("%s -> %s", key, text)
	fmt.Printf("%s│%s %s%s%s %s│%s\n",
		taskFunc.Cyan, taskFunc.Reset,
		taskFunc.Green, taskFunc.Pad(content, width-2), taskFunc.Reset,
		taskFunc.Cyan, taskFunc.Reset)
}

func PrintMenu() {

	w := taskFunc.FormWidth()

	fmt.Printf("\n%s┌%s┐%s\n", taskFunc.Cyan, strings.Repeat("─", w), taskFunc.Reset)
	menuItem(w, "0", "Показать список задач")
	menuItem(w, "1", "Создать задачу")
	menuItem(w, "2", "Изменить задачу")
	menuItem(w, "3", "Переключить статус")
	menuItem(w, "4", "Удалить задачу")
	menuItem(w, "5", "Показать активные задачи")
	menuItem(w, "6", "Выход")
	fmt.Printf("%s└%s┘%s\n", taskFunc.Cyan, strings.Repeat("─", w), taskFunc.Reset)
}
