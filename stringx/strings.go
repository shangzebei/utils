package stringx

import (
	"bytes"
	"fmt"
	"github.com/shangzebei/utils/optional"

	"sort"
)

func Distinct(s []string) []string {
	sort.Strings(s)
	var re []string
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] || i == 0 {
			re = append(re, s[i+1])
		}
	}
	return re
}

func Fprintf(s string, f ...interface{}) string {
	temp := bytes.NewBufferString("")
	optional.OfErrorable(fmt.Fprintf(temp, s, f...))
	return temp.String()
}
