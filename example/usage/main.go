package main

import (
	"encoding/json"
	"fmt"

	"example"
)

func main() {
	src := example.Person{
		Name: "Fufu@中 文",
		Age:  18,
	}
	fmt.Printf("src:\n%#v\n", src)

	// 序列化
	b, _ := src.Marshal(nil)
	fmt.Printf("src.Marshal:\n%v\nlen:%d\n", b, len(b))

	// 反序列化
	var dst example.Person
	n, _ := dst.Unmarshal(b)
	fmt.Printf("dst.Unmarshal:\n%#v\nlen:%d\n", dst, n)

	fmt.Println("------")

	// JSON 序列化对比
	b, _= json.Marshal(src)
	fmt.Printf("src.Marshal:\n%v\nlen:%d\n", b, len(b))

	_ = json.Unmarshal(b, &dst)
	fmt.Printf("dst.Unmarshal:\n%#v\n", dst)
}
