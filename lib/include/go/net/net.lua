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

---{{.netIPv4}}
---@param a byte
---@param b byte
---@param c byte
---@param d byte
---@return netIP
function net.IPv4(a, b, c, d)
end

---{{.netLookupNS}}
---@param name string
---@return any, err
function net.LookupNS(name)
end

---{{.netLookupTXT}}
---@param name string
---@return string[], err
function net.LookupTXT(name)
end

---{{.netPipe}}
---@return netConn, netConn
function net.Pipe()
end

---{{.netIPv4Mask}}
---@param a byte
---@param b byte
---@param c byte
---@param d byte
---@return netIPMask
function net.IPv4Mask(a, b, c, d)
end

---{{.netResolveIPAddr}}
---@param network string
---@param address string
---@return netIPAddr, err
function net.ResolveIPAddr(network, address)
end

---{{.netLookupPort}}
---@param network string
---@param service string
---@return number, err
function net.LookupPort(network, service)
end

---{{.netDialTCP}}
---@param network string
---@param laddr netTCPAddr
---@param raddr netTCPAddr
---@return netTCPConn, err
function net.DialTCP(network, laddr, raddr)
end

---{{.netDialUDP}}
---@param network string
---@param laddr netUDPAddr
---@param raddr netUDPAddr
---@return netUDPConn, err
function net.DialUDP(network, laddr, raddr)
end

---{{.netListenPacket}}
---@param network string
---@param address string
---@return netPacketConn, err
function net.ListenPacket(network, address)
end

---{{.netFilePacketConn}}
---@param f osFile
---@return netPacketConn, err
function net.FilePacketConn(f)
end

---{{.netResolveUDPAddr}}
---@param network string
---@param address string
---@return netUDPAddr, err
function net.ResolveUDPAddr(network, address)
end

---{{.netInterfaceAddrs}}
---@return any, err
function net.InterfaceAddrs()
end

---{{.netParseIP}}
---@param s string
---@return netIP
function net.ParseIP(s)
end

---{{.netTCPAddrFromAddrPort}}
---@param addr netipAddrPort
---@return netTCPAddr
function net.TCPAddrFromAddrPort(addr)
end

---{{.netListenMulticastUDP}}
---@param network string
---@param ifi netInterface
---@param gaddr netUDPAddr
---@return netUDPConn, err
function net.ListenMulticastUDP(network, ifi, gaddr)
end

---{{.netDialUnix}}
---@param network string
---@param laddr netUnixAddr
---@param raddr netUnixAddr
---@return netUnixConn, err
function net.DialUnix(network, laddr, raddr)
end

---{{.netInterfaceByIndex}}
---@param index number
---@return netInterface, err
function net.InterfaceByIndex(index)
end

---{{.netInterfaceByName}}
---@param name string
---@return netInterface, err
function net.InterfaceByName(name)
end

---{{.netCIDRMask}}
---@param ones number
---@param bits number
---@return netIPMask
function net.CIDRMask(ones, bits)
end

---{{.netLookupIP}}
---@param host string
---@return any, err
function net.LookupIP(host)
end

---{{.netLookupCNAME}}
---@param host string
---@return string, err
function net.LookupCNAME(host)
end

---{{.netListenUnix}}
---@param network string
---@param laddr netUnixAddr
---@return netUnixListener, err
function net.ListenUnix(network, laddr)
end

---{{.netListenUnixgram}}
---@param network string
---@param laddr netUnixAddr
---@return netUnixConn, err
function net.ListenUnixgram(network, laddr)
end

---{{.netFileListener}}
---@param f osFile
---@return netListener, err
function net.FileListener(f)
end

---{{.netParseCIDR}}
---@param s string
---@return netIP, netIPNet, err
function net.ParseCIDR(s)
end

---{{.netListenIP}}
---@param network string
---@param laddr netIPAddr
---@return netIPConn, err
function net.ListenIP(network, laddr)
end

---{{.netLookupAddr}}
---@param addr string
---@return string[], err
function net.LookupAddr(addr)
end

---{{.netListenTCP}}
---@param network string
---@param laddr netTCPAddr
---@return netTCPListener, err
function net.ListenTCP(network, laddr)
end

---{{.netLookupSRV}}
---@param service string
---@param proto string
---@param name string
---@return string, any, err
function net.LookupSRV(service, proto, name)
end

---{{.netLookupMX}}
---@param name string
---@return any, err
function net.LookupMX(name)
end

---{{.netLookupHost}}
---@param host string
---@return string[], err
function net.LookupHost(host)
end

---{{.netDialTimeout}}
---@param network string
---@param address string
---@param timeout timeDuration
---@return netConn, err
function net.DialTimeout(network, address, timeout)
end

