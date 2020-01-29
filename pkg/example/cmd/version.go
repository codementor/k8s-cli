package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	//TODO (kensipe): refactor packages
	"github.com/codementor/k8s-cli/pkg/version"
)

var (
	versionExample = `  # Print the current installed Example package version
  kubectl example version`
)

// newVersionCmd returns a new initialized instance of the version sub command
func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the current Example package version.",
		Long:    `Print the current installed Example package version.`,
		Example: versionExample,
		RunE:    VersionCmd,
	}

	return versionCmd
}

// VersionCmd performs the version sub command
func VersionCmd(cmd *cobra.Command, args []string) error {
	exampleVersion := version.Get()
	fmt.Printf("Example Version: %s\n", fmt.Sprintf("%#v", exampleVersion))
	return nil
}
