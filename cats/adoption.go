package cats

// Adoptable reports whether a cat is old enough to go home with someone.
func Adoptable(c Cat) bool {
	return c.Age >= 1 && c.Naps < 18
}
