package stringx

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"sort"
)

func Distinct(s []string) []string {
	sort.Strings(s)
	var re []string
	re = append(re, s[0])
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			re = append(re, s[i+1])
		}
	}
	return re
}

func HexTODecimal(s string) uint64 {
	if s[0:2] == "0x" || s[0:2] == "0X" {
		s = s[2:]
	}
	n, e := strconv.ParseUint(s, 16, 32)
	if e != nil {
		return 0
	}
	return n
}

func DecimalTOHex(de int64) string {
	return fmt.Sprintf("%x", de)
}

func HexTODecString(hexs string) string {
	return big.NewInt(0).SetBytes(FromHex(hexs)).String()
}

func DecStringToHex(ds string) string {
	v, _ := big.NewInt(0).SetString(ds, 10)
	return BigIntToHex(v)
}

func BigIntToHex(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}

func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
