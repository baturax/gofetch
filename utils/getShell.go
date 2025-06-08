package utils

import (
	"fmt"
	"os"
)

func GetShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return fmt.Sprintf("󰲒 sorry %v", shell)
	}

	return fmt.Sprintf("󰲒 %v", shell)

}
