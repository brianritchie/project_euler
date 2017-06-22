package main

import (
  "fmt"
)

func square_of_sum_func(limit int) int {
  
  var sum,a int
  
  for a = 1; a < limit +1; a++ {
    sum += a
  }
  
  return sum*sum
}

func sum_of_squares_func(limit int) int {
  
  var total,b int
  
  for b = 1; b < limit +1; b++ {
    total = total + b*b
  }
  
  return total
}

func main() {
    
  var natural_number_max = 100
  var difference int
  
  square_of_sum := square_of_sum_func(natural_number_max)
  fmt.Println("The Square of the sum is :", square_of_sum)
  
  sum_of_squares := sum_of_squares_func(natural_number_max)
  fmt.Println("The Sum of the squares is :", sum_of_squares)
  
  difference = square_of_sum - sum_of_squares
  fmt.Println("The difference is :", difference)
  
}