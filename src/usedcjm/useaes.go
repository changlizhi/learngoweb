package usedcjm

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

//上面通过调用函数aes.NewCipher(参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法)
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Aes_jm() {
	plain_text := []byte("my name is clz")
	if len(os.Args) > 1 {
		plain_text = []byte(os.Args[1])
	}
	key_text := "clzclzx12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	fmt.Println(len(key_text))
	//创建加密算法
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}
	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipher_text := make([]byte, len(plain_text))
	cfb.XORKeyStream(cipher_text, plain_text)
	fmt.Printf("%s 加密为：===%x\n", plain_text, cipher_text)

	//解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plain_text_copy := make([]byte, len(plain_text))
	cfbdec.XORKeyStream(plain_text_copy, cipher_text)
	fmt.Printf("%x解密为：========%s\n", cipher_text, plain_text_copy)
}
