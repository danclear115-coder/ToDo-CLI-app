package taskFunc

import (
	"strconv"
	"strings"

	database "servInTerm/createDb"
)

var taskColumns = []Column{
	{Header: "ID", Fixed: 4},
	{Header: "ЗАГОЛОВОК", Weight: 1, MinWidth: 10},
	{Header: "ОПИСАНИЕ", Weight: 2, MinWidth: 14},
	{Header: "ПРИОРИТЕТ", Fixed: 11},
	{Header: "СТАТУС", Fixed: 11},
}

func priorityColor(priority string) string {
	switch strings.ToLower(priority) {
	case "high", "высокий":
		return Red
	case "medium", "средний":
		return Yellow
	case "low", "низкий":
		return Green
	default:
		return Gray
	}
}

func renderTaskList(tasks []database.Task, emptyTitle, emptyText string) {
	if len(tasks) == 0 {
		w := BoxTop(emptyTitle, Yellow)
		BoxLine(w, Yellow, Yellow, emptyText)
		BoxBottom(w, Yellow)
		return
	}

	rows := make([][]Cell, len(tasks))
	for i, task := range tasks {
		status, statusColor := "В РАБОТЕ", Cyan
		if task.IsCompleted {
			status, statusColor = "ВЫПОЛНЕНО", Green
		}

		rows[i] = []Cell{
			{Text: strconv.Itoa(int(task.ID)), Color: Reset},
			{Text: task.Title, Color: Reset},
			{Text: task.Content, Color: Gray},
			{Text: task.Priority, Color: priorityColor(task.Priority)},
			{Text: status, Color: statusColor},
		}
	}

	RenderTable(Green, taskColumns, rows)
}
