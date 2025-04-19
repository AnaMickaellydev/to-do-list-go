package main

import (
	"fmt"
	"todo-list-go/task"
	"todo-list-go/utils"
)

func main() {
	var tasks []task.Task
	var idCounter int

	for {
		fmt.Println("\n1. Adicionar tarefa")
		fmt.Println("2. Listar tarefas")
		fmt.Println("3. Marcar tarefa como concluída")
		fmt.Println("4. Remover tarefa")
		fmt.Println("5. Sair")
		fmt.Print("Escolha uma opção: ")

		option := utils.ReadInput()

		switch option {
		case "1":

			idCounter++
			fmt.Printf("Digite o nome da tarefa:")
			name := utils.ReadInput()
			tasks = append(tasks, task.NewTask(idCounter, name))
			fmt.Println("Tarefa adicionada!")

		case "2":

			if len(tasks) == 0 {
				fmt.Println("Nenhuma tarefa cadastrada.")
			} else {
				for _, task := range tasks {
					task.Print()
				}
			}

		case "3":

			fmt.Print("Digite o ID da tarefa a ser concluída: ")
			id := utils.ReadInput()

			taskID := utils.StringToInt(id)
			found := false
			for i := range tasks {
				if tasks[i].ID == taskID {
					tasks[i].Complete = true
					fmt.Println("Tarefa concluída!")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Tarefa não encontrada!")
			}

		case "4":

			fmt.Print("Digite o ID da tarefa a ser removida: ")
			id := utils.ReadInput()
			taskID := utils.StringToInt(id)
			var updatedTasks []task.Task
			found := false
			for _, t := range tasks {
				if t.ID != taskID {
					updatedTasks = append(updatedTasks, t)
				} else {
					found = true
				}
			}
			if found {
				tasks = updatedTasks
				fmt.Println("Tarefa removida!")
			} else {
				fmt.Println("Tarefa não encontrada!")
			}

		case "5":

			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}

	}
}
