package cmd

import "testing"

func TestExecute_NoPanic(t *testing.T) {
	// Test de non-panique lors de Execute
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute panics: %v", r)
		}
	}()
	Execute()
}
