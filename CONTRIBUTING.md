# Contribution à springcli

Merci de contribuer à ce projet ! Pour garantir une collaboration efficace et un projet de qualité, merci de suivre ces recommandations :

## Processus de contribution

1. **Fork le repository** et clone-le en local.
2. **Crée une branche** dédiée (ex : `feature/ma-fonctionnalite`, `fix/bug-bizarre`).
3. **Code : Respecte le style Go et ajoute des tests unitaires** pour tout comportement non-trivial.
4. **Formate et lint le code** (`go fmt ./...`, `go vet ./...`, `golangci-lint run ./...`).
5. **Push ta branche** et ouvre **une Pull Request claire** sur `main`, décris :
   - Le but du changement
   - Comment tester la feature/bugfix
   - S’il y a des impacts de compatibilité
6. **Un reviewer du projet** validera la PR avant merge.

## Règles de code
- Utilise le modèle idiomatique Go (“table-driven tests”, gestion élégante des erreurs, logs explicites).
- Respecte la structure du CLI (noms de dossier, cobra Commands, separation business/UI).
- Mets à jour le README et le CHANGELOG si tu ajoutes/modifies significativement une fonctionnalité.
- Ajoute des messages clairs d’aide pour chaque commande ou option additionnelle.

## Astuces
- Lance la CI (`go test ./...`) pour vérifier l’intégration continue.
- Si tu ajoutes une dépendance, modifie le go.mod & go.sum.

---

Merci de rendre springcli encore plus utile à tous !