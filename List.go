package main

import (
	"time"
)

type todoList struct {
	todos map[int64]todoItem
	currentTodoCount int64
}

type todoItem struct {
	task      string
	completed bool
	dateDue time.Time 
}

//Not Receiver because creating
func createTodoList() todoList {
	myTodoList := todoList{
		currentTodoCount: 0,
		todos: map[int64]todoItem{},
	}
	return myTodoList
}

func (t* todoList) addTodoItem(newTask string, newDateDue time.Time) {
	newTodo := todoItem {
		task: newTask, 
		completed:  false,
		dateDue: newDateDue,
	}

	//Go automatically dereferences pointers that are attributes or methods inside struct
	t.todos[t.currentTodoCount] = newTodo
	t.currentTodoCount++
	// return *t
}

func (t* todoList) updateTodo(todoID int64, changedTask string, changedDateDue time.Time) {
	//Get copy of entry
	todoItem, isValInMap := t.todos[todoID]

	//If entry found, we change the copy and then add it back 
	if isValInMap == true {
		//change copy
		todoItem.task = changedTask
		todoItem.dateDue = changedDateDue
		//add back
		t.todos[todoID] = todoItem
	}
	// return *t
}

func (t* todoList) completeTodo(todoID int64) {
	myTodo, isValidTodo := t.todos[todoID]
	if isValidTodo == true {
		myTodo.completed = true
		t.todos[todoID]= myTodo
	}
}

func (t* todoList) removeTodo(todoID int64) {
	_, isValidTodo := t.todos[todoID]
	if isValidTodo == true {
		// remove Todo
		// decrement other keys above that are above the one that was deleted
		var tilEnd int64 = int64(len(t.todos))
		delete(t.todos, todoID)
		//check if deleted was already last one in map
		if (tilEnd-1 != todoID){
			//shifts todos
			for i := todoID+1; i < tilEnd; i++ {
				myTodo, _ := t.todos[i] //i == location of todo to shift 
				shifterTodo, _ := t.todos[i-1] //i-1 == shift location
				// fmt.Println("\nmyTodo: ", myTodo, "\nshifterTodo: ", shifterTodo)
				shifterTodo = myTodo
				//complete shift
				t.todos[i-1] = shifterTodo
			}
			delete(t.todos, int64(len(t.todos)-1))
		}
		t.currentTodoCount -= 1
	}
}