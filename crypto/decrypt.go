
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

// Decrypt assure le déchiffrement AES-256-GCM.
// Il attend une clé hexadécimale et le contenu brut du fichier (plus performant que l'hex).
func Decrypt(secretKeyHex string, encryptedData []byte) (string, error) {
	// 1. Décodage de la clé
	key, err := hex.DecodeString(secretKeyHex)
	if err != nil {
		return "", fmt.Errorf("clé invalide: %w", err)
	}

	// 2. Initialisation du bloc AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("erreur initialisation cipher: %w", err)
	}

	// 3. Initialisation du mode GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("erreur mode GCM: %w", err)
	}

	// 4. Extraction du Nonce
	// Le nonce est stocké au début du message chiffré
	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", errors.Join(errors.New("données corrompues : taille inférieure au nonce"))
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	// 5. Déchiffrement et vérification d'intégrité
	// GCM vérifie automatiquement si les données ont été altérées
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("échec du déchiffrement (clé incorrecte ou fichier altéré): %w", err)
	}

	return string(plaintext), nil
}
