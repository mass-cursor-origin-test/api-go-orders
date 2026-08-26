package orders

func Total(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

// ShippingQuote returns the delivery charge in cents for an order.
// Standard delivery is 899 cents, express is 1499 cents. Orders at or above
// 5000 cents get free standard shipping. Local destinations receive a 200
// cent discount, floored at zero.
func ShippingQuote(orderTotalCents int, express bool, destination string) int {
	rate := 899
	if express {
		rate = 1499
	}

	if !express && orderTotalCents >= 5000 {
		rate = 0
	}

	if destination == "local" {
		rate -= 200
		if rate < 0 {
			rate = 0
		}
	}

	return rate
}

// LoyaltyDiscount returns the discount in cents for a customer tier.
// Higher tiers receive a larger percent. Orders over 10000 cents get an
// extra 3 percentage points.
func LoyaltyDiscount(orderTotalCents int, tier string) int {
	discountPercent := 0

	switch tier {
	case "bronze":
		discountPercent = 5
	case "silver":
		discountPercent = 10
	case "gold":
		discountPercent = 25
	}

	if orderTotalCents > 10000 {
		discountPercent += 3
	}

	return orderTotalCents * discountPercent / 100
}
