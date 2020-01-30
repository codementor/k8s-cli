package cmd

import (
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/codementor/k8s-cli/pkg/example/env"
)

var (
	resourceExample = `  # Print the current api resources
  kubectl example resources`
)

// newResourcesCmd returns list of kubernetes api-resources
func newResourcesCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "resources",
		Short:   "List Kubernetes API Resources",
		Example: resourceExample,
		RunE:    ResourcesCmd,
	}

	return versionCmd
}

// ResourcesCmd list api resources
func ResourcesCmd(cmd *cobra.Command, args []string) error {
	client := env.NewClientSet(&Settings)
	lists, err := client.Discovery().ServerPreferredResources()
	if err != nil {
		return err
	}

	resources := []metav1.APIResource{}
	for _, item := range lists {
		if len(item.APIResources) == 0 {
			continue
		}
		if err != nil {
			continue
		}

		resources = append(resources, item.APIResources...)
	}
	table := uitable.New()
	table.AddRow("Name", "Namespaced", "Kind")
	for _, resource := range resources {
		table.AddRow(resource.Name, resource.Namespaced, resource.Kind)
	}
	fmt.Print(table)
	return nil
}
