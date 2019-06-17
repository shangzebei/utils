package optional

import (
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime/debug"
)

type Optional struct {
	//tv   interface{}
	fv   interface{}
	ev   error
	tagV map[int]interface{}
	ef   func(error)
	pass bool
}

var debugFlags = false

func OfNilable(t interface{}) *Optional {
	return &Optional{fv: t, tagV: make(map[int]interface{}), pass: false}
}

func (o *Optional) isNil(f interface{}) bool {
	if f == nil {
		if !o.pass {
			logrus.Tracef("nil check stack =%s", debugInfo())
		}
		o.pass = true
		return true
	}
	switch reflect.TypeOf(f).Kind() {
	case reflect.Ptr, reflect.Map:
		isnil := reflect.ValueOf(f).IsNil()
		if isnil {
			if !o.pass {
				logrus.Tracef("nil check stack =%s", debugInfo())
			}
			o.pass = true
		}
		return isnil
	case reflect.String:
		return f.(string) == ""
	default:

	}
	return false

}

func isErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func SetGlobDebug(f bool) {
	debugFlags = f
}

func Of(f func() interface{}) *Optional {
	return &Optional{fv: f(), tagV: make(map[int]interface{}), pass: false}
}

func (o *Optional) Then(f func(interface{}) interface{}) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		o.fv = f(o.fv)
	}
	return o
}

func debugInfo() string {
	if debugFlags {
		return string(debug.Stack())
	}
	return ""
}

func (o *Optional) ThenE(f func(interface{}) (interface{}, error)) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			return o
		}
	}
	return o
}

func (o *Optional) SetError(f func(error)) *Optional {
	o.ef = f
	return o
}

func (o *Optional) OfError(f func(error)) *Optional {
	if o.tagV == nil {
		o.tagV = make(map[int]interface{})
	}
	o.ef = f
	if isErr(o.ev) && o.ef != nil {
		f(o.ev)
		o.ef = nil
		return o
	}
	if o.isNil(o.fv) {
		f(errors.New("last value nul"))
		o.ef = nil
	}
	return o
}

func (o *Optional) ThenSet(tag int, f func(interface{}) interface{}) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		o.fv = f(o.fv)
		if o.isNil(o.fv) {
			o.error(errors.New("ThenSet return nil point"))
			return o
		}
		o.tagV[tag] = o.fv
	}
	return o
}

func (o *Optional) ThenSetE(tag int, f func(interface{}) (interface{}, error)) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			return o
		}
		if o.isNil(o.fv) {
			o.error(errors.New("ThenSet return nil point"))
			return o
		}
		o.tagV[tag] = o.fv
	}
	return o
}

func (o *Optional) error(err error) {
	o.pass = true
	logrus.Tracef("error %s stack %s", err.Error(), debugInfo())
	if o.ef != nil {
		o.ef(err)
	}
}

func (o *Optional) ThenGet(f func(interface{}) interface{}, tag ...int) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv = f(kk)
		if o.isNil(o.fv) {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	}
	return o
}

func (o *Optional) ThenGetE(f func(interface{}) (interface{}, error), tag ...int) *Optional {
	if o.pass {
		return o
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv, o.ev = f(kk)
		if o.isNil(o.fv) {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	}
	return o
}

func OfErrorable(a interface{}, err error) *Optional {
	if err != nil {
		logrus.Trace(err.Error())
	}
	return &Optional{ev: err, fv: a, tagV: make(map[int]interface{}), pass: false}
}

func (o *Optional) IsPrent() bool {
	switch o.fv.(type) {
	case string:
		if o.fv.(string) == "" {
			logrus.Tracef("IsPrent warn return type string = '' stack= %s", debugInfo())
		}
	default:

	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		return true
	} else {
		return false
	}
}

func (o *Optional) IfPrent(f func(interface{})) {
	if isErr(o.ev) && o.ef != nil {
		o.error(o.ev)
		return
	}
	if !o.isNil(o.fv) && !isErr(o.ev) {
		f(o.fv)
	}
}

func (o *Optional) OrElseGet(f func() interface{}) interface{} {
	if o.isNil(o.fv) && isErr(o.ev) {
		return f()
	} else {
		return o.fv
	}
}
func (o *Optional) Get() interface{} {
	return o.fv
}
