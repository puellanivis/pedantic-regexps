package pedantic

import (
	"regexp"
)

const (
	ipv4Address = `(?:` + numberByteDec + `\.` + numberByteDec + `\.` + numberByteDec + `\.` + numberByteDec + `)`

	ipv6Hex        = numberUint16Hex
	ipv6HexZero    = `0`
	ipv6HexNotZero = `(?:[1-9A-Fa-f]` + digitHex + `{0,3})`

	ipv6Follow        = `(?::` + ipv6Hex + `)`
	ipv6FollowNotZero = `(?::` + ipv6HexNotZero + `)`

	ipv6FullAddress = `(?:` + ipv6Hex + `(?::` + ipv6Hex + `){7})`
	ipv6ShortAddr0  = `(?:::` + ipv6HexNotZero + ipv6Follow + `{0,5})`
	ipv6ShortAddr1  = `(?:` + ipv6HexNotZero + `::` + ipv6HexNotZero + ipv6Follow + `{0,4})`
	ipv6ShortAddr2  = `(?:` + ipv6Hex + ipv6FollowNotZero + `::` + ipv6HexNotZero + ipv6Follow + `{0,3})`
	ipv6ShortAddr3  = `(?:` + ipv6Hex + ipv6Follow + ipv6FollowNotZero + `::` + ipv6HexNotZero + ipv6Follow + `{0,2})`
	ipv6ShortAddr4  = `(?:` + ipv6Hex + ipv6Follow + `{2}` + ipv6FollowNotZero + `::` + ipv6HexNotZero + ipv6Follow + `?)`
	ipv6ShortAddr5  = `(?:` + ipv6Hex + ipv6Follow + `{3}` + ipv6FollowNotZero + `::` + ipv6HexNotZero + `?)`
	ipv6ShortAddrA  = `(?:` + ipv6Hex + ipv6Follow + `{0,4}` + ipv6FollowNotZero + `::)`
	ipv6ShortAddrB  = `(?:` + ipv6HexNotZero + `::)`

	ipv6Address = `(?:(?:::)|` + ipv6FullAddress + `|` + ipv6ShortAddr0 + `|` + ipv6ShortAddr1 + `|` + ipv6ShortAddr2 + `|` + ipv6ShortAddr3 + `|` + ipv6ShortAddr4 + `|` + ipv6ShortAddr5 + `|` + ipv6ShortAddrA + `|` + ipv6ShortAddrB + `)`
)

// IPv4 is a regexp that matches a valid ipv4 address.
var IPv4 = regexp.MustCompile(`^` + ipv4Address + `$`)

// IPv6 is a regexp that matches a valid ipv6 address.
var IPv6 = regexp.MustCompile(`^` + ipv6Address + `$`)
