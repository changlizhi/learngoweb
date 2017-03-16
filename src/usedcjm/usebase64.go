package usedcjm

import (
	"encoding/base64"
	"fmt"
)

func my_base64_encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
func my_base64_decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func Use_base64() {
	hello := "hello world!"
	debyte := my_base64_encode([]byte(hello))
	fmt.Println(debyte)
	enbyte, err := my_base64_decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}
	fmt.Println(string(enbyte))

}
