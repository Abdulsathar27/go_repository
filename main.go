package main

import "fmt"

// Variadic function to sum numbers
func sum(numbers ...int) (int,bool) {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total,true
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
}

