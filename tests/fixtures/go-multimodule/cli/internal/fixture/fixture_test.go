package fixture

import "testing"

func TestName(t *testing.T) {
	if got := Name(); got != "cli" {
		t.Fatalf("Name() = %q, want cli", got)
	}
}
