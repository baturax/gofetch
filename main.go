package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	getUser()
	getSystem()
	getKernel()
	getMem()
}

func getMem() {
	m, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err.Error())
	}

	usedMem := m.Used / 1048576
	totalMem := m.Total / 1048576

	fmt.Printf("  %v / %v\n", usedMem, totalMem)
}

func getKernel() {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	kernel := h.KernelVersion

	fmt.Printf("  %v %v\n", platform, kernel)
}

func getSystem() {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	os := h.OS
	arch := h.KernelArch

	fmt.Printf("  %v %v %v\n", platform, os, arch)
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
