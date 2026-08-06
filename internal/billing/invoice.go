package billing

import "fmt"

// Invoice renders a total in euros, the way it prints on the receipt.
func Invoice(table int, cents int) string {
	return fmt.Sprintf("table %d: %d.%02d EUR", table, cents/100, cents%100)
}
