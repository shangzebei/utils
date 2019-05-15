package arrays

import (
	"fmt"
	"reflect"
)

type Arrays struct {
	arr interface{}
}

func Of(ar interface{}) *Arrays {
	return &Arrays{arr: ar}
}

func (ar *Arrays) Add(v interface{}) *Arrays {
	ar.arr = reflect.Append(reflect.ValueOf(ar.arr), reflect.ValueOf(v)).Interface()
	return ar
}

func (ar *Arrays) Out() interface{} {
	return ar.arr
}

func (ar *Arrays) Remove(v interface{}) *Arrays {
	valueOf := reflect.ValueOf(ar.arr)
	var temp []reflect.Value
	fmt.Println(valueOf.Len(), valueOf.Cap())
	for i := 0; i < valueOf.Len(); i++ {
		in := valueOf.Index(i)
		if in.Interface() != v {
			temp = append(temp, valueOf.Index(i))
		}
	}
	ar.arr = reflect.ValueOf(temp).Interface()
	return ar
}
