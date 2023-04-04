package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	Complete bool   `json: "complete"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {
		status := " "

		if task.Complete {
			status = "✓"
		}

		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	// Añade el elemento de newTask al arreglo inicial
	return append(tasks, newTask)
}

func SaveTasks(file *os.File, tasks []Task) {
	// Convierte a JSON las tareas
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// Seek se coloca al princio del archivo
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Luego borra todo el contenido del archivo
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// *NewWriter* escribe en el archivo
	writer := bufio.NewWriter(file)

	// Escribe lo que introdujo el usuario
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	// verifica que se haya escrito en el archivo
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func DeleteTasks(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			// Seleciona lo que esta antes y despues de i y lo une en un solo arreglo
			return append(tasks[:i], tasks[i+1])
		}
	}
	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
