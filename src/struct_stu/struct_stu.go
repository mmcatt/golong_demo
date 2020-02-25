package main

import "fmt"

// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
type student struct {
	id    int
	name  string
	age   int
	score float32
}
type class struct {
	Map map[int]*student
}

func (c *class) addStu() {
	var id int
	var name string
	var age int
	var score float32
	fmt.Println("输入id，必须为整数：")
	_, err := fmt.Scan(&id)
	fmt.Println("输入name：")
	_, err = fmt.Scan(&name)
	fmt.Println("输入age，必须为整数：")
	_, err = fmt.Scan(&age)
	fmt.Println("输入score，浮点数：")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("出错")
	}
	_, isSave := c.Map[id]
	if isSave {
		fmt.Println("id已存在，请更换")
		return
	}
	stu := &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
	c.Map[id] = stu
	fmt.Println("保存成功")
}
func (s *student) listStu() {}
func main()                 {}
