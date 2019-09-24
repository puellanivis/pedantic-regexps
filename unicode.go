package pedantic

const (
	unicodeLetter        = `\p{L}`
	unicodeLetterNumber  = `[\p{L}\p{N}]`
	unicodeC1Control     = `\x80-\x9f`
	unicodeSurrogateHigh = `\x{d800}-\x{dbff}`
	unicodeSurrogateLow  = `\x{dc00}-\x{dfff}`
	unicodeSurrogate     = `\x{d800}-\x{dfff}`
	unicodeReplacement   = `\x{fffd}`

	// every codepoint that is in XXfffe-XXffff is a set of noncharacters.
	unicodeNoncharacter = `\x{fdd0}-\x{fdef}\x{00fffe}-\x{00ffff}\x{01fffe}-\x{01ffff}\x{02fffe}-\x{02ffff}\x{03fffe}-\x{03ffff}\x{04fffe}-\x{04ffff}\x{05fffe}-\x{05ffff}\x{06fffe}-\x{06ffff}\x{07fffe}-\x{07ffff}\x{08fffe}-\x{08ffff}\x{09fffe}-\x{09ffff}\x{0afffe}-\x{0affff}\x{0bfffe}-\x{0bffff}\x{0cfffe}-\x{0cffff}\x{0dfffe}-\x{0dffff}\x{0efffe}-\x{0effff}\x{0ffffe}-\x{0fffff}\x{10fffe}-\x{10ffff}`

	unicodeBMP     = `\xa0-\x{d7ff}\x{e000}-\x{fdcf}\x{fdf0}-\x{fffc}`
	unicodeBlock1  = `\x{010000}-\x{01fffd}`
	unicodeBlock2  = `\x{020000}-\x{02fffd}`
	unicodeBlock3  = `\x{030000}-\x{03fffd}`
	unicodeBlock4  = `\x{040000}-\x{04fffd}`
	unicodeBlock5  = `\x{050000}-\x{05fffd}`
	unicodeBlock6  = `\x{060000}-\x{06fffd}`
	unicodeBlock7  = `\x{070000}-\x{07fffd}`
	unicodeBlock8  = `\x{080000}-\x{08fffd}`
	unicodeBlock9  = `\x{090000}-\x{09fffd}`
	unicodeBlock10 = `\x{0a0000}-\x{0afffd}`
	unicodeBlock11 = `\x{0b0000}-\x{0bfffd}`
	unicodeBlock12 = `\x{0c0000}-\x{0cfffd}`
	unicodeBlock13 = `\x{0d0000}-\x{0dfffd}`
	unicodeBlock14 = `\x{0e0000}-\x{0efffd}`
	unicodeBlock15 = `\x{0f0000}-\x{0ffffd}`
	unicodeBlock16 = `\x{100000}-\x{10fffd}`

	// skips Surrogate, Replacement and Noncharacters (U+FDD0-U+FDEF and U+xxFFF0-U+xxFFFF)
	unicodeInvalid = unicodeSurrogate + unicodeReplacement + unicodeNoncharacter
	// skips above and C1 Control
	unicodeExclude = unicodeC1Control + unicodeInvalid

	unicodeNonASCII = unicodeBMP + unicodeBlock1 + unicodeBlock2 + unicodeBlock3 + unicodeBlock4 + unicodeBlock5 + unicodeBlock6 + unicodeBlock7 + unicodeBlock8 + unicodeBlock9 + unicodeBlock10 + unicodeBlock11 + unicodeBlock12 + unicodeBlock13 + unicodeBlock14 + unicodeBlock15 + unicodeBlock16
)
