package main

import (
  "fmt"
)

func main() {
  
  // This solution will find the largest prime for the given number
  
  var a int
  
  // Create and initialize the maximum limit
  
  var number_to_be_factored = 13195
  
  // Create and initialize an empty slice to store all the factors
  factors := make([]int,1)
  factors[0] = 1
  
  primes := make([]int,1)
  primes[0] = 1
  
  for a = 2; a <= number_to_be_factored; a++ {
    if number_to_be_factored % a == 0 {
      factors = append(factors,a)
      fmt.Println(factors)
     
      
    }
  }
  
  
  
  fmt.Println(factors)
  //fmt.Println(primes)
  
}