//2、 以下代码有什么问题，说明原因
package main

import "fmt"

type student struct {
    Name string
    Age  int
}

func pase_student() map[string]*student {
    m := make(map[string]*student)
    stus := []student{
        {"zhou", 24},
        {"li", 23},
        {"wang", 22},
    }

    for _, stu := range stus {
//        fmt.Println(stu)
//        println(&stu)
        m[stu.Name] = &stu
    }

    // 正确
//    for i:=0;i<len(stus);i++  {
//        m[stus[i].Name] = &stus[i]
//    }
//    for i, _ := range stus {
//        stu := stus[i]//这里直接用key来取value赋值到stu变量中，这样stu的地址都是新的。
//        m[stu.Name] = &stu
//    }

    return m
}

func main() {
    students := pase_student()
    fmt.Println(students)
    for k, v := range students {
        fmt.Printf("key=%s,value=%v\n", k, v)
    }
}
