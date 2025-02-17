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

## 17/02/2025

- [x] Zero values
- [x] Go packages (Capitalize to export, etc.)
- [x] File read/write
- [x] Generate random values using time as seed
- [x] Writing test files
- [x] Structs
- [x] Pointers
- [x] Values Types and Reference Types

### Go packages

- Don't confuse the file structure of your project with the package structure
- Packages are logical groupings, but package names can be different to the folder name they live in. Eg: In the card deck project, the actual package is `deck`, but it is present in the `deck-of-cards` folder.
- Only files corresponding to a single package can exist within a directory (not nested). For eg:

  ```go
  // deck.go (in /deck-of-cards)
  package deck
  func blahblahDeck() {}

  // state.go (in /deck-of-cards)
  package state
  func blahblahState() {}
  ```

  will give a compiler and LSP error.

- When you want to expose functions in your package, their names must start with a capital letter, eg: `SaveFileToDeck`. Else they can't be accessed from other packages.

### File read/write

- Write

  ```go
  import "os"

  func saveToFile(filename string) error {
    // 0666 => perm mode: all users can read and write
    // (think chmod 666)
    return os.WriteFile(filename, []byte(d.ToString()), 0666)
  }
  ```

- Read
  ```go
  func readFromFile(filename string) deck {
    bs, err := os.ReadFile(filename)
    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
    }
    return deck(strings.Split(string(bs), ","))
  }
  ```

### Generate random values using time seed

Go documentation: [math/rand](https://pkg.go.dev/math/rand)

```go
func (d deck) Shuffle() {
	for i := range d {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// swapIdx := r.Int() % len(d) // Instead we can do this:
		swapIdx := r.Intn(len(d))
		d[i], d[swapIdx] = d[swapIdx], d[i]
	}
}
```

### Writing test files

- Naming rule: For testing a file called `deck.go`, name your file `deck_test.go`. Learn about testing in depth later.

```go
import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	expected_length := 16

	if len(d) != expected_length {
		t.Errorf("Expected deck length of %d, but got: %d", expected_length, len(d))
	}
}
```

### Structs

- Create a new struct

  ```go
  type ContactInfo struct {
    email string
    zip int
  }
  type Person struct {
    firstName string
    lastName string
    contact ContactInfo
  }
  ```

  We can also merge the property name and type, like so:

  ```go
  type Person struct {
    firstName string
    lastName string
    ContactInfo // property name: ContactInfo, of type: ContactInfo
  }
  ```

- Initiating a struct

  ```go
  // Method 1:
  // same order of properties as in struct declaration
  alex := Person{"Alex", "Anderson"}

  // Method 2:
  abhinav := Person{
  	firstName: "Abhinav",
  	lastName:  "LV",
  }

  // Method 3: (zero value initialization)
  var dhruv Person
  fmt.Printf("%+v", dhruv)
  ```

### Pointers

- Similar to C, like `&identifier` gives you the address and `*pointer` will give you the value.

- Go is a `pass-by-value` language. This means, EVERY time you pass any argument to a function (whether it's a pointer or a value or something else), it WILL GET COPIED.

- But SLICES (`[]string`) are custom data structures, like a slice is basically a structure that stores:

  - length
  - capacity
  - head address (of the array which stores the data in the slice)

  So when we pass a slice, the thing that gets copied is the slice structure, and not the content it's pointing to.

- When we pass a struct to a function, the entire struct gets copied, including all of its properties and nested structs.

- So when we need to update a struct, say using a receiver function, we need to use pointers like so:

  ```go
  func (personPointer *Person) updateName(newName string) {
    (*personPointer).firstName = newName;
    fmt.Println("New name:", personPointer.firstName)
  }

  func bruh() {
    jim := Person{"blah", "blah"}
    (&jim).updateName("jimmy") // can also use: jim.updateName("jimmy")
    fmt.Println(jim.firstName) // "jimmy"
  }
  ```

  As we see above, we can also use `jim.updateName` instead of `(&jim).updateName`, since Go will pass the required type based on the function signature.

### Value type and Reference type

Explained in pointers (read Slice). Refer to this table:

| Value Types    | Reference Types        |
| -------------- | ---------------------- |
| int            | \*int (`Pointer`)      |
| float64        | []int (`Slice`)        |
| bool           | map[string]int (`Map`) |
| string         | chan int (`Channels`)  |
| array ([5]int) | func() (`Function`)    |
| struct         | interface{}            |
