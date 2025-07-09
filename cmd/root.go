package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue   = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan   = "\033[36m"
    ColorWhite  = "\033[37m"
)

var rootCmd = &cobra.Command{
	Use:   "springcli",
	Short: "CLI pour créer et gérer des projets Spring Boot",
	Long: `SpringCLI est un outil en ligne de commande pour générer,
configurer et gérer des projets Spring Boot en utilisant Spring Initializr
et des modèles personnalisés.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
███████ ██████  ██████  ██ ███    ██  ██████      ██████ ██      ██ 
 ██      ██   ██ ██   ██ ██ ████   ██ ██           ██     ██      ██ 
 ███████ ██████  ██████  ██ ██ ██  ██ ██   ███     ██     ██      ██ 
      ██ ██      ██   ██ ██ ██  ██ ██ ██    ██     ██     ██      ██ 
 ███████ ██      ██   ██ ██ ██   ████  ██████      ██████ ███████ ██ `)
    fmt.Println("Bienvenue dans SpringCLI 🚀.\n")
    fmt.Printf("Utilisez %s--help%s pour voir les commandes disponibles.\n", ColorYellow, ColorReset)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(generateCmd)
	}
