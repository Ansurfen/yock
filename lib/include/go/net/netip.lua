-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

netip = {}

--- IPv6Unspecified returns the IPv6 unspecified address "::".
---@return netAddr
function netip.IPv6Unspecified() end

--- AddrPortFrom returns an AddrPort with the provided IP and port.
--- It does not allocate.
---@param ip netAddr
---@param port any
---@return netipAddrPort
function netip.AddrPortFrom(ip, port) end

--- ParsePrefix parses s as an IP address prefix.
--- The string can be in the form "192.168.1.0/24" or "2001:db8::/32",
--- the CIDR notation defined in RFC 4632 and RFC 4291.
--- IPv6 zones are not permitted in prefixes, and an error will be returned if a
--- zone is present.
---
--- Note that masked address bits are not zeroed. Use Masked for that.
---@param s string
---@return netipPrefix, err
function netip.ParsePrefix(s) end

--- IPv6LinkLocalAllNodes returns the IPv6 link-local all nodes multicast
--- address ff02::1.
---@return netAddr
function netip.IPv6LinkLocalAllNodes() end

--- MustParseAddr calls ParseAddr(s) and panics on error.
--- It is intended for use in tests with hard-coded strings.
---@param s string
---@return netAddr
function netip.MustParseAddr(s) end

--- PrefixFrom returns a Prefix with the provided IP address and bit
--- prefix length.
---
--- It does not allocate. Unlike Addr.Prefix, PrefixFrom does not mask
--- off the host bits of ip.
---
--- If bits is less than zero or greater than ip.BitLen, Prefix.Bits
--- will return an invalid value -1.
---@param ip netAddr
---@param bits number
---@return netipPrefix
function netip.PrefixFrom(ip, bits) end

--- AddrFrom16 returns the IPv6 address given by the bytes in addr.
--- An IPv4-mapped IPv6 address is left as an IPv6 address.
--- (Use Unmap to convert them if needed.)
---@param addr byte[]
---@return netAddr
function netip.AddrFrom16(addr) end

--- AddrFromSlice parses the 4- or 16-byte byte slice as an IPv4 or IPv6 address.
--- Note that a net.IP can be passed directly as the []byte argument.
--- If slice's length is not 4 or 16, AddrFromSlice returns Addr{}, false.
---@param slice byte[]
---@return netAddr, boolean
function netip.AddrFromSlice(slice) end

--- MustParsePrefix calls ParsePrefix(s) and panics on error.
--- It is intended for use in tests with hard-coded strings.
---@param s string
---@return netipPrefix
function netip.MustParsePrefix(s) end

--- MustParseAddrPort calls ParseAddrPort(s) and panics on error.
--- It is intended for use in tests with hard-coded strings.
---@param s string
---@return netipAddrPort
function netip.MustParseAddrPort(s) end

--- ParseAddrPort parses s as an AddrPort.
---
--- It doesn't do any name resolution: both the address and the port
--- must be numeric.
---@param s string
---@return netipAddrPort, err
function netip.ParseAddrPort(s) end

--- IPv6Loopback returns the IPv6 loopback address ::1.
---@return netAddr
function netip.IPv6Loopback() end

--- IPv4Unspecified returns the IPv4 unspecified address "0.0.0.0".
---@return netAddr
function netip.IPv4Unspecified() end

--- AddrFrom4 returns the address of the IPv4 address given by the bytes in addr.
---@param addr byte[]
---@return netAddr
function netip.AddrFrom4(addr) end

--- ParseAddr parses s as an IP address, returning the result. The string
--- s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
--- or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
---@param s string
---@return netAddr, err
function netip.ParseAddr(s) end

--- IPv6LinkLocalAllRouters returns the IPv6 link-local all routers multicast
--- address ff02::2.
---@return netAddr
function netip.IPv6LinkLocalAllRouters() end

--- Prefix is an IP address prefix (CIDR) representing an IP network.
---
--- The first Bits() of Addr() are specified. The remaining bits match any address.
--- The range of Bits() is [0,32] for IPv4 or [0,128] for IPv6.
---@class netipPrefix
local netipPrefix = {}

