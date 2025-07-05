package main

import (
	"fmt"
	"gofetch/utils"
	"math/rand"
)

func main() {
	display()
}

func display() {
	ascii := utils.GetAscii()
	cat := utils.Cats()
	mods := []string{
		utils.GetUser(),
		utils.GetSystem(),
		utils.GetKernel(),
		utils.GetMemory(),
		utils.GetUptime(),
		utils.GetShell(),
		utils.GetDesktop(),
		utils.GetPackages(),
	}

	ri := rand.Intn(len(colors()))
	re := colors()[ri]
	maxLines := max(len(ascii), len(mods), len(cat))

	for i := range maxLines {
		var asciiLine, modLine, catLine string

		if i < len(ascii) {
			asciiLine = ascii[i]
		} else {
			asciiLine = " "
		}

		if i < len(mods) {
			modLine = mods[i]
		} else {
			modLine = ""
		}

		if i < len(cat) {
			catLine = cat[i]
		} else {
			catLine = ""
		}

		fmt.Printf("%-22s  %-35s %s\n", re+asciiLine, modLine, catLine)
	}
}

func colors() []string {
	return []string{
		"\033[0m",
		"\033[31m",
		"\033[32m",
		"\033[33m",
		"\033[34m",
		"\033[35m",
		"\033[36m",
		"\033[97m",
	}
}
