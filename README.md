# Go Challenge for Codenation

### Context

[Codenation](https://www.codenation.dev) proposes learning oportunities to programmers based on market level programming challenges.
As part of the candidats selection for the **Full Stack Go + React** remote learning program.

### Instructions (resumed)

Propose an algorithm to decode a phrase encrypted by [Caesar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher) method.
The input data shall be read by an HTML request and sent by POST to a specific URL to each candidate.

The following rules must be respected :
- all uppercase letters should be converted to lower case
- numbers and ponctuation should be kept unencrypted

The coding language for the challenge is not specified.

### Chosen Aproach

First of all — as part of a Full Stack Go oportunity — I chose to tackle the problem with [golang](https://golang.org).
As my first time around with the language, some resources were helpfull [4] [5].

The language is quite straightforward and (with the help of the community), all parts of the solution were developed.

### Limitations

Although the total score for the solution was achieved, it still has some limitations.

* Bytes -> Runes
Text processing in Go seems to be based on Runes, an UTF-8 (int32) data type able to encode Unicode [6].
However, the solution assumes the text is a string is a set of bytes (uint8), so its not Rune-compatible.

* CLI
As an executable, the program does not support any change to its parameters, which are hardcoded.

### References

[1] : https://www.codenation.dev
[2] : https://en.wikipedia.org/wiki/Caesar_cipher
[3] : https://golang.org
[4] : [Effective Go](https://golang.org/doc/effective_go.html#names)
[5] : [Why should you learn Go?](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)
[6] : https://blog.golang.org/strings
