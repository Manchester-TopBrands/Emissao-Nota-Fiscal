package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	sha := sha256.New()
	sha.Write([]byte("Luccasantos10@"))
	rst := sha.Sum(nil)
	fmt.Println(base64.RawStdEncoding.EncodeToString(rst))

}
