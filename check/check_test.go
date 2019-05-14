package check

import (
	"fmt"
	"testing"
	"time"
)

func TestHTTPChecker(t *testing.T) {
	//check := &Check{Up: make(chan int), Down: make(chan int)}
	//check.AddHealthCheck("http://192.168.0.137:9443/v1/management/health", time.Second*3)
	//go func() {
	//	for {
	//		<-check.Up
	//		fmt.Println("done")
	//	}
	//}()

	AddCheck("http://127.0.0.1:8081", time.Second*3, func(i int) {
		fmt.Println("aaaaa", i)
	}, func(i int) {
		fmt.Println("bbbbb", i)
	})
	time.Sleep(time.Second * 10000)
	fmt.Println("done")
}
