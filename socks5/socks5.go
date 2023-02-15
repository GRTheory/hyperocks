package main

import "errors"

const (
	Version         = 5
	UserPassVersion = 1
)

const (
	MethodNoAuth uint8 = iota
	MethodGSSAPI
	MethodUserPass
	MethodNoAcceptable uint8 = 0xFF
)

// Commands
const (
	CmdConnect uint8 = iota + 1
	CmdBind
	CmdUDP
	CmdUDPOverTCP
)

const (
	AddrIPv4   uint8 = 1
	AddrDomain uint8 = 3
	AddrIPv6   uint8 = 4
)

const (
	Succeeded uint8 = iota
	Failure
	Allowed
	NetUnreachable
	HostUnreachable
	ConnRefused
	TTLExpired
	CmdUnsupported
	AddrUnsupported
)

// Errors
var (
	ErrBadVersion  = errors.New("bad version")
	ErrBadFormat   = errors.New("bad format")
	ErrBadAddrType = errors.New("bad address type")
	ErrShortBuffer = errors.New("short buffer")
	ErrBadMethod   = errors.New("bad method")
	ErrAuthFailure = errors.New("auth failure")
)

