package orders

import "github.com/mustafatamyapar/super-cool-test-project/internal/menu"

// Order is one customer's request, taken at the counter.
type Order struct {
	Table int
	Items []string
}

// Total adds up an order, skipping anything not on the menu.
func Total(o Order) int {
	sum := 0
	for _, name := range o.Items {
		if cents, ok := menu.Price(name); ok {
			sum += cents
		}
	}
	return sum
}
