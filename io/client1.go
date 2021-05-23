package main

import (
	"fmt"
	"io"
	"os"
)

type wt struct {
}

func (w *wt) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	w := new(wt)
	n, _ := io.Copy(w, os.Stdin)
	fmt.Println(n)
	fmt.Println("done")
}
