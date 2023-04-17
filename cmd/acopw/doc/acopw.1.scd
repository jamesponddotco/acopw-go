acopw(1)

# NAME

acopw - A cryptographically secure password generator

# SYNOPSIS

*acopw* [commands...] [options...]

# DESCRIPTION

*acopw* is an easy-to-use, versatile and secure command-line utility for
generating random passwords, passphrases, and PINs.

# OPTIONS

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

# COMMANDS

*help* <command>
	Help about any command.

*diceware* <options>
	Generate a random diceware password.

	Options are:

	*-l*, *--length*
		Set the number of words in the generated passphrase. Defaults to 8.

	*-s*, *--separator*
		Set the separator between words in the generated passphrase. Defaults to
		space.

	*-c*, *--capitalized*
		Capitalize a random word in the generated passphrase. Defaults to true.

*pin* <options>
	Generate a random numeric PIN.

	Options are:

	*-l*, *--length*
		Set the length of the generated PIN. Defaults to 6.

*random* <options>
	Generate a random password with optional character set restrictions.

	Options are:

	*-l*, *--length*
		Set the length of the generated password. Defaults to 64.

	*-n*, *--numbers*
		Include numbers in the generated password. Defaults to true.

	*-U*, *--uppercase*
		Include uppercase letters in the generated password. Defaults to true.

	*-L*, *--lowercase*
		Include lowercase letters in the generated password. Defaults to true.

	*-s*, *--symbols*
		Include symbols in the generated password. Defaults to true.

	*-e*, *--exclude*
		Exclude specific characters from the generated password.

# EXAMPLES

Generate a random password of length 16:

	acopw random -l 16

Generate a Diceware passphrase with 6 words:

	acopw diceware -l 6

Generate a 6-digit PIN:

	acopw pin -l 6

# AUTHORS

Maintained by James Pond <james@cipher.host>.
