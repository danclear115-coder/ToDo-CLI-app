package main

import (
	database "servInTerm/createDb"
	taskFunc "servInTerm/functions/tasks"
	uiFunc "servInTerm/functions/ui"
)

func main() {

	database.InitDB()

	if !taskFunc.RequireAuth() {
		return
	}
	
	uiFunc.PrintMenu()

	for {	

		choice := taskFunc.ReadLine("root@tasks:~# ")
		uiFunc.UserChoice(choice)

	}
	
}
