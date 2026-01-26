package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypt chiffre les données en utilisant AES-256-GCM.
// Retourne les données brutes sous forme de []byte (Nonce + Ciphertext + Tag).
func Encrypt(secretKeyHex string, plaintext []byte) ([]byte, error) {
	// 1. Décodage de la clé hexadécimale (doit être de 32 octets pour AES-256)
	key, err := hex.DecodeString(secretKeyHex)
	if err != nil {
		return nil, fmt.Errorf("clé hexadécimale invalide : %w", err)
	}

	// 2. Création du bloc de chiffrement AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du cipher : %w", err)
	}

	// 3. Initialisation du mode GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("erreur configuration GCM : %w", err)
	}

	// 4. Génération d'un Nonce (Number Used Once) unique
	// Le nonce est crucial pour la sécurité de GCM
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("erreur génération nonce : %w", err)
	}

	// 5. Chiffrement
	// Seal(dst, nonce, plaintext, additionalData)
	// On passe 'nonce' comme premier argument (dst) pour que le résultat final
	// contienne le nonce concaténé au début du ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}