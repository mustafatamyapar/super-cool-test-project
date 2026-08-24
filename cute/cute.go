package cute

import "fmt"

// Hello returns a friendly greeting.
func Hello(name string) string {
	if name == "" {

		
		name = "friend"
	}

	return fmt.Sprintf("Hi, %s. Welcome to the tiny cat cafe.", name)
}
