package main

import (
	"fmt"

	"github.com/ajaypp123/goutils/collections"
)

func main() {
	li := collections.LinkedList{}

	li.InsertAtEnd(5)

	data := li.GetFirst().(int)
	fmt.Println(data)
}
