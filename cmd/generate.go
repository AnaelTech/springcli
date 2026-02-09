// Package cmd : contient les commandes principales de springcli
package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"springcli/internal/generator"
	"springcli/internal/utils"

	"github.com/spf13/cobra"
	/* "github.com/charmbracelet/lipgloss" */)

// ===================== TEMPLATE ==============================
type TemplateData struct {
	PackageName string
	ClassName   string
	EntityName  string
	TableName   string
	Fields      []Field
}

type Field struct {
	Name     string
	Type     string
	JSONName string
}

type Relation struct {
	Name   string
	Type   string
	Target string
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	GroupId string   `xml:"groupId"`
	Parent  struct {
		GroupId string `xml:"groupId"`
	} `xml:"parent"`
}

// ===================== INIT ==================================
func init() {
	generateCmd.AddCommand(generateControllerCmd)
	generateCmd.AddCommand(generateServiceCmd)
	generateCmd.AddCommand(generateRepositoryCmd)
	generateCmd.AddCommand(generateEntityCmd)
	generateCmd.AddCommand(generateJwtCmd)
}

// ===================== GENERATE ==============================
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code from templates",
	Long:  `Generate code from templates`,
}

// ==================== GENERATE CONTROLLER ====================
var generateControllerCmd = &cobra.Command{
	Use:   "controller [controller-name]",
	Short: "Génère le code source d'un contrôleur à partir de modèles personnalisés.",
	Long: `Cette commande permet de générer automatiquement le code d'un contrôleur Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide de contrôleurs
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("🎮 GÉNÉRATEUR DE CONTRÔLEUR SPRING BOOT")
		if len(args) == 0 {
			utils.PrintError("Le nom du contrôleur est requis")
			os.Exit(1)
		}
		controllerName := args[0]
		utils.PrintInfo(fmt.Sprintf("Génération du contrôleur: %s", controllerName))
		generateController(controllerName)
	},
}

const controllerTemplate = `package {{.packageName}}.controller;
	
import {{.packageName}}.service.{{.serviceName}};
import {{.packageName}}.repository.{{.repositoryName}};
import {{.packageName}}.entity.{{.entityName}};
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class {{.controllerName}} {
    @Autowired
    private {{.serviceName}} {{.serviceName}}Service;
}`

func generateController(controllerName string) {
	params := map[string]string{
		"controllerName": controllerName + "Controller",
		"serviceName":    controllerName + "Service",
		"repositoryName": controllerName + "Repository",
		"entityName":     controllerName,
		"packageName":    strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}

	tmpl, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors du parsing du template: %v", err))
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'exécution du template: %v", err))
		os.Exit(1)
	}

	path := getJavaSourcePath() + "/controller"
	filename := controllerName + "Controller.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Erreur lors de la création du dossier: %v", err))
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		utils.PrintWarning(fmt.Sprintf("Le fichier %s existe déjà", filename))
		return
	}

	generateFile(path, filename, buf.Bytes())
}

// ==================== GENERATE SERVICE ====================
var generateServiceCmd = &cobra.Command{
	Use:   "service [service-name]",
	Short: "Génère le code source d'un service à partir de modèles personnalisés.",
	Long: `Cette commande permet de générer automatiquement le code d'un service Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide de services
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("🔧 GÉNÉRATEUR DE SERVICE SPRING BOOT")
		if len(args) == 0 {
			utils.PrintError("Le nom du service est requis")
			os.Exit(1)
		}
		serviceName := args[0]
		utils.PrintInfo(fmt.Sprintf("Génération du service: %s", serviceName))
		generateService(serviceName)
	},
}

const serviceTemplate = `package {{.packageName}}.service;

import {{.packageName}}.repository.{{.repositoryName}};
import {{.packageName}}.entity.{{.entityName}};

public interface {{.serviceName}} {

}`

