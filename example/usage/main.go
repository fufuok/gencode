package main

import (
	"encoding/json"
	"fmt"

	"example"
)

func main() {
	data := example.Person{
		Name: "Fufu@中 文",
		Age:  18,
	}
	fmt.Printf("data:\n%#v\n", data)

	// gencode 序列化
	b, _ := data.Marshal(nil)
	fmt.Printf("gencode.Marshal:\n%v\nlen:%d\n", b, len(b))

	// gencode 反序列化
	var resGencode example.Person
	n, _ := resGencode.Unmarshal(b)
	fmt.Printf("gencode.Unmarshal:\n%#v\nlen:%d\n", resGencode, n)

	fmt.Println("------")

	// msgp 序列化
	b, _ = data.MarshalMsg(nil)
	fmt.Printf("msgp.Marshal:\n%v\nlen:%d\n", b, len(b))

	// msgp 反序列化
	var resMsgp example.Person
	_, _ = resMsgp.UnmarshalMsg(b)
	fmt.Printf("msgp.Unmarshal:\n%#v\n", resMsgp)

	fmt.Println("------")

	// JSON 序列化
	b, _ = json.Marshal(data)
	fmt.Printf("json.Marshal:\n%v\nlen:%d\n", b, len(b))

	// JSON 反序列化
	var resJSON example.Person
	_ = json.Unmarshal(b, &resJSON)
	fmt.Printf("json.Unmarshal:\n%#v\n", resJSON)
}
