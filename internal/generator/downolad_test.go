package generator

import "testing"

func TestUnzip_Errors(t *testing.T) {
	dir := t.TempDir()
	// Si le fichier n'existe pas
	if err := unzip("/path/does/not/exist.zip", dir); err == nil {
		t.Error("unzip should error if zip does not exist")
	}
}

func TestDownloadSpringProject_BadParams(t *testing.T) {
	dir := t.TempDir()
	// Simule un paramètre impossible
	params := map[string]string{"type": "zzzzzz-INVALID"}
	err := DownloadSpringProject(params, dir)
	if err == nil {
		t.Error("Devrait échouer avec des paramètres impossibles")
	}
}
