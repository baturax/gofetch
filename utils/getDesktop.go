package utils

import (
	"fmt"
	"os"
)

func GetDesktop() string {

	desktop := os.Getenv("XDG_CURRENT_DESKTOP")

	if desktop == "" {
		return fmt.Sprintf("  sorry %v", desktop)
	} else {
		return fmt.Sprintf("  %v", desktop)
	}
}
