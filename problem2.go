package main

import "fmt"
//this function determines if a number is prime
func IsPrime(number int) bool {
    
    for i := 2; i < number; i++ {
        if (number % i == 0 && i != number){
			return false
		} 
    }
    return true
}
//main loops over 10001 prime numbers and finds the 10001th prime number
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
