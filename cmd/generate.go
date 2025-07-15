package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"text/template"
	//"path/filepath"
	"bytes"
	"springcli/internal/utils"
	"regexp"
	"encoding/xml"
	)

// ===================== TEMPLATE ============================== 
type TemplateData struct {
	PackageName string
	ClassName   string
	EntityName  string
	TableName   string
	Fields      []Field
}

type Field struct {
	Name string
	Type string
	JsonName string
}

type Relation struct {
	Name string
	Type string
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
	Short: "Génère le code source d’un contrôleur à partir de modèles personnalisés.",
	Long:  `Cette commande permet de générer automatiquement le code d’un contrôleur Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide de contrôleurs
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Error: controller name is required")
			os.Exit(1)
		}
		controllerName := args[0]
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
		"controllerName": controllerName,
		"serviceName": controllerName + "Service",
		"repositoryName": controllerName + "Repository",
		"entityName": controllerName + "Entity",
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}
	tmpl, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := getJavaSourcePath() + "/controller"
	filename := controllerName + "Controller.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			fmt.Println("❌ Erreur lors de la création du dossier:", err)
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		fmt.Println("❌ Le fichier " + filename + " existe déjà")
		return
	}
	generateFile(path, filename, buf.Bytes())
}	
// ==================== END CONTROLLER ================================


	
// ==================== GENERATE SERVICE ==================== 

var generateServiceCmd = &cobra.Command{
	Use:   "service [service-name]",
	Short: "Génère le code source d’un service à partir de modèles personnalisés.",
	Long:  `Cette commande permet de générer automatiquement le code d’un service Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide de services
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Error: service name is required")
			os.Exit(1)
		}
		serviceName := args[0]
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
		"serviceName": serviceName,
		"repositoryName": serviceName + "Repository",
		"entityName": serviceName + "Entity",
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := getJavaSourcePath() + "/service"
	filename := serviceName + "Service.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			fmt.Println("❌ Erreur lors de la création du dossier:", err)
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		fmt.Println("❌ Le fichier " + filename + " existe déjà")
		return
	}
	generateFile(path, filename, buf.Bytes())
}
// ==================== END SERVICE ================================


// ==================== GENERATE REPOSITORY ==================== 

