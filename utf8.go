package pedantic

const (
	utf8C1Control    = `\x80-\x9f`
	utf8Surrogate    = `\x{d800}-\x{dbff}`
	utf8Noncharacter = `\x{fdd0}-\x{fdef}`
	utf8Replacement  = `\x{fffd}`
	// every codepoint that is at XXfffe and XXffff is a second-set of noncharacters.
	utf8Special = `\x{00fffe}-\x{00ffff}\x{01fffe}-\x{01ffff}\x{02fffe}-\x{02ffff}\x{03fffe}-\x{03ffff}\x{04fffe}-\x{04ffff}\x{05fffe}-\x{05ffff}\x{06fffe}-\x{06ffff}\x{07fffe}-\x{07ffff}\x{08fffe}-\x{08ffff}\x{09fffe}-\x{09ffff}\x{0afffe}-\x{0affff}\x{0bfffe}-\x{0bffff}\x{0cfffe}-\x{0cffff}\x{0dfffe}-\x{0dffff}\x{0efffe}-\x{0effff}\x{0ffffe}-\x{0fffff}\x{10fffe}-\x{10ffff}`
	// skips Surrogates, Noncharacters (U+FDD0-U+FDEF), and Specials (U+xxFFF0-U+xxFFFF)
	utf8Invalid = utf8Surrogate + utf8Noncharacter + utf8Replacement + utf8Special
	// skips above and C1 Control
	utf8Exclude  = utf8C1Control + utf8Invalid
	utf8NonASCII = `\xa0-\x{d7ff}\x{e000}-\x{fdcf}\x{fdf0}-\x{fffc}\x{10000}-\x{1fffd}\x{20000}-\x{2fffd}\x{30000}-\x{3fffd}\x{40000}-\x{4fffd}\x{50000}-\x{5fffd}\x{60000}-\x{6fffd}\x{70000}-\x{7fffd}\x{80000}-\x{8fffd}\x{90000}-\x{9fffd}\x{a0000}-\x{afffd}\x{b0000}-\x{bfffd}\x{c0000}-\x{cfffd}\x{d0000}-\x{dfffd}\x{e0000}-\x{efffd}\x{f0000}-\x{ffffd}\x{100000}-\x{10fffd}`
)
