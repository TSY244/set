package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	data := NewSet[string]()
	data.Add("a")
	data.Add("b")
	data.Add("c")
	data.Add("d")
	data.Add("a")
	data.Add("b")
	if data.Size() != 4 {
		t.Error("[Add] Set size", data.Size())
	}
	fmt.Println("Data:", data.Values())
	data.AddAll([]string{"a", "b", "c", "d"})
	if data.Size() != 4 {
		t.Error("[AddAll] Set size", data.Size())
	}
	fmt.Println("Data:", data.Values())

	data.Clear()
	if data.Size() != 0 {
		t.Error("[Clear] Set size", data.Size())
	}

}
