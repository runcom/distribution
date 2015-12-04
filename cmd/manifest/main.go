package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "manifest IMAGE",
	Short: "manifest inspect a given image",
	Long:  "manifest inspect a given image.",
	Run: func(cmd *cobra.Command, args []string) {
		if format != "" {
			fmt.Println("with format")
		}
		fmt.Println("here I am" + strings.Join(args, " "))
	},
}

var format string

func init() {
	Cmd.Flags().StringVarP(&format, "format", "f", "", "format data returned")
}

func main() {
	Cmd.Execute()
}
