package menu

// Item is one thing you can order at the cafe.
type Item struct {
	Name  string
	Cents int
}

// Menu is today's list, in the order it is printed on the board.
var Menu = []Item{
	{"espresso", 220},
	{"flat white", 320},
	{"catpuccino", 380},
	{"tuna toastie", 540},
}

// Price returns the price of an item, and whether we sell it at all.
func Price(name string) (int, bool) {
	for _, it := range Menu {
		if it.Name == name {
			return it.Cents, true
		}
	}
	return 0, false
}
