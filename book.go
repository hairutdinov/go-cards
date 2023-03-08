package main

import "fmt"

type book string

func (b book) printTitle() {
	fmt.Println(b)
}
