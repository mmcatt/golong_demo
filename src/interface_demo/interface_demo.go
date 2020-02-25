package main

import "fmt"

//接口是一种类型，一种抽象的类型

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

//接口不管你是什么类型，只管你要实现什么方法
//定义一个类型，一个抽象的类型，只要实现了say()这个方法的类型
type sayer interface {
	say()
}

//打的函数
func hit(arg sayer) {
	arg.say() //不管传进来的是什么类型，都要打ta，打就会叫，就执行say的方法
}

//使用值接收者实现接口和使用指针接收者实现接口的区别

type mover interface {
	move()
}

type person struct {
	name string
	age  int8
}

type animal interface {
	mover
	sayer
}

// 使用值接收者实现接口：类型的值和类型的指针都能够保存到接口变量中
// func (p person) move() {
// 	fmt.Printf("%s在跑\n", p.name)
// }

//使用指针接收者实现接口：只有类型指针能够保存到接口变量as，类型的值无法保存到接口变量种
func (p *person) move() {
	fmt.Printf("%s在跑\n", p.name)
}
func (p *person) say() {
	fmt.Printf("%s在叫\n", p.name)
}

//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
//接口中没有定义任何方法时 该接口就是一个空接口
//任意类型都实现了空接口 --> 空接口变量可以存储任意类型的值 存储任意类型的值 存储任意类型的值

//空接口一般不需要提前定义
type xxx interface {
}

//空接口的应用
//1. 空接口类型作为函数的参数
//2.空接口可以作为map的value

func main() {
	c1 := cat{}
	c1.name = "小宝"
	hit(c1)
	d1 := dog{}
	hit(d1)
	var s sayer
	s = c1
	fmt.Println(s)

	var m mover
	p1 := person{
		name: "张三",
		age:  30,
	}
	p2 := &person{
		name: "王五",
		age:  50,
	}
	fmt.Println(p1)
	// m = p1 //无法赋值 因为p1是person类型的值 没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	var ss sayer //定义一个sayer类型的变量
	ss = p2
	ss.say()
	fmt.Println(ss)

	var a animal
	a = p2
	fmt.Printf("animal\n")
	a.move()
	a.say()

	var x interface{} //定义一个空接口变量x
	x = "ssss"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)
	x = p1
	fmt.Println(x)

	// var mm = make(map[string]interface{}, 16)
	// mm["name"] = "李四"
	// mm["age"] = 18
	// mm["cats"] = []string{"大喵", "小咪", "小宝"}
	// fmt.Println(mm)

	switch v := x.(type) {
	case string:
		fmt.Printf("string,value:%v\n", v)
	case bool:
		fmt.Printf("bool,value:%v\n", v)
	case int:
		fmt.Printf("int,value:%v\n", v)
	default:
		fmt.Printf("猜不到，value:%v\n", v)
	}
}
