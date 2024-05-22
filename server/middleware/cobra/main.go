package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "",
		Short: "简短描述",
		Long:  "详细描述",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	rootCmd.Execute()
}
