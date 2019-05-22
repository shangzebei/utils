package arrays

import (
	"fmt"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	a := []int64{1, 2}
	b := make([]string, 1)
	b[0] = "A"
	//for _, value := range b {
	//	fmt.Println("@", value)
	//}
	fmt.Println(Of(a).Add(int64(3)).Remove(1).Out())
	fmt.Println(Of(b).Add("C").Add("W").Remove("A").Out())
	//var name interface{}
	//fmt.Println(reflect.ValueOf(name).Addr().Kind())
	//reflect.ValueOf(name).Elem().Set(reflect.ValueOf(a))
	//reflect.ValueOf(&name).Elem().Set(reflect.ValueOf(a))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(Of([]string{"1"}).Add(strconv.Itoa(i)).Out())
	}
}

func TestV(t *testing.T) {
	fmt.Println(Of([]string{"1"}).Add("").Remove("1").Strings())
}
