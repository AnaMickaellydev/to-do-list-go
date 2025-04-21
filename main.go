package main

import (
	"fmt"
	"todo-list-go/task"
	"todo-list-go/utils"
)

func main() {
	const filename = "tasks.json"

	tasks, err := task.LoadTasksFromFile(filename)
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
		return
	}

	for {
		fmt.Println("\n1. Adicionar tarefa")
		fmt.Println("2. Listar tarefas")
		fmt.Println("3. Marcar tarefa como concluída")
		fmt.Println("4. Remover tarefa")
		fmt.Println("5. Sair")

		opcao := utils.InputInt("Escolha uma opção: ")

		switch opcao {
		case 1:
			name := utils.InputString("Nome da tarefa: ")
			newTask := task.NewTask(len(tasks)+1, name)
			tasks = append(tasks, newTask)
			_ = task.SaveTasksToFile(filename, tasks)

		case 2:
			for _, t := range tasks {
				t.Print()
			}

		case 3:
			id := utils.InputInt("ID da tarefa a marcar como concluída: ")
			for i, t := range tasks {
				if t.ID == id {
					tasks[i].Complete = true
					break
				}
			}
			_ = task.SaveTasksToFile(filename, tasks)

		case 4:
			id := utils.InputInt("ID da tarefa a remover: ")
			for i, t := range tasks {
				if t.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					break
				}
			}
			_ = task.SaveTasksToFile(filename, tasks)

		case 5:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
