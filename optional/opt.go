package optional

import (
	"errors"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

type Optional struct {
	//tv   interface{}
	fv   interface{}
	ev   error
	tagV map[int]interface{}
	ef   func(error)
}

func OfNilable(t interface{}) *Optional {
	return &Optional{fv: t, tagV: make(map[int]interface{})}
}

func Of(f func() interface{}) *Optional {
	return &Optional{fv: f(), tagV: make(map[int]interface{})}
}

func (o *Optional) Then(f func(interface{}) interface{}) *Optional {
	if o.fv != nil && o.ev == nil {
		o.fv = f(o.fv)
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
	}
	return o
}

func (o *Optional) ThenE(f func(interface{}) (interface{}, error)) *Optional {
	if o.fv != nil && o.ev == nil {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
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
	if o.fv == nil {
		f(errors.New("last value nul"))
	}
	return o
}

func (o *Optional) ThenSet(tag int, f func(interface{}) interface{}) *Optional {

	if o.fv != nil && o.ev == nil {
		o.fv = f(o.fv)
		if o.fv == nil {
			o.error(errors.New("ThenSet return nil point"))
			o.fv = nil
			return o
		}
		o.tagV[tag] = &o.fv
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
	}
	return o
}

func (o *Optional) ThenSetE(tag int, f func(interface{}) (interface{}, error)) *Optional {
	if o.fv != nil && o.ev == nil {
		var err error
		o.fv, err = f(o.fv)
		if err != nil {
			o.error(err)
			o.fv = nil
			return o
		}
		if o.fv == nil {
			o.error(errors.New("ThenSet return nil point"))
			o.fv = nil
			return o
		}
		o.tagV[tag] = o.fv
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
	}
	return o
}

func (o *Optional) error(err error) {
	if o.ef != nil {
		o.ef(err)
	} else {
		logrus.Tracef("Optional SetError not set,stack %s", string(debug.Stack()))
	}
}

func (o *Optional) ThenGet(f func(interface{}) interface{}, tag ...int) *Optional {
	if o.fv != nil && o.ev == nil {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv = f(kk)
		if o.fv == nil {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
	}
	return o
}

func (o *Optional) ThenGetE(f func(interface{}) (interface{}, error), tag ...int) *Optional {
	if o.fv != nil && o.ev == nil {
		var kk []interface{}
		for _, value := range tag {
			kk = append(kk, o.tagV[value])
		}
		o.fv, o.ev = f(kk)
		if o.fv == nil {
			o.ef(errors.New("ThenGet return nil point"))
			o.fv = nil
			return o
		}
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
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
			logrus.Tracef("IsPrent warn return type string = '' stack= %s", string(debug.Stack()))
		}
	default:

	}
	if o.fv != nil && o.ev == nil {
		return true
	} else {
		logrus.Tracef("Then fv = %p and error = %p stack =%s ", o.fv, o.ev, string(debug.Stack()))
		return false
	}
}

func (o *Optional) IfPrent(f func(interface{})) {
	if o.ev != nil && o.ef != nil {
		o.error(o.ev)
		o.fv = nil
		return
	}
	if o.fv != nil && o.ev == nil {
		f(o.fv)
	}
}

func (o *Optional) OrElseGet(f func() interface{}) interface{} {
	if o.fv == nil && o.ev != nil {
		return f()
	} else {
		return o.fv
	}
}
func (o *Optional) Get() interface{} {
	return o.fv
}
