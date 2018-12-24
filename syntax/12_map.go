package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["ok"] = 200
	m["error"] = 404

	//	delete(m, "ok")

	//	is, v:= m["error"]

	for is, v := range m {
		fmt.Print(v, is, "\n")
	}

	mp := map[string]int{"a": 10, "b": 20}
	fmt.Print(mp, "\n")
}
