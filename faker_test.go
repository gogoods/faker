package main

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	pass := "123456"
	enc := Enc(pass)
	fmt.Println(enc)
	dec := Dec(enc)
	fmt.Println(dec)

	conceal := Conceal(pass)
	fmt.Println(conceal)

	reveal := Reveal(conceal)
	fmt.Println(reveal)
}
