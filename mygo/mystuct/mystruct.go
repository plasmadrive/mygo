package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	Name string
	Age  int
}

type MyError struct {
	msg  string
	when time.Time
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %s occurred at time: %v", e.msg, e.when)
}

func main() {
	m := MyStruct{"Gareth", 51}
	fmt.Printf("struct is %v", m)
	n := MyStruct{Name: "Sandra", Age: 50}
	fmt.Printf("Other Struct is %v\n", n)
	err := MyError{"Nil Pointer", time.Now()}
	fmt.Printf("Error occured %v", err)

}
