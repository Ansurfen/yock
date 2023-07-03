-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

netip = {}

---{{.netipIPv6Loopback}}
---@return netAddr
function netip.IPv6Loopback()
end

---{{.netipIPv6LinkLocalAllNodes}}
---@return netAddr
function netip.IPv6LinkLocalAllNodes()
end

---{{.netipMustParseAddrPort}}
---@param s string
---@return netipAddrPort
function netip.MustParseAddrPort(s)
end

---{{.netipAddrPortFrom}}
---@param ip netAddr
---@param port any
---@return netipAddrPort
function netip.AddrPortFrom(ip, port)
end

---{{.netipParseAddrPort}}
---@param s string
---@return netipAddrPort, err
function netip.ParseAddrPort(s)
end

---{{.netipAddrFrom4}}
---@param addr byte[]
---@return netAddr
function netip.AddrFrom4(addr)
end

---{{.netipAddrFrom16}}
---@param addr byte[]
---@return netAddr
function netip.AddrFrom16(addr)
end

---{{.netipParseAddr}}
---@param s string
---@return netAddr, err
function netip.ParseAddr(s)
end

---{{.netipAddrFromSlice}}
---@param slice byte[]
---@return netAddr, boolean
function netip.AddrFromSlice(slice)
end

---{{.netipParsePrefix}}
---@param s string
---@return netipPrefix, err
function netip.ParsePrefix(s)
end

---{{.netipMustParsePrefix}}
---@param s string
---@return netipPrefix
function netip.MustParsePrefix(s)
end

---{{.netipMustParseAddr}}
---@param s string
---@return netAddr
function netip.MustParseAddr(s)
end

---{{.netipIPv6Unspecified}}
---@return netAddr
function netip.IPv6Unspecified()
end

---{{.netipIPv4Unspecified}}
---@return netAddr
function netip.IPv4Unspecified()
end

---{{.netipPrefixFrom}}
---@param ip netAddr
---@param bits number
---@return netipPrefix
function netip.PrefixFrom(ip, bits)
end

---{{.netipIPv6LinkLocalAllRouters}}
---@return netAddr
function netip.IPv6LinkLocalAllRouters()
end

---@class netipPrefix
local netipPrefix = {}

---{{.netipPrefixMasked}}
---@return netipPrefix
function netipPrefix:Masked()
end

---{{.netipPrefixContains}}
---@param ip netAddr
---@return boolean
function netipPrefix:Contains(ip)
end

---{{.netipPrefixAppendTo}}
---@param b byte[]
---@return byte[]
function netipPrefix:AppendTo(b)
end

---{{.netipPrefixMarshalText}}
---@return byte[], err
function netipPrefix:MarshalText()
end

---{{.netipPrefixAddr}}
---@return netAddr
function netipPrefix:Addr()
end

---{{.netipPrefixIsValid}}
---@return boolean
function netipPrefix:IsValid()
end

---{{.netipPrefixUnmarshalText}}
---@param text byte[]
---@return err
function netipPrefix:UnmarshalText(text)
end

---{{.netipPrefixMarshalBinary}}
---@return byte[], err
function netipPrefix:MarshalBinary()
end

---{{.netipPrefixIsSingleIP}}
---@return boolean
function netipPrefix:IsSingleIP()
end

---{{.netipPrefixString}}
---@return string
function netipPrefix:String()
end

---{{.netipPrefixBits}}
---@return number
function netipPrefix:Bits()
end

---{{.netipPrefixOverlaps}}
---@param o netipPrefix
---@return boolean
function netipPrefix:Overlaps(o)
end

---{{.netipPrefixUnmarshalBinary}}
---@param b byte[]
---@return err
function netipPrefix:UnmarshalBinary(b)
end

---@class netAddr
local netAddr = {}

---{{.netAddrUnmarshalBinary}}
---@param b byte[]
---@return err
function netAddr:UnmarshalBinary(b)
end

---{{.netAddrBitLen}}
---@return number
function netAddr:BitLen()
end

---{{.netAddrZone}}
---@return string
function netAddr:Zone()
end

