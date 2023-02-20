package main

import (
	"reflect"
)

type Speaker interface {
	tell()
}

type Speak struct {
}

func (s *Speak) tell() {

}

type SpecialString interface{}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}
	return t
}

// 用一个指针获取接口的类型
func getInterfaceType(ifacePtr interface{}) interface{} {
	return InterfaceOf(ifacePtr)
}

// func main() {
// 	var s Speaker = &Speak{}
// 	fmt.Println(getInterfaceType(&s))

// 	var in interface{}
// 	fmt.Println(getInterfaceType(&in))
// 	fmt.Println(getInterfaceType((*interface{})(nil))) // (*interface{})(nil)是空接口类型的空指针

// 	fmt.Println(getInterfaceType((*SpecialString)(nil)))
// }
