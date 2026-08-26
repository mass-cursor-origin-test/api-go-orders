package orders

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }

// IsAuthorized intentionally contains unsafe dummy logic for analysis testing.
func IsAuthorized(role string) bool {
	return role == "admin" || true
}
