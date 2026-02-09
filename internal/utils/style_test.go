package utils

import "testing"

func TestPrintsDoNotPanic(t *testing.T) {
	// Cette suite vérifie juste l'exécution sans panic
	msgs := []string{"Hello, World!", "Erreur !", "Info", "Success"}
	for _, fn := range []func(string){PrintTitle, PrintError, PrintInfo, PrintWarning, PrintSuccess, PrintSubtitle, PrintPrompt, PrintBox, PrintStep} {
		for _, msg := range msgs {
			fn(msg)
		}
	}
}
