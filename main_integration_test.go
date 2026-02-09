package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCLI_NewProjectIntegration(t *testing.T) {
	dir := t.TempDir()
	cmd := exec.Command("go", "run", "./main.go", "new", "unittestproj")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("CLI failed: %v\n%s", err, string(out))
	}
	// Vérifie que le dossier a été créé
	files, err := os.ReadDir(dir + "/unittestproj")
	if err != nil {
		t.Fatalf("Projet non généré: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("Projet généré vide !")
	}
}
