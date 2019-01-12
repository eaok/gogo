package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3)
	oldCap := cap(s)

	for i := 0; i < 99; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap:%d====>%d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
