package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

// Clear for windows
func Clear() {

	// Check os and use correct settings
	switch currentOS := runtime.GOOS; currentOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		fallthrough
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	default:

		log.Printf("Currently running on %s. No clear command for this type\n", currentOS)
	}
}
