package useos

import (
	"fmt"
	"os"
)

func MymakeDirs() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("axtaxie/tes1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie")
}

func MyFileUse() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("just a tess!\r\n"))
	}
}

func MyOpenFile() {
	userFile := "astaxie.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
