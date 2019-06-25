package lb

import (
	"fmt"
	"testing"
)

func TestLb(t *testing.T) {
	var l LB
	l = &LBPoll{}
	for i := 0; i < 1000; i++ {
		go fmt.Println(l.SelectOne([]string{"A", "b", "c", "d", "e", "f"}))
	}
}

func TestLbRand_SelectOne(t *testing.T) {
	poll := &LbRand{}
	for i := 0; i < 1000; i++ {
		fmt.Println(poll.SelectOne([]string{"a", "b", "c", "d", "e", "f"}))
	}
}

func TestLBPoll_SelectOne(t *testing.T) {
	poll := &LbRand{}
	for i := 0; i < 10; i++ {
		go fmt.Println(poll.SelectOne([]string{"a"}))
	}

}

func BenchmarkName(b *testing.B) {
	poll := &LBPoll{}
	for i := 0; i < b.N; i++ {
		poll.SelectOne([]string{"A", "b", "c", "d", "e", "f"})
	}
}

func BenchmarkNameB(b *testing.B) {
	poll := &LbRand{}
	for i := 0; i < b.N; i++ {
		poll.SelectOne([]string{"a", "b", "c", "d", "e", "f"})
	}
}
