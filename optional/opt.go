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
}

var debugFlags = false

func OfNilable(t interface{}) *Optional {
	return &Optional{fv: t, tagV: make(map[int]interface{})}
}

func isNil(f interface{}) bool {
	if reflect.TypeOf(f).Kind() == reflect.Ptr {
		return f == nil || reflect.ValueOf(f).IsNil()
	}
	return f == nil

}

func SetGlobDebug(f bool) {
	debugFlags = f
}

func Of(f func() interface{}) *Optional {
	return &Optional{fv: f(), tagV: make(map[int]interface{})}
}

func (o *Optional) Then(f func(interface{}) interface{}) *Optional {
	if !isNil(o.fv) && o.ev == nil {
		o.fv = f(o.fv)
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
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
	if !isNil(o.fv) && o.ev == nil {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
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
	if o.ev != nil && o.ef != nil {
		f(o.ev)
	}
	if isNil(o.fv) {
		f(errors.New("last value nul"))
	}
	return o
}

func (o *Optional) ThenSet(tag int, f func(interface{}) interface{}) *Optional {

	if !isNil(o.fv) && o.ev == nil {
		o.fv = f(o.fv)
		if isNil(o.fv) {
			o.error(errors.New("ThenSet return nil point"))
			o.fv = nil
			return o
		}
		o.tagV[tag] = &o.fv
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
	}
	return o
}

func (o *Optional) ThenSetE(tag int, f func(interface{}) (interface{}, error)) *Optional {
	if !isNil(o.fv) && o.ev == nil {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			o.fv = nil
			return o
		}
		if isNil(o.fv) {
			o.error(errors.New("ThenSet return nil point"))
			o.fv = nil
			return o
		}
		o.tagV[tag] = o.fv
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
	}
	return o
}

func (o *Optional) error(err error) {
	if o.ef != nil {
		o.ef(err)
	} else {
		logrus.Tracef("Optional SetError not set,stack %s", debugInfo())
	}
}

func (o *Optional) ThenGet(f func(interface{}) interface{}, tag ...int) *Optional {
	if !isNil(o.fv) && o.ev == nil {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv = f(kk)
		if isNil(o.fv) {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
	}
	return o
}

func (o *Optional) ThenGetE(f func(interface{}) (interface{}, error), tag ...int) *Optional {
	if !isNil(o.fv) && o.ev == nil {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv, o.ev = f(kk)
		if isNil(o.fv) {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
	}
	return o
}

func OfErrorable(a interface{}, err error) *Optional {
	if err != nil {
		logrus.Trace(err.Error())
	}
	return &Optional{ev: err, fv: a, tagV: make(map[int]interface{})}
}

func (o *Optional) IsPrent() bool {
	switch o.fv.(type) {
	case string:
		if o.fv.(string) == "" {
			logrus.Tracef("IsPrent warn return type string = '' stack= %s", debugInfo())
		}
	default:

	}
	if !isNil(o.fv) && o.ev == nil {
		return true
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, debugInfo())
		return false
	}
}

func (o *Optional) IfPrent(f func(interface{})) {
	if o.ev != nil && o.ef != nil {
		o.error(o.ev)
		o.fv = nil
		return
	}
	if !isNil(o.fv) && o.ev == nil {
		f(o.fv)
	}
}

func (o *Optional) OrElseGet(f func() interface{}) interface{} {
	if isNil(o.fv) && o.ev != nil {
		return f()
	} else {
		return o.fv
	}
}
func (o *Optional) Get() interface{} {
	return o.fv
}
