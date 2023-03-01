## staticrouter

The `staticrouter` program is a small utility to help people looking to
distribute static classless routes to network clients via DHCP using their
routers.

This is generally done via DHCP option 121 as defined in
[RFC3442](https://www.rfc-editor.org/rfc/rfc3442), and a number of advanced
home routers let end users set this. Unfortunately, the interface to set
this up is generally kind of clunky, and requires using a fairly obtuse,
hexidecimal based notation that is easy to get wrong. With `staticrouter`,
hopefully the potential for error can be significantly reduced.

The utility was inspired from
[this Javascript app](https://www.medo64.com/2018/01/configuring-classless-static-route-option/)
on Medo's homepage. It doesn't share any code, but the basic functionality
and output are essentially identical.

### Usage

The program reads route information from a text file on your computer. The
format of the text file should match:

```text
<default gateway ip>
<network 1> <gateway for network 1>
<network 2> <gateway for network 2>
<network 3> <gateway for network 3>
...
```

As an example, say that your network's default gateway is `192.168.3.1`.
Additionally, let's assume you have two VPNs into this network. The first
VPN uses network `10.2.3.0/24` with a gateway of `192.168.3.2`, and the
second VPN uses network `10.4.5.0/24` with a gateway of `192.168.3.3`. If
this was your setup, you would make the file `routes.txt` (which we'll be
using as an example here) as follows:

```text
192.168.3.1
10.2.3.0/24 192.168.3.2
10.4.5.0/24 192.168.3.3
```

You would then run the following from your shell:
```shell
$ staticrouter /path/to/routes.txt
```

The resulting output

```shell
00:c0:a8:03:01:18:0a:02:03:c0:a8:03:02:18:0a:04:05:c0:a8:03:03
```

is the hexidecimal format needed for Ubiquiti/Unifi and OPNsense routers for
those classless routes.

The utility also accepts a command line argument `-format` which can have the
following values:

* `ubiquiti`: Outputs format for Ubiquiti / Unifi routers (this is the default)
* `opnsense`: Identical to ubiquiti as the format is the same for these routers
* `dhcp`: Outputs the format defined in the RFC
* `mikrotik`: Outputs the specific commands needed for Mikrotik routers

For example:

```shell
$ staticrouter -format=dhcp /path/to/routes.txt
0x00C0A80301180A0203C0A80302180A0405C0A80303
```

```shell
$ staticrouter -format=mikrotik /path/to/routes.txt
/ip dhcp-server option
add code=121 name=classless-static-route-option value=0x00C0A80301180A0203C0A80302180A0405C0A80303
```