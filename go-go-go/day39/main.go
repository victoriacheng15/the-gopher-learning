package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var reverse bool

func reverseLines(lines string) string {
	runes := []rune(lines)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func runCmd(cmd *cobra.Command, args []string) {
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if reverse {
			line = reverseLines(line)
		}
		fmt.Println(line)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "readfile",
		Short: "Read a file and print its contents",
		Args:  cobra.ExactArgs(1),
		Run:   runCmd,
	}

	rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Reverse each line of the file")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
