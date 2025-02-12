# Learning Go

[Udemy Course by Stephen Grider](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797280#overview)

## 12/02/2024

- [x] Variables
- [x] Functions
- [x] Receiver functions
- [x] Multiple return values in Go functions

Notes:

- Variables can be initialized in two ways:

  ```go
  a := "Abhinav"
  var b string = "Rishith"
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
