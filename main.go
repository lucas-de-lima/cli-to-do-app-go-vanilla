package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Description string
	Done        bool
	StartData   time.Time
	EndData     time.Time
}

func main() {
	println("Bem vindo ao Gerenciador de tarefas!")

	var todoList []Todo

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Descrição da tarefa")
	if scanner.Scan() {
		description := strings.TrimSpace(scanner.Text())
		if description != "" {
			newTodo := Todo{
				Description: description,
				Done:        false,
				StartData:   time.Now(),
			}
			todoList = append(todoList, newTodo)
		}
	}

	fmt.Println("\nTarefas registradas:")
	for i, t := range todoList {
		fmt.Printf("%d. %s (Iniciada em %s)\n", i+1, t.Description, t.StartData.Format("02/01/2006 15:04:05"))
	}
}
