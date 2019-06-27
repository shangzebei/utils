package retry

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"runtime/debug"
	"syscall"
	"time"
)

type Retry struct {
	f       func(interface{}, error)
	vs      []reflect.Value
	times   int
	limits  int
	ps      []reflect.Value
	sleep   time.Duration
	isExit  bool
	isSleep bool
}

func NewRetry(limit int) *Retry {
	return &Retry{limits: limit, sleep: time.Second, isExit: true, isSleep: true}
}

func NewNoExitRetry(limit int) *Retry {
	return &Retry{limits: limit, sleep: time.Second, isExit: false, isSleep: true}
}

func NewNoSleepRetry(limit int) *Retry {
	return &Retry{limits: limit, sleep: time.Second, isExit: true, isSleep: false}
}

func NewNoSleepNoExitRetry(limit int) *Retry {
	return &Retry{limits: limit, sleep: time.Second, isExit: false, isSleep: false}
}

func (retry *Retry) call(f interface{}) {
	retry.vs = reflect.ValueOf(f).Call(retry.ps)
}

func (retry *Retry) Get(f interface{}, param ...interface{}) []interface{} {
	retry.getParams(param...)
	retry.call(f)
	for true {
		getError := retry.getError(retry.vs)
		if getError != nil {

			logrus.Warnf("[Retry find error] %s", getError.Error())
			logrus.Trace(string(debug.Stack()))
			retry.times++
			if retry.times > retry.limits {
				if retry.isExit {
					debug.PrintStack()
					process, _ := os.FindProcess(os.Getpid())
					process.Signal(syscall.SIGINT)
				} else {
					return retry.getVale()
				}
			}
			retry.call(f)
			retry.sleep = time.Duration(retry.times) * time.Second

		} else {
			return retry.getVale()
		}
		if retry.isSleep {
			logrus.Infof("try times %d and sleep %ds", retry.times, retry.sleep/time.Second)
			time.Sleep(retry.sleep)
		} else {
			logrus.Infof("try times %d", retry.times)
		}
	}
	return nil
}

func (retry *Retry) getParams(param ...interface{}) {
	if param != nil {
		for _, value := range param {
			retry.ps = append(retry.ps, reflect.ValueOf(value))
		}
	}

}

func (retry *Retry) getVale() []interface{} {
	var res []interface{}
	for _, value := range retry.vs {
		res = append(res, value.Interface())
	}
	return res
}

func (retry *Retry) getError(v []reflect.Value) error {
	err := v[len(v)-1]
	if reflect.TypeOf(err).Kind() == reflect.Struct {
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		if err.Type().Implements(errorInterface) {
			if err.Interface() != nil {
				return err.Interface().(error)
			} else {
				return nil
			}

		}

	}
	panic("error NewRetry")
	return nil
}
