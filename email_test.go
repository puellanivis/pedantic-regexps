package pedantic

import (
	"regexp"
	"regexp/syntax"
	"testing"
)

const emailRegexString = `^` + emailString + `$`

func TestEmail(t *testing.T) {
	input := emailRegexString

	r, err := syntax.Parse(input, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", input)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	tests := []test{
		// Simple 7bit ASCII:
		{"user@example.org", true},
		{"user", false},
		{"user@", false},
		{"@example.org", false},
		// Dotted forms:
		{"user.name@example.org", true},
		{".user@example.org", false},
		{"user.@example.org", false},
		{"user@example.org.", false},

		// Various special characters
		{" @example.org", false},
		{"!@example.org", true},
		{"\"@example.org", false},
		{"#@example.org", true},
		{"$@example.org", true},
		{"%@example.org", true},
		{"&@example.org", true},
		{"'@example.org", true},
		{"(@example.org", false},
		{")@example.org", false},
		{"*@example.org", true},
		{"+@example.org", true},
		{",@example.org", false},
		{"-@example.org", true},
		{".@example.org", false},
		{"/@example.org", true},
		{":@example.org", false},
		{";@example.org", false},
		{"<@example.org", false},
		{"=@example.org", true},
		{">@example.org", false},
		{"?@example.org", true},
		{"@@example.org", false},
		{"[@example.org", false},
		{"\\@example.org", false},
		{`"@"@example.org`, true},
		{"]@example.org", false},
		{"^@example.org", true},
		{"_@example.org", true},
		{"`@example.org", true},
		{"{@example.org", true},
		{"|@example.org", true},
		{"}@example.org", true},
		{"~@example.org", true},
		{"\x7f@example.org", false},

		// Basic Unicode test:
		{"\u00a0@example.org.com", true}, // UNSURE: standard says yes, but this is Unicode whitepace
		{"\u00a1@example.org.com", true},
		{"\ufffd@example.org.com", false}, // test that replacement character does not match

		// Various quote tests:
		{`""@example.org`, true},
		{`"@example.org`, false},
		{`"\""@example.org`, true},
		{`"\"@example.org`, false},
		{`"\\"@example.org`, true},
		{`"\\@example.org`, false},
		{`user@[]`, true},
		{`user@[`, false},
		{`user@]`, false},
		{`user@[\[]`, true}, // This is obsolete in RFC 5322. No escape in a domain-literal.
		{`user@\[]`, false},
		{`user@[\[`, false},
		{`user@[\]]`, true}, // This is obsolete in RFC 5322. No escape in a domain-literal.
		{`user@\]]`, false},
		{`user@[\]`, false},
		{`user@[\\]`, true}, // This is obsolete in RFC 5322. No escape in a domain-literal.
		{`user@\\]`, false},
		{`user@[\\`, false},
		{`user@[example.org]`, true}, // This is obsolete in RFC 5322. Must be `ipv4-addr`, or `IPv6:<ipv6-addr>`.

		// Test escaping whitespace in literals.
		{`"\ "@example.org`, true},     // This WSP is properly escaped.
		{`\ "@example.org`, false},     // This WSP is properly escaped, but no start quote.
		{`"\ @example.org`, false},     // This WSP is properly escaped, but no end quote.
		{"\"\\\t\"@example.org", true}, // This WSP is properly escaped. Yes, this is how it is encoded.
		{"\\\t\"@example.org", false},  // This WSP is properly escaped, but no start quote.
		{"\"\\\t@example.org", false},  // This WSP is properly escaped, but no end quote.
		{`user@[\ ]`, true},            // This WSP is properly escaped. This is obsolete in RFC 5322.
		{`user@\ ]`, false},            // This WSP is properly escaped, but no start bracket.
		{`user@[\ `, false},            // This WSP is properly escaped, but no end bracket.
		{"user@[\\\t]", true},          // This WSP is properly escaped. Yes, this is how it is encoded. This is obsolete in RFC 5322.
		{"user@\\\t]", false},          // This WSP is properly escaped, but no start bracket.
		{"user@[\\\t", false},          // This WSP is properly escaped, but no end bracket.

		// Test CR and LF not allowed:
		{"\"\n\"@example.org", false},
		{"\"\r\"@example.org", false},
		{"\"\\\n\"@example.org", false}, // even escaped, not allowed
		{"\"\\\r\"@example.org", false}, // even escaped, not allowed
		{"\"\r\n\"@example.org", false},
		{"\"\n\r\"@example.org", false},
		{"user@[\n]", false},
		{"user@[\r]", false},
		{"user@[\\\n]", false}, // even escaped, not allowed
		{"user@[\\\r]", false}, // even escaped, not allowed
		{"user@[\n\r]", false},
		{"user@[\r\n]", false},

		// Test Folding-White-Space (FWS) handling (it is invisible, and not part of the address):
		{"user @example.org", false}, // CANONICALLY: user@example.org
		{" user@example.org", false}, // CANONICALLY: user@example.org
		{"user@ example.org", false}, // CANONICALLY: user@example.org
		{"user@example.org ", false}, // CANONICALLY: user@example.org
		{"\" \"@example.org", true},
		{"\"  \"@example.org", false},    // CANONICALLY: " "@example.org
		{"\" \\  \"@example.org", true},  // quoted-string is: "   "
		{"\" \\\t \"@example.org", true}, // quoted-string is: " \t "
		{"\"\t\"@example.org", false},    // CANONICALLY: " "@example.org
		{" \"user\"@example.org", false}, // CANONICALLY: "user"@example.org
		{"\" user\"@example.org", true},
		{"\"user \"@example.org", true},
		{"\" user \"@example.org", true},
		{"\"\tuser\"@example.org", false},     // CANONICALLY: " user"@example.org
		{"\"user\t\"@example.org", false},     // CANONICALLY: "user "@example.org
		{"\"  user\"@example.org", false},     // CANONICALLY: " user"@example.org
		{"\"user  \"@example.org", false},     // CANONICALLY: "user "@example.org
		{"\"user\" @example.org", false},      // CANONICALLY: "user"@example.org
		{"user@ [example.org]", false},        // CANONICALLY: user@[example.org]
		{"user@[ example.org]", true},         // Obsolete in RFC 5322. Must be <ipv4-addr> or `IPv6:<ipv6-addr>`.
		{"user@[example.org ]", true},         // Obsolete in RFC 5322. Must be <ipv4-addr> or `IPv6:<ipv6-addr>`.
		{"user@[\texample.org]", false},       // CANONICALLY: user@[ example.org]
		{"user@[example.org\t]", false},       // CANONICALLY: user@[example.org ]
		{"user@[  example.org]", false},       // CANONICALLY: user@[ example.org]
		{"user@[example.org  ]", false},       // CANONICALLY: user@[example.org ]
		{"user@[example.org] ", false},        // CANONICALLY: user@[example.org]
		{"\"\r\n \"@example.org", false},      // CANONICALLY: " "@example.org
		{"\" \r\n \"@example.org", false},     // CANONICALLY: " "@example.org
		{"\"\r\n user\"@example.org", false},  // CANONICALLY: " user"@example.org
		{"\" \r\n user\"@example.org", false}, // CANONICALLY: " user"@example.org
		{"\"user\r\n \"@example.org", false},  // CANONICALLY: "user "@example.org
		{"\"user \r\n \"@example.org", false}, // CANONICALLY: "user "@example.org
		{"user@\r\n [example.org]", false},    // CANONICALLY: user@[example.org]
		{"user@ \r\n [example.org]", false},   // CANONICALLY: user@[example.org]
		{"user@[\r\n example.org]", false},    // CANONICALLY: user@[ example.org]
		{"user@[ \r\n example.org]", false},   // CANONICALLY: user@[ example.org]
		{"user@[example.org\r\n ]", false},    // CANONICALLY: user@[example.org ]
		{"user@[example.org \r\n ]", false},   // CANONICALLY: user@[example.org ]
		{"user@[example.org]\r\n ", false},    // CANONICALLY: user@[example.org]
		{"user@[example.org] \r\n ", false},   // CANONICALLY: user@[example.org]
		{"user@[ \\  ]", true},                // domain-literal is: "   ". Obsolete in RFC 5322.
		{"user@[ \\\t ]", true},               // domain-literal is: " \t ". Obsolete in RFC 5322.

		// Test Comment handling (it is invisible, and not part of the address):
		{"(comment)user@example.org", false},     // CANONICALLY: user@example.org
		{"user(comment)@example.org", false},     // CANONICALLY: user@example.org
		{"user@(comment)example.org", false},     // CANONICALLY: user@example.org
		{"user@example.org(comment)", false},     // CANONICALLY: user@example.org
		{"(comment)\"user\"@example.org", false}, // CANONICALLY: "user"@example.org
		{"\"(comment)user\"@example.org", true},  // Not a comment, but part of the local-part. Obsolete in RFC 5322.
		{"\"user(comment)\"@example.org", true},  // Not a comment, but part of the local-part. Obsolete in RFC 5322.
		{"\"user\"(comment)@example.org", false}, // CANONICALLY: "user"@example.org
		{"user@(comment)[example.org]", false},   // CANONICALLY: user@[example.org]
		{"user@[(comment)example.org]", true},    // Not a comment, but part of the domain-literal. Obsolete in RFC 5322.
		{"user@[example.org(comment)]", true},    // Not a comment, but part of the domain-literal. Obsolete in RFC 5322.
		{"user@[example.org](comment)", false},   // CANONICALLY: user@[example.org]

		// From: https://youtube.com/watch?v=mrGfahzt-4Q
		// Email vs Capitalism, or, Why We Can't Have Nice Things - Dylan Beattie - NDC Oslo 2023
		{"iron.man@avengers.com", true},
		{"spider-man@avengers.com", true},
		{"t'challa@avengers.com", true},
		{"rocket+groot@avengers.com", true},
		{`"Bruce 'The Hulk' Banner"@avengers.com`, true},
		{"vision@[IPv6:2001:db8:1ff::a0b:dbd0]", true},

		// From: https://e-mail.wtf/
		{"easy@example.com", true},
		{"easy+tag@example.com", true},
		{"easy@", false},
		{"@example.com", false},
		{"easy@example", true}, // Obsolete in RFC 2822
		{"what about spaces@example.com", false},
		{" maybe-like-this @example.com", false}, // CANONICALLY: maybe-like-this@example.com
		{"trailing-dot.@example.com", false},
		{".leading-dot@example.com", false}, // implied from trailing-dot
		{"middle.dot@example.com", true},
		{"fed-up-yet@ example.com ", false},         // CANONICALLY: fed-up-yet@example.com
		{"normal(wtf is this?)@example.com", false}, // CANONICALLY: normal@example.com
		{"(@)example.com", false},
		{`":(){ :|:& };:"@example.com`, true},
		{`""@example.com`, true},
		{"according-to-all-known-laws-of-aviation-there-is-no-way-a-bee-should-be-able-to-fly-its-wings-are-too-small-to-get-its-fat-little-body-off-the-ground-the-bee-of-course-flies-anyway-because-bees-don-t-care-what-humans-think-is-impossible-yellow-black-yellow-black-yellow-black-yellow-black-ooh-black-and-yellow-let-s-shake-it-up-a-little-barry-breakfast-is-ready-coming-hang-on-a-second-hello-barry-adam-can-you-believe-this-is-happening-i-can-t-i-ll-pick-you-up-looking-sharp-use-the-stairs-your-father-paid-good-money-for-those-sorry-i-m-excited-here-s-the-graduate-we-re-very-proud-of-you-son-a-perfect-report-card-all-b-s-very-proud-ma-i-got-a-thing-going-here-you-got-lint-on-your-fuzz-ow-that-s-me-wave-to-us-we-ll-be-in-row-118-000-bye-barry-i-told-you-stop-flying-in-the-house-hey-adam-hey-barry-is-that-fuzz-gel-a-little-special-day-graduation-never-thought-i-d-make-it-three-days-grade-school-three-days-high-school-those-were-awkward-three-days-college-i-m-glad-i-took-a-day-and-hitchhiked-around-the-hive-you-did-come-back-different-hi-barry-artie-growing-a-mustache-looks-good-hear-about-frankie-yeah-you-going-to-the-funeral-no-i-m-not-going-everybody-knows-sting-someone-you-die-don-t-waste-it-on-a-squirrel-such-a-hothead-i-guess-he-could-have-just-gotten-out-of-the-way-i-love-this-incorporating-an-amusement-park-into-our-day-that-s-why-we-don-t-need-vacations-boy-quite-a-bit-of-pomp-under-the-circumstances-well-adam-today-we-are-men-we-are-bee-men-amen-hallelujah-students-faculty-distinguished-bees-please-welcome-dean-buzzwell-welcome-new-hive-city-graduating-class-of-9-15-that-concludes-our-ceremonies-and-begins-your-career-at-honex-industries-will-we-pick-our-job-today-i-heard-it-s-just-orientation-heads-up-here-we-go-keep-your-hands-and-antennas-inside-the-tram-at-all-times-wonder-what-it-ll-be-like-a-little-scary-welcome-to-honex-a-division-of-honesco-and-a-part-of-the-hexagon-group-this-is-it-wow-wow-we-know-that-you-as-a-bee-have-worked-your-whole-life-to-get-to-the-point-where-you-can-work-for-your-whole-life-honey-begins-when-our-valiant-pollen-jocks-bring-the-nectar-to-the-hive-our-top-secret-formula-is-automatically-color-corrected-scent-adjusted-and-bubble-contoured-into-this-soothing-sweet-syrup-with-its-distinctive-golden-glow-you-know-as-honey-that-girl-was-hot-she-s-my-cousin-she-is-yes-we-re-all-cousins-right-you-re-right-at-honex-we-constantly-strive-to-improve-every-aspect-of-bee-existence-these-bees-are-stress-testing-a-new-helmet-technology-what-do-you-think-he-makes-not-enough-here-we-have-our-latest-advancement-the-krelman-what-does-that-do-catches-that-little-strand-of-honey-that-hangs-after-you-pour-it-saves-us-millions-can-anyone-work-on-the-krelman-of-course-most-bee-jobs-are-small-ones@example.com", true}, // Too long, length limit 998, cannot be enforced by regex.
		{"magic@[::1]", true},
		{"poop@[💩]", true},
		{"👉@👈", true},
		{`"@"@[@]`, true},
		{`"'()'"("''")@example.com`, false}, // CANONICALLY: "'()'"@example.com
		{"i...wonder@example.com", false},
		{"c̷̨̈́i̵̮̅l̶̠̐͊͝ȁ̷̠̗̆̍̍n̷͖̘̯̍̈͒̅t̶͍͂͋ř̵̞͈̓ȯ̷̯̠-̸͚̖̟͋s̴͉̦̭̔̆̃͒û̵̥̪͆̒̕c̸̨̨̧̺̎k̵̼͗̀s̸̖̜͍̲̈́͋̂͠@example.com", true},
	}

	re := regexp.MustCompile(input)

	for _, tt := range tests {
		got := re.MatchString(tt.s)
		if got != tt.match {
			switch tt.match {
			case true:
				t.Errorf("expected %q to match, but it did not", tt.s)
			case false:
				t.Errorf("expected %q to not match, but it did", tt.s)
			}
		}
	}
}
