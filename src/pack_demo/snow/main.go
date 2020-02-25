package main

import (
	"fmt"
	"pack_demo/clac"
)

func main() {
	var a, b float32
	a = 1
	b = 2
	s := clac.Add(a, b)
	fmt.Printf("两数之和：%v\n", s)
	s = clac.Sub(a, b)
	fmt.Printf("两数之差：%v\n", s)
	s = clac.Multiply(a, b)
	fmt.Printf("两数之乘：%v\n", s)
	s = clac.Division(a, b)
	fmt.Printf("两数相除：%v\n", s)
}
