package check

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Check struct {
	Up     chan int
	Down   chan int
	status map[int]bool
}

type serverStatus struct {
	Interval time.Duration
	Id       int
	Addr     string
	Down     func(int)
	Up       func(int)
}

var pool sync.Map

func GetAddr(i int) string {
	v, b := pool.Load(i)
	if b {
		return v.(string)
	}
	return ""
}

func AddCheck(addr string, interval time.Duration, up func(int), down func(int)) int {

	t := 0
	exit := false
	//get max id if exist return
	pool.Range(func(key, value interface{}) bool {
		if key.(int) > t {
			t = key.(int)
		}
		if value == addr {
			exit = true
			t = key.(int)
			logrus.Tracef("url %s check has exist", addr)
			return false
		}
		return true
	})
	if exit {
		return t
	}
	pool.Store(t+1, addr)
	check := &Check{
		Up:     make(chan int),
		Down:   make(chan int),
		status: make(map[int]bool),
	}
	check.AddHealthCheck(addr, time.Second*3, t+1)
	go func() {
		for {
			select {
			case a := <-check.Up:
				up(a)
			case b := <-check.Down:
				down(b)
			}
		}

	}()
	logrus.Debugf("add check service %s return %d", addr, t+1)
	return t + 1
}

func (c *Check) AddHealthCheck(addr string, interval time.Duration, id int) {
	if c.status == nil {
		c.status = make(map[int]bool)
	}
	c.CheckRunInterval(serverStatus{
		Interval: interval,
		Id:       id,
		Addr:     addr,
		Down: func(i int) { //false
			logrus.Tracef("%d down", i)
			v := c.status[i]
			if v {
				c.Down <- i
			}
			c.status[i] = false
		},
		Up: func(i int) { //true
			logrus.Tracef("%d up", i)
			v := c.status[i]
			if !v {
				c.Up <- i
			}
			c.status[i] = true
		},
	})
}

func (c *Check) CheckRunInterval(s serverStatus) {
	go func() {
		t := time.NewTicker(s.Interval)
		for {
			<-t.C
			//logrus.Tracef("Check pool %s", pool)
			checker := c.HTTPChecker(s.Addr, 200, time.Second*5, nil)
			if checker != nil {
				logrus.Tracef("check addr %s id %d err=[%s]", s.Addr, s.Id, checker.Error())
				s.Down(s.Id)
			} else {
				s.Up(s.Id)
			}

		}
	}()
}

func (*Check) HTTPChecker(r string, statusCode int, timeout time.Duration, headers http.Header) error {
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("HEAD", r, nil)
	if err != nil {
		return errors.New("error creating request: " + r)
	}
	for headerName, headerValues := range headers {
		for _, headerValue := range headerValues {
			req.Header.Add(headerName, headerValue)
		}
	}
	response, err := client.Do(req)
	if err != nil {
		return errors.New("error while checking: " + r)
	}
	if response.StatusCode != statusCode {
		return errors.New("downstream service returned unexpected status: " + strconv.Itoa(response.StatusCode))
	}
	return nil
}

// TCPChecker attempts to open a TCP connection.
func (*Check) TCPChecker(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return errors.New("connection to " + addr + " failed")
	}
	conn.Close()
	return nil

}
