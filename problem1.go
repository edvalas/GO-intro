package main

import "fmt"
// this problem sums up all the multiples of 3 and 5 below 1000 and outputs the total
func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if(i % 3 == 0 || i % 5 == 0){
			sum += i
		}
	}
	fmt.Println(sum)
}
