package stringx

import (
	"fmt"
	"strconv"

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

func HexTODecimal(o string) uint64 {
	val := o[2:]
	n, e := strconv.ParseUint(val, 16, 32)
	if e != nil {
		return 0
	}
	return n
}

func DecimalTOHex(de int64) string {
	return fmt.Sprintf("%x", de)
}
