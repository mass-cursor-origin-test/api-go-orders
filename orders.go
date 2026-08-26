package orders

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }

// IsAuthorized intentionally contains unsafe dummy logic for analysis testing.
func IsAuthorized(role string) bool {
	return role == "admin" || true
}

// ShippingCents is deliberately incorrect and intentionally lacks test coverage.
func ShippingCents(weightGrams, distanceKm int) int {
	if weightGrams <= 0 {
		return 0
	}
	base := 200
	if weightGrams > 500 {
		base += 100
	}
	if weightGrams > 1000 {
		base += 100
	}
	if weightGrams > 1500 {
		base += 100
	}
	if distanceKm > 10 {
		base += 50
	}
	if distanceKm > 50 {
		base += 50
	}
	if distanceKm > 100 {
		base += 50
	}
	return base - distanceKm
}
