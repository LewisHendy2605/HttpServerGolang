package request

/*
1.2. Syntax Notation

This specification uses the Augmented Backus-Naur Form (ABNF) notation of [RFC5234], extended with the notation for case-sensitivity in strings defined in [RFC7405].

It also uses a list extension, defined in Section 5.6.1 of [HTTP], that allows for compact definition of comma-separated lists using a "#" operator (similar to how the "*" operator indicates repetition).

Appendix A shows the collected grammar with all list operators expanded to standard ABNF notation.

As a convention, ABNF rule names prefixed with "obs-" denote obsolete grammar rules that appear for historical reasons.

The following core rules are included by reference, as defined in [RFC5234],
Appendix B.1: ALPHA (letters),

	CR (carriage return),
	CRLF (CR LF),
	CTL (controls),
	DIGIT (decimal 0-9),
	DQUOTE (double quote),
	HEXDIG (hexadecimal 0-9/A-F/a-f),
	HTAB (horizontal tab), LF (line feed),
	OCTET (any 8-bit sequence of data),
	SP (space),
	and VCHAR (any visible [USASCII] character).

The rules below are defined in [HTTP]:

	BWS           = <BWS, see [HTTP], Section 5.6.3>
	OWS           = <OWS, see [HTTP], Section 5.6.3>
	RWS           = <RWS, see [HTTP], Section 5.6.3>
	absolute-path = <absolute-path, see [HTTP], Section 4.1>
	field-name    = <field-name, see [HTTP], Section 5.1>
	field-value   = <field-value, see [HTTP], Section 5.5>
	obs-text      = <obs-text, see [HTTP], Section 5.6.4>
	quoted-string = <quoted-string, see [HTTP], Section 5.6.4>
	token         = <token, see [HTTP], Section 5.6.2>
	transfer-coding =
	                <transfer-coding, see [HTTP], Section 10.1.4>

The rules below are defined in [URI]:

	absolute-URI  = <absolute-URI, see [URI], Section 4.3>
	authority     = <authority, see [URI], Section 3.2>
	uri-host      = <host, see [URI], Section 3.2.2>
	port          = <port, see [URI], Section 3.2.3>
	query         = <query, see [URI], Section 3.4>
*/
var SP = " "
var CR = "\r"
var LF = "\n"
var CRLF = "\r\n"
var SLASH = "/"
var DOT = "."
var COLON = ":"
var SEMICOLON = ";"
var HTAB = "\t"
