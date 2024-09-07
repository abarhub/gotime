package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	// os.Args contient tous les arguments passés au programme, y compris le nom du programme lui-même
	argsCmd := os.Args

	var program string
	var args []string
	var exitCode int

	if len(argsCmd) > 1 {
		program = argsCmd[1]
		if len(argsCmd) > 2 {
			args = argsCmd[2:]
		}
	} else {
		log.Fatalf("Aucun argument utilisateur fourni.")
	}

	// Préparer la commande avec cmd /c start
	cmd := exec.Command("cmd", append([]string{"/c", program}, args...)...)

	// Rediriger la sortie standard et la sortie d'erreur vers la console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Capture le temps de début
	startTime := time.Now()

	// Exécuter la commande
	err := cmd.Start() // Démarre la commande sans bloquer
	if err != nil {
		log.Fatalf("Erreur lors de l'exécution de la commande: %v", err)
	}

	// Attendre la fin du programme
	err = cmd.Wait()
	// Capture le temps de fin
	endTime := time.Now()
	if err != nil {

		// Si une erreur survient, obtenir le code de retour
		exitError, ok := err.(*exec.ExitError)
		if ok {
			// Récupérer le code de retour du processus
			exitCode = exitError.ExitCode()
		} else {
			// Autre erreur
			log.Fatalf("Erreur lors de l'attente de la commande : %v", err)
		}

	}

	// Calcul de la durée écoulée
	duration := endTime.Sub(startTime)

	// Affichage de la durée écoulée
	fmt.Printf("Duree = %v, Code sortie = %v\n", duration, exitCode)
}
