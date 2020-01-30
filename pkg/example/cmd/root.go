package cmd

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/codementor/k8s-cli/pkg/example/clog"
	"github.com/codementor/k8s-cli/pkg/example/env"
	"github.com/codementor/k8s-cli/pkg/version"
)

var (
	// Settings defines global flags and settings
	Settings env.Settings
)

// NewExampleCmd creates a new root command for example CLI
func NewExampleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubectl-example",
		Short: "CLI to for kubectl plugin examples.",
		Long: `Example CLI for working with kubectl CLI plugin.
`,
		SilenceUsage: true,
		Example: `  # List pods
  kubectl example pod list

  # add pod 
  kubectl example pod add foo

  # View Example version
  kubectl example version
`,
		Version: version.Get().GitVersion,
	}

	// create pod
	// get pods
	// get api-resources
	// get CRD?
	cmd.AddCommand(newPodCmd(cmd.OutOrStdout()))
	cmd.AddCommand(newResourcesCmd())
	cmd.AddCommand(newVersionCmd())

	initGlobalFlags(cmd, cmd.OutOrStdout())

	return cmd
}

func initGlobalFlags(cmd *cobra.Command, out io.Writer) {
	flags := cmd.PersistentFlags()
	Settings.AddFlags(flags)
	clog.InitWithFlags(flags, out)
}
