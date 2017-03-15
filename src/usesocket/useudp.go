package usesocket

import (
	"fmt"
	"net"
	"os"
	"time"
)

func Udp_request() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(0)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("anything"))
	checkErr(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkErr(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func Udp_service() {
	service := ":7777"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handle_client_udp(conn)
	}
}
func handle_client_udp(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