func generateService(serviceName string) {
	params := map[string]string{
		"serviceName":    serviceName + "Service",
		"repositoryName": serviceName + "Repository",
		"entityName":     serviceName,
		"packageName":    strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors du parsing du template: %v", err))
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'exécution du template: %v", err))
		os.Exit(1)
	}

	path := getJavaSourcePath() + "/service"
	filename := serviceName + "Service.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Erreur lors de la création du dossier: %v", err))
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		utils.PrintWarning(fmt.Sprintf("Le fichier %s existe déjà", filename))
		return
	}

	generateFile(path, filename, buf.Bytes())
}

// ==================== GENERATE REPOSITORY ====================
var generateRepositoryCmd = &cobra.Command{
	Use:   "repository [repository-name]",
	Short: "Génère le code source d'une interface de dépôt à partir de modèles personnalisés.",
	Long: `Cette commande permet de générer automatiquement le code d'une interface de dépôt Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide d'interfaces de dépôt
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("🗄️  GÉNÉRATEUR DE REPOSITORY SPRING BOOT")
		if len(args) == 0 {
			utils.PrintError("Le nom du repository est requis")
			os.Exit(1)
		}
		repositoryName := args[0]
		utils.PrintInfo(fmt.Sprintf("Génération du repository: %s", repositoryName))
		generateRepository(repositoryName)
	},
}

const repositoryTemplate = `package {{.packageName}}.repository;

import {{.packageName}}.entity.{{.entityName}};
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface {{.repositoryName}} extends JpaRepository<{{.entityName}}, Long> {

}`

func generateRepository(repositoryName string) {
	params := map[string]string{
		"repositoryName": repositoryName + "Repository",
		"entityName":     repositoryName,
		"packageName":    strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}

	tmpl, err := template.New("repository").Parse(repositoryTemplate)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors du parsing du template: %v", err))
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'exécution du template: %v", err))
		os.Exit(1)
	}

	path := getJavaSourcePath() + "/repository"
	filename := repositoryName + "Repository.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Erreur lors de la création du dossier: %v", err))
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		utils.PrintWarning(fmt.Sprintf("Le fichier %s existe déjà", filename))
		return
	}

	generateFile(path, filename, buf.Bytes())
}

// ==================== GENERATE ENTITY ======================
var generateEntityCmd = &cobra.Command{
	Use:   "entity [entity-name] [fields...]",
	Short: "Génère le code source d'une entité à partir de modèles personnalisés.",
	Long: `Génère le code source d'une entité Spring Boot en utilisant des templates adaptés.
Elle facilite la création rapide d'entités structurées et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("🏗️  GÉNÉRATEUR D'ENTITÉ SPRING BOOT")
		if len(args) == 0 {
			utils.PrintError("Le nom de l'entité est requis")
			os.Exit(1)
		}
		entityName := args[0]
		var fields []Field
		var relations []Relation

		path := getJavaSourcePath() + "/entity"
		filename := entityName + ".java"
		fullPath := path + "/" + filename

		if utils.Exists(fullPath) {
			utils.PrintInfo(fmt.Sprintf("L'entité %s existe déjà", entityName))
			utils.PrintSubtitle("Que voulez-vous ajouter à cette entité ?")
			fields, relations = askFieldsAndRelations()
			updateEntity(entityName, fields, relations)
			return
		}

		utils.PrintInfo(fmt.Sprintf("Création de l'entité: %s", entityName))
		if len(args) == 1 {
			fields, relations = askFieldsAndRelations()
		} else {
			fields = parseFields(args[1:])
			relations = parseRelations(args[1:])
		}

		generateEntity(entityName, fields, relations)
	},
}

const entityTemplate = `package {{.packageName}}.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "{{.tableName}}")
public class {{.entityName}} {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
			
	{{range .fields}}
	private {{.Type}} {{.Name}};
	{{end}}
	
	{{range .relations}}
	{{.Type}}
	private {{.Target}} {{.Name}};
	{{end}}
}`

