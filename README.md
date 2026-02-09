# springcli

Un CLI moderne pour générer et gérer facilement des projets Spring Boot avec Go, en utilisant [Spring Initializr](https://start.spring.io/) et des templates personnalisés (contrôleurs, services, JWT, etc.).

![build](https://github.com/AnaelTech/springcli/actions/workflows/go.yml/badge.svg)

## Installation

### 1. Cloner le dépôt et compiler
```bash
git clone https://github.com/AnaelTech/springcli.git
cd springcli
go build -o springcli
```

### 2. Installer le binaire dans votre PATH pour n'utiliser que `springcli`
#### Linux/macOS
```bash
mkdir -p ~/.local/bin
cp springcli ~/.local/bin/
# Vérifier que ~/.local/bin est dans votre PATH :
echo $PATH | grep ".local/bin" || echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
# Rechargez le shell ou sourcez ~/.bashrc
```
Ensuite, tapez simplement :
```bash
springcli --help
```

#### Windows 
```powershell
# Déplacer le binaire springcli.exe vers %USERPROFILE%\go\bin ou %USERPROFILE%\.local\bin
mkdir $env:USERPROFILE\go\bin -ea 0
move springcli.exe $env:USERPROFILE\go\bin\springcli.exe
# Ajoutez ce dossier à votre PATH si besoin, puis ouvrez une nouvelle fenêtre de terminal :
# $env:PATH += ";$env:USERPROFILE\go\bin"
# Pour tester :
springcli.exe --help
```

---

## Utilisation rapide

```bash
springcli new mon-projet
springcli generate controller UserController
springcli generate service User
```

(...)

© 2026 AnaelTech | Licence MIT
