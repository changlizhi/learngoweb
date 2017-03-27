package usejiami

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func Use_three_jiami() {
	h := sha256.New()
	io.WriteString(h, "his money is twiced")
	fmt.Printf("sha256加密结果：=====%x", h.Sum(nil))
	fmt.Println()
	h2 := sha1.New()
	io.WriteString(h2, "his money is twiced")
	fmt.Printf("sha1加密结果：=====%x", h2.Sum(nil))
	fmt.Println()
	h3 := md5.New()
	io.WriteString(h3, "his money is twiced")
	fmt.Printf("md5加密结果：=====%x", h3.Sum(nil))
	fmt.Println()

	h4 := md5.New()
	io.WriteString(h, "his money is twiced")
	newmd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println()

	salt1 := "@#$%"
	salt2 := "^&*()"
	io.WriteString(h4, salt1)
	io.WriteString(h4, "abc")
	io.WriteString(h4, salt2)
	io.WriteString(h4, newmd5)
	last := fmt.Sprintf("%x", h4.Sum(nil))
	fmt.Printf("加盐结果：=======%x", last)
}
