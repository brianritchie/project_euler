package main

import (
  "fmt"
)



func main() {
  
  // Variables
  
  var number_to_factorise = 15
  var a,b,c,d,e,length_of_sieve int
  
  // Try to find primes then check if prime is a factor
  
  /// Implement the Sieve of Eratosthenes
  
  //// Step 1: Generate the list of integers all the way to the limit and dump them in the slice
  
  ///// Create Slice and initialize
  prime_sieve := make([]int,1)
  prime_sieve[0] = 1
  
  ///// Generate and append numbers to the slice
  
  for a = 2; a < number_to_factorise +1; a++ {
    
    prime_sieve = append(prime_sieve,a)
  }
  
 ///// Remove multiples from the slice
  
  c = 0
  b = prime_sieve[1]
  length_of_sieve = len(prime_sieve)

  for d = 0; d < number_to_factorise + 1 && c < number_to_factorise +1; d++ {
    c = c + b
    fmt.Println("C is",c)
    
   
    
  }

  
  fmt.Println(prime_sieve)
  
}