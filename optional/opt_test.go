package optional

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(OfErrorable(a()).Get())

}

func a() (string, error) {
	return "aaaaaa", nil
}
func b() (string, string) {
	return "aaaaaa", "bbbb"
}

func TestNil(t *testing.T) {
	OfErrorable(nil, nil).
		OfError(func(e error) {
			fmt.Println("err", e)
		}).Then(func(i interface{}) interface{} {
		fmt.Println("one", i)
		return i
	})
}

func TestNil1(t *testing.T) {
	OfErrorable(nil, nil).
		Then(func(i interface{}) interface{} {
			fmt.Println("one", i)
			return i
		})
}