func generateEntity(entityName string, fields []Field, relations []Relation) {
	params := map[string]interface{}{
		"entityName":  entityName,
		"tableName":   strings.ToLower(entityName),
		"fields":      fields,
		"relations":   relations,
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}

	tmpl, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors du parsing du template: %v", err))
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'exécution du template: %v", err))
		os.Exit(1)
	}

	path := getJavaSourcePath() + "/entity"
	filename := entityName + ".java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Erreur lors de la création du dossier: %v", err))
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		utils.PrintWarning(fmt.Sprintf("Le fichier %s existe déjà", filename))
		return
	}

	generateFile(path, filename, buf.Bytes())
}

func updateEntity(entityName string, fields []Field, relations []Relation) {
	path := getJavaSourcePath() + "/entity"
	filename := entityName + ".java"
	fullPath := path + "/" + filename

	existingContent, err := os.ReadFile(fullPath)
	var existingFields []Field
	var existingRelations []Relation

	if err == nil {
		utils.PrintInfo(fmt.Sprintf("Mise à jour du fichier %s...", filename))
		existingFields = extractFields(string(existingContent))
		existingRelations = extractRelations(string(existingContent))
	} else if !os.IsNotExist(err) {
		utils.PrintError(fmt.Sprintf("Erreur lors de la lecture du fichier existant: %v", err))
		os.Exit(1)
	}

	mergedFields := mergeFields(existingFields, fields)
	mergedRelations := mergeRelations(existingRelations, relations)

	params := map[string]interface{}{
		"entityName":  entityName,
		"tableName":   strings.ToLower(entityName),
		"fields":      mergedFields,
		"relations":   mergedRelations,
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}

	tmpl, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors du parsing du template: %v", err))
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'exécution du template: %v", err))
		os.Exit(1)
	}

	err = os.WriteFile(fullPath, buf.Bytes(), 0o644)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'écriture du fichier: %v", err))
		os.Exit(1)
	}

	utils.PrintSuccess(fmt.Sprintf("Fichier %s mis à jour avec succès", filename))
}

func extractFields(javaContent string) []Field {
	fieldRegexp := regexp.MustCompile(`(?m)^\s*private\s+(\w+)\s+(\w+);`)
	matches := fieldRegexp.FindAllStringSubmatch(javaContent, -1)

	var fields []Field
	for _, m := range matches {
		fields = append(fields, Field{
			Type: m[1],
			Name: m[2],
		})
	}

	return fields
}

func mergeFields(existing, added []Field) []Field {
	fieldMap := make(map[string]Field)
	for _, f := range existing {
		if strings.ToLower(f.Name) == "id" {
			continue
		}
		fieldMap[f.Name] = f
	}
	for _, f := range added {
		if strings.ToLower(f.Name) == "id" {
			continue
		}
		fieldMap[f.Name] = f
	}

	merged := make([]Field, 0, len(fieldMap))
	for _, f := range fieldMap {
		merged = append(merged, f)
	}

	return merged
}

func askFieldsAndRelations() ([]Field, []Relation) {
	var fields []Field
	var relations []Relation

	utils.PrintSubtitle("Configuration des champs de l'entité")
	utils.PrintBox("Tapez 'relations' comme nom de propriété pour configurer les relations JPA")

	for {
		var name, typ string
		utils.PrintPrompt("Nom de la propriété (laisser vide pour finir): ")
		_, _ = fmt.Scanln(&name)
		if name == "" {
			break
		} else if name == "relations" {
			rs := askRelations()
			relations = append(relations, rs...)
			continue
		}

		for {
			utils.PrintPrompt("Type du champ (tapez '?' pour voir la liste): ")
			_, _ = fmt.Scanln(&typ)
			if typ == "?" {
				utils.PrintSubtitle("Types disponibles:")
				for _, t := range allTypes() {
					fmt.Println(utils.ListItemStyle.Render(t))
				}
				continue
			}
			break
		}

		fields = append(fields, Field{
			Name:     name,
			Type:     javaType(typ),
			JSONName: name,
		})
		utils.PrintSuccess(fmt.Sprintf("Champ ajouté: %s (%s)", name, javaType(typ)))
	}

	return fields, relations
}

