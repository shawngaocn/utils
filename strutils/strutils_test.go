package strutils

import (
	"log"
	"testing"
)

func TestNonceString(t *testing.T) {
	for index := 0; index < 10; index++ {
		str := NonceString()
		log.Println(str)
	}
}
