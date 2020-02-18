package main

import (
	"errors"
	"fmt"
	"strings"
)

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		fmt.Printf("before +=, x=%v, y=%v\n", x, y)
		x += y
		fmt.Printf("after +=, x=%v, y=%v\n", x, y)
		return x
	}
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func countCoin(name string) int {
	var sum int
	sum = 0
	if &name != nil {
		for _, ch := range strings.Split(name, "") {
			switch {
			case ch == "e" || ch == "E":
				sum++
			case ch == "i" || ch == "I":
				sum += 2
			case ch == "o" || ch == "O":
				sum += 3
			case ch == "u" || ch == "U":
				sum += 4

			}
		}
	}
	return sum
}
func dispatchCoin() int {
	for _, name := range users {
		coin := countCoin(name)
		distribution[name] = coin
		fmt.Printf("%s的金币数为%v\n", name, coin)
	}
	result := 0
	for _, value := range distribution {
		result += value
	}
	fmt.Printf("总共使用的金币数为%v", result)
	return coins - result
}
func main() {
	// var s string
	// s = "-"
	// var op func(int, int) int
	// op, _ = do(s)
	// fmt.Println(op(10, 20))

	// f := adder()
	// fmt.Println(f(10))
	// fmt.Println(f(15))
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
