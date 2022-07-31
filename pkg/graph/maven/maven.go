package maven

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/hb-chen/deps/pkg/log"
	"github.com/pkg/errors"
)

type Graph struct {
	GraphName    string        `json:"graphName"`
	Artifacts    []*Artifact   `json:"artifacts"`
	Dependencies []*Dependency `json:"dependencies"`
}

type Artifact struct {
	Id         string   `json:"id"`
	NumericId  int      `json:"numericId"`
	GroupId    string   `json:"groupId"`
	ArtifactId string   `json:"artifactId"`
	Version    string   `json:"version"`
	Optional   bool     `json:"optional"`
	Scopes     []string `json:"scopes"`
	Types      []string `json:"types"`
}

type Dependency struct {
	From        string `json:"from"`
	To          string `json:"to"`
	NumericFrom int    `json:"numericFrom"`
	NumericTo   int    `json:"numericTo"`
	Resolution  string `json:"resolution"`
}

type Maven struct {
	opts *graph.Options
}

func (m *Maven) Graphs() ([]*graph.Dependency, error) {
	var err error
	err = os.Chdir(m.opts.ProjectPath)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("./mvnw", "depgraph:graph")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.Logger.Debug("run cmd:\n" + cmd.String())
	if err = cmd.Run(); err != nil {
		log.Logger.Error("cmd err:\n", err)
		log.Logger.Error("cmd stderr:\n", stderr.String())
		return nil, errors.Wrap(err, stderr.String())
	} else {
		log.Logger.Debug("cmd stdout:\n", stdout.String())
	}

	graphJsonPath := filepath.Join(m.opts.ProjectPath + "/target/dependency-graph.json")
	graphJson, err := ioutil.ReadFile(graphJsonPath)
	if err != nil {
		return nil, err
	}

	depGraph := Graph{}
	if err := json.Unmarshal(graphJson, &depGraph); err != nil {
		return nil, err
	}

	graphs := make([]*graph.Dependency, 0, len(depGraph.Artifacts))
	for _, art := range depGraph.Artifacts {
		// TODO direct
		pkg := graph.Dependency{
			Package: &graph.Package{
				System: graph.PkgSysMaven,
			},
			Requirement: &graph.Package{
				System:  graph.PkgSysMaven,
				Name:    art.GroupId + ":" + art.ArtifactId,
				Version: art.Version,
			},
		}

		graphs = append(graphs, &pkg)
	}

	return graphs, nil
}

func NewGraph(opts ...graph.Option) graph.Graph {
	o := graph.NewOptions(opts...)
	return &Maven{
		opts: o,
	}
}
