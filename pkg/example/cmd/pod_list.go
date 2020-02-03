package cmd

import (
	"fmt"
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
	// status boolean
	out    io.Writer
}

// newPodListCmd lists pods
func newPodListCmd(out io.Writer) *cobra.Command {

	pkg := &podListCmd{out: out}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   podListDesc,
		Example: podListExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := pkg.run(); err != nil {
				return err
			}
			return nil
		},
	}

	// status flag

	return cmd
}

// newPodList2Cmd lists pods
func newPodList2Cmd(out io.Writer) *cobra.Command {

	pkg := &podListCmd{out: out}
	cmd := &cobra.Command{
		Use:     "list2",
		Short:   podListDesc,
		Example: podListExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := pkg.run2(); err != nil {
				return err
			}
			return nil
		},
	}

	// status flag

	return cmd
}

// run 1st approach at list pods
func (p *podListCmd) run() error {

	fmt.Printf("add pod list code using direct object references\n")
	return nil
}

// run2 2nd approach at list pods
func (p *podListCmd) run2() error {

	//REST Client approach

	fmt.Printf("add pod list code using the rest client\n")
	return nil
}
