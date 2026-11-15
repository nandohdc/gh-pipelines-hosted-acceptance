package fixture

import "testing"

func TestName(t *testing.T) {
	if got := Name(); got != "api" {
		t.Fatalf("Name() = %q, want api", got)
	}
}
