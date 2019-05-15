package util

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
	"sort"
	"strconv"
)

func RemoveMapVale(s string, m map[string][]string) {
	for key, value := range m {
		//var re []string
		for index, v := range value {
			if v == s {
				m[key] = append(value[:index], value[index+1:]...)
				return
			}
		}
	}

}

func AddMapVale(name string, m map[string][]string, v string) {
	for _, value := range m[name] {
		if value == v {
			return
		}
	}
	m[name] = append(m[name], v)
}

func GetOutboundIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Md5Bytes(b []byte) string {
	h := md5.New()
	if _, err := io.Copy(h, bytes.NewReader(b)); err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%x", h.Sum(nil))
	return buf.String()
}

func GetIPAndPort(addr string) (string, int) {
	host, port, _ := net.SplitHostPort(addr)
	iPort, _ := strconv.Atoi(port)
	return host, iPort
}

func SortSlice(f interface{}) {
	sort.Slice(f, func(i, j int) bool {
		a := reflect.ValueOf(f).Index(i).MethodByName("Order").Call(nil)[0].Int()
		b := reflect.ValueOf(f).Index(j).MethodByName("Order").Call(nil)[0].Int()
		return a < b
	})
}
