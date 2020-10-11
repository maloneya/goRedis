package main

import "fmt"

type A interface {
	one()
}

type B interface {
	one()
	two()
}

type T struct {
}

func (t T) one() {
	fmt.Printf("%v", t)
}

func (t T) two() {
	fmt.Printf("%v", t)
}

func foo(i interface{}) {
	switch i.(type) {
	case A:
		fmt.Print("A\n")
	case B:
		fmt.Println("B\n")
	}

}

func bar(i interface{}) {
	switch i.(type) {
	case B:
		fmt.Print("B\n")
	case A:
		fmt.Println("A\n")
	}

}

func main() {
	t := T{}
	foo(t)
	bar(t)
}
