package signalx

import (
	"os"
	"os/signal"
)

func HandSignal(f func(), sigs ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, sigs...)
	go func() {
		<-c
		f()
	}()
}
