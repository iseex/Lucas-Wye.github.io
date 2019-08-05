/*
  author: Lustrew
  update: 2019.07.29  
  function: selection sort algorithm
*/

package main
import "fmt"

func main() {
  array := []int{11, 12, 31, 4, 1}
  fmt.Println("Initial array is:", array)

  selectionSort(array)
  fmt.Println("Sorted array: ", array)
}

func selectionSort(array []int) {
  var min int
  // 从小到大排序
  for i := 0; i < len(array) - 1; i++ {
    min = i
    for j := i + 1; j < len(array); j++{
      if array[j] < array[min] {
        min = j
      }
    }
    Swap(array,i,min)
  }
}

func Swap(array []int, i, j int) {
  array[i] ^= array[j]
  array[j] ^= array[i]
  array[i] ^= array[j]
}
