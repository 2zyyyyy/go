package main

import "fmt"

// 方法集

type T struct {
	int
}

func (t T) tFunc() {
	fmt.Println("类型T方法集包含所有receiver T的方法.")
}

func (t *T) pFunc() {
	fmt.Println("类型*T方法集包含所有receiver *T的方法.")
}

func main() {
	t1 := T{
		100,
	}
	t2 := &t1
	fmt.Printf("t2=%v\n", t2)
	t2.tFunc()
	t2.pFunc()
}
