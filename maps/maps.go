package maps

import "fmt"

func MapEntry() {
	// DeclaringMethodsForMap(3)
	// DeletingKeys()
	IterateOverMap(GetBasicColorsMap())
}

func GetBasicColorsMap() map[string]string {
	colors := map[string]string {
		"red": "hot",
		"yellow": "dirty",
		"green": "tabtabitab",
		"blue": "go",
		"orange": "bajrangbali",
	}
	return colors
}

func DeclaringMethodsForMap(method int) {
	switch method {
	case 1:
		var colors map[string]string
		fmt.Println("Zero value initialization using `var` keyword:", colors)
	case 2:
		colors := make(map[string]string)
		fmt.Println("Zero value initialization using `make` function:", colors)
	case 3:
		colors := make(map[string]string)
		colors["red"] = "hot"
		fmt.Println("Added key after initialization:", colors)
	default:
		fmt.Println("Invalid option.")
	}
}

func DeletingKeys() {
	colors := GetBasicColorsMap()
	delete(colors, "red")
	fmt.Println("Colors map after deleting `red` key:", colors)
}

func IterateOverMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, ":", value)
	}

}