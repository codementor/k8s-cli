package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

const (
	podAddDesc = `Adds a pod to a cluster.
`
	podAddExample = `  kubectl example pod add foo --image=nginx
`
)

type podAddCmd struct {
	interactive bool
	out         io.Writer
}

// newPodAddCmd adds a pod to the cluster
func newPodAddCmd(out io.Writer) *cobra.Command {

	pkg := &podAddCmd{out: out}
	cmd := &cobra.Command{
		Use:     "add",
		Short:   "adds a pod to the cluster",
		Long:    podAddDesc,
		Example: podAddExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateArgs(args); err != nil {
				return err
			}
			if err := pkg.run(args); err != nil {
				return err
			}
			return nil
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&pkg.interactive, "interactive", "i", false, "Interactive mode.")
	return cmd
}

func validateArgs(args []string) error {
	if len(args) == 1 {
		return errors.New("expecting 1 argument - name of pod`")
	}
	return nil
}

// run returns the errors associated with cmd env
func (pkg *podAddCmd) run(args []string) error {
	return nil
}
