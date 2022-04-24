package mod

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/hb-chen/deps/pkg/log"
	"github.com/pkg/errors"
)

type Graph struct {
	Module      *Module
	Requirement *Module
}

type Module struct {
	Path    string
	Version string
}

func ModGraph() ([]*Graph, error) {
	var err error
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
	graphs := make([]*Graph, 0, len(out))
	for _, line := range out {
		log.Logger.Debug("\n", line)
		if len(line) > 0 {
			graph := graphParse(line)
			graphs = append(graphs, graph)
		}
	}

	return graphs, nil
}

func graphParse(str string) *Graph {
	split := strings.Split(str, " ")
	module := strings.Split(split[0], "@")
	requirement := strings.Split(split[1], "@")

	if len(module) == 1 {
		module = append(module, "")
	}

	return &Graph{
		Module:      &Module{Path: module[0], Version: module[1]},
		Requirement: &Module{Path: requirement[0], Version: requirement[1]},
	}
}
