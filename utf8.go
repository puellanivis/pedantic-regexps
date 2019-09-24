package pedantic

const (
	utf8Letter       = unicodeLetter
	utf8LetterNumber = unicodeLetterNumber
	utf8C1Control    = unicodeC1Control
	utf8Surrogate    = unicodeSurrogate
	utf8Replacement  = unicodeReplacement
	utf8Noncharacter = unicodeNoncharacter

	// Excludes Surrogates
	utf8BMP = `\xa0-\x{d7ff}\x{e000}-\x{fdcf}\x{fdf0}-\x{fffc}`

	// skips Surrogates, Replacement, Noncharacters (U+FDD0-U+FDEF and U+xxFFF0-U+xxFFFF)
	utf8Invalid = utf8Surrogate + utf8Replacement + utf8Noncharacter
	// skips above and C1 Control
	utf8Exclude = utf8C1Control + utf8Invalid

	utf8NonASCII = utf8BMP + unicodeBlock1 + unicodeBlock2 + unicodeBlock3 + unicodeBlock4 + unicodeBlock5 + unicodeBlock6 + unicodeBlock7 + unicodeBlock8 + unicodeBlock9 + unicodeBlock10 + unicodeBlock11 + unicodeBlock12 + unicodeBlock13 + unicodeBlock14 + unicodeBlock15 + unicodeBlock16
)
