package orders
import "testing"
func TestTotal(t *testing.T) { if Total([]int{2, 3}) != 5 { t.Fatal("unexpected total") } }
