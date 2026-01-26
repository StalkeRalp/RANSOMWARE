package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// KeyGenerator génère une clé cryptographique aléatoire sécurisée.
// Il retourne la clé sous forme de chaîne hexadécimale et une erreur éventuelle.
func KeyGenerator(longByte int) (string, error) {
	// 1. Allocation de la tranche d'octets
	key := make([]byte, longByte)

	// 2. Lecture d'octets aléatoires sécurisés (CSPRNG)
	// Utilise crypto/rand pour une sécurité de niveau cryptographique.
	if _, err := rand.Read(key); err != nil {
		return "", fmt.Errorf("échec de génération de l'entropie : %w", err)
	}

	// 3. Conversion en hexadécimal
	// L'hexadécimal facilite le stockage et la saisie manuelle de la clé.
	strKey := hex.EncodeToString(key)

	return strKey, nil
}