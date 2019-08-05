/*
  author: Lustrew
  update: 2019.07.29  
  function: insertion sort algorithm
*/

package main
import "fmt"

func main() {
  array := []int{2, 1, 4, 3, 5, 9, 7, 6, 8, 0}
  fmt.Println("Initial array is:", array)
  
  insertionSort(array)
  fmt.Println("Sorted array: ", array)
}

func insertionSort(array []int) {  
  for out := 1; out < len(array); out++ {
    temp := array[out]
    in := out
    for ; in > 0 && array[in-1] >= temp; in-- {
      array[in] = array[in-1]
    }
    array[in] = temp  
  }
}