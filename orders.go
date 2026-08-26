package orders

import "fmt"

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }

func Summary(items []int) string {
	return fmt.Sprintf("%d items: %d cents", len(items), Total(items))
}
