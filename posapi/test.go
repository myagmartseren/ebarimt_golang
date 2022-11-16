package posapi

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
*/
import (
	"C"
)
import (
	"fmt"
)

func Test() {
	handle := C.dlopen(C.CString("./libPosAPI.so"), C.RTLD_LAZY)
	bar := C.dlsym(handle, C.CString("checkApi"))
	fmt.Printf("bar is at %v\n", bar)
	// C.checkApi()
}
