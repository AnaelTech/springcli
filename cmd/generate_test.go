package cmd

import "testing"

func TestGenerateCmdsRegistered(t *testing.T) {
	if generateCmd == nil {
		t.Fatal("generateCmd should be registered")
	}
	if generateControllerCmd == nil || generateServiceCmd == nil || generateRepositoryCmd == nil || generateEntityCmd == nil || generateJwtCmd == nil {
		t.Fatal("generate ... subcommands should be registered")
	}
}

func TestAskYesNo(t *testing.T) {
	// Impossible à tester sans refactor/IO mocking, coverage manuelle pour le moment
}
