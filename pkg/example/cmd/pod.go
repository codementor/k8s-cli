package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

const podDesc = `
This command consists of multiple sub-commands to interact with Pods.
`

const podExamples = `  kubectl example pod list [flags]
  kubectl example pod list2 [flags]
  kubectl example pod add [name] [flags]
`

// newPodCmd for working with pods
func newPodCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pod [SUBCOMMAND] [FLAGS] [ARGS]",
		Short:   "example of working with pods",
		Long:    podDesc,
		Example: podExamples,
	}

	cmd.AddCommand(newPodAddCmd(out))
	cmd.AddCommand(newPodListCmd(out))
	cmd.AddCommand(newPodList2Cmd(out))

	return cmd
}
