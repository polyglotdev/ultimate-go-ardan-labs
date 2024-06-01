package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Exercise: Variables")

	var name string
	var age int

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	fmt.Println("--------------------")

	aa := "10"
	bb := "22"

	fmt.Println("aa:", reflect.TypeOf(aa))
	fmt.Println("bb:", reflect.TypeOf(bb))

	fmt.Println("--------------------")
	newAA, err := strconv.ParseInt(aa, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	newBB, err := strconv.ParseInt(bb, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("newAA:", reflect.TypeOf(newAA))
	fmt.Println("newBB:", reflect.TypeOf(newBB))
}
