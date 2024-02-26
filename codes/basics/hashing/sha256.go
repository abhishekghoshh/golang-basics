package hashing

import (
	"crypto/sha256"
	"fmt"
)

// https://gobyexample.com/sha256-hashes
func SHA256() {
	str := "sha256 this string"
	hsh := sha256.New()
	hsh.Write([]byte(str))

	bs := hsh.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs)
}