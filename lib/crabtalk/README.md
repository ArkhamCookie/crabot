# crabtalk

Translate English [^1] to crab.

## How it Works

### Convert to Morse Code

`crabot/crabtalk` first translates your message to [Morse code](https://wikipedia.org/wiki/Morse_code).
We don't worry about converting punctuation/symbols or [prosigns](https://wikipedia.org/wiki/Prosigns_for_Morse_code).
But instead of dits (also known as dashes "-") and dahs (also known as dots ".");

*For reference, you can check Wikipedia's [Letters, numbers, punctuation, prosigns for Morse code and non-Latin variants section](https://wikipedia.org/wiki/Morse_code#Letters,_numbers,_punctuation,_prosigns_for_Morse_code_and_non-Latin_variants)*.

### Convert to Crab

"click" = "."
"clack" = "-"

---

[^1]: More languages planned!
  See issue [#foo](https://github.com/ArkhamCookie/crabot/issues/) <!-- issue not created yet -->
  This means we will also translate [non-Latin extensions/alphabets](https://wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets).