---{{.netAddrIs6}}
---@return boolean
function netAddr:Is6()
end

---{{.netAddrAppendTo}}
---@param b byte[]
---@return byte[]
function netAddr:AppendTo(b)
end

---{{.netAddrUnmap}}
---@return netAddr
function netAddr:Unmap()
end

---{{.netAddrIsLoopback}}
---@return boolean
function netAddr:IsLoopback()
end

---{{.netAddrIsMulticast}}
---@return boolean
function netAddr:IsMulticast()
end

---{{.netAddrString}}
---@return string
function netAddr:String()
end

---{{.netAddrIsValid}}
---@return boolean
function netAddr:IsValid()
end

---{{.netAddrIsInterfaceLocalMulticast}}
---@return boolean
function netAddr:IsInterfaceLocalMulticast()
end

---{{.netAddrAs4}}
---@return byte[]
function netAddr:As4()
end

---{{.netAddrNext}}
---@return netAddr
function netAddr:Next()
end

---{{.netAddrMarshalBinary}}
---@return byte[], err
function netAddr:MarshalBinary()
end

---{{.netAddrIs4In6}}
---@return boolean
function netAddr:Is4In6()
end

---{{.netAddrIsUnspecified}}
---@return boolean
function netAddr:IsUnspecified()
end

---{{.netAddrStringExpanded}}
---@return string
function netAddr:StringExpanded()
end

---{{.netAddrLess}}
---@param ip2 netAddr
---@return boolean
function netAddr:Less(ip2)
end

---{{.netAddrIs4}}
---@return boolean
function netAddr:Is4()
end

---{{.netAddrIsLinkLocalMulticast}}
---@return boolean
function netAddr:IsLinkLocalMulticast()
end

---{{.netAddrIsGlobalUnicast}}
---@return boolean
function netAddr:IsGlobalUnicast()
end

---{{.netAddrPrefix}}
---@param b number
---@return netipPrefix, err
function netAddr:Prefix(b)
end

---{{.netAddrCompare}}
---@param ip2 netAddr
---@return number
function netAddr:Compare(ip2)
end

---{{.netAddrIsLinkLocalUnicast}}
---@return boolean
function netAddr:IsLinkLocalUnicast()
end

---{{.netAddrAs16}}
---@return byte[]
function netAddr:As16()
end

---{{.netAddrMarshalText}}
---@return byte[], err
function netAddr:MarshalText()
end

---{{.netAddrWithZone}}
---@param zone string
---@return netAddr
function netAddr:WithZone(zone)
end

---{{.netAddrIsPrivate}}
---@return boolean
function netAddr:IsPrivate()
end

---{{.netAddrAsSlice}}
---@return byte[]
function netAddr:AsSlice()
end

---{{.netAddrPrev}}
---@return netAddr
function netAddr:Prev()
end

---{{.netAddrUnmarshalText}}
---@param text byte[]
---@return err
function netAddr:UnmarshalText(text)
end

---@class netipAddrPort
local netipAddrPort = {}

---{{.netipAddrPortAppendTo}}
---@param b byte[]
---@return byte[]
function netipAddrPort:AppendTo(b)
end

---{{.netipAddrPortMarshalText}}
---@return byte[], err
function netipAddrPort:MarshalText()
end

---{{.netipAddrPortUnmarshalText}}
---@param text byte[]
---@return err
function netipAddrPort:UnmarshalText(text)
end

---{{.netipAddrPortMarshalBinary}}
---@return byte[], err
function netipAddrPort:MarshalBinary()
end

---{{.netipAddrPortUnmarshalBinary}}
---@param b byte[]
---@return err
function netipAddrPort:UnmarshalBinary(b)
end

---{{.netipAddrPortString}}
---@return string
function netipAddrPort:String()
end

---{{.netipAddrPortPort}}
---@return any
function netipAddrPort:Port()
end

---{{.netipAddrPortIsValid}}
---@return boolean
function netipAddrPort:IsValid()
end

---{{.netipAddrPortAddr}}
---@return netAddr
function netipAddrPort:Addr()
end
