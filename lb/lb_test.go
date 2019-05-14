package lb

import (
	"fmt"
	"testing"
)

func TestLb(t *testing.T) {
	poll := &LBPoll{}
	for i := 0; i < 1000; i++ {
		fmt.Println(poll.SelectOne([]string{"A", "b", "c", "d", "e", "f"}))
	}
}

func TestLbRand_SelectOne(t *testing.T) {
	poll := &LbRand{}
	for i := 0; i < 1000; i++ {
		fmt.Println(poll.SelectOne([]string{"a", "b", "c", "d", "e", "f"}))
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
