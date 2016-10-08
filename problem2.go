package main

import "fmt"

func IsPrime(number int) bool {
    
    for i := 2; i < number; i++ {
        if (number % i == 0 && i != number){
			return false
		} 
    }
    return true
}

func main() {
	var ans int
	count := 0
	
	for i := 2; count < 10001; i++ {
		if(IsPrime(i)){
			count++
			ans = i
		}
	}
	
	fmt.Printf("Count %v, Answer %v", count, ans)
}