package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Add(1, 2)
		fmt.Println(got) // output: 3
	})
}

func ExampleAdd() {
	got := Add(1, 2)
	fmt.Println(got) // output: 3
}
