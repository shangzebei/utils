package stringx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, []string{"ehe", "walletapi"}, Distinct([]string{"ehe", "walletapi", "walletapi"}))
	assert.Equal(t, []string{"ehe", "walletapi"}, Distinct([]string{"ehe", "ehe", "walletapi"}))
	assert.Equal(t, []string{"ehe", "walletapi"}, Distinct([]string{"ehe", "walletapi", "ehe"}))
	assert.Equal(t, []string{"ehe"}, Distinct([]string{"ehe", "ehe", "ehe"}))
	assert.Equal(t, []string{"ehe", "gateway", "tsh", "walletapi"}, Distinct([]string{"ehe", "ehe", "tsh", "tsh", "walletapi", "walletapi", "gateway", "gateway", "walletapi"}))
}

func TestHexTODecimal(t *testing.T) {
	fmt.Println(HexTODecimal("0x1e7d4c"))
	fmt.Println(DecimalTOHex(1998156))
}

func TestHex(t *testing.T) {
	assert.Equal(t, "1998156", HexTODecString("0x1e7d4c"))
	assert.Equal(t, "1e7d4c", DecStringToHex("1998156"))
	assert.Equal(t, "dcf5", BigIntToHex(big.NewInt(56565)))
	assert.Equal(t, "200000000000000000000", HexTODecString("0xad78ebc5ac6200000"))
}

func TestBIg(t *testing.T) {
	//00020c49ba5e353f7ced916872b020c49ba5e353f7ced916872b020c49ba5e35
	fmt.Println(HexTODecString("0x020c49ba5e353f7ced916872b020c49ba5e353f7ced916872b020c49ba5e35"))
	fmt.Println(HexTODecString("00020c49ba5e353f7ced916872b020c49ba5e353f7ced916872b020c49ba5e35"))
	fmt.Println(FromHex("020c49ba5e353f7ced916872b020c49ba5e353f7ced916872b020c49ba5e35"))

}
