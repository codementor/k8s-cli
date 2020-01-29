package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

const (
	podListDesc = `List pods.
`
	podListExample = `  kubectl example pod list
`
)

type podListCmd struct {
	interactive bool
	out         io.Writer
}

// newPodListCmd lists pods
func newPodListCmd(out io.Writer) *cobra.Command {

	pkg := &podListCmd{out: out}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   podListDesc,
		Example: podListExample,
		RunE: func(cmd *cobra.Command, args []string) error {
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


// run returns the errors associated with cmd env
func (pkg *podListCmd) run(args []string) error {
	return nil
}
