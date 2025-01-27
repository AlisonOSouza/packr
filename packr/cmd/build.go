package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/AlisonOSouza/packr"
	"github.com/AlisonOSouza/packr/builder"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:                "build",
	Short:              "Wraps the go build command with packr",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		defer builder.Clean(input)
		b := builder.New(context.Background(), input)
		err := b.Run()
		if err != nil {
			return err
		}

		cargs := []string{"build"}
		cargs = append(cargs, args...)
		cp := exec.Command(packr.GoBin(), cargs...)
		cp.Stderr = os.Stderr
		cp.Stdin = os.Stdin
		cp.Stdout = os.Stdout

		return cp.Run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
