package main

/*
#cgo CFLAGS: -I${SRCDIR}/ctestlib
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/ctestlib
#cgo LDFLAGS: -L${SRCDIR}/ctestlib -ltest

#include <test.h>
*/
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

type User struct {
	Username string
	Visits   int
}

type Status int

const (
	Pending Status = iota
)

func evenNumberCallback(num int) {
	fmt.Print("odd number: ", num, " ")
}

func userCallback(user unsafe.Pointer) {
	(*User)(user).Visits++
}

func convertCStringToGoString(c *C.uchar) string {
	return C.GoString((*C.char)(unsafe.Pointer(c)))
}

func main() {
	fmt.Println("== convertCStringToGoString")
	fmt.Println(convertCStringToGoString(C.get_unsigned_char()))

	fmt.Println("== Numbers")
	sum := int(C.sum(C.int(1), C.int(2)))
	fmt.Print(sum, "\n")

	fmt.Println("== Get string")
	fmt.Println(C.GoString(C.get_string()))
	stringBytes := C.GoBytes(unsafe.Pointer(C.get_string()), 24)
	fmt.Println(stringBytes[0:bytes.Index(stringBytes, []byte{0})])

	fmt.Println("== Send string")
	cStr := C.CString("lorem ipsum")
	C.print_string(cStr)
	C.free(unsafe.Pointer(cStr))

	fmt.Println("== Send byte array")
	data := []byte{1, 4, 2}
	cBytes := (*C.uchar)(unsafe.Pointer(&data[0]))
	cBytesLength := C.size_t(len(data))
	fmt.Print("bytes: ")
	C.print_buffer(cBytes, cBytesLength)

	fmt.Println("== Get and pass struct")
	point := C.struct_point{}
	point.x = 0
	point.y = 2
	fmt.Println(point)
	fmt.Print(C.point_diff(point), "\n")

	// Arbitrary data: unsafe.Pointer to void pointer
	fmt.Println("== Pass void pointer")
	C.pass_void_pointer(unsafe.Pointer(&point.y))

	// Enum
	fmt.Println("== Access enum")
	fmt.Print(C.enum_status(Pending) == C.PENDING, C.PENDING, C.DONE, "\n")

	fmt.Println("== Pass callback")
	c := registerCallback(evenNumberCallback, nil)
	C.generate_numbers(5, c)
	fmt.Println()
	unregisterCallback(c)

	// Callback with params
	user := User{
		Username: "johndoe",
	}
	cWithParams := registerCallback(userCallback, unsafe.Pointer(&user))
	C.user_action(cWithParams)
	unregisterCallback(cWithParams)
	fmt.Println(user)
}
