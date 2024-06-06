package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include <string.h>

typedef int (*func_ptr)();

void* load_code(unsigned char* code, size_t code_size) {
    void* mem = mmap(NULL, code_size, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_ANON | MAP_PRIVATE, -1, 0);
    if (mem == MAP_FAILED) {
        perror("mmap");
        exit(EXIT_FAILURE);
    }
    memcpy(mem, code, code_size);
    return mem;
}

int execute_code(void* code) {
    func_ptr func = (func_ptr)code;
    return func();
}
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"unsafe"
)

func main() {
	code := []byte{
		0x48, 0xC7, 0xC0, 0x2A, 0x00, 0x00, 0x00, // mov rax, 42
		0xC3, // ret
	}

	codePtr := C.load_code((*C.uchar)(unsafe.Pointer(&code[0])), C.size_t(len(code)))
	if codePtr == nil {
		fmt.Println("Failed to load code into mem")
		return
	}

	memContent := (*[8]byte)(codePtr)
	fmt.Printf("Memory content: %x\n", *memContent)

	result := C.execute_code(codePtr)
	fmt.Printf("Result from in-memory exec: %d\n", result)

	objCode, err := os.ReadFile("asm/routine.bin")
	if err != nil {
		log.Fatalf("Failed to read obj file: %v", err)
	}

	objCodePtr := C.load_code((*C.uchar)(unsafe.Pointer(&objCode[0])), C.size_t(len(objCode)))
	if objCodePtr == nil {
		log.Fatalf("Failed to load obj file into mem")
	}

	objResult := C.execute_code(objCodePtr)
	fmt.Printf("Result from obj file in mem exec: %d\n", objResult)
}
