package mathlib_test

import (
	"testing"

	"github.com/ajaypp123/goutils"
)

func TestAdd(t *testing.T) {
	x, y := 10, 5
	expected := 15
	result := goutils.AddTwoNum(x, y)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", x, y, result, expected)
	}
}
