package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func echo(c io.Writer, s string, d time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(d)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(d)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintln(os.Stderr, "%v", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, "%v", err)
		}
		go handleConn(conn)
	}

}
