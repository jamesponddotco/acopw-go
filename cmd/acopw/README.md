# acopw

[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/acopw-go.svg)](https://builds.sr.ht/~jamesponddotco/acopw-go?)

**acopw** is an easy-to-use, versatile and cryptographically secure
command-line utility for generating random passwords, passphrases, and
PINs.

## Installation

### From source

First install the dependencies:

- Go 1.20 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```console
$ acopw --help
Generate cryptographically secure random passwords, passphrases, and PINs.

Usage:
  acopw [command]

Available Commands:
  diceware    Generate a random diceware password.
  help        Help about any command
  pin         Generate a random numeric PIN.
  random      Generate a random password.

Flags:
  -h, --help      help for acopw
  -v, --version   version for acopw

Use "acopw [command] --help" for more information about a command.

$ acopw random --length 32
OY5UyaBCt1Vrm5ukPrpsY0SvPBOxVD!X

$ acopw diceware
beneficial massive bored corners rumored INSTITUTIONS labourers julie
```

See _acopw(1)_ after installing for more information.

## Contributing

Anyone can help make `acopw` better. Check out [the contribution
guidelines](https://git.sr.ht/~jamesponddotco/acopw-go/tree/master/item/CONTRIBUTING.md)
for more information.

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/acopw-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/acopw-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/acopw).

---

Released under the [MIT License](LICENSE.md).
