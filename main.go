package main

import (
	"fmt"

	"github.com/myagmartseren/posapi_golang/posapi"
)

func main() {
	posapi, err := posapi.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// defer posapi.Close()
	// fmt.Println("response ", posapi.SendData())
	fmt.Println("response ", posapi.CheckAPI())

	// fmt.Println("response", posapi.GetInformation())
}
