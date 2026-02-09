package utils

import "testing"

func TestExistsAndCreateFolder(t *testing.T) {
	dir := t.TempDir()
	if !Exists(dir) {
		t.Error("Temp dir should exist")
	}

	notExists := dir + "/notyet"
	if Exists(notExists) {
		t.Error("Should not exist before creation")
	}
	if err := CreateFolder(notExists); err != nil {
		t.Fatalf("CreateFolder error: %v", err)
	}
	if !Exists(notExists) {
		t.Error("Should exist after creation")
	}
	// Permissions/invalid path: dépend du système, souvent root/protected only
}
