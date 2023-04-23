package test

import (
	"testing"

	"github.com/ajaypp123/goutils/collections"
)

func TestList(t *testing.T) {
	li := collections.LinkedList{}

	li.InsertAtEnd(5)

	data := li.GetFirst().(int)
	if data != 5 {
		t.Errorf("Failed")
	}
}
