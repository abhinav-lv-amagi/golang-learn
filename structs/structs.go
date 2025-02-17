package structs

import "fmt"

/*
 * NOTE: In Go, identifiers (functions, structs, variables)
 * are exported if they start with an uppercase letter.
 */

type Person struct {
	firstName string
	lastName  string
}

type ContactInfo struct {
	email string
	zip int
}

type NewPerson struct {
	firstName string
	lastName string
	contact ContactInfo
	// instead of the above way, we can also merge the 
	// names of the property and type together, for eg:
	// ContactInfo // This translates to ContactInfo ContactInfo
}

func (p NewPerson) introduce() {
	fmt.Println(p.firstName, p.lastName,"is a wonderful person!")
	fmt.Println("Here's how you can contact them:")
	fmt.Println("Email address:", p.contact.email)
	fmt.Println("Zip code:", p.contact.zip)
}

func AlexAbhinav() {
	// Method 1
	alex := Person{"Alex", "Anderson"}

	// Method 2
	abhinav := Person{
		firstName: "Abhinav",
		lastName:  "LV",
	}
	fmt.Println(alex, abhinav)
	
	// Zero value (what will it be for a custom struct?)
	// Ans: Zero values of the fields inside (ultimately primitive data types)
	var dhruv Person
	// fmt.Println(dhruv)
	fmt.Printf("%+v\n", dhruv)
}

func EmbeddedStructs() {
	jim := NewPerson {
		firstName: "Jim",
		lastName: "Party",	// Need to give ',' in multi line structs
		contact: ContactInfo{
			email: "jim@gmail.com",
			zip: 94000,
		},
	}

	// fmt.Printf("%+v\n", jim)
	jim.introduce()
}

func Execute() {
	// EmbeddedStructs()
	WhyPointers()
}

func (p NewPerson) updateName(newName string) {
	p.firstName = newName;
	fmt.Println("New name:", p.firstName);
}

func (pointerToPerson *NewPerson) actuallyUpdateName(newName string) {
	(*pointerToPerson).firstName = newName;
	fmt.Println("New name:", pointerToPerson.firstName)
}

func WhyPointers() {
	max := NewPerson{
		firstName: "Max",
		lastName: "Verstappen",
		contact: ContactInfo{
			email: "landoisajoke@verstappen.com",
			zip: 696969,
		},
	}
	max.introduce()
	fmt.Println("-----------------")
	max.updateName("Maxipad") // This won't update this instance, unlike primitive types
	fmt.Println("-----------------")
	max.introduce()

	fmt.Println("-----------------")
	// But this will work:
	(&max).actuallyUpdateName("Maxipad")
	fmt.Println("-----------------")
	max.introduce() // Maxipad verstappen! When will RPM fucking upload :(
}