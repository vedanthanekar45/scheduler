package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Conformation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [Y/n]: ", prompt)
		response, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading input: \n%v\n", err)
			os.Exit(1)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "Yes" || response == "" {
			return true
		} else if response == "no" || response == "n" {
			return false
		}
	}
}
