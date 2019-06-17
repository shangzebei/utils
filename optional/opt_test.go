package optional

import (
	"errors"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(OfErrorable(a()).Get())

}

type B struct {
}
type A struct {
	bb *B
}

func TestOfNilableOnce(t *testing.T) {
	OfErrorable(nil, errors.New("error")).OfError(func(e error) {
		fmt.Println("aaaaa")
	}).Then(func(i interface{}) interface{} {
		return i
	})
	//OutPut aaaaa
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
	//OutPut err last value nul
}

func TestNil1(t *testing.T) {
	OfErrorable(nil, nil).Then(func(i interface{}) interface{} {
		fmt.Println("not invoke", i)
		return i
	})
}
