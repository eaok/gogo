package main

import (
    "fmt"
    "strings"
)

type Person struct {
    firstName   string
    lastName    string
}

func upPerson(p *Person) {
    p.firstName = strings.ToUpper(p.firstName)
    p.lastName = strings.ToUpper(p.lastName)
}

func printPerson(pers Person) {
    fmt.Printf("名字是 %s%s\n", pers.firstName, pers.lastName)
}

func printPersonP(pers *Person) {
    fmt.Printf("名字是 %s%s\n", pers.firstName, pers.lastName)
}

func main() {
    // 作为值类型
    var pers1 Person
    pers1.firstName = "张"
    pers1.lastName = "三"
    upPerson(&pers1)
    printPerson(pers1)
    printPersonP(&pers1)
    fmt.Printf("名字是 %s%s\n", pers1.firstName, pers1.lastName)

    // 作为指针
    pers2 := new(Person)
    pers2.firstName = "张"
    pers2.lastName = "三"
    //(*pers2).lastName = "三" // 这也是合法的
    upPerson(pers2)
    fmt.Printf("名字是 %s%s\n", pers2.firstName, pers2.lastName)

    // 作为字面量
    pers3 := &Person{"张", "三"}
    upPerson(pers3)
    fmt.Printf("名字是 %s%s\n", pers3.firstName, pers3.lastName)
}
