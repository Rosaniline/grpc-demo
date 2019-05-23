package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	person "github.com/Rosaniline/grpc-demo/proto"
	"github.com/golang/protobuf/proto"
)

func create() {
	p := &person.Person{
		Name: "唐伯虎",
		Id: 9527,
		Email: "9527@huafu.com",
	}

	out, _ := proto.Marshal(p)

	_ = ioutil.WriteFile("proto_message", out, 0644)

	out, _ = json.Marshal(p)

	_ = ioutil.WriteFile("json_message.json", out, 0644)
}

func read() {
	in, _ := ioutil.ReadFile("proto_message")

	p := &person.Person{}
	_ = proto.Unmarshal(in, p)

	fmt.Println(p)
}

func main() {
	read()
}