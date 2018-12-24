package main

import (
	"strings"

	"github.com/gogoods/x/crypto/aes"
	"github.com/gogoods/x/crypto/base64"
)

const (
	Key        = "YOu aRa AweSoME!"
	FakePrefix = "~"
)

func Reveal(c string) string {
	if strings.HasPrefix(c, FakePrefix) {
		dec := Dec(strings.Replace(c, FakePrefix, "", 1))
		if dec != "" {
			return dec
		}
	}

	return c
}

func Conceal(c string) string {
	return FakePrefix + Enc(c)
}

func Dec(v string) string {
	bs, _ := base64.DiyDecodeToBytes(v)
	bytes, _ := aes.Decrypt(bs, []byte(Key))
	return string(bytes)
}

func Enc(v string) string {
	bytes, _ := aes.Encrypt([]byte(v), []byte(Key))
	return base64.DiyEncodeBytes(bytes)
}