---{{.netListen}}
---@param network string
---@param address string
---@return netListener, err
function net.Listen(network, address)
end

---{{.netFileConn}}
---@param f osFile
---@return netConn, err
function net.FileConn(f)
end

---{{.netJoinHostPort}}
---@param host string
---@param port string
---@return string
function net.JoinHostPort(host, port)
end

---{{.netSplitHostPort}}
---@param hostport string
---@return string, err
function net.SplitHostPort(hostport)
end

---{{.netResolveTCPAddr}}
---@param network string
---@param address string
---@return netTCPAddr, err
function net.ResolveTCPAddr(network, address)
end

---{{.netListenUDP}}
---@param network string
---@param laddr netUDPAddr
---@return netUDPConn, err
function net.ListenUDP(network, laddr)
end

---{{.netUDPAddrFromAddrPort}}
---@param addr netipAddrPort
---@return netUDPAddr
function net.UDPAddrFromAddrPort(addr)
end

---{{.netDial}}
---@param network string
---@param address string
---@return netConn, err
function net.Dial(network, address)
end

---{{.netInterfaces}}
---@return any, err
function net.Interfaces()
end

---{{.netDialIP}}
---@param network string
---@param laddr netIPAddr
---@param raddr netIPAddr
---@return netIPConn, err
function net.DialIP(network, laddr, raddr)
end

---{{.netParseMAC}}
---@param s string
---@return netHardwareAddr, err
function net.ParseMAC(s)
end

---{{.netResolveUnixAddr}}
---@param network string
---@param address string
---@return netUnixAddr, err
function net.ResolveUnixAddr(network, address)
end

---@class any
local any = {}

---@class netNS
---@field Host string
local netNS = {}

---@class netFlags
local netFlags = {}

---{{.netFlagsString}}
---@return string
function netFlags:String()
end

---@class netConn
local netConn = {}

---@class netAddrError
---@field Err string
---@field Addr string
local netAddrError = {}

---{{.netAddrErrorError}}
---@return string
function netAddrError:Error()
end

---{{.netAddrErrorTimeout}}
---@return boolean
function netAddrError:Timeout()
end

---{{.netAddrErrorTemporary}}
---@return boolean
function netAddrError:Temporary()
end

---@class any
local any = {}

---@class netDNSError
---@field Err string
---@field Name string
---@field Server string
---@field IsTimeout boolean
---@field IsTemporary boolean
---@field IsNotFound boolean
local netDNSError = {}

---{{.netDNSErrorError}}
---@return string
function netDNSError:Error()
end

---{{.netDNSErrorTimeout}}
---@return boolean
function netDNSError:Timeout()
end

---{{.netDNSErrorTemporary}}
---@return boolean
function netDNSError:Temporary()
end

---@class netBuffers
local netBuffers = {}

---{{.netBuffersWriteTo}}
---@param w ioWriter
---@return number, err
function netBuffers:WriteTo(w)
end

---{{.netBuffersRead}}
---@param p byte[]
---@return number, err
function netBuffers:Read(p)
end

---@class netOpError
---@field Op string
---@field Net string
---@field Source netAddr
---@field Addr netAddr
---@field Err err
local netOpError = {}

---{{.netOpErrorUnwrap}}
---@return err
function netOpError:Unwrap()
end

---{{.netOpErrorError}}
---@return string
function netOpError:Error()
end

---{{.netOpErrorTimeout}}
---@return boolean
function netOpError:Timeout()
end

---{{.netOpErrorTemporary}}
---@return boolean
function netOpError:Temporary()
end

---@class netUnknownNetworkError
local netUnknownNetworkError = {}

---{{.netUnknownNetworkErrorError}}
---@return string
function netUnknownNetworkError:Error()
end

---{{.netUnknownNetworkErrorTimeout}}
---@return boolean
function netUnknownNetworkError:Timeout()
end

---{{.netUnknownNetworkErrorTemporary}}
---@return boolean
function netUnknownNetworkError:Temporary()
end

---@class netAddr
local netAddr = {}

---@class netError
local netError = {}

---@class any
local any = {}

---@class netListenConfig
---@field Control any
---@field KeepAlive any
local netListenConfig = {}

---{{.netListenConfigListen}}
---@param ctx contextContext
---@param network string
---@param address string
---@return netListener, err
function netListenConfig:Listen(ctx, network, address)
end

---{{.netListenConfigListenPacket}}
---@param ctx contextContext
---@param network string
---@param address string
---@return netPacketConn, err
function netListenConfig:ListenPacket(ctx, network, address)
end

---@class netSRV
---@field Target string
---@field Port any
---@field Priority any
---@field Weight any
local netSRV = {}

---@class netListener
local netListener = {}

---@class netIPNet
---@field IP netIP
---@field Mask netIPMask
local netIPNet = {}

