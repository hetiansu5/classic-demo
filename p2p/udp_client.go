package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"strings"
	"strconv"
	"log"
)

func main() {
	var remoteAddr *net.UDPAddr

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9982} // 注意端口必须固定
	serverAddr := parseAddr("106.14.162.184:11000")
	conn, err := net.DialUDP("udp", srcAddr, serverAddr)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(conn.LocalAddr().String(), conn.RemoteAddr());
	localAddr := parseAddr(conn.LocalAddr().String())

	var i int
	for {
		i = i + 1
		if conn != nil {
			serverMsg := "to server " + strconv.Itoa(i)
			_, err := conn.Write([]byte(serverMsg))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("send", serverMsg)

			var msg [100]byte
			n, _ := conn.Read(msg[0:])
			s := string(msg[0:n])
			if n > 0 {
				fmt.Println("receive", s)
			}

			if strings.HasPrefix(s, "old:") {
				remoteAddr = parseAddr(strings.TrimPrefix(s, "old:"))
			}

			if strings.HasPrefix(s, "new:") {
				remoteAddr = parseAddr(strings.TrimPrefix(s, "new:"))
			}
		}

		if remoteAddr != nil {
			conn.Close()
			conn = nil
			break
		}

		time.Sleep(3 * time.Second)
	}

	digHole(localAddr, remoteAddr)

}

func parseAddr(addr string) *net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return &net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

func digHole(localAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", localAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 向另一个peer发送一条udp消息(对方peer的nat设备会丢弃该消息,非法来源),用意是在自身的nat设备打开一条可进入的通道,这样对方peer就可以发过来udp消息
	if _, err = conn.Write([]byte("hand_shake")); err != nil {
		log.Println("send handshake:", err)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err = conn.Write([]byte("from [1]")); err != nil {
				log.Println("send msg fail", err)
			}
		}
	}()
	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("收到数据:%s\n", data[:n])
		}
	}
}
