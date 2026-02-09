package main

import "testing"

func TestMainEntry(t *testing.T) {
	// Ce test vérifie simplement qu'une exécution du main ne panic pas
	// (ne pas tester l'effet complet de cmd.Execute ici)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main panics: %v", r)
		}
	}()
	main()
}
