package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task // slice of type task
}

type task struct {
	name        string
	description string
	completed   bool
}

func (l *taskList) addTask(t *task) {
	l.tasks = append(l.tasks, t)
}

func (l *taskList) deleteTask(i int) {
	l.tasks = append(l.tasks[:i], l.tasks[i+1:]...) // ... allows execute operatio
}

func (l *taskList) printList() {
	for _, element := range l.tasks {
		fmt.Println(element.name)
	}
}

func (t *task) masrkAsComplete() { // receiver function, *x value of direction memory of x not copy becasuse functions make copies of every variable that use
	t.completed = true
}

func (t *task) setDescription(description string) {
	t.description = description
}

func (t *task) setName(name string) { // *x, value of reference address x
	t.name = name
}

func main() {

	t := &task{
		name:        "Proof 1",
		description: "Description of proof 1",
		completed:   false,
	}

	t2 := &task{
		name:        "Proof 2",
		description: "Description of proof 2",
		completed:   false,
	}

	fmt.Println(&t)         // memory address
	fmt.Printf("%+v\n", *t) // showing and interface key:value
	t.masrkAsComplete()
	fmt.Printf("%+v\n", *t)

	list := &taskList{
		tasks: []*task{ // comma for end structure code
			t,
		},
	}

	//fmt.Println(list.tasks[0])
	list.addTask(t2)
	// list.deleteTask(1)
	fmt.Println(len(list.tasks))

	for i := 0; i < len(list.tasks); i++ {
		fmt.Println(list.tasks[i])
	}

	for _, element := range list.tasks { // for index, element := range
		fmt.Println(element)
	}

	// break: breaks the loop, continue breaks in an especific index
}
