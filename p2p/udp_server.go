package main

import (
	"os"
	"fmt"
	"net"
	"strconv"
)

var addrList map[string]*net.UDPAddr

func init() {
	addrList = make(map[string]*net.UDPAddr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn) {
	var buf [100]byte

	n, rAddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("receive", string(buf[0:n]))

	key := rAddr.IP.String() + ":" + strconv.Itoa(rAddr.Port)
	if _, ok := addrList[key]; !ok {
		fmt.Println("new addr join:", rAddr)
		_, err = conn.WriteToUDP([]byte("welcome!"), rAddr)
		checkError(err)

		//给在线客户端发送已经在线的客户端IP、端口
		for k, v := range addrList {
			_, err = conn.WriteToUDP([]byte("new:"+key), v)
			if err != nil {
				delete(addrList, k)
			}

			_, err = conn.WriteToUDP([]byte("old:"+k), rAddr)
			checkError(err)
		}

		addrList[key] = rAddr
	}

	//WriteToUDP
	//func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, err = conn.WriteToUDP([]byte("to client "+string(buf[0:n])), rAddr)
	checkError(err)
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:11000")
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)

	for {
		recvUDPMsg(conn)
	}
}
