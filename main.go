package main

import (
	"fmt"

	awsw "github.com/ManabuSeki/awsw/cmd"
	prompt "github.com/c-bata/go-prompt"
)

var (
	version  string
	revision string
)

func main() {
	fmt.Printf("awsw %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")

	p := prompt.New(
		awsw.Executor,
		awsw.Completer,
		prompt.OptionTitle("awsw"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionPrefixTextColor(prompt.Green),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)
	p.Run()
	defer fmt.Println("Bye!")
}
