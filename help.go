package main

func helpMessage() string {
	var s string

	s = `
staticrouter: A utility to help configure classless static routes.

Usage:

$ staticrouter /path/to/routes.txt

The program takes a -format option that can be one of:

* ubiquiti: The hex notation used by Ubiquiti / Unifi equipment
* opnsense: The same as ubiquiti, as OPNsense routers use the same format
* dhcp: The literal hexidecimal representation defined in the RFC
* mikrotik: The explicit commands needed for a Mikrotik router

The routes.txt file has a specific but easy format.

The first line is the IP address of the default network gateway.

Each subsequent line has the CIDR of a classless network, followed by
whitespace, followed by the IP address of that CIDR's gateway.

A three line example file might look like:

192.168.3.1
10.2.3.0/24 192.168.3.2
10.4.5.0/24 192.168.3.3

For more examples, please reference:

    https://github.com/chrislea/staticrouter/tree/main/examples/

`
	return s
}
