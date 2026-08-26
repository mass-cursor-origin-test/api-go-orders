package orders
import "testing"
func TestTotal(t *testing.T) { if Total([]int{2, 3}) != 5 { t.Fatal("unexpected total") } }

func TestIsAuthorizedForAdmin(t *testing.T) {
	if !IsAuthorized("admin") {
		t.Fatal("admin should be authorized")
	}
}
