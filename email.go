package pedantic

import (
	"regexp"
)

const (
	emailFWS           = `\x20?`
	emailQtext         = emailFWS + `[^` + asciiWhitespace + `"\\` + utf8Invalid + `]+`
	emailQuotedPair    = emailFWS + `\\[^\n\r` + utf8Invalid + `]`
	emailQcontent      = `(?:` + emailQtext + `|` + emailQuotedPair + `)`
	emailQuotedString  = `"` + emailQcontent + `*` + emailFWS + `"`
	emailAtext         = `[!#\$%&'\*\+\-/` + asciiNumeric + `=\?` + asciiUpper + `\^_\x60` + asciiLower + `{|}~` + utf8NonASCII + `]+`
	emailDotAtom       = `(?:` + emailAtext + `(?:\.` + emailAtext + `)*)`
	emailLocalPart     = `(?:` + emailDotAtom + `|` + emailQuotedString + `)`
	emailDtext         = emailFWS + `[^` + asciiWhitespace + `\[\\\]` + utf8Invalid + `]+`
	emailDcontent      = `(?:` + emailDtext + `|` + emailQuotedPair + `)`
	emailDomainLiteral = `\[` + emailDcontent + `*` + emailFWS + `\]`
	emailDomain        = `(?:` + hostnameString + `|` + emailDomainLiteral + `)`
)

const emailString = emailLocalPart + `@` + emailDomain

// Email is a regexp that matches a valid email.
var Email = regexp.MustCompile(`^` + emailString + `$`)
