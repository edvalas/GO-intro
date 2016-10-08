package main

import (
	"fmt"
	"os"
	"strconv"
)


func collatz(n uint){
	for ; n != 1; {
		fmt.Print(n, " ")
		if (n % 2 == 0){
			n = n / 2
		}else{
			n = (3 * n) + 1
		}
	}
	fmt.Println(n)
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "[unit] [unit] ...")
	}else{
		for i:= 1; i < len(os.Args); i++{
			if u, err := strconv.ParseUint(os.Args[i], 10, 64); err == nil {
				collatz(uint(u))
			}else {
				fmt.Println(os.Args[i], "is not a positive number.")
			}	
		}
	}
}