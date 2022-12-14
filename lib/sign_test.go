package lib

import (
	"crypto/rand"
	"testing"

	"golang.org/x/crypto/nacl/sign"
)

func TestSign(t *testing.T) {
	pub, priv, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		t.Error(err)
	}

	crypt := Sign(priv, []byte("Hello World"))
	msg, b := Verify(pub, crypt)

	if string(msg) != "Hello World" || !b {
		t.Error("Signature verification failed")
	}
}
