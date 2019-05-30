package stringx

import (
	"fmt"
	"math/big"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Distinct([]string{"ehe", "ehe", "tsh", "tsh", "walletapi", "walletapi", "gateway", "gateway", "walletapi"}))

}

func TestHexTODecimal(t *testing.T) {
	fmt.Println(HexTODecimal("0x1e7d4c"))
	fmt.Println(DecimalTOHex(1998156))
}

func TestHex(t *testing.T) {
	fmt.Println(HexTODecString("0x1e7d4c"))
	fmt.Println(DecStringToHex("1998156"))
	fmt.Println(BigIntToHex(big.NewInt(56565)))
}
