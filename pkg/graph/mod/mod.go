package mod

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/hb-chen/deps/pkg/log"
	"github.com/pkg/errors"
)

type Mod struct {
	opts *graph.Options
}

func (m *Mod) Graphs() ([]*graph.Dependency, error) {
	err := os.Chdir(m.opts.ProjectPath)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "mod", "graph")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.Logger.Debug("\n" + cmd.String())
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, stderr.String())
	} else {
		log.Logger.Debug("\n", stdout.String())
	}

	out := strings.Split(stdout.String(), "\n")
	graphs := make([]*graph.Dependency, 0, len(out))
	for _, line := range out {
		log.Logger.Debug("\n", line)
		if len(line) > 0 {
			dep := dependencyParse(line)
			graphs = append(graphs, dep)
		}
	}

	return graphs, nil
}

func dependencyParse(str string) *graph.Dependency {
	split := strings.Split(str, " ")
	module := strings.Split(split[0], "@")
	requirement := strings.Split(split[1], "@")

	if len(module) == 1 {
		module = append(module, "")
	}

	// true: github.com/hb-chen/deps github.com/hashicorp/errwrap@v1.0.0
	// false: github.com/spf13/cobra@v1.4.0 gopkg.in/yaml.v2@v2.4.0
	direct := false
	if module[1] == "" {
		direct = true
	}

	return &graph.Dependency{
		Package: &graph.Package{
			System:  graph.PkgSysGo,
			Name:    module[0],
			Version: module[1],
		},
		Requirement: &graph.Package{
			System:  graph.PkgSysGo,
			Name:    requirement[0],
			Version: requirement[1],
			Direct:  direct,
		},
	}
}

func NewGraph(opts ...graph.Option) graph.Graph {
	o := graph.NewOptions(opts...)
	return &Mod{
		opts: o,
	}
}
