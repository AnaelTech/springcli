package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// ===================== COULEURS CENTRALISÉES ==============================
var (
	// Couleurs principales
	PrimaryColor   = lipgloss.Color("#00D4AA")
	SecondaryColor = lipgloss.Color("#5A67D8")
	AccentColor    = lipgloss.Color("#F59E0B")
	ErrorColor     = lipgloss.Color("#F56565")
	SuccessColor   = lipgloss.Color("#48BB78")
	WarningColor   = lipgloss.Color("#ED8936")
	InfoColor      = lipgloss.Color("#4299E1")
	MutedColor     = lipgloss.Color("#9CA3AF")
	WhiteColor     = lipgloss.Color("#FFFFFF")
	GrayColor      = lipgloss.Color("#6B7280")
)

// ===================== STYLES POUR LES TEXTES ==============================
var (
	// Style pour les titres principaux
	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor)

	// Style pour les sous-titres
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true).
			Margin(1, 0)

	// Style pour le titre principal de l'app
	MainTitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Margin(0, 0, 1, 0)

	// Style pour les messages d'erreur
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	// Style pour les messages de succès
	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	// Style pour les messages d'avertissement
	WarningStyle = lipgloss.NewStyle().
			Foreground(WarningColor).
			Bold(true)

	// Style pour les messages d'information
	InfoStyle = lipgloss.NewStyle().
			Foreground(InfoColor)

	// Style pour les étapes de progression
	StepStyle = lipgloss.NewStyle().
			Foreground(InfoColor).
			Bold(true)

	// Style pour les prompts
	PromptStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	// Style pour les valeurs/paramètres
	ValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E8F0")).
			Bold(true)

	// Style pour les labels
	LabelStyle = lipgloss.NewStyle().
			Foreground(MutedColor)

	// Style pour les commandes
	CommandStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	// Style pour les descriptions de commandes
	CommandDescStyle = lipgloss.NewStyle().
				Foreground(MutedColor)

	// Style pour les en-têtes de commandes
	CommandHeaderStyle = lipgloss.NewStyle().
				Foreground(SecondaryColor).
				Bold(true).
				Margin(1, 0, 0, 0)

	// Style pour la version
	VersionStyle = lipgloss.NewStyle().
			Foreground(MutedColor).
			Italic(true)

	// Style pour le message de bienvenue
	WelcomeStyle = lipgloss.NewStyle().
			Foreground(WhiteColor).
			Bold(true).
			Margin(1, 0)

	// Style pour la description
	DescriptionStyle = lipgloss.NewStyle().
				Foreground(MutedColor).
				Margin(1, 0).
				Width(80)

	// Style pour les conseils
	TipStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true)

	// Style pour les icônes
	IconStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	// Style pour les séparateurs
	SeparatorStyle = lipgloss.NewStyle().
			Foreground(MutedColor).
			Faint(true)
)

// ===================== STYLES POUR LES BOÎTES ==============================
var (
	// Style pour les boîtes d'information principales
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2).
			Margin(1, 0)

	// Style pour les boîtes d'aide
	HelpBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2).
			Margin(1, 0).
			Width(80)

	// Style pour les boîtes de fonctionnalités
	FeatureBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(SecondaryColor).
			Padding(1, 2).
			Margin(1, 0).
			Width(38)

	// Style pour les boîtes de configuration
	ConfigBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(InfoColor).
			Padding(1, 2).
			Margin(1, 0)

	// Style pour les boîtes d'erreur
	ErrorBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ErrorColor).
			Padding(1, 2).
			Margin(1, 0)

	// Style pour les boîtes de succès
	SuccessBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(SuccessColor).
			Padding(1, 2).
			Margin(1, 0)
)

// ===================== STYLES POUR LES TABLEAUX ==============================
var (
	// Style pour les en-têtes de tableau
	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(WhiteColor).
				Background(PrimaryColor).
				Bold(true).
				Padding(0, 1).
				Align(lipgloss.Center)

	// Style pour les cellules de tableau
	TableCellStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("#404040"))

	// Style pour les lignes de tableau alternées
	TableRowStyle = lipgloss.NewStyle().
			Foreground(MutedColor)

	// Style pour les lignes de tableau actives
	TableRowActiveStyle = lipgloss.NewStyle().
				Foreground(WhiteColor).
				Background(lipgloss.Color("#374151"))
)

// ===================== STYLES POUR LES LISTES ==============================
var (
	// Style pour les éléments de liste
	ListItemStyle = lipgloss.NewStyle().
			Foreground(InfoColor).
			SetString("  • ")

	// Style pour les éléments de liste numérotés
	ListNumberStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	// Style pour les éléments de liste avec icônes
	ListIconStyle = lipgloss.NewStyle().
			Foreground(AccentColor)
)

// ===================== STYLES SPÉCIAUX ==============================
var (
	// Style pour le logo ASCII
	LogoStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Align(lipgloss.Center).
			Margin(1, 0)

	// Style pour les barres de progression
	ProgressBarStyle = lipgloss.NewStyle().
				Foreground(SuccessColor)

	// Style pour les badges
	BadgeStyle = lipgloss.NewStyle().
			Foreground(WhiteColor).
			Background(PrimaryColor).
			Padding(0, 1).
			Bold(true)

	// Style pour les liens
	LinkStyle = lipgloss.NewStyle().
			Foreground(InfoColor).
			Underline(true)

	// Style pour le code inline
	CodeStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Background(lipgloss.Color("#1F2937")).
			Padding(0, 1)
)

// Fonctions d'affichage stylisé
func PrintTitle(message string) {
	fmt.Println(TitleStyle.Render(message))
}

func PrintError(message string) {
	fmt.Println(ErrorStyle.Render(message))
}

func PrintInfo(message string) {
	fmt.Println(InfoStyle.Render(message))
}

func PrintWarning(message string) {
	fmt.Println(WarningStyle.Render(message))
}

func PrintSuccess(message string) {
	fmt.Println(SuccessStyle.Render(message))
}

func PrintSubtitle(message string) {
	fmt.Println(SubtitleStyle.Render(message))
}

func PrintPrompt(message string) {
	fmt.Print(PromptStyle.Render(message))
}

func PrintBox(content string) {
	fmt.Println(BoxStyle.Render(content))
}

func PrintStep(message string) {
	fmt.Println(StepStyle.Render("➤ " + message))
}
