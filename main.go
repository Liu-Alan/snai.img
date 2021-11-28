package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func fileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	imga, _ := ioutil.ReadFile("aa.jpg")
	imgb := append(imga, ' ')
	ioutil.WriteFile("bb.jpg", imgb, 0666)

	imgamd5 := fileHash(imga)

	imgb, _ = ioutil.ReadFile("bb.jpg")
	imgbmd5 := fileHash(imgb)

	fmt.Println("imga:", imgamd5)
	fmt.Println("imgb:", imgbmd5)
}
