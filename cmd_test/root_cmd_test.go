package cmd_test

import (
	"github.com/intelliguy/cli/cmd"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	is := is.New(t)

	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)

	err := root.Execute()

	is.NoErr(err)
}
