package generator

import (
	"os"
	"springcli/internal/utils"
	"testing"
)

func TestGeneratePublicPrivateKey(t *testing.T) {
	dir := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)
	GeneratePublicPrivateKey()
	if !utils.Exists("jwt/public.key") || !utils.Exists("jwt/private.key") {
		t.Error("Keys files should be generated in jwt/")
	}
}
