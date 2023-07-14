-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class net
---@field FlagUp any
---@field FlagBroadcast any
---@field FlagLoopback any
---@field FlagPointToPoint any
---@field FlagMulticast any
---@field FlagRunning any
---@field IPv4len any
---@field IPv6len any
---@field IPv4bcast any
---@field IPv4allsys any
---@field IPv4allrouter any
---@field IPv4zero any
---@field IPv6zero any
---@field IPv6unspecified any
---@field IPv6loopback any
---@field IPv6interfacelocalallnodes any
---@field IPv6linklocalallnodes any
---@field IPv6linklocalallrouters any
---@field DefaultResolver any
---@field ErrWriteToConnected any
---@field ErrClosed any
net = {}

--- LookupHost looks up the given host using the local resolver.
--- It returns a slice of that host's addresses.
---
--- LookupHost uses context.Background internally; to specify the context, use
--- Resolver.LookupHost.
---@param host string
---@return string[], err
function net.LookupHost(host) end

--- ResolveTCPAddr returns an address of TCP end point.
---
--- The network must be a TCP network name.
---
--- If the host in the address parameter is not a literal IP address or
--- the port is not a literal port number, ResolveTCPAddr resolves the
--- address to an address of TCP end point.
--- Otherwise, it parses the address as a pair of literal IP address
--- and port number.
--- The address parameter can use a host name, but this is not
--- recommended, because it will return at most one of the host name's
--- IP addresses.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param network string
---@param address string
---@return netTCPAddr, err
function net.ResolveTCPAddr(network, address) end

--- ListenTCP acts like Listen for TCP networks.
---
--- The network must be a TCP network name; see func Dial for details.
---
--- If the IP field of laddr is nil or an unspecified IP address,
--- ListenTCP listens on all available unicast and anycast IP addresses
--- of the local system.
--- If the Port field of laddr is 0, a port number is automatically
--- chosen.
---@param network string
---@param laddr netTCPAddr
---@return netTCPListener, err
function net.ListenTCP(network, laddr) end

--- UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false,
--- then the returned UDPAddr will contain a nil IP field, indicating an
--- address family-agnostic unspecified address.
---@param addr netipAddrPort
---@return netUDPAddr
function net.UDPAddrFromAddrPort(addr) end

--- DialUDP acts like Dial for UDP networks.
---
--- The network must be a UDP network name; see func Dial for details.
---
--- If laddr is nil, a local address is automatically chosen.
--- If the IP field of raddr is nil or an unspecified IP address, the
--- local system is assumed.
---@param network string
---@param laddr netUDPAddr
---@param raddr netUDPAddr
---@return netUDPConn, err
function net.DialUDP(network, laddr, raddr) end

--- ListenUDP acts like ListenPacket for UDP networks.
---
--- The network must be a UDP network name; see func Dial for details.
---
--- If the IP field of laddr is nil or an unspecified IP address,
--- ListenUDP listens on all available IP addresses of the local system
--- except multicast IP addresses.
--- If the Port field of laddr is 0, a port number is automatically
--- chosen.
---@param network string
---@param laddr netUDPAddr
---@return netUDPConn, err
function net.ListenUDP(network, laddr) end

--- Interfaces returns a list of the system's network interfaces.
---@return any, err
function net.Interfaces() end

--- IPv4Mask returns the IP mask (in 4-byte form) of the
--- IPv4 mask a.b.c.d.
---@param a byte
---@param b byte
---@param c byte
---@param d byte
---@return netIPMask
function net.IPv4Mask(a, b, c, d) end

--- DialUnix acts like Dial for Unix networks.
---
--- The network must be a Unix network name; see func Dial for details.
---
--- If laddr is non-nil, it is used as the local address for the
--- connection.
---@param network string
---@param laddr netUnixAddr
---@param raddr netUnixAddr
---@return netUnixConn, err
function net.DialUnix(network, laddr, raddr) end

--- FilePacketConn returns a copy of the packet network connection
--- corresponding to the open file f.
--- It is the caller's responsibility to close f when finished.
--- Closing c does not affect f, and closing f does not affect c.
---@param f osFile
---@return netPacketConn, err
function net.FilePacketConn(f) end

