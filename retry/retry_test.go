package retry

import (
	"errors"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(NewNoSleepNoExitRetry(5).Get(Do, 1))
}

var kk int

func Do(int2 int) (int64, int64, error) {
	fmt.Println("do", int2)
	kk++
	if kk == 3 {
		return 1, 2, nil
	}
	return 0, 1, errors.New("aaaaaaa")
}

func TestFuncName(t *testing.T) {

}