---{{.netIPNetContains}}
---@param ip netIP
---@return boolean
function netIPNet:Contains(ip)
end

---{{.netIPNetNetwork}}
---@return string
function netIPNet:Network()
end

---{{.netIPNetString}}
---@return string
function netIPNet:String()
end

---@class netParseError
---@field Type string
---@field Text string
local netParseError = {}

---{{.netParseErrorTimeout}}
---@return boolean
function netParseError:Timeout()
end

---{{.netParseErrorTemporary}}
---@return boolean
function netParseError:Temporary()
end

---{{.netParseErrorError}}
---@return string
function netParseError:Error()
end

---@class any
local any = {}

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

---{{.netDialerDial}}
---@param network string
---@param address string
---@return netConn, err
function netDialer:Dial(network, address)
end

---{{.netDialerDialContext}}
---@param ctx contextContext
---@param network string
---@param address string
---@return netConn, err
function netDialer:DialContext(ctx, network, address)
end

---@class netMX
---@field Host string
---@field Pref any
local netMX = {}

---@class netIPMask
local netIPMask = {}

---{{.netIPMaskSize}}
---@return number
function netIPMask:Size()
end

---{{.netIPMaskString}}
---@return string
function netIPMask:String()
end

---@class netIP
local netIP = {}

---{{.netIPIsPrivate}}
---@return boolean
function netIP:IsPrivate()
end

---{{.netIPTo4}}
---@return netIP
function netIP:To4()
end

---{{.netIPMask}}
---@param mask netIPMask
---@return netIP
function netIP:Mask(mask)
end

---{{.netIPUnmarshalText}}
---@param text byte[]
---@return err
function netIP:UnmarshalText(text)
end

---{{.netIPIsInterfaceLocalMulticast}}
---@return boolean
function netIP:IsInterfaceLocalMulticast()
end

---{{.netIPIsLinkLocalUnicast}}
---@return boolean
function netIP:IsLinkLocalUnicast()
end

---{{.netIPDefaultMask}}
---@return netIPMask
function netIP:DefaultMask()
end

---{{.netIPString}}
---@return string
function netIP:String()
end

---{{.netIPEqual}}
---@param x netIP
---@return boolean
function netIP:Equal(x)
end

---{{.netIPIsLoopback}}
---@return boolean
function netIP:IsLoopback()
end

---{{.netIPIsMulticast}}
---@return boolean
function netIP:IsMulticast()
end

---{{.netIPIsGlobalUnicast}}
---@return boolean
function netIP:IsGlobalUnicast()
end

---{{.netIPMarshalText}}
---@return byte[], err
function netIP:MarshalText()
end

---{{.netIPIsUnspecified}}
---@return boolean
function netIP:IsUnspecified()
end

---{{.netIPIsLinkLocalMulticast}}
---@return boolean
function netIP:IsLinkLocalMulticast()
end

---{{.netIPTo16}}
---@return netIP
function netIP:To16()
end

---@class netInvalidAddrError
local netInvalidAddrError = {}

---{{.netInvalidAddrErrorError}}
---@return string
function netInvalidAddrError:Error()
end

---{{.netInvalidAddrErrorTimeout}}
---@return boolean
function netInvalidAddrError:Timeout()
end

---{{.netInvalidAddrErrorTemporary}}
---@return boolean
function netInvalidAddrError:Temporary()
end

---@class any
local any = {}

---@class netHardwareAddr
local netHardwareAddr = {}

---{{.netHardwareAddrString}}
---@return string
function netHardwareAddr:String()
end

---@class netPacketConn
local netPacketConn = {}

---@class any
local any = {}

---{{.anySetUnlinkOnClose}}
---@param unlink boolean
function any:SetUnlinkOnClose(unlink)
end

---@class any
local any = {}

---@class netInterface
---@field Index number
---@field MTU number
---@field Name string
---@field HardwareAddr netHardwareAddr
---@field Flags netFlags
local netInterface = {}

---{{.netInterfaceAddrs}}
---@return any, err
function netInterface:Addrs()
end

---{{.netInterfaceMulticastAddrs}}
---@return any, err
function netInterface:MulticastAddrs()
end

---@class any
local any = {}

---@class any
local any = {}

---@class any
local any = {}

---@class any
local any = {}

---@class netDNSConfigError
---@field Err err
local netDNSConfigError = {}

---{{.netDNSConfigErrorUnwrap}}
---@return err
function netDNSConfigError:Unwrap()
end

---{{.netDNSConfigErrorError}}
---@return string
function netDNSConfigError:Error()
end

---{{.netDNSConfigErrorTimeout}}
---@return boolean
function netDNSConfigError:Timeout()
end

---{{.netDNSConfigErrorTemporary}}
---@return boolean
function netDNSConfigError:Temporary()
end
