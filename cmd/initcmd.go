package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"springcli/internal/generator"
)

// ==================== INIT ===========================

func init() {

	newCmd.Flags().StringP("group-id", "g", "com.example", "Group ID for the project")
	newCmd.Flags().StringP("artifact-id", "a", "demo", "Artifact ID for the project")
	newCmd.Flags().StringP("type", "t", "maven-project", "Type of project to create")
	newCmd.Flags().StringP("spring-boot-version", "s", "3.1.0", "Spring Boot version")
	newCmd.Flags().StringP("java-version", "j", "17", "Java version")
	
	rootCmd.AddCommand(newCmd)
}


// ==================== NEW PROJECT ====================
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Spring Boot project",
	Long:  `Create a new Spring Boot project with sensible defaults`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		groupId, err := cmd.Flags().GetString("group-id")
		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}
		artifactId, err := cmd.Flags().GetString("artifact-id")
		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}
		typeName, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}
		springBootVersion, err := cmd.Flags().GetString("spring-boot-version")
		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}
		javaVersion, err := cmd.Flags().GetString("java-version")
		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}
		createNewProject(projectName, groupId, artifactId, typeName, springBootVersion, javaVersion)
	},
}

func createNewProject(projectName, groupId, artifactId, typeName, bootVersion, javaVersion string) {
	if artifactId == "" {
		artifactId = projectName
	}

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
		"dependencies": "web", // tu peux rendre ça dynamique plus tard
	}

	fmt.Println("📦 Downloading Spring Boot project...")
	err := generator.DownloadSpringProject(params, "./"+projectName)
	if err != nil {
		fmt.Println("❌ Failed to create project:", err)
		return
	}
	fmt.Println("✅ Project created successfully at", filepath.Join(".", projectName))
}


