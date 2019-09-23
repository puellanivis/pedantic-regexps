package pedantic

import (
	"regexp"
)

const (
	emailQtext         = `[^\t\n\r "\\]+`
	emailQuotedPair    = `\\[^\n\r]`
	emailQuotedString  = `"(?:` + emailQtext + `|` + emailQuotedPair + `)*"`
	emailAtext         = `[!#\$%&'\*\+\-/0-9=\?A-Z\^_\x60a-z{|}~` + utf8NonASCII + `]+`
	emailDotAtom       = `(?:` + emailAtext + `(?:\.` + emailAtext + `)*)`
	emailLocalPart     = `(?:` + emailDotAtom + `|` + emailQuotedString + `)`
	emailDtext         = `[^\t\n\r \[\\\]]+`
	emailDcontent      = `(?:` + emailDtext + `|` + emailQuotedPair + `)`
	emailDomainLiteral = `\[` + emailDcontent + `*\]`
	emailDomain        = `(?:` + hostnameString + `|` + emailDomainLiteral + `)`
)

const emailString = emailLocalPart + `@` + emailDomain

// Email is a regexp that matches a valid email.
var Email = regexp.MustCompile(`^` + emailString + `$`)
