[BITS 64]

section .text
global return_number

return_number:
    mov rax, 42
    ret