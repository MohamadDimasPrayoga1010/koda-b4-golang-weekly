package utils

import (
	"fmt"
	"time"
)


func Alert(text string){
	for i := range text{
		fmt.Print(string(text[i]))
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(500 * time.Millisecond)
}