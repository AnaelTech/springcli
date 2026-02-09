# Changelog

Toutes les modifications majeures du projet sont documentées ici, dans un format simple inspiré de Keep a Changelog.

## [Unreleased]

### Ajouté
- Intégration GitHub Actions pour build, test, format, lint (go.yml)
- Tests d’intégration CLI: génération projet Spring Boot (main_integration_test.go)
- CONTRIBUTING.md & README.md enrichis
- Tests unitaires pour tous les modules principaux

### Corrigé
- Correction fmt.Sprintf dans jwt.go (erreurs runtime build)

---

## [0.1.0] – 2024-06
- Première version stable publique du CLI springcli
