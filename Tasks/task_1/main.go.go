package task_2

import "fmt"

type Human struct {
	Name string
	Age  int
	Job  string
}

func (h *Human) Work() {
	fmt.Printf(`Working %s...`, h.Job)
}

func (h *Human) Say(word string) {
	fmt.Println(word)
}

type Action struct {
	Human
}

func Run() {
	human := Human{Name: "Jake", Age: 22, Job: "teacher"}

	action := Action{human}

	action.Work()
}
