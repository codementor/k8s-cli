module github.com/codementor/k8s-cli

go 1.13

require (
	github.com/Masterminds/semver v1.5.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gosuri/uitable v0.0.4
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190923035154-9ee001bba392 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20191025023517-2077df36852e // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20191106092431-e228e37189d3 // indirect
	k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apiextensions-apiserver v0.0.0-20191016113550-5357c4baaf65
	k8s.io/apimachinery v0.0.0-20191028221656-72ed19daf4bb
	k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.18.0-alpha.1.0.20191220033320-6b257a9d6f46
	sigs.k8s.io/controller-tools v0.2.4
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
