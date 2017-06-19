package main

import (
 // "io/ioutil"
  "fmt"
)

// This function accepts two integer inputs and the maximum range from the command line

func input() (int,int,int) {
  
  var a,b,c int
  
  fmt.Println("Please input the first integer: ")
  fmt.Scan(&a)
  fmt.Println("Please input the second integer: ")
  fmt.Scan(&b)
  fmt.Println("Please input the maximum range for all multiples: ")
  fmt.Scan(&c)
  
  return a, b, c
}

// This function computes the multiples of each number below the range

func mathmagic(first int,second int,limit int) int {

  var total_sum_of_multiples, total_sum_of_error_multiple int
  
  //Declare an empty slive to store all the multiples
  multiple_slice := make([]int, 2)
  
  //Initialize slice with the two integers who are multiples of themselves
  multiple_slice[0]=first
  multiple_slice[1]=second
  
  //fmt.Println(multiple_slice)
  
  // Generate range of natural numbers from first int to limit
  for a := first; a < limit+1; a++ {
    //fmt.Println(a)
    
    // Check if the  natural number is a multiple of the first integer
    if a%first == 0 && a != first && a!= limit {
      multiple_slice = append(multiple_slice, a)
      //fmt.Println(multiple_slice)
      }
    
  } 
  
   for b := second; b < limit+1; b++ {
    //fmt.Println(b)
    
    // Check if the natural number is a multiple of the second integer
    if b%second == 0 && b != second && b!= limit {
      multiple_slice = append(multiple_slice, b)
      //fmt.Println(multiple_slice)
      }
    
  }
  
  
  // Sum all integers in the slice generated
  
  for c:=0; c < len(multiple_slice); c++ {
   
    total_sum_of_multiples += multiple_slice[c]
  }
  
   //Computing the multiples of 15 that need to be removed from the sum
  
    // Declare empty slice for 15 and initialize
    error_multiple := make([]int,1)
    error_multiple[0] = 15
  
    //Compute the multiples for 15
    for d:=15; d < limit+1; d++ {
      if d%15 ==0 && d != 15 && d != limit {
        error_multiple = append(error_multiple,d)
      }
    }
  
    //Sum all the integers for the error multiple
    for e:=0; e < len(error_multiple); e++ {
   
    total_sum_of_error_multiple += error_multiple[e]
  }
  
  total_sum_of_multiples = total_sum_of_multiples - total_sum_of_error_multiple
  
  return total_sum_of_multiples
}


// The function below accepts two numbers and a maximum limit, finds all the multiples between the range and sums them up
func main() {
  
  fmt.Println("Natural Numbers Sum Between Two #'s and a Range")
  first_integer, second_integer, range_limit :=input()
  //fmt.Println("You said - ", first_integer, "and", second_integer, "the maximum range limit of", range_limit)
  euler_solution :=mathmagic(first_integer, second_integer, range_limit)
  fmt.Println("The total sum of all multiples of",first_integer,"and",second_integer,"below",range_limit,"is", euler_solution)
}