package deps

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/hb-chen/deps/pkg/graph/maven"
	"github.com/hb-chen/deps/pkg/graph/mod"
	"github.com/hb-chen/deps/pkg/log"
	"github.com/pkg/errors"
)

type Dependency struct {
	Package    string
	Licenses   []string
	Direct     bool
	Advisories []*Advisory
}

func Deps(system, project string) error {
	if strings.ToLower(system) == "auto" {
		// TODO 根据目录文件自动判断 go:mod.go maven:pom.xml
		return errors.New("system=auto unsupported")
	}
	log.Logger.Debugf("system type %v", project)

	if project == "" {
		project = "./"
	}
	if path, err := filepath.Abs(project); err == nil {
		project = path
	} else {
		return err
	}
	log.Logger.Debugf("project path %v", project)

	opts := []graph.Option{
		graph.WithProjectPath(project),
	}

	var g graph.Graph
	switch strings.ToLower(system) {
	case "maven":
		g = maven.NewGraph(opts...)
	case "mod":
		g = mod.NewGraph(opts...)
	// more: gradle,dep,glide...
	default:
		return errors.New("system=" + system + " unsupported")
	}

	graphs, err := g.Graphs()
	if err != nil {
		log.Logger.Error(err)
	} else {
		deps := make(map[string]*Dependency)
		for _, graph := range graphs {
			if _, ok := deps[graph.Requirement.Name]; ok {
				continue
			}

			if err := limiter.Wait(context.TODO()); err != nil {
				log.Logger.Error(err)
				continue
			}

			if info, err := info(graph.Requirement); err == nil {

				deps[graph.Requirement.Name] = &Dependency{
					Package:    graph.Requirement.Name,
					Licenses:   info.Version.Licenses,
					Direct:     graph.Requirement.Direct,
					Advisories: info.Version.Advisories,
				}

				log.Logger.Infof("Pkg: %v, Licenses: %v", info.Package.Name, strings.Join(info.Version.Licenses, ","))
			} else {
				log.Logger.Errorw(err.Error(), "p", graph.Requirement.Name, "v", graph.Requirement.Version)
			}
		}

		// TODO output
		fmt.Println("Dependencies:")
		for _, dep := range deps {
			fmt.Printf("Pkg: %v, Licenses: %v , Direct: %v ,Advisories: %v \n", dep.Package,
				strings.Join(dep.Licenses, ","), dep.Direct,
				len(dep.Advisories))
		}
	}

	return nil
}
