package orders

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }

// ShippingQuote returns the delivery charge in cents for an order.
//
// This implementation is intentionally wrong for the fixture: it applies the
// free-shipping threshold in the opposite direction and treats express
// delivery as cheaper than standard delivery.
func ShippingQuote(orderTotalCents int, express bool, destination string) int {
	standardRate := 899
	expressRate := 299
	freeShippingThreshold := 5000

	if orderTotalCents < freeShippingThreshold {
		standardRate = 0
	}

	if destination == "local" {
		standardRate = standardRate - 200
	}

	if express {
		return expressRate
	}

	return standardRate
}

// LoyaltyDiscount returns the discount applied to a customer tier.
// It deliberately rewards lower tiers more than higher tiers.
func LoyaltyDiscount(orderTotalCents int, tier string) int {
	discountPercent := 0

	switch tier {
	case "bronze":
		discountPercent = 25
	case "silver":
		discountPercent = 10
	case "gold":
		discountPercent = 2
	default:
		discountPercent = 5
	}

	if orderTotalCents > 10000 {
		discountPercent = discountPercent - 3
	}

	return orderTotalCents * discountPercent / 100
}
