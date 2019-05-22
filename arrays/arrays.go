package arrays

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type Arrays struct {
	arr interface{}
}

func Of(ar interface{}) *Arrays {
	if reflect.TypeOf(ar).Kind() != reflect.Slice {
		logrus.Error("of input is not slice")
		return nil
	}
	return &Arrays{arr: ar}
}

func (ar *Arrays) Add(v interface{}) *Arrays {
	reflect.ValueOf(&ar.arr).Elem().Set(reflect.Append(reflect.ValueOf(ar.arr), reflect.ValueOf(v)))
	return ar
}

func (ar *Arrays) Out() interface{} {
	return ar.arr

}

func (ar *Arrays) Remove(v interface{}) *Arrays {
	valueOf := reflect.ValueOf(ar.arr)
	for i := 0; i < valueOf.Len(); i++ {
		in := valueOf.Index(i)
		if in.Interface() == v {
			reflect.ValueOf(&ar.arr).Elem().Set(reflect.AppendSlice(valueOf.Slice(0, i),
				valueOf.Slice(i+1, valueOf.Len())))
		}
	}
	return ar
}

func (ar *Arrays) Strings() []string {
	return ar.arr.([]string)
}

func (ar *Arrays) Ints() []int64 {
	return ar.arr.([]int64)
}
