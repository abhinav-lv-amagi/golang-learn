# Learning Go

[Udemy Course by Stephen Grider](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797280#overview)

## 12/02/2025

- [x] Packages
- [x] Variables
- [x] Functions
- [x] Receiver functions
- [x] Multiple return values in Go functions

Notes:

- Packages are groupings of code, but don't actually require separate folders. For eg, we can use code from the same package, but from different files - without importing the necessary functions / variables.

  ```go
  // main.go
  package main

  import "fmt"

  func main() {
    var state string = getState() // defined in 'state.go'
    fmt.Println(state)
  }

  // state.go
  package main  // present in the same 'main' package

  func getState() string {
    return "stateful widget"
  }

  ```

- Variables can be initialized in two ways:

  ```go
  // Standard way
  var b string = "Rishith"

  // Note: This cannot be used at the package level, only inside functions
  a := "Abhinav"
  ```

- A slice is kind of a dynamic array

  ```go
  nums := []int{1, 2, 3, 4}
  nums = append(nums, 5)
  fmt.Println(nums)
  ```

- Functions can return multiple values

  ```go
  func getRGB() (string, string, string) {
    return "red", "green", "blue"
  }
  ```

- We can make custom types as such:

  ```go
  type deck []string
  var cards deck = deck{"Three of Spades", "Five of Diamonds"}
  ```

- We can attach receiver functions to custom types as such:

  ```go
  func (d deck) print() {
    for _, card: range d {
        // here 'card' is of type 'string'
        fmt.Println(card)
    }
  }
  var cards deck = deck{"Spades", "Clubs"}
  // Call the function from a deck instance
  cards.print()
  ```

  > See if we can do this to primitive types in Go as well.
