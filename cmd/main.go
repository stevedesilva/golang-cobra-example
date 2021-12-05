package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	localRootFlag   bool
	times           int
	rootCmd         = &cobra.Command{
		Use:   "example usage",
		Short: "An example cobra program",
		Long: `This is a simple example of a cobra program.,
It will have several subcommands and flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running the root command")
		},
	}

	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	printCmd = &cobra.Command{
		Use:   "print [strings to echo]",
		Short: "prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Printing: " + strings.Join(args, " "))
		},
	}
	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given strings to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return fmt.Errorf("times cannot be 0")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
			return nil
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent  root flag")
	rootCmd.LocalFlags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(printCmd)
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout ")
	timesCmd.MarkFlagRequired("times")
	echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
