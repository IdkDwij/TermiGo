package main
// Made by IDKDwij
import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("TermiGo")
	askcommand()

}

func askcommand() {
	var command string
	fmt.Printf("Command: ")
	fmt.Scan(&command)
	commandfunc(command)
}

func commandfunc(getcommand string) {
	if getcommand == "hostname" {
		println(os.Hostname())
	} else if getcommand == "ip" {
		getip()
	} else if getcommand == "" {
		print("\n")
		print("Command: ")
	} else if getcommand != "" {
		cmderror(getcommand)
	}
	askcommand()
}
func getip() {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(ipAddress)
}
func cmderror(commandname string) {
	println("'" + commandname + "'" + "is not recognised as a command")
}
