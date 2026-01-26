package readwrite

import (
	"fmt"
	"os"
)

// ReadFile lit le contenu d'un fichier et retourne les octets bruts.
// Utiliser []byte au lieu de string est beaucoup plus performant pour la cryptographie.
func ReadFile(path string) ([]byte, error) {
	// 1. Utilisation de os.ReadFile (standard moderne depuis Go 1.16+)
	data, err := os.ReadFile(path)

	if err != nil {
		// Retourne une erreur contextuelle pour faciliter le debug professionnel
		return nil, fmt.Errorf("erreur de lecture du fichier [%s]: %w", path, err)
	}

	// 2. Retourne directement les données brutes
	return data, nil
}