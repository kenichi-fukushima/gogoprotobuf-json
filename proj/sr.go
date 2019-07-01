package main

import (
	"bufio"
	"fmt"
	"os"

	// "github.com/golang/protobuf/jsonpb"
	"github.com/gogo/protobuf/jsonpb"

	"example.com/api"
)

func main() {
	file, err := os.Open("./msg.json")
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}
	plan := &api.MyMessage{}
	err = jsonpb.Unmarshal(bufio.NewReader(file), plan)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}
	fmt.Println("msg " + plan.String())
	return
}
