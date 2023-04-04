package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/fabianmolinab/crud-cli/tasks"
)

func main() {
	// Lea o crea tasks.json
	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	// defer: se usa para asegurarse que algunas operaciones se realicen antes
	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		// Lee el archivo y lo debuelve como un slice bits
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		// Deserializa el JSON en una variable tasks, que una slice de tasks
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		// Si esta vacio devuelve una tarea vacia
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		// Lee la entrada del usuario desde la linea de comandos
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es tu tarea?")
		name, err := reader.ReadString('\n')
		// Manejamos el error
		if err != nil {
			panic(err)
		}
		// Elimina los espacios en blanco
		name = strings.TrimSpace(name)

		// Agrega la tarea
		tasks = task.AddTask(tasks, name)

		// Guarda la tarea
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar el id de la tarea")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}
		tasks = task.DeleteTasks(tasks, id)
		task.SaveTasks(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar el id de la tarea")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)

		fmt.Println("Su tarea marcada como completada")
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}
