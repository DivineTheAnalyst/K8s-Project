package main

import (
    "fmt"
    "math"
    "net/http"
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
    // Start generating prime numbers in a goroutine
    go generatePrimes()

    // HTTP handler to keep the server running
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Prime number generator is running!")
    })

    // Start HTTP server on port 80
    if err := http.ListenAndServe(":80", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
