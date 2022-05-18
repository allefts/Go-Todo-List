package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var scanner *bufio.Scanner  = bufio.NewScanner(os.Stdin)

func main() {
	myTodoList := createTodoList()
	
	fmt.Println("Allef's Todo List: ")
	userChoice := ""

	for userChoice != "6" {
		fmt.Printf("\n1.) Add a Todo Item\n")
		fmt.Printf("2.) View Todos\n")
		fmt.Printf("3.) Edit a Todo Item\n")
		fmt.Printf("4.) Remove a Todo Item\n")
		fmt.Printf("5.) Complete a Todo Item\n")
		fmt.Printf("6.) Exit\n")
		fmt.Printf("Input: ")
		scanner.Scan()
		userChoice = scanner.Text()
		fmt.Printf("\n")

		switch userChoice {
		case "1": myTodoList.handleAdd()
		case "2": myTodoList.viewTodos()
		case "3": myTodoList.editTodos()
		case "4": myTodoList.handleDelete()
		case "5": myTodoList.handleComplete()
		case "6": //Exit 
		default: fmt.Printf("PLEASE CHOOSE A VALID INPUT\n") 
		}
	}
}	

func (t* todoList) handleAdd() {
	fmt.Printf("Adding a todo....\n")
	fmt.Printf("Enter your task: ")
	scanner.Scan()
	newTask := scanner.Text()
	fmt.Printf("Enter the date and time due(YYYY-MM-DD HH:MM): ")
	scanner.Scan()
	newTaskDateDue := scanner.Text()

	if newTask == "" || newTaskDateDue == "" {
		fmt.Printf("\nPLEASE ENTER A TASK AND DATE DUE\n")
		return
	} else {
		parsedDateDue, err := time.Parse("2006-01-02 15:04", newTaskDateDue)

		if err != nil {
			fmt.Printf("\nPLEASE ENTER A VALID DATE\n")
			return
		}
		
		//Backend add
		t.addTodoItem(newTask, parsedDateDue)
		fmt.Printf("New task added: %s due on %v, TASKID: %d\n", newTask , parsedDateDue.Format(time.RFC850), t.currentTodoCount-1)
	}
}

func (t* todoList) viewTodos() {
	fmt.Printf("Current Todos: \n")
	fmt.Printf("%-25s %-25s %-25s %-25s\n", "ID", "Task", "Status", "Due")
	for todoID, todoItem := range(t.todos) {
		fmt.Printf("%-25d %-25s %-25t %-25v\n", todoID, todoItem.task, todoItem.completed, todoItem.dateDue)
	}
}

func (t* todoList) editTodos() {
	t.viewTodos()
	fmt.Printf("Enter the ID of the task you would like to change: ")
	scanner.Scan()
	taskIDStr := scanner.Text()
	taskID, err := strconv.Atoi(taskIDStr) 
	if err != nil {
		fmt.Printf("PLEASE ENTER A VALID ID")
		return
	}

	taskToEdit, notFound := t.todos[int64(taskID)]

	if notFound != true {
		fmt.Printf("COULD NOT FIND ID")
	} else {
		fmt.Printf("Enter your new task, or press Enter if you don't want to change: ")
		scanner.Scan()
		newTask := scanner.Text() 

		fmt.Printf("Enter your new date(YYYY-MM-DD HH:MM), or press Enter if you don't want to change: ")
		scanner.Scan()
		newDate := scanner.Text()

		if checkEmpty(newTask) == true {
			if checkEmpty(newDate) == true {
				//Task Empty and Date Empty
				return
			} else {
				//Task Empty and Date Changed
				parsedNewDate, err := time.Parse("2006-01-02 15:04", newDate)
				taskToEdit.dateDue = parsedNewDate

				if err != nil {
					fmt.Printf("\nPLEASE ENTER A VALID DATE\n")
					return
				}
			}
		} else {
			if checkEmpty(newDate) == true {
				//Task Changed and Date Empty
				taskToEdit.task = newTask
			} else {
				//Task Changed and Date Changed
				taskToEdit.task = newTask
				parsedNewDate, err := time.Parse("2006-01-02 15:04", newDate)
				taskToEdit.dateDue = parsedNewDate

				if err != nil {
					fmt.Printf("\nPLEASE ENTER A VALID DATE\n")
					return
				}
			}
		}
		t.todos[int64(taskID)] = taskToEdit
	}
	
}

func checkEmpty(dateOrtask string) bool {
	if dateOrtask == "" {
		return true 
	} else {
		return false
	}
}

func (t* todoList) handleDelete() {
	fmt.Printf("Enter the ID of the task you would like to delete: ")
	scanner.Scan()
	taskIDStr := scanner.Text()
	taskID, err := strconv.Atoi(taskIDStr) 

	if err != nil {
		fmt.Printf("PLEASE ENTER A VALID ID")
		return
	}

	_, notFound := t.todos[int64(taskID)]

	if notFound != true {
		fmt.Printf("COULD NOT FIND ID")
	} else {
		t.removeTodo(int64(taskID))
	}
}

func (t* todoList) handleComplete() {
	fmt.Printf("Enter the ID of the task you would like to delete: ")
	scanner.Scan()
	taskIDStr := scanner.Text()
	taskID, err := strconv.Atoi(taskIDStr) 

	if err != nil {
		fmt.Printf("PLEASE ENTER A VALID ID")
		return
	}

	_, notFound := t.todos[int64(taskID)]

	if notFound != true {
		fmt.Printf("COULD NOT FIND ID")
	} else {
		t.completeTodo(int64(taskID))
	}
}