package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var format string
	var printHelp bool

	flag.StringVar(&format, "format", "ubiquiti", "The output format")
	flag.BoolVar(&printHelp, "help", false, "Print help message")

	flag.Parse()

	if printHelp {
		fmt.Print(helpMessage())
		return
	}

	routes := parseInputFile(flag.Arg(0))
	var hexoutput []string

	for _, line := range routes {
		tmp := strings.Fields(line)
		network, destination := tmp[0], tmp[1]
		hexoutput = append(hexoutput, network2hex(network)...)
		hexoutput = append(hexoutput, ip2hex(destination)...)
	}

	switch format {
	case "opnsense":
		fallthrough
	case "opensense":
		fallthrough
	case "ubiquiti":
		fmt.Println(strings.Join(hexoutput, ":"))
	case "dhcp":
		fmt.Println("0x" + strings.ToUpper(strings.Join(hexoutput, "")))
	case "mikrotik":
		fmt.Println("/ip dhcp-server option")
		fmt.Println("add code=121 name=classless-static-route-option value=0x" +
			strings.ToUpper(strings.Join(hexoutput, "")))
	}
}
