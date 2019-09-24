package pedantic

const (
	digitHex      = `[0-9A-Fa-f]`
	numberByteDec = `(?:0|1(?:[0-9][0-9]?)?|2(?:[0-4][0-9]?|5[0-5]?|[6-9])?|[3-9][0-9]?)`
	numberByteHex = `(?:` + digitHex + digitHex + `)`
)
