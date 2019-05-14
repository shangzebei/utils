package stringx

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Distinct([]string{"ehe", "ehe", "tsh", "tsh", "walletapi", "walletapi", "gateway", "gateway", "walletapi"}))

}

func TestHexTODecimal(t *testing.T) {
	fmt.Println(HexTODecimal("0x1e7d4c"))
	fmt.Println(DecimalTOHex(1998156))
}
