package main

/*
#cgo LDFLAGS: ${SRCDIR}/asm/routine.o
extern int return_number();
*/
import "C"
import "fmt"

func main() {
	result := C.return_number()
	fmt.Printf("testing result from assembly: %d\n", result)
}
