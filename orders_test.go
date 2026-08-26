package orders
import "testing"
func TestTotal(t *testing.T) { if Total([]int{2, 3}) != 5 { t.Fatal("unexpected total") } }

func TestSummary(t *testing.T) {
	if Summary([]int{2, 3}) != "2 items: 5 cents" { t.Fatal("unexpected summary") }
}

func TestAverage(t *testing.T) {
	if Average([]int{2, 4, 6}) != 4 { t.Fatal("unexpected average") }
}