--- CIDRMask returns an IPMask consisting of 'ones' 1 bits
--- followed by 0s up to a total length of 'bits' bits.
--- For a mask of this form, CIDRMask is the inverse of IPMask.Size.
---@param ones number
---@param bits number
---@return netIPMask
function net.CIDRMask(ones, bits) end

--- JoinHostPort combines host and port into a network address of the
--- form "host:port". If host contains a colon, as found in literal
--- IPv6 addresses, then JoinHostPort returns "[host]:port".
---
--- See func Dial for a description of the host and port parameters.
---@param host string
---@param port string
---@return string
function net.JoinHostPort(host, port) end

--- LookupSRV tries to resolve an SRV query of the given service,
--- protocol, and domain name. The proto is "tcp" or "udp".
--- The returned records are sorted by priority and randomized
--- by weight within a priority.
---
--- LookupSRV constructs the DNS name to look up following RFC 2782.
--- That is, it looks up _service._proto.name. To accommodate services
--- publishing SRV records under non-standard names, if both service
--- and proto are empty strings, LookupSRV looks up name directly.
---
--- The returned service names are validated to be properly
--- formatted presentation-format domain names. If the response contains
--- invalid names, those records are filtered out and an error
--- will be returned alongside the remaining results, if any.
---@param service string
---@param proto string
---@param name string
---@return string, any, err
function net.LookupSRV(service, proto, name) end

--- ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
--- IP over InfiniBand link-layer address using one of the following formats:
---
---	00:00:5e:00:53:01
---	02:00:5e:10:00:00:00:01
---	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
---	00-00-5e-00-53-01
---	02-00-5e-10-00-00-00-01
---	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
---	0000.5e00.5301
---	0200.5e10.0000.0001
---	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
---@param s string
---@return netHardwareAddr, err
function net.ParseMAC(s) end

--- DialTCP acts like Dial for TCP networks.
---
--- The network must be a TCP network name; see func Dial for details.
---
--- If laddr is nil, a local address is automatically chosen.
--- If the IP field of raddr is nil or an unspecified IP address, the
--- local system is assumed.
---@param network string
---@param laddr netTCPAddr
---@param raddr netTCPAddr
---@return netTCPConn, err
function net.DialTCP(network, laddr, raddr) end

--- Listen announces on the local network address.
---
--- The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
---
--- For TCP networks, if the host in the address parameter is empty or
--- a literal unspecified IP address, Listen listens on all available
--- unicast and anycast IP addresses of the local system.
--- To only use IPv4, use network "tcp4".
--- The address can use a host name, but this is not recommended,
--- because it will create a listener for at most one of the host's IP
--- addresses.
--- If the port in the address parameter is empty or "0", as in
--- "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
--- The Addr method of Listener can be used to discover the chosen
--- port.
---
--- See func Dial for a description of the network and address
--- parameters.
---
--- Listen uses context.Background internally; to specify the context, use
--- ListenConfig.Listen.
---@param network string
---@param address string
---@return netListener, err
function net.Listen(network, address) end

--- FileListener returns a copy of the network listener corresponding
--- to the open file f.
--- It is the caller's responsibility to close ln when finished.
--- Closing ln does not affect f, and closing f does not affect ln.
---@param f osFile
---@return netListener, err
function net.FileListener(f) end

--- LookupAddr performs a reverse lookup for the given address, returning a list
--- of names mapping to that address.
---
--- The returned names are validated to be properly formatted presentation-format
--- domain names. If the response contains invalid names, those records are filtered
--- out and an error will be returned alongside the remaining results, if any.
---
--- When using the host C library resolver, at most one result will be
--- returned. To bypass the host resolver, use a custom Resolver.
---
--- LookupAddr uses context.Background internally; to specify the context, use
--- Resolver.LookupAddr.
---@param addr string
---@return string[], err
function net.LookupAddr(addr) end

