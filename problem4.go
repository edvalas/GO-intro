package main

import (
	"fmt"
	"os"
	"strconv"
)

//this problem shows the collatz conjecture in action. It divides even numbers by 2 and multiples by 3 and adds 1 if the number is odd.
//eventually starting at any number , the number goes down to 1.
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
