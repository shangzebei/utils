package arrays

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := []string{"A", "B"}
	b := make([]string, 1)
	b[0] = "A"
	for _, value := range b {
		fmt.Println("@", value)
	}
	fmt.Println(Of(a).Add("C").Remove("A").Out())
	fmt.Println(Of(b).Add("C").Add("W").Remove("A").Out())

}
