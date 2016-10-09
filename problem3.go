package main

import "fmt"
//this problem takes in input from keyboard from the user and reverses the input string
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)
	
    fmt.Print(Reverse(input))
}
