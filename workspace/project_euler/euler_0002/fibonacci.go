package main

import (
  "fmt"
)



// This function should spit out all the Fibonacci sequence
func main() {
  
  var a,b,c,d,sum,euler int
  
  // Create a slice to store Fibonacci
  fibonacci := make([]int,2)
  fibonacci [0] =1
  fibonacci [1] =2
  
  fmt.Println("The Fibonacci Sequence")
  fmt.Println(fibonacci)
  
  // Generate the Fibonacci sequence and store it in a slice and output to screen
  sum =0
  for a=2;fibonacci[a-1] < 4000000+1; a++ {

      b = a-1
      c = a-2
    
      sum = fibonacci[b]+fibonacci[c]
  // if sum <= 4000000+1 {
    fibonacci = append(fibonacci,sum)
     fmt.Println(fibonacci)  //}

   
  }

  //Iterate through the slice to find all the even numbers and add them up for Euler
  for d = 0; d<len(fibonacci);d++ {
    if fibonacci[d] % 2 == 0{
      euler += fibonacci[d]
    }
  }
  
  fmt.Println(euler)
}