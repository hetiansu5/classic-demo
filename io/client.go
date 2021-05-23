package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func MustCopy(dst io.Writer, c io.Reader) {
	n, err := io.Copy(dst, c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// go MustCopy(os.Stdout, conn)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	MustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}