package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class string
}

func main() {
	var jingbinh Student
	jingbinh.Name = "jingbinh"
	jingbinh.Age = 24
	jingbinh.Class = "LUT CS-N1"
	momo := Student{"baimenghao", 23, "LUT CS-N1"}
	fmt.Println(momo.Age)
	fmt.Println(momo.Class)
	fmt.Println(momo.Name)
}
