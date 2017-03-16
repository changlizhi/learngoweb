package usehttp

import (
	"net/http"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func zqg_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bytes, err := ioutil.ReadFile("texts/zhengqige")
		if err != nil {
			log.Fatal("os.Open error")
		}
		key_text := "aaaabbbbccccddddaaaabbbbccccdddd"
		fmt.Println(len(key_text))
		c, err := aes.NewCipher([]byte(key_text))
		if err != nil {
			fmt.Printf("Error : NewCipher(%d bytes) = %s", len(key_text), err)
		}
		cfb := cipher.NewCFBEncrypter(c, commonIV)
		cipher_text := make([]byte, len(bytes))
		cfb.XORKeyStream(cipher_text, bytes)
		fmt.Println(string(cipher_text))
		fmt.Fprintf(w, "%s", string(cipher_text))
	}
}

func Map_zqg() {
	http.HandleFunc("/zqg", zqg_request)
	http.ListenAndServe(":9090", nil)
}