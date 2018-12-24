//写出打印的结果。
package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
	//    Name string `json:"name"`
}

func main() {
	js := `{
        "name":"11"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

//p中的属性值为空的原因是因为，name的首字母小写，修改成大写，重新运行即可。
