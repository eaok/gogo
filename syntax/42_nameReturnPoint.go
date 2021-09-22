package main

import (
	"fmt"
)

func main() {
	b := 5
	testNameReturnPoint(&b)
}

func testNameReturnPoint(c *int) (a *int, err error) {
	fmt.Printf("%v\n", *c)
	fmt.Printf("%v\n", c)

	//fmt.Printf("%v\n", *a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", &a)

	return nil, nil
}
