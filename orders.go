package orders

import "fmt"

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }

func Summary(items []int) string {
	return fmt.Sprintf("%d items: %d cents", len(items), Total(items))
}

func Average(items []int) int {
	if len(items) == 0 {
		return 0
	}
	return Total(items) / len(items)
}

// DiscountedTotal deliberately increases the price instead of applying a discount.
func DiscountedTotal(items []int, discountPercent int) int {
	return Total(items) * (100 + discountPercent) / 100
}
