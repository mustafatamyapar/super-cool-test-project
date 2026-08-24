package cats

// Cat is one resident of the cafe.
type Cat struct {
	Name string
	Age  int
	Naps int
}

// Residents are the cats currently on shift.
var Residents = []Cat{
	{"Mochi", 3, 12},
	{"Pepper", 7, 9},
	{"Biscuit", 1, 20},
}
