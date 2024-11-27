// backend/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terraform-generator",
	Short: "A CLI for generating Terraform templates dynamically",
}

func Execute() error {
	return rootCmd.Execute()
}
