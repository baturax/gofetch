package mods

import (
	"fmt"
	"os"
)

func GetDesktop() string {
	s := os.Getenv("XDG_CURRENT_DESKTOP")
	if s == "" {
		return fmt.Sprintf("  sorry %v", s)
	}
	return fmt.Sprintf("  %v", s)

}