--- Dial connects to the address on the named network.
---
--- Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
--- "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
--- (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
--- "unixpacket".
---
--- For TCP and UDP networks, the address has the form "host:port".
--- The host must be a literal IP address, or a host name that can be
--- resolved to IP addresses.
--- The port must be a literal port number or a service name.
--- If the host is a literal IPv6 address it must be enclosed in square
--- brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
--- The zone specifies the scope of the literal IPv6 address as defined
--- in RFC 4007.
--- The functions JoinHostPort and SplitHostPort manipulate a pair of
--- host and port in this form.
--- When using TCP, and the host resolves to multiple IP addresses,
--- Dial will try each IP address in order until one succeeds.
---
--- Examples:
---
---	Dial("tcp", "golang.org:http")
---	Dial("tcp", "192.0.2.1:http")
---	Dial("tcp", "198.51.100.1:80")
---	Dial("udp", "[2001:db8::1]:domain")
---	Dial("udp", "[fe80::1%lo0]:53")
---	Dial("tcp", ":80")
---
--- For IP networks, the network must be "ip", "ip4" or "ip6" followed
--- by a colon and a literal protocol number or a protocol name, and
--- the address has the form "host". The host must be a literal IP
--- address or a literal IPv6 address with zone.
--- It depends on each operating system how the operating system
--- behaves with a non-well known protocol number such as "0" or "255".
---
--- Examples:
---
---	Dial("ip4:1", "192.0.2.1")
---	Dial("ip6:ipv6-icmp", "2001:db8::1")
---	Dial("ip6:58", "fe80::1%lo0")
---
--- For TCP, UDP and IP networks, if the host is empty or a literal
--- unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
--- TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
--- assumed.
---
--- For Unix networks, the address must be a file system path.
---@param network string
---@param address string
---@return netConn, err
function net.Dial(network, address) end

--- ListenPacket announces on the local network address.
---
--- The network must be "udp", "udp4", "udp6", "unixgram", or an IP
--- transport. The IP transports are "ip", "ip4", or "ip6" followed by
--- a colon and a literal protocol number or a protocol name, as in
--- "ip:1" or "ip:icmp".
---
--- For UDP and IP networks, if the host in the address parameter is
--- empty or a literal unspecified IP address, ListenPacket listens on
--- all available IP addresses of the local system except multicast IP
--- addresses.
--- To only use IPv4, use network "udp4" or "ip4:proto".
--- The address can use a host name, but this is not recommended,
--- because it will create a listener for at most one of the host's IP
--- addresses.
--- If the port in the address parameter is empty or "0", as in
--- "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
--- The LocalAddr method of PacketConn can be used to discover the
--- chosen port.
---
--- See func Dial for a description of the network and address
--- parameters.
---
--- ListenPacket uses context.Background internally; to specify the context, use
--- ListenConfig.ListenPacket.
---@param network string
---@param address string
---@return netPacketConn, err
function net.ListenPacket(network, address) end

--- InterfaceByIndex returns the interface specified by index.
---
--- On Solaris, it returns one of the logical network interfaces
--- sharing the logical data link; for more precision use
--- InterfaceByName.
---@param index number
---@return netInterface, err
function net.InterfaceByIndex(index) end

--- ResolveUDPAddr returns an address of UDP end point.
---
--- The network must be a UDP network name.
---
--- If the host in the address parameter is not a literal IP address or
--- the port is not a literal port number, ResolveUDPAddr resolves the
--- address to an address of UDP end point.
--- Otherwise, it parses the address as a pair of literal IP address
--- and port number.
--- The address parameter can use a host name, but this is not
--- recommended, because it will return at most one of the host name's
--- IP addresses.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param network string
---@param address string
---@return netUDPAddr, err
function net.ResolveUDPAddr(network, address) end

--- DialIP acts like Dial for IP networks.
---
--- The network must be an IP network name; see func Dial for details.
---
--- If laddr is nil, a local address is automatically chosen.
--- If the IP field of raddr is nil or an unspecified IP address, the
--- local system is assumed.
---@param network string
---@param laddr netIPAddr
---@param raddr netIPAddr
---@return netIPConn, err
function net.DialIP(network, laddr, raddr) end

--- FileConn returns a copy of the network connection corresponding to
--- the open file f.
--- It is the caller's responsibility to close f when finished.
--- Closing c does not affect f, and closing f does not affect c.
---@param f osFile
---@return netConn, err
function net.FileConn(f) end

--- IPv4 returns the IP address (in 16-byte form) of the
--- IPv4 address a.b.c.d.
---@param a byte
---@param b byte
---@param c byte
---@param d byte
---@return netIP
function net.IPv4(a, b, c, d) end

--- ListenIP acts like ListenPacket for IP networks.
---
--- The network must be an IP network name; see func Dial for details.
---
--- If the IP field of laddr is nil or an unspecified IP address,
--- ListenIP listens on all available IP addresses of the local system
--- except multicast IP addresses.
---@param network string
---@param laddr netIPAddr
---@return netIPConn, err
function net.ListenIP(network, laddr) end

--- SplitHostPort splits a network address of the form "host:port",
--- "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
--- host%zone and port.
---
--- A literal IPv6 address in hostport must be enclosed in square
--- brackets, as in "[::1]:80", "[::1%lo0]:80".
---
--- See func Dial for a description of the hostport parameter, and host
--- and port results.
---@param hostport string
---@return string, err
function net.SplitHostPort(hostport) end

--- LookupNS returns the DNS NS records for the given domain name.
---
--- The returned name server names are validated to be properly
--- formatted presentation-format domain names. If the response contains
--- invalid names, those records are filtered out and an error
--- will be returned alongside the remaining results, if any.
---
--- LookupNS uses context.Background internally; to specify the context, use
--- Resolver.LookupNS.
---@param name string
---@return any, err
function net.LookupNS(name) end

--- DialTimeout acts like Dial but takes a timeout.
---
--- The timeout includes name resolution, if required.
--- When using TCP, and the host in the address parameter resolves to
--- multiple IP addresses, the timeout is spread over each consecutive
--- dial, such that each is given an appropriate fraction of the time
--- to connect.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param network string
---@param address string
---@param timeout timeDuration
---@return netConn, err
function net.DialTimeout(network, address, timeout) end

--- InterfaceAddrs returns a list of the system's unicast interface
--- addresses.
---
--- The returned list does not identify the associated interface; use
--- Interfaces and Interface.Addrs for more detail.
---@return any, err
function net.InterfaceAddrs() end

--- LookupPort looks up the port for the given network and service.
---
--- LookupPort uses context.Background internally; to specify the context, use
--- Resolver.LookupPort.
---@param network string
---@param service string
---@return number, err
function net.LookupPort(network, service) end

--- LookupCNAME returns the canonical name for the given host.
--- Callers that do not care about the canonical name can call
--- LookupHost or LookupIP directly; both take care of resolving
--- the canonical name as part of the lookup.
---
--- A canonical name is the final name after following zero
--- or more CNAME records.
--- LookupCNAME does not return an error if host does not
--- contain DNS "CNAME" records, as long as host resolves to
--- address records.
---
--- The returned canonical name is validated to be a properly
--- formatted presentation-format domain name.
---
--- LookupCNAME uses context.Background internally; to specify the context, use
--- Resolver.LookupCNAME.
---@param host string
---@return string, err
function net.LookupCNAME(host) end

--- LookupTXT returns the DNS TXT records for the given domain name.
---
--- LookupTXT uses context.Background internally; to specify the context, use
--- Resolver.LookupTXT.
---@param name string
---@return string[], err
function net.LookupTXT(name) end

--- LookupIP looks up host using the local resolver.
--- It returns a slice of that host's IPv4 and IPv6 addresses.
---@param host string
---@return any, err
function net.LookupIP(host) end

--- Pipe creates a synchronous, in-memory, full duplex
--- network connection; both ends implement the Conn interface.
--- Reads on one end are matched with writes on the other,
--- copying data directly between the two; there is no internal
--- buffering.
---@return netConn, netConn
function net.Pipe() end

--- TCPAddrFromAddrPort returns addr as a TCPAddr. If addr.IsValid() is false,
--- then the returned TCPAddr will contain a nil IP field, indicating an
--- address family-agnostic unspecified address.
---@param addr netipAddrPort
---@return netTCPAddr
function net.TCPAddrFromAddrPort(addr) end

--- ParseIP parses s as an IP address, returning the result.
--- The string s can be in IPv4 dotted decimal ("192.0.2.1"), IPv6
--- ("2001:db8::68"), or IPv4-mapped IPv6 ("::ffff:192.0.2.1") form.
--- If s is not a valid textual representation of an IP address,
--- ParseIP returns nil.
---@param s string
---@return netIP
function net.ParseIP(s) end

--- ResolveIPAddr returns an address of IP end point.
---
--- The network must be an IP network name.
---
--- If the host in the address parameter is not a literal IP address,
--- ResolveIPAddr resolves the address to an address of IP end point.
--- Otherwise, it parses the address as a literal IP address.
--- The address parameter can use a host name, but this is not
--- recommended, because it will return at most one of the host name's
--- IP addresses.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param network string
---@param address string
---@return netIPAddr, err
function net.ResolveIPAddr(network, address) end

--- ListenMulticastUDP acts like ListenPacket for UDP networks but
--- takes a group address on a specific network interface.
---
--- The network must be a UDP network name; see func Dial for details.
---
--- ListenMulticastUDP listens on all available IP addresses of the
--- local system including the group, multicast IP address.
--- If ifi is nil, ListenMulticastUDP uses the system-assigned
--- multicast interface, although this is not recommended because the
--- assignment depends on platforms and sometimes it might require
--- routing configuration.
--- If the Port field of gaddr is 0, a port number is automatically
--- chosen.
---
--- ListenMulticastUDP is just for convenience of simple, small
--- applications. There are golang.org/x/net/ipv4 and
--- golang.org/x/net/ipv6 packages for general purpose uses.
---
--- Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option
--- to 0 under IPPROTO_IP, to disable loopback of multicast packets.
---@param network string
---@param ifi netInterface
---@param gaddr netUDPAddr
---@return netUDPConn, err
function net.ListenMulticastUDP(network, ifi, gaddr) end

--- ListenUnixgram acts like ListenPacket for Unix networks.
---
--- The network must be "unixgram".
---@param network string
---@param laddr netUnixAddr
---@return netUnixConn, err
function net.ListenUnixgram(network, laddr) end

--- LookupMX returns the DNS MX records for the given domain name sorted by preference.
---
--- The returned mail server names are validated to be properly
--- formatted presentation-format domain names. If the response contains
--- invalid names, those records are filtered out and an error
--- will be returned alongside the remaining results, if any.
---
--- LookupMX uses context.Background internally; to specify the context, use
--- Resolver.LookupMX.
---@param name string
---@return any, err
function net.LookupMX(name) end

--- ResolveUnixAddr returns an address of Unix domain socket end point.
---
--- The network must be a Unix network name.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param network string
---@param address string
---@return netUnixAddr, err
function net.ResolveUnixAddr(network, address) end

--- ListenUnix acts like Listen for Unix networks.
---
--- The network must be "unix" or "unixpacket".
---@param network string
---@param laddr netUnixAddr
---@return netUnixListener, err
function net.ListenUnix(network, laddr) end

--- InterfaceByName returns the interface specified by name.
---@param name string
---@return netInterface, err
function net.InterfaceByName(name) end

--- ParseCIDR parses s as a CIDR notation IP address and prefix length,
--- like "192.0.2.0/24" or "2001:db8::/32", as defined in
--- RFC 4632 and RFC 4291.
---
--- It returns the IP address and the network implied by the IP and
--- prefix length.
--- For example, ParseCIDR("192.0.2.1/24") returns the IP address
--- 192.0.2.1 and the network 192.0.2.0/24.
---@param s string
---@return netIP, netIPNet, err
function net.ParseCIDR(s) end


---@class any
local any = {}

--- OpError is the error type usually returned by functions in the net
--- package. It describes the operation, network type, and address of
--- an error.
---@class netOpError
---@field Op string
---@field Net string
---@field Source netAddr
---@field Addr netAddr
---@field Err err
local netOpError = {}


---@return string
function netOpError:Error() end


---@return boolean
function netOpError:Timeout() end


---@return boolean
function netOpError:Temporary() end


---@return err
function netOpError:Unwrap() end

--- Buffers contains zero or more runs of bytes to write.
---
--- On certain machines, for certain types of connections, this is
--- optimized into an OS-specific batch write operation (such as
--- "writev").
---@class netBuffers
local netBuffers = {}

--- WriteTo writes contents of the buffers to w.
---
--- WriteTo implements io.WriterTo for Buffers.
---
--- WriteTo modifies the slice v as well as v[i] for 0 <= i < len(v),
--- but does not modify v[i][j] for any i, j.
---@param w ioWriter
---@return number, err
function netBuffers:WriteTo(w) end

--- Read from the buffers.
---
--- Read implements io.Reader for Buffers.
---
--- Read modifies the slice v as well as v[i] for 0 <= i < len(v),
--- but does not modify v[i][j] for any i, j.
---@param p byte[]
---@return number, err
function netBuffers:Read(p) end

--- An IPNet represents an IP network.
---@class netIPNet
---@field IP netIP
---@field Mask netIPMask
local netIPNet = {}

--- Contains reports whether the network includes ip.
---@param ip netIP
---@return boolean
function netIPNet:Contains(ip) end

--- Network returns the address's network name, "ip+net".
---@return string
function netIPNet:Network() end

--- String returns the CIDR notation of n like "192.0.2.0/24"
--- or "2001:db8::/48" as defined in RFC 4632 and RFC 4291.
--- If the mask is not in the canonical form, it returns the
--- string which consists of an IP address, followed by a slash
--- character and a mask expressed as hexadecimal form with no
--- punctuation like "198.51.100.0/c000ff00".
---@return string
function netIPNet:String() end

--- A Listener is a generic network listener for stream-oriented protocols.
---
--- Multiple goroutines may invoke methods on a Listener simultaneously.
---@class netListener
local netListener = {}

--- PacketConn is a generic packet-oriented network connection.
---
--- Multiple goroutines may invoke methods on a PacketConn simultaneously.
---@class netPacketConn
local netPacketConn = {}

--- A ParseError is the error type of literal network address parsers.
---@class netParseError
---@field Type string
---@field Text string
local netParseError = {}


---@return string
function netParseError:Error() end


---@return boolean
function netParseError:Timeout() end


---@return boolean
function netParseError:Temporary() end


---@class netInvalidAddrError
local netInvalidAddrError = {}


---@return string
function netInvalidAddrError:Error() end


---@return boolean
function netInvalidAddrError:Timeout() end


---@return boolean
function netInvalidAddrError:Temporary() end

--- A Dialer contains options for connecting to an address.
---
--- The zero value for each field is equivalent to dialing
--- without that option. Dialing with the zero value of Dialer
--- is therefore equivalent to just calling the Dial function.
---
--- It is safe to call Dialer's methods concurrently.
---@class netDialer
---@field Timeout any
---@field Deadline any
---@field LocalAddr netAddr
---@field DualStack boolean
---@field FallbackDelay any
---@field KeepAlive any
---@field Resolver netResolver
---@field Cancel any
---@field Control any
---@field ControlContext any
local netDialer = {}

--- Dial connects to the address on the named network.
---
--- See func Dial for a description of the network and address
--- parameters.
---
--- Dial uses context.Background internally; to specify the context, use
--- DialContext.
---@param network string
---@param address string
---@return netConn, err
function netDialer:Dial(network, address) end

--- DialContext connects to the address on the named network using
--- the provided context.
---
--- The provided Context must be non-nil. If the context expires before
--- the connection is complete, an error is returned. Once successfully
--- connected, any expiration of the context will not affect the
--- connection.
---
--- When using TCP, and the host in the address parameter resolves to multiple
--- network addresses, any dial timeout (from d.Timeout or ctx) is spread
--- over each consecutive dial, such that each is given an appropriate
--- fraction of the time to connect.
--- For example, if a host has 4 IP addresses and the timeout is 1 minute,
--- the connect to each single address will be given 15 seconds to complete
--- before trying the next one.
---
--- See func Dial for a description of the network and address
--- parameters.
---@param ctx contextContext
---@param network string
---@param address string
---@return netConn, err
function netDialer:DialContext(ctx, network, address) end


---@class netFlags
local netFlags = {}


---@return string
function netFlags:String() end

--- DNSConfigError represents an error reading the machine's DNS configuration.
--- (No longer used; kept for compatibility.)
---@class netDNSConfigError
---@field Err err
local netDNSConfigError = {}


---@return err
function netDNSConfigError:Unwrap() end


---@return string
function netDNSConfigError:Error() end


---@return boolean
function netDNSConfigError:Timeout() end


---@return boolean
function netDNSConfigError:Temporary() end


---@class netUnknownNetworkError
local netUnknownNetworkError = {}


---@return string
function netUnknownNetworkError:Error() end


---@return boolean
function netUnknownNetworkError:Timeout() end


---@return boolean
function netUnknownNetworkError:Temporary() end


---@class any
local any = {}


---@class any
local any = {}

--- A HardwareAddr represents a physical hardware address.
---@class netHardwareAddr
local netHardwareAddr = {}


---@return string
function netHardwareAddr:String() end

--- Conn is a generic stream-oriented network connection.
---
--- Multiple goroutines may invoke methods on a Conn simultaneously.
---@class netConn
local netConn = {}

--- Addr represents a network end point address.
---
--- The two methods Network and String conventionally return strings
--- that can be passed as the arguments to Dial, but the exact form
--- and meaning of the strings is up to the implementation.
---@class netAddr
local netAddr = {}


---@class any
local any = {}

--- SetUnlinkOnClose sets whether the underlying socket file should be removed
--- from the file system when the listener is closed.
---
--- The default behavior is to unlink the socket file only when package net created it.
--- That is, when the listener and the underlying socket file were created by a call to
--- Listen or ListenUnix, then by default closing the listener will remove the socket file.
--- but if the listener was created by a call to FileListener to use an already existing
--- socket file, then by default closing the listener will not remove the socket file.
---@param unlink boolean
function any:SetUnlinkOnClose(unlink) end

--- An IP is a single IP address, a slice of bytes.
--- Functions in this package accept either 4-byte (IPv4)
--- or 16-byte (IPv6) slices as input.
---
--- Note that in this documentation, referring to an
--- IP address as an IPv4 address or an IPv6 address
--- is a semantic property of the address, not just the
--- length of the byte slice: a 16-byte slice can still
--- be an IPv4 address.
---@class netIP
local netIP = {}

--- IsUnspecified reports whether ip is an unspecified address, either
--- the IPv4 address "0.0.0.0" or the IPv6 address "::".
---@return boolean
function netIP:IsUnspecified() end

--- IsMulticast reports whether ip is a multicast address.
---@return boolean
function netIP:IsMulticast() end

--- IsLinkLocalMulticast reports whether ip is a link-local
--- multicast address.
---@return boolean
function netIP:IsLinkLocalMulticast() end

--- String returns the string form of the IP address ip.
--- It returns one of 4 forms:
---   - "<nil>", if ip has length 0
---   - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
---   - IPv6 conforming to RFC 5952 ("2001:db8::1"), if ip is a valid IPv6 address
---   - the hexadecimal form of ip, without punctuation, if no other cases apply
---@return string
function netIP:String() end

--- IsLoopback reports whether ip is a loopback address.
---@return boolean
function netIP:IsLoopback() end

--- IsGlobalUnicast reports whether ip is a global unicast
--- address.
---
--- The identification of global unicast addresses uses address type
--- identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
--- the exception of IPv4 directed broadcast addresses.
--- It returns true even if ip is in IPv4 private address space or
--- local IPv6 unicast address space.
---@return boolean
function netIP:IsGlobalUnicast() end

--- DefaultMask returns the default IP mask for the IP address ip.
--- Only IPv4 addresses have default masks; DefaultMask returns
--- nil if ip is not a valid IPv4 address.
---@return netIPMask
function netIP:DefaultMask() end

--- Equal reports whether ip and x are the same IP address.
--- An IPv4 address and that same address in IPv6 form are
--- considered to be equal.
---@param x netIP
---@return boolean
function netIP:Equal(x) end

--- To4 converts the IPv4 address ip to a 4-byte representation.
--- If ip is not an IPv4 address, To4 returns nil.
---@return netIP
function netIP:To4() end

--- To16 converts the IP address ip to a 16-byte representation.
--- If ip is not an IP address (it is the wrong length), To16 returns nil.
---@return netIP
function netIP:To16() end

--- MarshalText implements the encoding.TextMarshaler interface.
--- The encoding is the same as returned by String, with one exception:
--- When len(ip) is zero, it returns an empty slice.
---@return byte[], err
function netIP:MarshalText() end

--- Mask returns the result of masking the IP address ip with mask.
---@param mask netIPMask
---@return netIP
function netIP:Mask(mask) end

--- UnmarshalText implements the encoding.TextUnmarshaler interface.
--- The IP address is expected in a form accepted by ParseIP.
---@param text byte[]
---@return err
function netIP:UnmarshalText(text) end

--- IsPrivate reports whether ip is a private address, according to
--- RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
---@return boolean
function netIP:IsPrivate() end

--- IsInterfaceLocalMulticast reports whether ip is
--- an interface-local multicast address.
---@return boolean
function netIP:IsInterfaceLocalMulticast() end

--- IsLinkLocalUnicast reports whether ip is a link-local
--- unicast address.
---@return boolean
function netIP:IsLinkLocalUnicast() end

--- An IPMask is a bitmask that can be used to manipulate
--- IP addresses for IP addressing and routing.
---
--- See type IPNet and func ParseCIDR for details.
---@class netIPMask
local netIPMask = {}

--- Size returns the number of leading ones and total bits in the mask.
--- If the mask is not in the canonical form--ones followed by zeros--then
--- Size returns 0, 0.
---@return number
function netIPMask:Size() end

--- String returns the hexadecimal form of m, with no punctuation.
---@return string
function netIPMask:String() end


---@class any
local any = {}


---@class any
local any = {}

--- DNSError represents a DNS lookup error.
---@class netDNSError
---@field Err string
---@field Name string
---@field Server string
---@field IsTimeout boolean
---@field IsTemporary boolean
---@field IsNotFound boolean
local netDNSError = {}


---@return string
function netDNSError:Error() end

--- Timeout reports whether the DNS lookup is known to have timed out.
--- This is not always known; a DNS lookup may fail due to a timeout
--- and return a DNSError for which Timeout returns false.
---@return boolean
function netDNSError:Timeout() end

--- Temporary reports whether the DNS error is known to be temporary.
--- This is not always known; a DNS lookup may fail due to a temporary
--- error and return a DNSError for which Temporary returns false.
---@return boolean
function netDNSError:Temporary() end

--- An MX represents a single DNS MX record.
---@class netMX
---@field Host string
---@field Pref any
local netMX = {}

--- An NS represents a single DNS NS record.
---@class netNS
---@field Host string
local netNS = {}

--- Interface represents a mapping between network interface name
--- and index. It also represents network interface facility
--- information.
---@class netInterface
---@field Index number
---@field MTU number
---@field Name string
---@field HardwareAddr netHardwareAddr
---@field Flags netFlags
local netInterface = {}

--- Addrs returns a list of unicast interface addresses for a specific
--- interface.
---@return any, err
function netInterface:Addrs() end

--- MulticastAddrs returns a list of multicast, joined group addresses
--- for a specific interface.
---@return any, err
function netInterface:MulticastAddrs() end


---@class netAddrError
---@field Err string
---@field Addr string
local netAddrError = {}


---@return string
function netAddrError:Error() end


---@return boolean
function netAddrError:Timeout() end


---@return boolean
function netAddrError:Temporary() end


---@class any
local any = {}


---@class any
local any = {}

--- ListenConfig contains options for listening to an address.
---@class netListenConfig
---@field Control any
---@field KeepAlive any
local netListenConfig = {}

--- ListenPacket announces on the local network address.
---
--- See func ListenPacket for a description of the network and address
--- parameters.
---@param ctx contextContext
---@param network string
---@param address string
---@return netPacketConn, err
function netListenConfig:ListenPacket(ctx, network, address) end

--- Listen announces on the local network address.
---
--- See func Listen for a description of the network and address
--- parameters.
---@param ctx contextContext
---@param network string
---@param address string
---@return netListener, err
function netListenConfig:Listen(ctx, network, address) end

--- An SRV represents a single DNS SRV record.
---@class netSRV
---@field Target string
---@field Port any
---@field Priority any
---@field Weight any
local netSRV = {}

--- An Error represents a network error.
---@class execError
local execError = {}


---@class any
local any = {}


---@class any
local any = {}


---@class any
local any = {}
