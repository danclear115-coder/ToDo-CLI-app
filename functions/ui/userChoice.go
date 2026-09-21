package uiFunc

import (
	"fmt"
	taskFunc "servInTerm/functions/tasks"
)

func UserChoice(choice string) {
	
	switch choice {
		case "0":
			taskFunc.GetTasks()
		case "1":
			taskFunc.CreateTask()
		case "2":
			taskFunc.ChangeTask()
		case "3":
			taskFunc.CompleteTask()
		case "4":
			taskFunc.DeleteTask()
		case "5":
			taskFunc.GetUncompletedTasks()
		case "6":
			fmt.Printf("\n%s[!] Соединение закрыто. Пока!%s\n", taskFunc.Green, taskFunc.Reset)
			return
		default:
			fmt.Printf("\n%s[!] Неизвестная команда.%s\n", taskFunc.Red, taskFunc.Reset)
		}

}
