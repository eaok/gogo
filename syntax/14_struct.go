package main

import (
	"fmt"
)

func main() {
	type info struct {
		id   int
		name struct {
			firstName string
			lastName  string
		}
	}

	var p *info
	p = new(info)
	p.id = 12
	p.name.firstName = "jian"
	p.name.lastName = "hu"

	infoo := new(info)
	infoo.id = 13
	infoo.name.firstName = "haha"
	infoo.name.lastName = "lulu"

	fmt.Println(p)
	fmt.Println(infoo)
}
