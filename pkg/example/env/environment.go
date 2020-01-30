package env

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	v1 "k8s.io/api/core/v1"
)

var DefaultKubeConfig = filepath.Join(homedir.HomeDir(), "/.kube/config")

func KubeConfigHome() string {
	if val, ok := os.LookupEnv("KUBECONFIG"); ok {
		return val
	}
	return DefaultKubeConfig
}

// Settings defines global variables and settings
type Settings struct {
	// KubeConfig is the path to an explicit kubeconfig file. This overwrites the value in $KUBECONFIG
	KubeConfig string
	// Namespace used when working with Kubernetes
	Namespace string
	// RequestTimeout is the timeout value (in seconds) when making API calls
	RequestTimeout int64
}

// DefaultSettings initializes the settings to its defaults
var DefaultSettings = &Settings{
	Namespace:      "default",
	RequestTimeout: 0,
}

// AddFlags binds flags to the given flagset.
func (s *Settings) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.KubeConfig, "kubeconfig", KubeConfigHome(), "Path to your Kubernetes configuration file.")
	fs.StringVarP(&s.Namespace, "namespace", "n", "default", "Target namespace for the object.")
}

// OverrideDefault used for deviations from global defaults
func (s *Settings) OverrideDefault(fs *pflag.FlagSet, name, value string) string {
	if fs.Changed(name) {
		return s.Namespace
	}

	return value
}

// NewClientSet is a helper function that takes the Settings struct and returns a new kube Client
func NewClientSet(s *Settings) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", s.KubeConfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func NewRestClient(s *Settings) *rest.RESTClient {
	config, err := clientcmd.BuildConfigFromFlags("", s.KubeConfig)
	if err != nil {
		panic(err)
	}
	// defaulting configuration
	config.GroupVersion = &v1.SchemeGroupVersion
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	return client
}
