package main

import (
	"fmt"
	"github.com/Doozers/BasicPublicKeyEncryption/lib"
	"os"
)

func main() {
	err := lib.GenKeys("private1.key", "public1.key")
	if err != nil {
		return
	}

	privKey, err := os.ReadFile("public1.key")
	if err != nil {
		return
	}
	pubKey, err := os.ReadFile("private1.key")
	if err != nil {
		return
	}

	crypt := lib.Sign((*[64]byte)(privKey), []byte("Hello World"))
	msg, b := lib.Verify((*[32]byte)(pubKey), crypt)

	fmt.Println(string(msg), b)
}
