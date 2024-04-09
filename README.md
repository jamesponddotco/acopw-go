# acopw

[![Go Documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go?status.svg)](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go)
[![Go Report Card](https://goreportcard.com/badge/git.sr.ht/~jamesponddotco/acopw-go)](https://goreportcard.com/report/git.sr.ht/~jamesponddotco/acopw-go)
[![Coverage Report](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://git.sr.ht/~jamesponddotco/acopw-go/tree/trunk/item/cover.out)
[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/acopw-go.svg)](https://builds.sr.ht/~jamesponddotco/acopw-go?)

Package `acopw` provides a simple, efficient, and secure way to generate
random passwords, passphrases, and PINs using Go. It leverages the speed
of `math/rand/v2` with the cryptographic security of `ChaCha8` for
generating random data, ensuring the highest level of randomness,
security, and performance.

When generating diceware passwords, it uses a [curated list with **over
23 thousand
words**](https://git.sr.ht/~jamesponddotco/acopw-go/blob/trunk/words/word-list.txt),
one of the largest word lists out there.

**Sample output:**

```console
(#lR?xdVe^o#;|{K>k%Y$,SXnn?nLl[=+|^cf|AWCtA}YoP(Vb=G^rwj]f;u@~Py
u{AQTrcOcHG#/.K>j{?P=\=jm%O>)hC;.Y%l,~fE'v];^@AY!?I}=DzyKlE@GEKb
hefty_spacetime_ENVELOPE_hearing_trend_fossils_unusable
deplored-desert-victory-runtime-coupland-costly-CLASSICS
728079
996388
```

## Installation

To install `acopw` and use it in your project, run:

```console
go get git.sr.ht/~jamesponddotco/acopw-go@latest
```

## Usage

### Random passwords

To generate a random password, use `Random` and call the `Generate()` method.

```go
package main

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func main() {
	random := &acopw.Random{
		Length:     16,
		UseLower:   true,
		UseUpper:   true,
		UseNumbers: true,
		UseSymbols: true,
	}

	log.Println(random.Generate())
}
```

### Diceware passwords

To generate a diceware password, use `Diceware` and call the `Generate()` method.

```go
package main

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func main() {
	diceware := &acopw.Diceware{
		Separator:  "-",
		Length:     6,
		Capitalize: true,
	}

	log.Println(diceware.Generate())
}
```

### PINs

To generate a PIN, use `PIN` and call the `Generate()` method.

```go
package main

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func main() {
	pin := &acopw.PIN{
		Length: 6,
	}

	log.Println(pin.Generate())
}
```

## Contributing

Anyone can help make `acopw` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/acopw-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/acopw).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the MIT License](LICENSE.md).

## Acknowledgements

- Tests were mostly written using a combination of Claude 3 and GPT-4.
- Big thanks to the EFF for providing [some word
  lists](https://www.eff.org/dice), which were complimented by me
  [crawling Wikipedia](https://sr.ht/~jamesponddotco/wikiextract/).
- Big thanks to [Christopher Wellons](https://nullprogram.com/) for
  reviewing and auditing the underlying cryptographic implementation for
  biases and security issues. He also helped with many of the
  performance optimizations for the `v1.0.0` release.

## Resources

The following resources are available:

- [Package documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go).
- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/acopw-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/acopw-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/acopw).

---

Released under the [MIT License](LICENSE.md).
