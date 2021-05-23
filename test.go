package main

import (
	"fmt"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
)

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}

	// ①
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	fmt.Println(mutexLocked, mutexWoken, mutexStarving)
}

type options struct {
	Level int
}

type Server struct {
	ops options
}

type Option func(*options)

func LevelOption(level int) Option {
	return func(o *options) {
		o.Level = level
	}
}

func newServer(ops ...Option) {
	s := new(Server)
	for _, option := range ops {
		option(&s.ops)
	}
}
