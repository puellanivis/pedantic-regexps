# pedantic-regexps

The regexps exported by this package are overly strictly correct, and you should not use them.
They are provided essentially to give a baseline for the tests to validate against.
With valid tests in hand, one can test any particular validation regexp by injecting it into this library,
and seeing if it passes the tests.
With test results in hand, you can then evaluate if the failing test cases fall outside of your desired behavior.

## E-Mail Addresses

Some particular points of interest here.
[RFC-2822](https://tools.ietf.org/html/rfc2822) covering email addresses notes that some whitespace can occur in some quoted elements,
some validation libraries have interpreted this to mean that these whitespace are to be considered part of the email address.
However, definitively, any `CRLF` appearing in Folding-White-Space is semantically “invisible” and therefore not part of the quoted-string.
This definitively means that no email address may contain any `CR` or `LF`, as the `quote-pair` also does not allow `CR` or `LF` to be escaped.
Where any whitespace can occur in a quote, they are folded to a single whitespace,
this means that an email address of `"␣␣a"@example.com` can be definitely excluded,
because it would be unaddressable as it would fold to `"␣a"@example.com`. (Canonically, such a mailbox would have to be addressed with `"␣\␣a"@example.com` to prevent whitespace folding.)

## Hostnames

Hostname specifications say that each label should be no longer than 63 characters,
and that the whole hostname should be no longer than 255.
This is not something that can easily be checked with regexps,
and is not strictly enforced.

## Notes / Definitions

Unless otherwise noted,
where `␣` appears,
it is intended to represent the whitespace character at `U+0020` (SPACE).
It appears as this replacement character where folding-whitespace in Markdown would remove one or more space characters.
