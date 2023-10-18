package crabtalk

import (
	"github.com/gSpera/morse"
)

// Morse code map using default for letters
// but not converting anything else
var crabMorse = morse.EncodingMap{
	// Predefined Letters
	'A': morse.A,
	'B': morse.B,
	'C': morse.C,
	'D': morse.D,
	'E': morse.E,
	'F': morse.F,
	'G': morse.G,
	'H': morse.H,
	'I': morse.I,
	'J': morse.J,
	'K': morse.K,
	'L': morse.L,
	'M': morse.M,
	'N': morse.N,
	'O': morse.O,
	'P': morse.P,
	'Q': morse.Q,
	'R': morse.R,
	'S': morse.S,
	'T': morse.T,
	'U': morse.U,
	'V': morse.V,
	'W': morse.W,
	'X': morse.X,
	'Y': morse.Y,
	'Z': morse.Z,

	// Predefined Numbers
	'1': "1",
	'2': "2",
	'3': "3",
	'4': "4",
	'5': "5",
	'6': "6",
	'7': "7",
	'8': "8",
	'9': "9",
	'0': "0",

	// Predefined Punctuation
	'.':  ".",
	',':  ",",
	':':  ":",
	'?':  "?",
	'\'': "\\",
	'-':  "-",
	'/':  "/",
	'(':  "(",
	')':  ")",
	'“':  "“",
	'=':  "=",
	'+':  "+",
	'@':  "@",

	// Misc. Predefined
	' ':  morse.Space,

	// Undefined Punctuation
	'!': "!",
}
