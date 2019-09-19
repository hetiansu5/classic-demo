package main

import "fmt"

type Config struct {
	name   string
	status int
}

type Option func(*Config)

func SetName(name string) Option {
	return func(config *Config) {
		config.name = name
	}
}

func SetStatus(status int) Option {
	return func(config *Config) {
		config.status = status
	}
}

//通过使用函数类型的数据结构，配置函数返回函数数据结构，设置配置
func main() {
	c := new(Config)
	fmt.Println(c)
	op1 := SetName("hts")
	op2 := SetStatus(1)
	op1(c)
	op2(c)
	fmt.Println(c)
}
