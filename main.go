package main

import "time"

type Todo struct {
	Description string
	Done        bool
	StartData   time.Time
	EndData     time.Time
}

func main() {
	println("Bom vindo ao Gerenciador de tarefas!")

	var todo []Todo
}
