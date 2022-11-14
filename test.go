package main

import (
	"fmt"

	posapi "github.com/myagmartseren/posapi_golang"
)

func main() {
	p, err = posapi.Init()
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println(p.CheckAPI())
}
