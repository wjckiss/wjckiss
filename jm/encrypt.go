package jm

import (
	"crypto/md5"
	"encoding/hex"
	//	"crypto/aes"
	//	"crypto/cipher"
	//	"fmt"
	//	"os"
)

func MD5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	//	fmt.Println(md.BlockSize(), md.Size())
	return hex.EncodeToString(md.Sum(nil))
}
