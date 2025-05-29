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

func getKernel() {

}

func getUser() {
	u, uerr := user.Current()

	if uerr != nil {
		log.Fatal(uerr.Error())
	}

	h, herr := host.Info()

	if herr != nil {
		log.Fatal(herr.Error())
	}

	user := u.Username
	hostName := h.Hostname
	fmt.Printf("%v@%v\n", user, hostName)
}

func getDistro() {
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
