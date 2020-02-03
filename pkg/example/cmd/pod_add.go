package cmd

import (
	"errors"
	"fmt"
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
	image string
	out   io.Writer
}

// newPodAddCmd adds a pod to the cluster
func newPodAddCmd(out io.Writer) *cobra.Command {

	p := &podAddCmd{out: out}
	cmd := &cobra.Command{
		Use:     "add",
		Short:   "adds a pod to the cluster",
		Long:    podAddDesc,
		Example: podAddExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateArgs(args); err != nil {
				return err
			}
			if err := p.run(args[0]); err != nil {
				return err
			}
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVar(&p.image, "image", "nginx", "image to be used in creation of pod")
	return cmd
}

func validateArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("expecting 1 argument - name of pod")
	}
	return nil

}

// run returns the errors associated with cmd env
func (p *podAddCmd) run(name string) error {

	fmt.Printf("adding a pod\n")
	return nil
}
