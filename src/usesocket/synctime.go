package usesocket

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func Dia_time() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handClient(conn)
	}
}
func handClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
	request := make([]byte, 28)
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break
		} else if string(request) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128)

	}

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
