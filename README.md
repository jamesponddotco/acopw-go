# acopw

[![Go Documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go?status.svg)](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go)
[![Go Report Card](https://goreportcard.com/badge/git.sr.ht/~jamesponddotco/acopw-go)](https://goreportcard.com/report/git.sr.ht/~jamesponddotco/acopw-go)
[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/acopw-go.svg)](https://builds.sr.ht/~jamesponddotco/acopw-go?)

> **Note**: This library has not been reviewed by a security expert yet.

Package `acopw`—accio password, get it?—provides a simple way to generate cryptographically secure random and diceware passwords, and PINs.

## Features

- Generate passwords with customizable character sets.
- Generate secure PINs of any length.
- Simple and easy-to-use API.
- Cryptographically secure random number generation using
  [crypto/rand](https://godocs.io/crypto/rand)

## Installation

To install `acopw`, run:

```sh
go get git.sr.ht/~jamesponddotco/acopw-go
```

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

	password := diceware.Generate()

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