func allTypes() []string {
	return []string{"string", "int", "bool", "double", "Long", "LocalDate", "LocalDateTime"}
}

func javaType(t string) string {
	switch t {
	case "string":
		return "String"
	case "int":
		return "int"
	case "bool":
		return "boolean"
	case "double":
		return "double"
	case "Long":
		return "Long"
	case "LocalDate":
		return "LocalDate"
	case "LocalDateTime":
		return "LocalDateTime"
	case "?":
		utils.PrintSubtitle("Types disponibles:")
		for _, t := range allTypes() {
			fmt.Println(utils.ListItemStyle.Render(t))
		}
		return ""
	default:
		return t // fallback, could be another entity for relation
	}
}

func parseFields(fieldArgs []string) []Field {
	fields := make([]Field, 0)
	for _, arg := range fieldArgs {
		parts := strings.SplitN(arg, ":", 2)
		if len(parts) == 2 {
			fields = append(fields, Field{
				Name:     parts[0],
				Type:     javaType(parts[1]),
				JSONName: parts[0],
			})
		}
	}
	return fields
}

func parseRelations(fieldArgs []string) []Relation {
	relations := make([]Relation, 0)
	for _, arg := range fieldArgs {
		parts := strings.SplitN(arg, ":", 3)
		if len(parts) == 3 {
			relations = append(relations, Relation{
				Name:   parts[0],
				Type:   relationsType(parts[1]),
				Target: parts[2],
			})
		}
	}
	return relations
}

func askRelations() []Relation {
	var relations []Relation

	utils.PrintSubtitle("Configuration des relations JPA")

	for {
		var name, typ string
		utils.PrintPrompt("Nom de la relation (laisser vide pour finir): ")
		_, _ = fmt.Scanln(&name)
		if name == "" {
			break
		}

		utils.PrintSubtitle("Types de relations disponibles:")
		fmt.Println(formatRelationsTable())
		utils.PrintPrompt("Type de la relation: ")
		_, _ = fmt.Scanln(&typ)

		relations = append(relations, Relation{
			Name:   name,
			Type:   relationsType(typ),
			Target: askTarget(),
		})
		utils.PrintSuccess(fmt.Sprintf("Relation ajoutée: %s (%s)", name, typ))
	}

	return relations
}

func askTarget() string {
	var target string
	utils.PrintPrompt("Nom de la classe cible: ")
	_, _ = fmt.Scanln(&target)
	return target
}

func formatRelationsTable() string {
	headers := []string{"Nom", "Description"}
	rows := [][]string{
		{"OneToOne", "Relation 1-1 : chaque entité A a une entité B"},
		{"OneToMany", "1-N : une entité A a plusieurs B"},
		{"ManyToOne", "N-1 : plusieurs entités A pour une entité B"},
		{"ManyToMany", "N-N : plusieurs A pour plusieurs B"},
	}

	var table strings.Builder
	// En-têtes
	headerRow := ""
	for _, header := range headers {
		headerRow += utils.TableHeaderStyle.Width(25).Render(header)
	}
	table.WriteString(headerRow + "\n")

	// Lignes
	for _, row := range rows {
		rowStr := ""
		for _, cell := range row {
			rowStr += utils.TableCellStyle.Width(25).Render(cell)
		}
		table.WriteString(rowStr + "\n")
	}

	return utils.BoxStyle.Render(table.String())
}

func relationsType(typ string) string {
	switch typ {
	case "OneToOne":
		return "@OneToOne"
	case "OneToMany":
		return "@OneToMany"
	case "ManyToOne":
		return "@ManyToOne"
	case "ManyToMany":
		return "@ManyToMany"
	default:
		return typ // fallback, could be another entity for relation
	}
}