--- Addr returns p's IP address.
---@return netAddr
function netipPrefix:Addr() end

--- IsSingleIP reports whether p contains exactly one IP.
---@return boolean
function netipPrefix:IsSingleIP() end

--- Contains reports whether the network p includes ip.
---
--- An IPv4 address will not match an IPv6 prefix.
--- An IPv4-mapped IPv6 address will not match an IPv4 prefix.
--- A zero-value IP will not match any prefix.
--- If ip has an IPv6 zone, Contains returns false,
--- because Prefixes strip zones.
---@param ip netAddr
---@return boolean
function netipPrefix:Contains(ip) end

--- Overlaps reports whether p and o contain any IP addresses in common.
---
--- If p and o are of different address families or either have a zero
--- IP, it reports false. Like the Contains method, a prefix with an
--- IPv4-mapped IPv6 address is still treated as an IPv6 mask.
---@param o netipPrefix
---@return boolean
function netipPrefix:Overlaps(o) end

--- AppendTo appends a text encoding of p,
--- as generated by MarshalText,
--- to b and returns the extended buffer.
---@param b byte[]
---@return byte[]
function netipPrefix:AppendTo(b) end

--- MarshalText implements the encoding.TextMarshaler interface,
--- The encoding is the same as returned by String, with one exception:
--- If p is the zero value, the encoding is the empty string.
---@return byte[], err
function netipPrefix:MarshalText() end

--- IsValid reports whether p.Bits() has a valid range for p.Addr().
--- If p.Addr() is the zero Addr, IsValid returns false.
--- Note that if p is the zero Prefix, then p.IsValid() == false.
---@return boolean
function netipPrefix:IsValid() end

--- UnmarshalText implements the encoding.TextUnmarshaler interface.
--- The IP address is expected in a form accepted by ParsePrefix
--- or generated by MarshalText.
---@param text byte[]
---@return err
function netipPrefix:UnmarshalText(text) end

--- MarshalBinary implements the encoding.BinaryMarshaler interface.
--- It returns Addr.MarshalBinary with an additional byte appended
--- containing the prefix bits.
---@return byte[], err
function netipPrefix:MarshalBinary() end

--- Bits returns p's prefix length.
---
--- It reports -1 if invalid.
---@return number
function netipPrefix:Bits() end

--- Masked returns p in its canonical form, with all but the high
--- p.Bits() bits of p.Addr() masked off.
---
--- If p is zero or otherwise invalid, Masked returns the zero Prefix.
---@return netipPrefix
function netipPrefix:Masked() end

--- UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
--- It expects data in the form generated by MarshalBinary.
---@param b byte[]
---@return err
function netipPrefix:UnmarshalBinary(b) end

--- String returns the CIDR notation of p: "<ip>/<bits>".
---@return string
function netipPrefix:String() end

--- Addr represents an IPv4 or IPv6 address (with or without a scoped
--- addressing zone), similar to net.IP or net.IPAddr.
---
--- Unlike net.IP or net.IPAddr, Addr is a comparable value
--- type (it supports == and can be a map key) and is immutable.
---
--- The zero Addr is not a valid IP address.
--- Addr{} is distinct from both 0.0.0.0 and ::.
---@class netAddr
local netAddr = {}

--- Prefix keeps only the top b bits of IP, producing a Prefix
--- of the specified length.
--- If ip is a zero Addr, Prefix always returns a zero Prefix and a nil error.
--- Otherwise, if bits is less than zero or greater than ip.BitLen(),
--- Prefix returns an error.
---@param b number
---@return netipPrefix, err
function netAddr:Prefix(b) end

--- As16 returns the IP address in its 16-byte representation.
--- IPv4 addresses are returned as IPv4-mapped IPv6 addresses.
--- IPv6 addresses with zones are returned without their zone (use the
--- Zone method to get it).
--- The ip zero value returns all zeroes.
---@return byte[]
function netAddr:As16() end

