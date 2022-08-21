package stdout

import (
	"fmt"
	"strings"

	"github.com/hb-chen/deps/pkg/output"
)

type stdout struct{}

func (*stdout) Generate(deps map[string]*output.Dependency) error {
	fmt.Println("Dependencies:")
	for _, dep := range deps {
		fmt.Printf("Pkg: %v, Licenses: %v , Direct: %v ,Advisories: %v \n", dep.Package,
			strings.Join(dep.Licenses, ","), dep.Direct,
			len(dep.Advisories))
	}

	return nil
}

func NewOutput() output.Output {
	return &stdout{}
}