func extractRelations(javaContent string) []Relation {
	// regex pour trouver les relations de type @ManyToOne private User user;
	relationRegexp := regexp.MustCompile(`@(\w+)\s+private\s+(\w+)\s+(\w+);`)
	matches := relationRegexp.FindAllStringSubmatch(javaContent, -1)

	var relations []Relation
	for _, m := range matches {
		relations = append(relations, Relation{
			Type:   "@" + m[1],
			Target: m[2],
			Name:   m[3],
		})
	}

	return relations
}

func mergeRelations(existing, added []Relation) []Relation {
	relationMap := make(map[string]Relation)
	for _, r := range existing {
		key := r.Name + "|" + r.Type + "|" + r.Target
		relationMap[key] = r
	}
	for _, r := range added {
		key := r.Name + "|" + r.Type + "|" + r.Target
		relationMap[key] = r
	}

	merged := make([]Relation, 0, len(relationMap))
	for _, r := range relationMap {
		merged = append(merged, r)
	}

	return merged
}

//====================== END ENTITY =========================================================
// ====================== START JWT =========================================================

var generateJwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Génère la clé publique et privée RSA pour JWT",
	Long:  `Génère la clé publique et privée RSA pour JWT`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("🔧 GÉNÉRATEUR JWT SPRING BOOT")
		if len(args) != 0 {
			utils.PrintError("Cette commande ne prend pas d'arguments")
			os.Exit(1)
		}

		// Check if the keys already exist and folder jwt exists
		if utils.Exists("jwt/public.key") && utils.Exists("jwt/private.key") {
			utils.PrintWarning("Les clés RSA existent déjà. Voulez-vous les écraser ?")
			if !AskYesNo() {
				os.Exit(1)
			}
		}

		generator.GeneratePublicPrivateKey()
	},
}

// ===================== END JWT ===============================================================

// =======================FUNCIONS UTILES =====================================================

func getJavaSourcePath() string {
	base := "src/main/java/" + getPackageName()
	entries, err := os.ReadDir(base)
	if err != nil {
		// fallback to base groupId path
		return base
	}

	// Check if there's exactly one subfolder (typical in Spring apps)
	for _, entry := range entries {
		if entry.IsDir() {
			return base + "/" + entry.Name()
		}
	}

	return base
}

func getPackageName() string {
	pomPath := "./pom.xml"
	strict := true // Si strict est vrai, on ne prend pas le groupId du parent

	data, err := os.ReadFile(pomPath)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Impossible de lire %s: %v", pomPath, err))
		os.Exit(1)
	}

	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		utils.PrintError(fmt.Sprintf("Impossible de parser %s: %v", pomPath, err))
		os.Exit(1)
	}

	// Priorité au groupId défini dans <project>
	if project.GroupId != "" {
		return strings.ReplaceAll(project.GroupId, ".", "/")
	}

	// Fallback : utiliser <parent><groupId> si strict mode désactivé
	if !strict && project.Parent.GroupId != "" {
		utils.PrintWarning(fmt.Sprintf("groupId non trouvé dans <project>. Utilisation du groupId parent: %s", project.Parent.GroupId))
		return strings.ReplaceAll(project.Parent.GroupId, ".", "/")
	}

	// Sinon erreur
	utils.PrintError("groupId non trouvé dans pom.xml (<project> ou <parent>)")
	os.Exit(1)
	return ""
}

func generateFile(path string, filename string, content []byte) {
	err := os.WriteFile(path+"/"+filename, content, 0o644)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de l'écriture du fichier: %v", err))
		os.Exit(1)
	}

	utils.PrintSuccess(fmt.Sprintf("Fichier %s généré avec succès", filename))

	// Afficher des informations supplémentaires
	utils.PrintInfo(fmt.Sprintf("Emplacement: %s/%s", path, filename))
	utils.PrintInfo(fmt.Sprintf("Taille: %d octets", len(content)))
}

func AskYesNo() bool {
	var answer string
	utils.PrintPrompt("Voulez-vous continuer ? (y/n): ")
	_, _ = fmt.Scanln(&answer)
	return answer == "y"
}