--- As4 returns an IPv4 or IPv4-in-IPv6 address in its 4-byte representation.
--- If ip is the zero Addr or an IPv6 address, As4 panics.
--- Note that 0.0.0.0 is not the zero Addr.
---@return byte[]
function netAddr:As4() end

--- Compare returns an integer comparing two IPs.
--- The result will be 0 if ip == ip2, -1 if ip < ip2, and +1 if ip > ip2.
--- The definition of "less than" is the same as the Less method.
---@param ip2 netAddr
---@return number
function netAddr:Compare(ip2) end

--- IsMulticast reports whether ip is a multicast address.
---@return boolean
function netAddr:IsMulticast() end

--- Is6 reports whether ip is an IPv6 address, including IPv4-mapped
--- IPv6 addresses.
---@return boolean
function netAddr:Is6() end

--- Prev returns the IP before ip.
--- If there is none, it returns the IP zero value.
---@return netAddr
function netAddr:Prev() end

--- AppendTo appends a text encoding of ip,
--- as generated by MarshalText,
--- to b and returns the extended buffer.
---@param b byte[]
---@return byte[]
function netAddr:AppendTo(b) end

--- MarshalBinary implements the encoding.BinaryMarshaler interface.
--- It returns a zero-length slice for the zero Addr,
--- the 4-byte form for an IPv4 address,
--- and the 16-byte form with zone appended for an IPv6 address.
---@return byte[], err
function netAddr:MarshalBinary() end

--- IsPrivate reports whether ip is a private address, according to RFC 1918
--- (IPv4 addresses) and RFC 4193 (IPv6 addresses). That is, it reports whether
--- ip is in 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or fc00::/7. This is the
--- same as net.IP.IsPrivate.
---@return boolean
function netAddr:IsPrivate() end

--- IsInterfaceLocalMulticast reports whether ip is an IPv6 interface-local
--- multicast address.
---@return boolean
function netAddr:IsInterfaceLocalMulticast() end

--- IsLinkLocalMulticast reports whether ip is a link-local multicast address.
---@return boolean
function netAddr:IsLinkLocalMulticast() end

--- IsLoopback reports whether ip is a loopback address.
---@return boolean
function netAddr:IsLoopback() end

--- IsValid reports whether the Addr is an initialized address (not the zero Addr).
---
--- Note that "0.0.0.0" and "::" are both valid values.
---@return boolean
function netAddr:IsValid() end

--- Zone returns ip's IPv6 scoped addressing zone, if any.
---@return string
function netAddr:Zone() end

--- Less reports whether ip sorts before ip2.
--- IP addresses sort first by length, then their address.
--- IPv6 addresses with zones sort just after the same address without a zone.
---@param ip2 netAddr
---@return boolean
function netAddr:Less(ip2) end

--- Is4In6 reports whether ip is an IPv4-mapped IPv6 address.
---@return boolean
function netAddr:Is4In6() end

--- IsUnspecified reports whether ip is an unspecified address, either the IPv4
--- address "0.0.0.0" or the IPv6 address "::".
---
--- Note that the zero Addr is not an unspecified address.
---@return boolean
function netAddr:IsUnspecified() end

--- Next returns the address following ip.
--- If there is none, it returns the zero Addr.
---@return netAddr
function netAddr:Next() end

--- BitLen returns the number of bits in the IP address:
--- 128 for IPv6, 32 for IPv4, and 0 for the zero Addr.
---
--- Note that IPv4-mapped IPv6 addresses are considered IPv6 addresses
--- and therefore have bit length 128.
---@return number
function netAddr:BitLen() end

--- AsSlice returns an IPv4 or IPv6 address in its respective 4-byte or 16-byte representation.
---@return byte[]
function netAddr:AsSlice() end

--- UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
--- It expects data in the form generated by MarshalBinary.
---@param b byte[]
---@return err
function netAddr:UnmarshalBinary(b) end

--- Is4 reports whether ip is an IPv4 address.
---
--- It returns false for IPv4-mapped IPv6 addresses. See Addr.Unmap.
---@return boolean
function netAddr:Is4() end

