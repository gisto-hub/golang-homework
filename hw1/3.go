package main

import (
	"fmt"
	"os"
)

func Fibonacci2(n int) int {
    if n < 3 { 
    	return 1 
    }
    return Fibonacci2(n - 2) + Fibonacci2(n - 1)
}

func main(){
	var n int
    fmt.Fscan(os.Stdin, &n)
	fmt.Println(Fibonacci2(n))
}


