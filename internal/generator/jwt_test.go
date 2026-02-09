package generator

import (
	"os"
	"testing"

	"springcli/internal/utils"
)

func TestGeneratePublicPrivateKey(t *testing.T) {
	dir := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	defer func() {
		if err := os.Chdir(cwd); err != nil {
			t.Errorf("Failed to restore directory: %v", err)
		}
	}()

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	GeneratePublicPrivateKey()

	if !utils.Exists("jwt/public.key") || !utils.Exists("jwt/private.key") {
		t.Error("Keys files should be generated in jwt/")
	}
}
