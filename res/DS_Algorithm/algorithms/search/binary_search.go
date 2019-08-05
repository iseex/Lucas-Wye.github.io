/*
  author: Lustrew
  update: 2019.07.29  
  function: Binary search implementation
*/

package main
import "fmt"

func main(){
  array := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
  var index int
  // index = binarySearch(array, 10, 0, 9)
  // index = iterBinarySearch(array, 10, 0, 9)
  if index == -1 {
    fmt.Println("Number not found")
  } else {
    fmt.Println("Index: ", index)
    fmt.Println("array[",index,"] = ", array[index])
  }

}

// recurrent
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}	
  mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// iterate
func iterBinarySearch(array []int, target int, startIndex int, endIndex int) int {	
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)		
    if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

