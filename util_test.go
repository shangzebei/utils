package util

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	a := map[string][]string{"aaaa": {"aa", "bb"}}
	RemoveMapVale("aa", a)
	fmt.Println(a)
}

func TestAdd(t *testing.T) {
	a := map[string][]string{"aaaa": {"aa", "bb"}}
	AddMapVale("aaaa", a, "ccc")
	fmt.Println(a)
}

func TestIP(t *testing.T) {
	fmt.Println(GetOutboundIP())
}

func TestMd5Bytes(t *testing.T) {
	fmt.Println(Md5Bytes([]byte("shangzebei")))
}

func TestSortSlice(t *testing.T) {
	a := &BB{3}
	b := &AA{2}
	ar := []order{b, a}
	SortSlice(ar)
	for i := 0; i < len(ar); i++ {
		fmt.Println(ar[i])
	}
}

type order interface {
	Order() int
}

type AA struct {
	A int
}

func (t *AA) Order() int {
	return t.A
}
func (*AA) String() int {
	return 1
}

type BB struct {
	A int
}

func (t *BB) Order() int {
	return t.A
}
func (*BB) String() int {
	return 2
}
