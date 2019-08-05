/*
  author: Lustrew
  update: 2019.07.29  
  function: selection sort algorithm
*/

package main
import (
  "fmt"
  "math/rand"
)

func main() {
	arr := RandomArray(11)
  fmt.Println("Initial array is:", arr)
  
  shellSort(arr)
  fmt.Println("Sorted array is: ", arr) 	
}

func shellSort(arr []int){
  for d := int(len(arr)/2); d > 0; d /= 2 {
    for i := d; i < len(arr); i++ {
      for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
        arr[j], arr[j-d] = arr[j-d], arr[j]
      }
    }
  }
}

func RandomArray(n int) []int {
  arr := make([]int, n)
  for i := 0; i <= n - 1; i++ {
    arr[i] = rand.Intn(n)
  }
  return arr
}