var generateRepositoryCmd = &cobra.Command{
	Use:   "repository [repository-name]",
	Short: "Génère le code source d’une interface de dépôt à partir de modèles personnalisés.",
	Long:  `Cette commande permet de générer automatiquement le code d’une interface de dépôt Spring Boot
en utilisant des templates adaptés. Elle facilite la création rapide d’interfaces de dépôt
structurés et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Error: repository name is required")
			os.Exit(1)
		}
		repositoryName := args[0]
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
		"repositoryName": repositoryName,
		"entityName": repositoryName + "Entity",
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}
	tmpl, err := template.New("repository").Parse(repositoryTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := getJavaSourcePath() + "/repository"
	filename := repositoryName + "Repository.java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			fmt.Println("❌ Erreur lors de la création du dossier:", err)
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		fmt.Println("Le fichier " + filename + " existe déjà")
		return
	}
	generateFile(path, filename, buf.Bytes())
}

// ==================== END REPOSITORY ======================


// ==================== GENERATE ENTITY ====================== 

var generateEntityCmd = &cobra.Command{
	Use:   "entity [entity-name] [fields...]",
	Short: "Génère le code source d’une entité à partir de modèles personnalisés.",
	Long:  `Génère le code source d’une entité Spring Boot en utilisant des templates adaptés.
Elle facilite la création rapide d’entités structurées et conformes aux bonnes pratiques du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Error: entity name is required")
			os.Exit(1)
		} 
		entityName := args[0]
		var fields []Field
		var relations []Relation
		path := getJavaSourcePath() + "/entity"
		filename := entityName + ".java"
		fullPath := path + "/" + filename
		
		
		if utils.Exists(fullPath) {
			fmt.Println("Que voulez ajouter à l'entité "+entityName+" ?")
	    fields, relations = askFieldsAndRelations()
			updateEntity(entityName, fields, relations)
      return
		}

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
		@{{.Type}}
		private {{.Target}} {{.Name}};
		{{end}}
}`

func generateEntity(entityName string, fields []Field, relations []Relation) {
		params := map[string]interface{}{
		"entityName": entityName,
		"tableName": strings.ToLower(entityName),
		"fields": fields,
		"relations": relations,
		"packageName": strings.ReplaceAll(getJavaSourcePath()[len("src/main/java/"):], "/", "."),
	}
	tmpl, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := getJavaSourcePath() + "/entity"
	filename := entityName + ".java"
	fullPath := path + "/" + filename

	// Crée le dossier s'il n'existe pas
	if !utils.Exists(path) {
		err := utils.CreateFolder(path)
		if err != nil {
			fmt.Println("❌ Erreur lors de la création du dossier:", err)
			os.Exit(1)
		}
	}

	// Vérifie si le fichier existe déjà
	if utils.Exists(fullPath) {
		fmt.Println("❌ Le fichier " + filename + " existe déjà")
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
		fmt.Println("Le fichier " + filename + " existe déjà, mise à jour en cours...")
		existingFields = extractFields(string(existingContent))
		existingRelations = extractRelations(string(existingContent))
	} else if !os.IsNotExist(err) {
		fmt.Println("Erreur lors de la lecture du fichier existant:", err)
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
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, params); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(fullPath, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier:", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Fichier %s mis à jour avec succès.\n", filename)
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

func generateFields(entityName string) string {
	fields := make([]Field, 0)
	for _, field := range strings.Split(entityName, ",") {
		fields = append(fields, Field{
			Name: field,
			Type: "String",
			JsonName: field,
		})
	}
	return generateFieldsTemplate(fields)
}

func askFieldsAndRelations() ([]Field , []Relation) {
	var fields []Field
	var relations []Relation
	for {
		var name, typ string
		fmt.Print("Nom de la propriété (laisser vide pour finir): ")
		fmt.Scanln(&name)
		if name == "" {
			break
		} else if name == "relations" {
			rs := askRelations()
			relations = append(relations, rs...)
			continue
		}
		for {
		fmt.Print("Type du champ tapez (?) pour voir la liste des types disponibles: ")
		fmt.Scanln(&typ)
		if typ == "?" {
			fmt.Println("Types disponibles:")
			for _, t := range allTypes() {
				fmt.Println(t)
			}
			continue
		}
			break
		}
		fields = append(fields, Field{
			Name: name,
			Type: javaType(typ),
			JsonName: name,
		})
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
		fmt.Println("Types disponibles:")
		for _, t := range allTypes() {
			fmt.Println(t)
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
				Name: parts[0],
				Type: javaType(parts[1]),
				JsonName: parts[0],
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
	for {
		var name, typ string
		fmt.Print("Nom de la relation (laisser vide pour finir): ")
		fmt.Scanln(&name)
		if name == "" {
			break 
		}
		fmt.Println("Type de relations ")
		fmt.Println(allRelations())
		fmt.Print("Type de la relation: ")
		fmt.Scanln(&typ)
		relations = append(relations, Relation{
			Name: name,
			Type: typ,
			Target: askTarget(),
		})
	}
	return relations
}

func askTarget() string {
	var target string
	fmt.Print("Nom de la classe cible: ")
	fmt.Scanln(&target)
	return target
}

func allRelations() string {
		return `
+-------------+---------------------------------------------------+
|   Nom       | Description                                       |
+-------------+---------------------------------------------------+
| OneToOne    | Relation 1-1 : chaque entité A a une entité B     |
| OneToMany   | 1-N : une entité A a plusieurs B                  |
| ManyToOne   | N-1 : plusieurs entités A pour une entité B       |
| ManyToMany  | N-N : plusieurs A pour plusieurs B                |
+-------------+---------------------------------------------------+
`
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


//=======================FUNCIONS UTILES =====================================================

func generateFieldsTemplate(fields []Field) string {
	var buffer bytes.Buffer
	t := template.Must(template.New("fields").Parse("{{range .}}{{.Name}} {{.Type}};\n{{end}}"))
	err := t.Execute(&buffer, fields)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buffer.String()
}

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
		fmt.Printf("❌ Failed to read %s: %v\n", pomPath, err)
		os.Exit(1)
	}

	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		fmt.Printf("❌ Failed to parse %s: %v\n", pomPath, err)
		os.Exit(1)
	}

	// Priorité au groupId défini dans <project>
	if project.GroupId != "" {
		return strings.ReplaceAll(project.GroupId, ".", "/")
	}

	// Fallback : utiliser <parent><groupId> si strict mode désactivé
	if !strict && project.Parent.GroupId != "" {
		fmt.Printf("⚠️  groupId not found in <project>. Using parent groupId: %s\n", project.Parent.GroupId)
		return strings.ReplaceAll(project.Parent.GroupId, ".", "/")
	}

	// Sinon erreur
	fmt.Println("❌ groupId not found in pom.xml (<project> or <parent>)")
	os.Exit(1)
	return ""
}


//func getPackageName() string {
//	wd, err := os.Getwd()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	return strings.Replace(filepath.Base(wd), "cmd", "", 1)
//}

func generateFile(Path string, filename string, content []byte) {
	err := os.WriteFile(Path+"/"+filename, content, 0644)
	if err != nil {
		fmt.Println("❌ Erreur lors de l'écriture du fichier:", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Fichier %s généré avec succès.\n", filename)
}

//======================== END UTILS =========================================================
