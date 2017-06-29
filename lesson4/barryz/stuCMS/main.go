/*
完成学生信息管理系统

实现如下4个指令

1. add name id，添加一个学生的信息，如果name有重复，报错
2. list, 列出所有的学生信息
3. save filename，保存所有的学生信息到filename指定的文件中
4. load filename, 从filename指定的文件中加载学生信息

Code Example:

...
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StudentSet struct {
	M map[string]*Student `json:"data"`
}

// NewStudentSets initial a new studentsets
func NewStudentSets() *StudentSet {
	return &StudentSet{M: make(map[string]*Student, 0)}
}

func (s *StudentSet) Add(id int, name string) (err error) {
	_, ok := s.M[name]
	if ok {
		err = fmt.Errorf("student %s already exists.", name)
		return
	}

	s.M[name] = &Student{Id: id, Name: name}

	return
}

// list list all student information as a string
func (s *StudentSet) list() string {
	var str string = "Id\t\tName\n"
	for k, v := range s.M {
		str += fmt.Sprintf("%d\t\t%s\n", v.Id, k)
	}
	return str
}

// String implements string method
func (s *StudentSet) String() string {
	return s.list()
}

// Dump student info to the file that specified
func Dump(fileName string, stu *StudentSet) (err error) {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	bs, err := json.Marshal(stu)
	if err != nil {
		return
	}

	_, err = fd.Write(bs)
	if err != nil {
		return
	}

	return
}

// Load student info to the file that specified
func Load(fileName string) (stu *StudentSet, err error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(bs, &stu)
	if err != nil {
		return
	}

	return
}

func main() {

	var cmd string
	var name string
	var file string
	var id int
	var line string
	var err error
	stu := NewStudentSets()
	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println(stu)
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			err = stu.Add(id, name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("add done.\n")
			}
		case "save":
			fmt.Sscan(line, &cmd, &file)
			err = Dump(file, stu)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("save student info to %s done\n", file)
			}
		case "load":
			fmt.Sscan(line, &cmd, &file)
			stu, err = Load(file)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("load student info from %s done\n", file)
			}
		case "exit", "quit", "q":
			os.Exit(0)
		default:
			fmt.Println("list|add|save|load|exit")
		}
	}

}
