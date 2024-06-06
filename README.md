# GoAsmLab

## Getting Started

### Prerequisites

- Go (Golang) installed
- NASM (Netwide Assembler) installed

### Setting Up the Project

1. Clone the repository:
    ```sh
    git clone https://github.com/HaoLiHaiO/GoAsmLab
    cd GoAsmLab
    ```

2. Assemble the assembly code:
    ```sh
    nasm -f bin -o asm/routine.bin asm/routine.asm
    ```

3. Build the Go project:
    ```sh
    go build -o main main.go asm/routine.o
    ```

4. Run the compiled program:
    ```sh
    ./main
    ```

## License

This project is licensed under the MIT License.