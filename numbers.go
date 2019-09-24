package pedantic

const (
	digitBin = `[01]`
	digitOct = `[0-7]`
	digitDec = asciiNumeric
	digitHex = `[` + digitDec + `A-Fa-f]`

	numberByteBin = `(?:0|1[01]{0,7})`
	numberByteOct = `(?:0|[1-3](?:[0-7][0-7]?)?|[4-7][0-7]?)`
	numberByteDec = `(?:0|1(?:[0-9][0-9]?)?|2(?:[0-4][0-9]?|5[0-5]?|[6-9])?|[3-9][0-9]?)`
	numberByteHex = `(?:0|[1-9A-Fa-f][0-9A-Fa-f]?)`

	numberByteBinFull = `(?:` + digitBin + digitBin + digitBin + digitBin + digitBin + digitBin + digitBin + digitBin + `)`
	numberByteOctFull = `(?:[0-3][0-7][0-7])`
	numberByteDecFull = `(?:[01][0-9][0-9]|2[0-4][0-9]|25[0-5])`
	numberByteHexFull = `(?:` + digitHex + digitHex + `)`
)
