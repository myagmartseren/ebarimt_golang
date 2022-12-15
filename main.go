package main

import (
	"fmt"

	"github.com/myagmartseren/posapi_golang/ebarimt"
)

func main() {
	temp, err := ebarimt.Init()
	if err != nil {
		panic(err)
	}
	result, err := temp.CallFunction("regNo", "АА00112233")
	if err != nil {
		fmt.Println("panic at result")
		panic(err)
	}
	fmt.Println(result)
}
