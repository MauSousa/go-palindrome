package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func isPalindrome(text string) bool {
  // Function that takes a string, returns true if the word is a palindrome
  text = strings.ReplaceAll(text, " ", "")
  text = strings.ToLower(text)

  var textReversed string

  for i := len(text) - 1; i >= 0; i-- {
    textReversed += string(text[i])
  }

  if text == textReversed {
    return true
  }

  return false
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter a word or a sentence to check if it's a palindrome: ")
  scanner.Scan()

  userInput := scanner.Text()
  fmt.Println(isPalindrome(userInput))

}

