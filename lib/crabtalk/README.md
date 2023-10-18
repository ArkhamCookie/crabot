# crabtalk

Translate English [^1] to crab.

## How it Works

### Convert to Morse Code

`crabot/crabtalk` first translates your message to [Morse code](https://wikipedia.org/wiki/Morse_code).
We don't worry about converting numbers, punctuation, symbols, or [prosigns](https://wikipedia.org/wiki/Prosigns_for_Morse_code).

*For reference, you can check Wikipedia's [Letters, numbers, punctuation, prosigns for Morse code and non-Latin variants section](https://wikipedia.org/wiki/Morse_code#Letters,_numbers,_punctuation,_prosigns_for_Morse_code_and_non-Latin_variants)*.

### Convert to Crab

Instead of dits (also known as dots ".") and dahs (also known as dashes "-"), we use "clicks" and "clacks".
Here is an example of crabtalk.

```text
"Hello World"
.... . .-.. .-.. --- / .-- --- .-. .-.. -..
clickclickclickclick click clickclackclickclick clickclackclickclick clackclackclack   clickclackclack clackclackclack clickclackclick clickclackclickclick clackclickclick
```

---

[^1]: More languages planned!
  See issue [#foo](https://github.com/ArkhamCookie/crabot/issues/) <!-- issue not created yet -->
  This means we will also translate [non-Latin extensions/alphabets](https://wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets).
