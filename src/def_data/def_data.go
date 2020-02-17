package main

import (
	"fmt"
	"strings"
	"unicode"
)

// PrtVars 定义各种类型的变量并打印
func PrtVars(a int, b float32, t bool, s1 string) {
	s := "编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。"
	fmt.Println(s)

	fmt.Printf("%d的类型是%T\n", a, a)
	fmt.Printf("%f的类型是%T\n", b, b)
	fmt.Printf("%v的类型是%T\n", t, t)
	fmt.Printf("%s的类型是%T\n", s1, s1)
}

//CountHan 计算字符串里的汉字数
func CountHan() {
	fmt.Println("--------CountHan start-------")
	//go语言中的字符串实际上是类型为byte的只读切片。或者说一个字符串就是一堆字节。
	//这意味着，当我们将字符存储在字符串中时，实际存储的是这个字符的字节。
	//一个字符串包含了任意个byte，它并不限定Unicode，UTF-8或者任何其他预定义的编码。
	s := "hello祝大家春天快乐"
	// fmt.Printf("普通的遍历")
	// for c := range s {
	// 	fmt.Println(s[c])
	// 	//fmt.Printf("%s\n", s[c])
	// }
	// fmt.Printf("unicode遍历")
	var count int
	count = 0
	for _, ch1 := range s {
		//ctype := reflect.Typeof(ch1)
		// fmt.Printf("%c\n", ch1)
		if unicode.Is(unicode.Scripts["Han"], ch1) {
			count++
		}
	}
	fmt.Printf("字符串中，中文的数量为%d\n", count)
}

//ArraySum 计算数组和
func ArraySum() int {
	var a1 = [...]int{1, 3, 5, 7, 8}
	fmt.Println(a1)
	var sum int
	for _, i := range a1 {
		fmt.Println(i)
		sum += i
	}
	return sum
}

func qiepian() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a[5])
	fmt.Println(a)
}

//countMap
func countMap(str string) {
	splitStr := strings.Split(str, " ")
	l := len(splitStr)
	fmt.Println(splitStr, len(splitStr))
	countmap := make(map[string]int, l)
	for _, k := range splitStr {
		fmt.Println(k)
		_, ok := countmap[k]
		if ok {
			//如果存在就值自增
			countmap[k]++
		} else {
			//如果不存在就加入这个key，值为1
			countmap[k] = 1
		}
	}
	fmt.Println(countmap)

}

//structDemo
func structDemo() {
	type person struct {
		name, color string
		age         int8
	}
	var p1, p2, p3 person
	p1.name = "小猪猪"
	p1.color = "gray"
	p1.age = 1
	p2.name = "小咪"
	p2.color = "lihua"
	p2.age = 2
	p3.name = "大喵"
	p3.color = "white+lihua"
	p3.age = 3
	fmt.Printf("p1:%v\n", p1)
	fmt.Printf("p2:%#v\n", p2)
	fmt.Printf("p3:%#v\n", p3)

	var p4 = new(person)
	p4.name = "小黑"
	p4.color = "black"
	p4.age = 5
	fmt.Printf("p4:%#v\n", p4)

	var yemao struct{ name, color string }
	yemao.name = "大白"
	yemao.color = "white"
	fmt.Printf("yemao:%v\n", yemao)

}

func main() {
	fmt.Printf("---demo PrtVars\n")
	var (
		a  int
		b  float32
		t  bool
		s1 string
	)
	a = 1
	b = 1.2
	t = true
	s1 = "abc"
	PrtVars(a, b, t, s1)
	//CountHan()
	// qiepian()

	fmt.Printf("---demo ArraySum\n")
	sum := ArraySum()
	fmt.Printf("数组的元素和为%d\n", sum)

	fmt.Printf("---demo countMap\n")
	str := "how do you do"
	countMap(str)

	fmt.Printf("---demo structDemo\n")
	structDemo()
}
