package pedantic

import (
	"regexp"
)

const (
	emailFWS           = `[ ]?`
	emailQtext         = emailFWS + `[^\t\n\r "\\` + utf8Invalid + `]+`
	emailQuotedPair    = emailFWS + `\\[^\n\r` + utf8Invalid + `]`
	emailQcontent      = `(?:` + emailQtext + `|` + emailQuotedPair + `)`
	emailQuotedString  = `"` + emailQcontent + `*` + emailFWS + `"`
	emailAtext         = `[!#\$%&'\*\+\-/0-9=\?A-Z\^_\x60a-z{|}~` + utf8NonASCII + `]+`
	emailDotAtom       = `(?:` + emailAtext + `(?:\.` + emailAtext + `)*)`
	emailLocalPart     = `(?:` + emailDotAtom + `|` + emailQuotedString + `)`
	emailDtext         = emailFWS + `[^\t\n\r \[\\\]` + utf8Invalid + `]+`
	emailDcontent      = `(?:` + emailDtext + `|` + emailQuotedPair + `)`
	emailDomainLiteral = `\[` + emailDcontent + `*` + emailFWS + `\]`
	emailDomain        = `(?:` + hostnameString + `|` + emailDomainLiteral + `)`
)

const emailString = emailLocalPart + `@` + emailDomain

// Email is a regexp that matches a valid email.
var Email = regexp.MustCompile(`^` + emailString + `$`)
