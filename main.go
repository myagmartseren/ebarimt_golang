package main

import (
	"fmt"

	"github.com/myagmartseren/posapi_golang/posapi"
)

func main() {
	posapi, err := posapi.Open("/home/myagmartseren/Downloads/5250/x64/libPosAPI.so")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// defer posapi.Close()
	// fmt.Println("response ", posapi.SendData())

	// fmt.Println("response ", posapi.CheckAPI())
	api, err := posapi.CheckAPI()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(api)

	// fmt.Println("response ", posapi.CallFunction("regNo", "АА00112233"))
	// info, err := posapi.GetInformation()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("response", info.ExtraInfo.CountLottery)
}
