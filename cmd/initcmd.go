package cmd

import (
	"fmt"
	"os"
	/* "path/filepath" */
	"strings"
	"time"
	
	/* "github.com/charmbracelet/lipgloss" */
	"github.com/spf13/cobra"
	"springcli/internal/generator"
	"springcli/internal/utils"
)

// ==================== INIT ===========================
func init() {
	newCmd.Flags().StringP("group-id", "g", "com.example", "Group ID for the project")
	newCmd.Flags().StringP("artifact-id", "a", "demo", "Artifact ID for the project")
	newCmd.Flags().StringP("type", "t", "maven-project", "Type of project to create")
	newCmd.Flags().StringP("spring-boot-version", "s", "3.4.0", "Spring Boot version")
	newCmd.Flags().StringP("java-version", "j", "21", "Java version")
}

// ==================== NEW PROJECT ====================
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Spring Boot project",
	Long:  `Create a new Spring Boot project with sensible defaults`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		// Affichage du titre avec style
		fmt.Println()
		fmt.Println(utils.TitleStyle.Render("🍃 Spring Boot Project Generator"))
		fmt.Println()
		
		// Récupération des flags avec gestion d'erreur stylisée
		groupId, err := cmd.Flags().GetString("group-id")
		if err != nil {
			utils.PrintError("Failed to get group-id flag")
			os.Exit(1)
		}
		
		artifactId, err := cmd.Flags().GetString("artifact-id")
		if err != nil {
			utils.PrintError("Failed to get artifact-id flag")
			os.Exit(1)
		}
		
		typeName, err := cmd.Flags().GetString("type")
		if err != nil {
			utils.PrintError("Failed to get type flag")
			os.Exit(1)
		}
		
		springBootVersion, err := cmd.Flags().GetString("spring-boot-version")
		if err != nil {
			utils.PrintError("Failed to get spring-boot-version flag")
			os.Exit(1)
		}
		
		javaVersion, err := cmd.Flags().GetString("java-version")
		if err != nil {
			utils.PrintError("Failed to get java-version flag")
			os.Exit(1)
		}
		
		createNewProject(projectName, groupId, artifactId, typeName, springBootVersion, javaVersion)
	},
}

func createNewProject(projectName, groupId, artifactId, typeName, bootVersion, javaVersion string) {
	if artifactId == "" {
		artifactId = projectName
	}
	
	// Affichage de la configuration du projet
	displayProjectConfig(projectName, groupId, artifactId, typeName, bootVersion, javaVersion)
	
	defaultDependencies := "web,data-jpa,validation,actuator"
	params := map[string]string{
		"type":        typeName,
		"language":    "java",
		"bootVersion": bootVersion,
		"baseDir":     projectName,
		"groupId":     groupId,
		"artifactId":  artifactId,
		"name":        projectName,
		"packageName": fmt.Sprintf("%s.%s", groupId, projectName),
		"javaVersion": javaVersion,
		"dependencies": defaultDependencies,
	}
	
	// Animation de téléchargement
	utils.PrintStep("Initializing project generation...")
	time.Sleep(500 * time.Millisecond)
	
	utils.PrintStep("Downloading Spring Boot project...")
	
	err := generator.DownloadSpringProject(params, "./"+projectName)
	if err != nil {
		utils.PrintError("Failed to create project")
		return
	}
	
	utils.PrintStep("Configuring project structure...")
	time.Sleep(300 * time.Millisecond)
	
	utils.PrintStep("Setting up dependencies...")
	time.Sleep(200 * time.Millisecond)
	
	// Message de succès final
	fmt.Println()
	utils.PrintSuccess("Project created successfully!")
	
	// Affichage des informations finales
	displayProjectSummary(projectName, groupId, artifactId)
}

// ==================== FONCTIONS D'AFFICHAGE STYLISÉES ====================
//
// func printError(message string, err error) {
// 	errorMsg := errorStyle.Render("❌ " + message)
// 	if err != nil {
// 		errorMsg += ": " + errorStyle.Render(err.Error())
// 	}
// 	fmt.Println(errorMsg)
// }
//
// func printSuccess(message string) {
// 	fmt.Println(successStyle.Render("✅ " + message))
// }
//
// func printWarning(message string) {
// 	fmt.Println(warningStyle.Render("⚠️  " + message))
// }
//
// func printInfo(message string) {
// 	fmt.Println(infoStyle.Render("ℹ️  " + message))
// }
//
// func printStep(message string) {
// 	fmt.Println(stepStyle.Render("➤ " + message))
// }
//
func displayProjectConfig(projectName, groupId, artifactId, typeName, bootVersion, javaVersion string) {
	fmt.Println(utils.SubtitleStyle.Render("📋 Project Configuration"))
	fmt.Println()
	
	configBox := strings.Builder{}
	configBox.WriteString(formatConfigLine("Project Name", projectName))
	configBox.WriteString(formatConfigLine("Group ID", groupId))
	configBox.WriteString(formatConfigLine("Artifact ID", artifactId))
	configBox.WriteString(formatConfigLine("Project Type", typeName))
	configBox.WriteString(formatConfigLine("Spring Boot", bootVersion))
	configBox.WriteString(formatConfigLine("Java Version", javaVersion))
	configBox.WriteString(formatConfigLine("Dependencies", "web, data-jpa, validation, actuator"))
	
	fmt.Println(utils.BoxStyle.Render(configBox.String()))
}

func displayProjectSummary(projectName, groupId, artifactId string) {
	fmt.Println(utils.SubtitleStyle.Render("📁 Project Summary"))
	fmt.Println()
	
	summaryBox := strings.Builder{}
	summaryBox.WriteString(formatConfigLine("Location", "./"+projectName))
	summaryBox.WriteString(formatConfigLine("Package", fmt.Sprintf("%s.%s", groupId, projectName)))
	summaryBox.WriteString("\n")
	summaryBox.WriteString(utils.InfoStyle.Render("Next steps:"))
	summaryBox.WriteString("\n")
	summaryBox.WriteString(utils.LabelStyle.Render("  1. ") + "cd " + utils.ValueStyle.Render(projectName))
	summaryBox.WriteString("\n")
	summaryBox.WriteString(utils.LabelStyle.Render("  2. ") + "./mvnw spring-boot:run")
	summaryBox.WriteString("\n")
	summaryBox.WriteString(utils.LabelStyle.Render("  3. ") + "Open " + utils.ValueStyle.Render("http://localhost:8080"))
	
	fmt.Println(utils.BoxStyle.Render(summaryBox.String()))
	fmt.Println()
}

func formatConfigLine(label, value string) string {
	return fmt.Sprintf("%-16s %s %s\n", 
		utils.LabelStyle.Render(label+":"), 
		utils.SeparatorStyle.Render("│"), 
		utils.ValueStyle.Render(value))
}