--- Unmap returns ip with any IPv4-mapped IPv6 address prefix removed.
---
--- That is, if ip is an IPv6 address wrapping an IPv4 address, it
--- returns the wrapped IPv4 address. Otherwise it returns ip unmodified.
---@return netAddr
function netAddr:Unmap() end

--- WithZone returns an IP that's the same as ip but with the provided
--- zone. If zone is empty, the zone is removed. If ip is an IPv4
--- address, WithZone is a no-op and returns ip unchanged.
---@param zone string
---@return netAddr
function netAddr:WithZone(zone) end

--- IsLinkLocalUnicast reports whether ip is a link-local unicast address.
---@return boolean
function netAddr:IsLinkLocalUnicast() end

--- IsGlobalUnicast reports whether ip is a global unicast address.
---
--- It returns true for IPv6 addresses which fall outside of the current
--- IANA-allocated 2000::/3 global unicast space, with the exception of the
--- link-local address space. It also returns true even if ip is in the IPv4
--- private address space or IPv6 unique local address space.
--- It returns false for the zero Addr.
---
--- For reference, see RFC 1122, RFC 4291, and RFC 4632.
---@return boolean
function netAddr:IsGlobalUnicast() end

--- String returns the string form of the IP address ip.
--- It returns one of 5 forms:
---
---   - "invalid IP", if ip is the zero Addr
---   - IPv4 dotted decimal ("192.0.2.1")
---   - IPv6 ("2001:db8::1")
---   - "::ffff:1.2.3.4" (if Is4In6)
---   - IPv6 with zone ("fe80:db8::1%eth0")
---
--- Note that unlike package net's IP.String method,
--- IPv4-mapped IPv6 addresses format with a "::ffff:"
--- prefix before the dotted quad.
---@return string
function netAddr:String() end

--- MarshalText implements the encoding.TextMarshaler interface,
--- The encoding is the same as returned by String, with one exception:
--- If ip is the zero Addr, the encoding is the empty string.
---@return byte[], err
function netAddr:MarshalText() end

--- UnmarshalText implements the encoding.TextUnmarshaler interface.
--- The IP address is expected in a form accepted by ParseAddr.
---
--- If text is empty, UnmarshalText sets *ip to the zero Addr and
--- returns no error.
---@param text byte[]
---@return err
function netAddr:UnmarshalText(text) end

--- StringExpanded is like String but IPv6 addresses are expanded with leading
--- zeroes and no "::" compression. For example, "2001:db8::1" becomes
--- "2001:0db8:0000:0000:0000:0000:0000:0001".
---@return string
function netAddr:StringExpanded() end

--- AddrPort is an IP and a port number.
---@class netipAddrPort
local netipAddrPort = {}


---@return string
function netipAddrPort:String() end

--- UnmarshalText implements the encoding.TextUnmarshaler
--- interface. The AddrPort is expected in a form
--- generated by MarshalText or accepted by ParseAddrPort.
---@param text byte[]
---@return err
function netipAddrPort:UnmarshalText(text) end

--- UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
--- It expects data in the form generated by MarshalBinary.
---@param b byte[]
---@return err
function netipAddrPort:UnmarshalBinary(b) end

--- Addr returns p's IP address.
---@return netAddr
function netipAddrPort:Addr() end

--- IsValid reports whether p.Addr() is valid.
--- All ports are valid, including zero.
---@return boolean
function netipAddrPort:IsValid() end

--- AppendTo appends a text encoding of p,
--- as generated by MarshalText,
--- to b and returns the extended buffer.
---@param b byte[]
---@return byte[]
function netipAddrPort:AppendTo(b) end

--- MarshalText implements the encoding.TextMarshaler interface. The
--- encoding is the same as returned by String, with one exception: if
--- p.Addr() is the zero Addr, the encoding is the empty string.
---@return byte[], err
function netipAddrPort:MarshalText() end

--- MarshalBinary implements the encoding.BinaryMarshaler interface.
--- It returns Addr.MarshalBinary with an additional two bytes appended
--- containing the port in little-endian.
---@return byte[], err
function netipAddrPort:MarshalBinary() end

--- Port returns p's port.
---@return any
function netipAddrPort:Port() end
