package orders
import "testing"
func TestTotal(t *testing.T) { if Total([]int{2, 3}) != 5 { t.Fatal("unexpected total") } }

func TestIntentionalFailure(t *testing.T) {
	if Total([]int{2, 3}) != 999 {
		t.Fatal("intentional failure for CI testing")
	}
}
