/*
  author: Lustrew
  update: 2019.07.29
  function: 字符串查找
    This algorithm tries to search the given pattern in the given text.
    If pattern is found from index position i in the text, 
    it is added to positions.
    
    Time Complexity : O(n*m)
      n = length of text
      m = length of pattern 
*/

package main

import (  
    "fmt"
    "flag"
)

func naivePatternSearch(text string, pattern string) []int {
  var positions []int 

  for i := 0;i < len(text);i++ {    
    match := true

    for j := 0; j < len(pattern); j++ {
      if text[i+j] != pattern[j] {
        match = false
        break
      }
    }

    if match {
      positions = append(positions, i)
    }
  }

  return positions
}

func main() {  
  text    := flag.String("t", "ABAAABCDBBABCDDEBCABC", "text")
  pattern := flag.String("p", "ABC", "pattern")
  flag.Parse()

  var positions []int
  positions = naivePatternSearch(*text, *pattern)

  if len(positions) == 0 {
    fmt.Printf("Pattern not found in given text!")
  } else {
    fmt.Printf("Pattern found in following position:\n")
    fmt.Printf("%v", positions)
  }

}