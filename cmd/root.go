package cmd

import (
	"fmt"
	"os"
	"strings"

	/* "github.com/charmbracelet/lipgloss" */
	"springcli/internal/utils"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "springcli",
	Short: "CLI pour créer et gérer des projets Spring Boot",
	Long: `SpringCLI est un outil en ligne de commande pour générer,
configurer et gérer des projets Spring Boot en utilisant Spring Initializr
et des modèles personnalisés.`,
	Run: func(cmd *cobra.Command, args []string) {
		ver, _ := cmd.Flags().GetBool("version")
		if ver {
			displayVersion()
			os.Exit(0)
		}
		displayWelcomeScreen()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(utils.ErrorStyle.Render("❌ " + err.Error()))
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.Flags().BoolP("version", "v", false, "Affiche la version de SpringCLI")
}

// ==================== FONCTIONS D'AFFICHAGE ====================

func displayVersion() {
	fmt.Println()
	versionBox := strings.Builder{}
	versionBox.WriteString(utils.IconStyle.Render("🍃 ") + utils.MainTitleStyle.Render("SpringCLI"))
	versionBox.WriteString("\n")
	versionBox.WriteString(utils.VersionStyle.Render("Version " + version))
	versionBox.WriteString("\n\n")
	versionBox.WriteString(utils.CommandDescStyle.Render("Un outil moderne pour gérer vos projets Spring Boot"))

	fmt.Println(utils.HelpBoxStyle.Render(versionBox.String()))
}

func displayWelcomeScreen() {
	fmt.Println()

	// Logo ASCII stylisé
	logo := `
 ███████ ██████  ██████  ██ ███    ██  ██████      ██████ ██      ██ 
 ██      ██   ██ ██   ██ ██ ████   ██ ██           ██     ██      ██ 
 ███████ ██████  ██████  ██ ██ ██  ██ ██   ███     ██     ██      ██ 
      ██ ██      ██   ██ ██ ██  ██ ██ ██    ██     ██     ██      ██ 
 ███████ ██      ██   ██ ██ ██   ████  ██████      ██████ ███████ ██ `

	fmt.Println(utils.LogoStyle.Render(logo))

	// Message de bienvenue
	fmt.Println(utils.WelcomeStyle.Render("🚀 Bienvenue dans SpringCLI"))
	fmt.Println(utils.VersionStyle.Render("Version " + version))
	fmt.Println()

	// Description
	fmt.Println(utils.DescriptionStyle.Render(
		"SpringCLI est un outil en ligne de commande moderne pour générer, " +
			"configurer et gérer des projets Spring Boot. Il utilise Spring Initializr " +
			"et propose des modèles personnalisés pour accélérer votre développement."))

	// Fonctionnalités principales
	// displayFeatures()

	// Commandes disponibles
	displayAvailableCommands()

	// Aide
	displayHelp()
}

// func displayFeatures() {
// 	fmt.Println(utils.CommandHeaderStyle.Render("✨ Fonctionnalités principales"))
//
// 	// Création de deux colonnes pour les fonctionnalités
// 	leftFeatures := strings.Builder{}
// 	leftFeatures.WriteString(utils.IconStyle.Render("🏗️  ") + "Génération de projets\n")
// 	leftFeatures.WriteString("   Spring Boot rapide\n\n")
// 	leftFeatures.WriteString(utils.IconStyle.Render("⚙️  ") + "Configuration avancée\n")
// 	leftFeatures.WriteString("   Dépendances et plugins\n\n")
// 	leftFeatures.WriteString(utils.IconStyle.Render("📦 ") + "Templates personnalisés\n")
// 	leftFeatures.WriteString("   Modèles prêts à l'emploi")
//
// 	rightFeatures := strings.Builder{}
// 	rightFeatures.WriteString(utils.IconStyle.Render("🔧 ") + "Outils de développement\n")
// 	rightFeatures.WriteString("   Scripts et utilitaires\n\n")
// 	rightFeatures.WriteString(utils.IconStyle.Render("🚀 ") + "Déploiement simplifié\n")
// 	rightFeatures.WriteString("   Configuration Docker\n\n")
// 	rightFeatures.WriteString(utils.IconStyle.Render("📊 ") + "Monitoring intégré\n")
// 	rightFeatures.WriteString("   Actuator et métriques")
//
// 	leftBox := utils.FeatureBoxStyle.Render(leftFeatures.String())
// 	rightBox := utils.FeatureBoxStyle.Render(rightFeatures.String())
//
// 	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox))
// }

func displayAvailableCommands() {
	fmt.Println(utils.CommandHeaderStyle.Render("📋 Commandes disponibles"))

	commandsBox := strings.Builder{}

	// Commande new
	commandsBox.WriteString(formatCommand("new", "[project-name]", "Créer un nouveau projet Spring Boot"))
	commandsBox.WriteString("\n")

	// Commande generate
	commandsBox.WriteString(formatCommand("generate", "[type]", "Générer des composants"))
	commandsBox.WriteString("\n")

	// Commande version
	commandsBox.WriteString(formatCommand("--version, -v", " ", "Afficher la version de SpringCLI"))
	commandsBox.WriteString("\n")

	// Commande help
	commandsBox.WriteString(formatCommand("--help, -h", " ", "Afficher l'aide détaillée"))

	fmt.Println(utils.HelpBoxStyle.Render(commandsBox.String()))
}

func displayHelp() {
	fmt.Println(utils.CommandHeaderStyle.Render("💡 Exemples d'utilisation"))

	examplesBox := strings.Builder{}
	examplesBox.WriteString(utils.IconStyle.Render("➤ ") + "Créer un projet simple:\n")
	examplesBox.WriteString("  " + utils.CommandStyle.Render("springcli new mon-projet") + "\n\n")

	examplesBox.WriteString(utils.IconStyle.Render("➤ ") + "Projet avec configuration personnalisée:\n")
	examplesBox.WriteString("  " + utils.CommandStyle.Render("springcli new mon-app -g com.monentreprise -j 17") + "\n\n")

	examplesBox.WriteString(utils.IconStyle.Render("➤ ") + "Générer un contrôleur:\n")
	examplesBox.WriteString("  " + utils.CommandStyle.Render("springcli generate controller UserController") + "\n\n")

	examplesBox.WriteString(utils.TipStyle.Render("💡 Astuce: ") +
		"Utilisez " + utils.CommandStyle.Render("springcli [commande] --help") +
		" pour plus d'options sur chaque commande.")

	fmt.Println(utils.HelpBoxStyle.Render(examplesBox.String()))

	fmt.Println()
	fmt.Println(utils.SeparatorStyle.Render("────────────────────────────────────────────────────────────────────────────────"))
	fmt.Println(utils.CommandDescStyle.Render("Pour plus d'informations, visitez: ") +
		utils.CommandStyle.Render("https://github.com/AnaelTech/springcli"))
	fmt.Println()
}

func formatCommand(command, args, description string) string {
	cmdText := utils.CommandStyle.Render(command)
	if args != "" {
		cmdText += " " + utils.CommandDescStyle.Render(args)
	}

	// Padding pour aligner les descriptions
	padding := strings.Repeat(" ", 30-len(command+args))

	return fmt.Sprintf("  %s%s %s %s",
		cmdText,
		padding,
		utils.SeparatorStyle.Render("│"),
		utils.CommandDescStyle.Render(description))
}
