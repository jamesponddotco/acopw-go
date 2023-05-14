# acopw

[![Go Documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go?status.svg)](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go)
[![Go Report Card](https://goreportcard.com/badge/git.sr.ht/~jamesponddotco/acopw-go)](https://goreportcard.com/report/git.sr.ht/~jamesponddotco/acopw-go)
[![Coverage Report](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://git.sr.ht/~jamesponddotco/acopw-go/tree/trunk/item/cover.out)
[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/acopw-go.svg)](https://builds.sr.ht/~jamesponddotco/acopw-go?)

> **Note**: Neither this project nor the underlying cryptograhic
> implementation have been independently audited. This project is a work
> in progress.

Package `acopw` provides an easy-to-use, versatile and cryptographically
secure way to generate cryptographically secure random passwords,
passphrases, and PINs.

**Samples for what this package may generate:**

```console
(#lR?xdVe^o#;|{K>k%Y$,SXnn?nLl[=+|^cf|AWCtA}YoP(Vb=G^rwj]f;u@~Py
u{AQTrcOcHG#/.K>j{?P=\=jm%O>)hC;.Y%l,~fE'v];^@AY!?I}=DzyKlE@GEKb
728079
996388
hefty_spacetime_ENVELOPE_hearing_trend_fossils_unusable
deplored-desert-victory-runtime-coupland-costly-CLASSICS
```

The packages uses [crypto/rand](https://godocs.io/crypto/rand) and a list with **over 23 thousand** words for added randomness.

## Installation

To install `acopw`, run:

```console
go get git.sr.ht/~jamesponddotco/acopw-go
```

You can also [install `acopw` the command-line
application](https://git.sr.ht/~jamesponddotco/acopw-go/tree/trunk/item/cmd/acopw/README.md).

## Usage

### Random passwords

To generate a random password, use the `Random` struct and call the `Generate()` method.

```go
package main

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func main() {
	random := &acopw.Random{
		Length: 16,
		UseLower: true,
		UseUpper: true,
		UseNumbers: true,
		UseSymbols: true,
	}

	password := random.Generate()

	log.Println(password)
}
```

### Diceware passwords

To generate a diceware password, use the `Diceware` struct and call the `Generate()` method.

```go
package main

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func main() {
	diceware := &acopw.Diceware{
		Separator: "-",
		Length: 6,
		Capitalize: true,
	}

	password,err := diceware.Generate()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(password)
}
```

### PINs

To generate a PIN, use the `PIN` struct and call the `Generate()` method.

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

	password := pin.Generate()

	log.Println(password)
}
```

## Contributing

Anyone can help make `acopw` better. Check out [the contribution
guidelines](https://git.sr.ht/~jamesponddotco/acopw-go/tree/master/item/CONTRIBUTING.md)
for more information.

## Credits

- The algorithm used is based on something [András Belicza](https://github.com/icza) wrote on [Stack Overflow](https://stackoverflow.com/a/31832326), so credits goes to them.
- Tests were mostly written using GPT-4.
- Big thanks to the EFF for providing [some word lists](https://www.eff.org/dice), which were complimented by me crawling Wikipedia.

## Resources

The following resources are available:

- [Package documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go).
- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/acopw-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/acopw-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/acopw).

---

Released under the [MIT License](LICENSE.md).
