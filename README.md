# springcli

Cli en go pour du Spring et API Rest

## Description

Cli en go pour du Spring et API Rest

## Installation

```bash
# Clone le dépôt
git clone https://github.com/votre-utilisateur/springcli.git
cd springcli

# Compile le binaire
go build -o springcli ./cmd

# Donner les permissions d'exécution
chmod +x springcli

# (Optionnel) Ajoute le binaire à ton PATH
sudo mv springcli /usr/local/bin/
```

## Utilisation

```bash
# Créer un nouveau projet Spring Boot
springcli new monprojet

# Générer une entité
springcli generate entity User name:string age:int

# Générer un service
springcli generate service User

# Générer un contrôleur
springcli generate controller User

# Voir toutes les commandes disponibles
springcli --help
```

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :

1. Fork le projet
2. Créer une branche pour votre fonctionnalité
3. Commit vos changements
4. Push vers la branche
5. Ouvrir une Pull Request

## Licence

Ce projet est sous licence [MIT](LICENSE).

---

> Généré avec le script GitHub Repository Creator
