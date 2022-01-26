package cmd_test

import (
	"bytes"
	"errors"
	"github.com/intelliguy/cli/cmd"
	"strings"
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

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestRootCmd2(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: nil,
			err:  errors.New("not ok"),
		},
		{
			args: []string{"-t"},
			err:  nil,
			out:  "ok",
		},
		{
			args: []string{"--toggle"},
			err:  nil,
			out:  "ok",
		},
	}

	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.args...)

		is.Equal(tc.err, err)

		if tc.err == nil {
			is.Equal(tc.out, out)
		}
	}
}
