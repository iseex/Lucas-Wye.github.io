/*
  author: Lustrew
  update: 2019.07.29  
  function: merge sort algorithm
*/

package main
import "fmt"

func main () {
  array := []int{11, 12, 31, 4, 1}
  fmt.Println("Initial array is:", array)

  Mergesort(array)
  fmt.Println("Sorted array: ", array)  
}

func merge(a []int, b []int) []int {
  r := make([]int, len(a) + len(b))
  i, j := 0, 0  

  for i < len(a) && j < len(b) {
    if a[i] <= b[j] { 
      r[i+j] = a[i]; i++; 
    } else { 
      r[i+j] = b[j]; j++; 
    } 
  }

  for i < len(a) { r[i+j] = a[i]; i++ }
  for j < len(b) { r[i+j] = b[j]; j++ }

  return r  
}

func Mergesort(items []int) []int {
  if len(items) < 2 {
    return items    
  }

  middle := len(items) / 2
  a := Mergesort(items[:middle])
  b := Mergesort(items[middle:])
  
  return merge(a, b)  
}
