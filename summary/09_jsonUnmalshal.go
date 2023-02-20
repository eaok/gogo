package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://www.sojson.com/
// https://mholt.github.io/json-to-go/
type Response struct {
	Code     int             `json:"code"`
	Message  string          `json:"message"`
	JsonData json.RawMessage `json:"data"`
}
type Items struct {
	ID2   int    `json:"Id2,omitempty"`
	Title string `json:"Title"`
	ID    int    `json:"Id,omitempty"`
	Test  int    `json:"Test,omitempty"`
}

type JsonData struct {
	Items []*Items `json:"items"`
}

var jsonStr = `{
	"code": 0,
	"message": "操作成功",
	"data": {
		"items": [{
			"Id": 100,
			"Title": "木华黎"
		}, {
			"Id": 200,
			"Title": "耶律楚才"
		}, {
			"Id": 300,
			"Title": "纳呀啊",
			"Test": 100
		}]
	}
}`

func jsonUnmalshal() {
	res := Response{}
	data := JsonData{}

	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonStr)
	fmt.Printf("%#v\n", res)

	err = json.Unmarshal(res.JsonData, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.TypeOf(data.Items), data.Items, *data.Items[0])
}

func jsonMarshal() {
	structJsonData := JsonData{
		[]*Items{
			{
				ID:    100,
				Title: "木华黎",
			},
			{
				ID:    101,
				Title: "木",
			},
		},
	}

	jsonRawData, err := json.Marshal(structJsonData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonRawData))

	structResponseData := Response{
		0,
		"操作成功",
		jsonRawData,
	}
	jsonRwaResponse, err := json.Marshal(structResponseData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonRwaResponse))
}

// func main() {
// 	// jsonUnmalshal()
// 	jsonMarshal()
// }
