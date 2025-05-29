package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
)

func main() {
	getUser()
}

func getUser() {
	u, err := user.Current()

	if err != nil {
		log.Fatal(err.Error())
	}

	user := u.Username
	fmt.Println(user)
}

func getSystem() {
	h, _ := host.Info()

	distro := h.Platform

	if strings.Contains(distro, "arch") {
		fmt.Println("      /\\          ")
		fmt.Println("     /  \\         ")
		fmt.Println("    /    \\        ")
		fmt.Println("   /      \\       ")
		fmt.Println("  /   ,,   \\      ")
		fmt.Println(" /   |  |   \\     ")
		fmt.Println("/_-''    ''-_\\    ")
	} else if strings.Contains(distro, "void") {
		fmt.Println("    _______      ")
		fmt.Println(" _ \\______ -     ")
		fmt.Println("| \\  ___  \\ |    ")
		fmt.Println("| | /   \\ | |    ")
		fmt.Println("| | \\___/ | |    ")
		fmt.Println("| \\______ \\_|    ")
		fmt.Println(" -_______\\       ")
	}
}
