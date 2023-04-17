// Package app is where the core logic for acopw lives.
package app

import (
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"git.sr.ht/~jamesponddotco/acopw-go/cmd/acopw/internal/build"
	"git.sr.ht/~jamesponddotco/acopw-go/cmd/acopw/internal/output"
	"github.com/spf13/cobra"
)

func Run() int {
	rootCmd := &cobra.Command{
		Use:               build.Name,
		Short:             "acopw is an easy-to-use and secure utility for generating random passwords, passphrases, and PINs.",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
		Version:           build.Version,
	}

	addDicewareCommand(rootCmd)
	addPinCommand(rootCmd)
	addRandomCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		output.Error(err)

		os.Exit(1)
	}

	return 0
}

func addDicewareCommand(rootCmd *cobra.Command) {
	dicewareCmd := &cobra.Command{
		Use:   "diceware",
		Short: "Generate a random diceware password.",
		Run: func(cmd *cobra.Command, args []string) {
			length, err := cmd.Flags().GetInt("length")
			if err != nil {
				output.Error("Error getting length:", err)

				return
			}

			separator, err := cmd.Flags().GetString("separator")
			if err != nil {
				output.Error("Error getting separator:", err)

				return
			}

			capitalized, err := cmd.Flags().GetBool("capitalized")
			if err != nil {
				output.Error("Error getting capitalized:", err)

				return
			}

			generator := &acopw.Diceware{
				Length:     length,
				Separator:  separator,
				Capitalize: capitalized,
			}

			password := generator.Generate()

			output.Password(password)
		},
	}

	dicewareCmd.Flags().IntP("length", "l", acopw.DefaultDicewareLength, "The number of words in the password.")
	dicewareCmd.Flags().StringP("separator", "s", " ", "The string used to separate words in the password.")
	dicewareCmd.Flags().BoolP("capitalized", "c", true, "Indicates whether a random word should be capitalized.")

	rootCmd.AddCommand(dicewareCmd)
}

func addPinCommand(rootCmd *cobra.Command) {
	pinCmd := &cobra.Command{
		Use:   "pin",
		Short: "Generate a random numeric PIN.",
		Run: func(cmd *cobra.Command, args []string) {
			length, err := cmd.Flags().GetInt("length")
			if err != nil {
				output.Error("Error getting length:", err)

				return
			}

			generator := &acopw.PIN{
				Length: length,
			}

			pin := generator.Generate()

			output.Password(pin)
		},
	}

	pinCmd.Flags().IntP("length", "l", acopw.DefaultPINLength, "The length of the generated PIN.")

	rootCmd.AddCommand(pinCmd)
}

func addRandomCommand(rootCmd *cobra.Command) {
	randomCmd := &cobra.Command{
		Use:   "random",
		Short: "Generate a random password.",
		Run: func(cmd *cobra.Command, args []string) {
			length, err := cmd.Flags().GetInt("length")
			if err != nil {
				output.Error("Error getting length:", err)

				return
			}

			numbers, err := cmd.Flags().GetBool("numbers")
			if err != nil {
				output.Error("Error getting numbers:", err)

				return
			}

			uppercase, err := cmd.Flags().GetBool("uppercase")
			if err != nil {
				output.Error("Error getting uppercase:", err)

				return
			}

			lowercase, err := cmd.Flags().GetBool("lowercase")
			if err != nil {
				output.Error("Error getting lowercase:", err)

				return
			}

			symbols, err := cmd.Flags().GetBool("symbols")
			if err != nil {
				output.Error("Error getting symbols:", err)

				return
			}

			excludeSet, err := cmd.Flags().GetString("exclude")
			if err != nil {
				output.Error("Error getting exclude:", err)

				return
			}

			generator := &acopw.Random{
				Length:          length,
				UseLower:        lowercase,
				UseUpper:        uppercase,
				UseNumbers:      numbers,
				UseSymbols:      symbols,
				ExcludedCharset: []string{excludeSet},
			}

			password := generator.Generate()

			output.Password(password)
		},
	}

	randomCmd.Flags().IntP("length", "l", 64, "The length of the generated password.")
	randomCmd.Flags().BoolP("numbers", "n", true, "Whether to include numbers in the generated password.")
	randomCmd.Flags().BoolP("uppercase", "U", true, "Whether to include uppercase letters in the generated password.")
	randomCmd.Flags().BoolP("lowercase", "L", true, "Whether to include lowercase letters in the generated password.")
	randomCmd.Flags().BoolP("symbols", "s", true, "Whether to include symbols in the generated password.")
	randomCmd.Flags().StringP("exclude", "e", "", "A string of characters to exclude from the generated password.")

	rootCmd.AddCommand(randomCmd)
}
