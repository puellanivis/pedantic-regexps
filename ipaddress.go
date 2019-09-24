package pedantic

import (
	"regexp"
)

const (
	ipv4Number  = `(?:0|1(?:[0-9][0-9]?)?|2(?:[0-4][0-9]?|5[0-5]?|[6-9])?|[3-9][0-9]?)`
	ipv4Address = `(?:` + ipv4Number + `\.` + ipv4Number + `\.` + ipv4Number + `\.` + ipv4Number + `)`
)

// IPv4 is a regexp that matches a valid ipv4 address.
var IPv4 = regexp.MustCompile(`^` + ipv4Address + `$`)
