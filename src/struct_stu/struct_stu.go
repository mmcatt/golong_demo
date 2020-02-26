package main

import "fmt"

// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

//1. add
func addStudent(s *studentMgr, newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2. modify
func modifyStudent(s *studentMgr, newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	//
	fmt.Println("在系统中没有找到对应的学号！")
}

//3. show
func showStudent(s *studentMgr) {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d; 姓名 %s; 班级 %s\n", v.id, v.name, v.class)
	}
}

//4. delete
func delStudent(s *studentMgr, id int) {
	for i, v := range s.allStudents {
		if v.id == id {
			fmt.Println("要删除的id，当前id，第几轮循环", id, i)
			s.allStudents = append(s.allStudents[:i], s.allStudents[i+1:]...)
			return
		}
	}
	fmt.Println("在系统中没有找到对应的学号！")
}

func showMenu() {
	fmt.Printf("请选择功能：\n")
	fmt.Printf("1. 展示学生列表\n")
	fmt.Printf("2. 添加学生\n")
	fmt.Printf("3. 编辑学生信息\n")
	fmt.Printf("4. 删除学生\n")
}
func getID() int {
	var id int
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	return id
}
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr()
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i), "萌萌哒")
		sm.allStudents = append(sm.allStudents, tmpStu)
	}
	var choice int
	for {
		showMenu()
		//输入
		fmt.Println("请输入你要操作的序号")
		fmt.Scanf("%d\n", &choice)
		fmt.Println("你输入的是:", choice)
		switch choice {
		case 1:
			//list
			showStudent(sm)
		case 2:
			//add添加
			stu := getInput()
			addStudent(sm, stu)
		case 3:
			//edit
			stu := getInput()
			modifyStudent(sm, stu)
		case 4:
			//delete
			id := getID()
			delStudent(sm, id)

		default:
			fmt.Printf("请重新选择\n")
		}
	}
}
