package actions

import "testing"

func TestLXDOperationID(t *testing.T) {
	got := lxdOperationID("/1.0/operations/abcd-1234/")
	if got != "abcd-1234" {
		t.Fatalf("lxdOperationID() = %q, want %q", got, "abcd-1234")
	}
}
