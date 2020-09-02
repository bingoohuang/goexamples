package textworker

import "fmt"

type Worker struct{}

func (w Worker) Do() {
	fmt.Println("I am working")
}
