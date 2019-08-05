/*
  author: Lustrew
  update: 2019.07.29  
  function: Caesar encrytion and decryption.
*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

func main() {
  // flag用法  
  // go run *.go -c 参数  
  cipherKey := flag.Int("c", 1, "Cipher shift amount is (-26, 26) ") // 若无"-c 参数"则取默认值 1
  input := flag.String("i", "zaft", "Input string")
  flag.Parse()

  fmt.Println("cipherKey is ",*cipherKey)
  fmt.Println("input is ",*input)

  if (*cipherKey > 26 || *cipherKey < -26) {
    flag.PrintDefaults()
    fmt.Println("Input again!")
  } else {
    fmt.Println("Start calculation.")
    fmt.Println(caesarCipher(*input, *cipherKey))
  }

}

func caesarCipher(input string, key int) string {
  var outputBuffer bytes.Buffer
  if key < 0 {
    key += 26
  }

  for _, r := range strings.ToLower(input) {
  	newByte := int(r)  	
    if newByte >= 'a' && newByte <= 'z' {
  		newByte -= 'a'
      newByte = (newByte + key) % 26
      newByte += 'a'
  	}
    outputBuffer.WriteString(string(newByte))
  }
  return outputBuffer.String()
}

