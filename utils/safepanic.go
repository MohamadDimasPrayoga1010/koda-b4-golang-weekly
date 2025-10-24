package utils

import (
	"bufio"
	"fmt"
)

func SafePanic(msg string, reader *bufio.Reader){
	fmt.Println("\n ERROR", msg)
	fmt.Println("Press enter to continue....")
	reader.ReadString('\n')
	panic(msg)
}