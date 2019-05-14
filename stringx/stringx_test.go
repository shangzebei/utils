package stringx

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Distinct([]string{"ehe", "ehe", "tsh", "tsh", "walletapi", "walletapi", "gateway", "gateway", "walletapi"}))

}
