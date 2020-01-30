package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/codementor/k8s-cli/pkg/example/env"
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
	client := env.NewClientSet(&Settings)

	podsClient := client.CoreV1().Pods(apiv1.NamespaceDefault)

	pod := &apiv1.Pod{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: map[string]string{"app": "demo"},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:  name,
					Image: p.image,
				},
			},
		},
	}

	pp, err := podsClient.Create(pod)
	if err != nil {
		return err
	}

	fmt.Fprintf(p.out, "Pod %v created with rev: %v\n", pp.Name, pp.ResourceVersion)
	return nil
}
