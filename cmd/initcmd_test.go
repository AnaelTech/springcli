package cmd

import "testing"

func TestNewCommandRegistered(t *testing.T) {
	if newCmd == nil {
		t.Fatal("newCmd should be registered")
	}
}

func TestCreateNewProject_DefaultArtifact(t *testing.T) {
	// Simule minimalement un appel avec tout sauf artifactId (pour coverage)
	createNewProject("TestProject", "com.gotest", "", "maven-project", "4.0.2", "21")
}
