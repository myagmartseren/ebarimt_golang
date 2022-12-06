package main

import "C"

import (
	"fmt"

	"github.com/myagmartseren/posapi_golang/posapi"
)

func main() {

	// export_name := "checkApi"
	// lib_path := "/usr/lib/libPosAPI.so"

	// //Loading .so
	// handle := C.dlopen(C.CString(lib_path), C.RTLD_LAZY)
	// if handle == nil {
	// 	fmt.Println(lib_path + ":\tNOT FOUND")
	// 	return
	// } else {
	// 	fmt.Println(lib_path + ":\tSUCCESS")
	// }

	// //looking for function address
	// func_pointer := C.dlsym(handle, C.CString(export_name))
	// if func_pointer == nil {
	// 	fmt.Println(export_name + ":\tNOT FOUND")
	// 	return
	// } else {
	// 	fmt.Println(export_name+":\t", func_pointer)
	// }

	// fmt.Printf("%f", C.checkApi(func_pointer, 1))

	posapi, err := posapi.Open()
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
