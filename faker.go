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

func Reveal(c string, prefix ...string) string {
	p := FakePrefix
	if len(prefix) > 0 {
		p = prefix[0]
	}
	if strings.HasPrefix(c, p) {
		dec := Dec(strings.Replace(c, p, "", 1))
		if dec != "" {
			return dec
		}
	}

	return c
}

func Conceal(c string, prefix ...string) string {
	p := FakePrefix
	if len(prefix) > 0 {
		p = prefix[0]
	}
	return p + Enc(c)
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
