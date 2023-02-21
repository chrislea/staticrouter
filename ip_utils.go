package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// trivial function that takes a string that represents an integer and returns
// the hex value. in practice, the string will be an octet from an ipv4 IP
// address.
func hexify(s string) string {
	i, _ := strconv.Atoi(s)

	if i < 16 {
		return fmt.Sprintf("0%x", i)
	}

	return fmt.Sprintf("%x", i)
}

// takes a network in CIDR notation and returns a slice containing the hex
// components needed in the output.
//
// we let Go parse the CIDR which has the convenient effect of letting us
// be a little loose with the network definition. Go will extract the
// actual network for us if it can. we also get the netmask as `mask`,
// which will be slice with four values, one for each octet of the netmask.
// that is used towards the end of the function to let us know if we need
// to keep parsing the network bits or not, since once the mask octet is
// zero (going "left to right"), we know we've defined everything we need to.
func network2hex(network string) []string {
	var ret []string

	_, ipv4Net, err := net.ParseCIDR(network)
	if err != nil {
		log.Fatal(err)
	}

	// the use of "mask" and "netmask" here isn't great in terms of naming,
	// but I'm not sure what else would be better.
	//
	// Here "mask" is a slice that contains the four octets of a netmask,
	// i.e., for a /24, it would represent the octets of 255.255.255.0
	//
	// the "netmask", for the same /24 network, would just be "24". we'll
	// need both.
	mask := ipv4Net.Mask
	tmp := strings.Split(ipv4Net.String(), "/")
	ip, netmask := tmp[0], tmp[1]

	octets := strings.Split(ip, ".")

	// the syntax we eventually want is basically
	//
	// <netmask><network bits of CIDR IP><destination IP>
	//
	// but every octet is in hex. so we start with the hex value of the
	// netmask, then we add *just* the hex values for the network portion
	// of the IP. adding the hex value for the destination IP happens in
	// a different function
	ret = append(ret, hexify(netmask))

	// we only do this when the mask isn't zero to make sure we only get
	// the network portion of the IP
	for i := range mask {
		if mask[i] > 0 {
			ret = append(ret, hexify(octets[i]))
		}
	}

	return ret
}

// this function simply breaks an ipv4 IP into octets and returns
// a slice with the hex representation of each octet.
func ip2hex(ip string) []string {
	var ret []string

	// sanity check ip address
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		log.Fatal("IP address " + ip + " doesn't look valid... exiting")
	}

	octets := strings.Split(ip, ".")

	for _, octet := range octets {
		ret = append(ret, hexify(octet))
	}

	return ret
}
