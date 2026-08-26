package orders

func Total(items []int) int { total := 0; for _, item := range items { total += item }; return total }
