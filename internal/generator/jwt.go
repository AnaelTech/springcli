package generator

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"springcli/internal/utils"
)

func GeneratePublicPrivateKey() {
	// Mettre les clés dans un dossier jwt
	if !utils.Exists("jwt") {
		err := utils.CreateFolder("jwt")
		if err != nil {
			utils.PrintError(fmt.Sprintf("Erreur lors de la création du dossier: %v", err))
			os.Exit(1)
		}
	}

	// Generate a new RSA private key with 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA private key:", err)
		os.Exit(1)
	}

	// Encode the private key to the PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyFile, err := os.Create("jwt/private.key")
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de la création du fichier: %v", err))
		os.Exit(1)
	}

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		_ = privateKeyFile.Close()
		utils.PrintError(fmt.Sprintf("Erreur lors de l'encodage de la clé privée: %v", err))
		os.Exit(1)
	}

	if err := privateKeyFile.Close(); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de la fermeture du fichier: %v", err))
		os.Exit(1)
	}

	// Extract the public key from the private key
	publicKey := &privateKey.PublicKey

	// Encode the public key to the PEM format
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	publicKeyFile, err := os.Create("jwt/public.key")
	if err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de la création du fichier: %v", err))
		os.Exit(1)
	}

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		_ = publicKeyFile.Close()
		utils.PrintError(fmt.Sprintf("Erreur lors de l'encodage de la clé publique: %v", err))
		os.Exit(1)
	}

	if err := publicKeyFile.Close(); err != nil {
		utils.PrintError(fmt.Sprintf("Erreur lors de la fermeture du fichier: %v", err))
		os.Exit(1)
	}

	utils.PrintSuccess("Les clés RSA ont été générées avec succès")
}

// func generateJwtPublicKey() {
//
// }
//
// func generateJwtPrivateKey() {
//
// }
