package billing

// vatPermille is the German reduced rate for food, in permille.
const vatPermille = 70

// WithVAT adds VAT to a net amount in cents.
func WithVAT(cents int) int {
	return cents + cents*vatPermille/1000
}
