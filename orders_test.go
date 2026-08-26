package orders

import "testing"

func TestTotal(t *testing.T) {
	if Total([]int{2, 3}) != 5 {
		t.Fatal("unexpected total")
	}
}

func TestShippingQuote(t *testing.T) {
	tests := []struct {
		name        string
		total       int
		express     bool
		destination string
		want        int
	}{
		{name: "standard below free shipping", total: 4999, express: false, destination: "remote", want: 899},
		{name: "standard at free shipping threshold", total: 5000, express: false, destination: "remote", want: 0},
		{name: "standard above free shipping", total: 8000, express: false, destination: "remote", want: 0},
		{name: "express below threshold still charged", total: 1000, express: true, destination: "remote", want: 1499},
		{name: "express above threshold still charged", total: 8000, express: true, destination: "remote", want: 1499},
		{name: "local standard below threshold", total: 1000, express: false, destination: "local", want: 699},
		{name: "local standard with free shipping", total: 5000, express: false, destination: "local", want: 0},
		{name: "local express", total: 1000, express: true, destination: "local", want: 1299},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShippingQuote(tt.total, tt.express, tt.destination)
			if got != tt.want {
				t.Fatalf("ShippingQuote(%d, %v, %q) = %d, want %d", tt.total, tt.express, tt.destination, got, tt.want)
			}
		})
	}
}

func TestLoyaltyDiscount(t *testing.T) {
	tests := []struct {
		name  string
		total int
		tier  string
		want  int
	}{
		{name: "bronze", total: 2000, tier: "bronze", want: 100},
		{name: "silver", total: 2000, tier: "silver", want: 200},
		{name: "gold", total: 2000, tier: "gold", want: 500},
		{name: "unknown tier", total: 2000, tier: "none", want: 0},
		{name: "gold with large-order bonus", total: 20000, tier: "gold", want: 5600},
		{name: "bronze with large-order bonus", total: 11000, tier: "bronze", want: 880},
		{name: "unknown tier with large-order bonus", total: 11000, tier: "none", want: 330},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoyaltyDiscount(tt.total, tt.tier)
			if got != tt.want {
				t.Fatalf("LoyaltyDiscount(%d, %q) = %d, want %d", tt.total, tt.tier, got, tt.want)
			}
		})
	}
}
