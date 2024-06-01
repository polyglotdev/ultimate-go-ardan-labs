package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Elijah struct {
	flag    bool
	counter int16
	pi      float32
}

type Ezra struct {
	flag    bool
	counter int16
	pi      float32
}

// example represents a struct type with three fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	fmt.Println("Hello, Struct Types!")

	var e1 example

	fmt.Printf("%+v\n", e1)
	fmt.Println("------------------------------")

	e1 = example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	fmt.Printf("%+v\n", e1)
	fmt.Printf("Flag: %t\n", e1.flag)
	fmt.Printf("Counter: %d\n", e1.counter)
	fmt.Printf("Pi: %f\n", e1.pi)

	fmt.Println("------------------------------")
	// Print the name of each field and its size in bytes
	t := reflect.TypeOf(e1)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Name: %s, Size: %d bytes\n", field.Name, unsafe.Sizeof(field.Type))
	}

	fmt.Println("------------------------------")

	totalSize := unsafe.Sizeof(e1)
	fmt.Printf("Total size of e1: %d bytes\n", totalSize)

	fmt.Println("------------------------------")
	e2 := example{
		flag:    false,
		counter: 10,
		pi:      3.14159,
	}

	fmt.Println("e2")
	fmt.Printf("%+v\n", e2)
	fmt.Printf("Flag: %t\n", e2.flag)
	fmt.Printf("Counter: %d\n", e2.counter)

	fmt.Println("------------------------------")
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	fmt.Println("e3")
	fmt.Printf("%+v\n", e3)
	fmt.Printf("Flag: %t\n", e3.flag)
	fmt.Printf("Counter: %d\n", e3.counter)
	fmt.Printf("Pi: %f\n", e3.pi)
	fmt.Println("------------------------------")
	el := Elijah{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	ez := Ezra{
		flag:    false,
		counter: 20,
		pi:      3.14159,
	}

	fmt.Println("Elijah")
	fmt.Printf("%+v\n", el)
	fmt.Printf("Elijah.Flag: %t\n", el.flag)
	fmt.Printf("Elijah.Counter: %d\n", el.counter)
	fmt.Printf("Elijah.Pi: %f\n", el.pi)

	fmt.Println("------------------------------")
	fmt.Println("Ezra")
	fmt.Printf("%+v\n", ez)
	fmt.Printf("Ezra.Flag: %t\n", ez.flag)
	fmt.Printf("Ezra.Counter: %d\n", ez.counter)
	fmt.Printf("Ezra.Pi: %f\n", ez.pi)
}
