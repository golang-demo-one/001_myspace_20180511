package main

import (
	"encoding/json"
	"fmt"
)

/*
	https://zhidao.baidu.com/question/459318125920057725.html
*/

type Project struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	//Docs string `json:"docs,omitempty"` //omitempty 为空则不输出
	Docs string `json:"docs"`
}

func main() {
	p1 := Project{
		Name: "CleverGo高性能框架",
		Url:  "https://github.com/headwindfly/clevergo",
	}

	data, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	p2 := Project{
		Name: "CleverGo高性能框架",
		Url:  "https://github.com/headwindfly/clevergo",
		Docs: "https://github.com/headwindfly/clevergo/tree/master/docs",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)
}
