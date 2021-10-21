package db

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	client1 := New()
	client2 := New()

	// object should be same
	if client1 != client2 {
		t.Error("Objects have different references.")
	}
}