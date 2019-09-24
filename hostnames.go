package pedantic

import (
	"regexp"
)

const (
	hostnameLetter = `[` + asciiAlphaNumeric + utf8NonASCII + `]+`
	hostnameLabel  = hostnameLetter + `(?:-+` + hostnameLetter + `)*`
	hostnameString = hostnameLabel + `(?:\.` + hostnameLabel + `)*`
)

// Hostname is a regexp that matches a valid hostname.
var Hostname = regexp.MustCompile(`^` + hostnameString + `$`)
