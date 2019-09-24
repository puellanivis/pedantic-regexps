package pedantic

const (
	asciiC0Control           = `\x00-\x1f\x7f`
	asciiControlNoWhitespace = `\x00-\x08\x0b\x0c\x0e-\x1f\x7f`
	asciiWhitespace          = `\t\n\r\x20`
	asciiUpper               = `A-Z`
	asciiLower               = `a-z`
	asciiAlpha               = asciiUpper + asciiLower
	asciiNumeric             = "0-9"
	asciiAlphaNumeric        = asciiAlpha + asciiNumeric
)
