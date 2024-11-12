package main

import (
    "fmt"
    "math"
    "time"
)

// Function to check if a number is prime
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Function to calculate and print prime numbers
func generatePrimes() {
    num := 2
    for {
        if isPrime(num) {
            fmt.Println(num)
        }
        num++
    }
}

func main() {
    // Run prime number generation in a loop to stress the CPU
    go generatePrimes()
    
    // Keep the main goroutine running to avoid premature exit
    for {
        time.Sleep(time.Second)
    }
}
