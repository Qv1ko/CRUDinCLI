package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/Qv1ko/CRUDinCLI/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What's your task?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		tasks = task.AddTasks(tasks, name)
		task.SaveTasks(file, tasks)
	case "list":
		task.ListTasks(tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide an ID to complete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The ID must be a number")
			return
		}
		tasks = task.CompleteTasks(tasks, id)
		task.SaveTasks(file, tasks)
	case "incomplete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide an ID to incomplete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The ID must be a number")
			return
		}
		tasks = task.IncompleteTasks(tasks, id)
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide an ID to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The ID must be a number")
			return
		}
		tasks = task.DeleteTasks(tasks, id)
		task.SaveTasks(file, tasks)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: CRUDinCLi.exe [add|list|complete|incomplete|delete]")
}
