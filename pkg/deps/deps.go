package deps

import (
	"path/filepath"
	"strings"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/hb-chen/deps/pkg/graph/maven"
	"github.com/hb-chen/deps/pkg/graph/mod"
	"github.com/hb-chen/deps/pkg/log"
	"github.com/hb-chen/deps/pkg/output"
	"github.com/hb-chen/deps/pkg/output/template"
	"github.com/hb-chen/deps/pkg/scrape"
	"github.com/pkg/errors"
)

func filepathAbs(path string) string {
	if path, err := filepath.Abs(path); err != nil {
		log.Logger.Fatal(err)
		return ""
	} else {
		return path
	}

}

func Deps(system, project string) error {
	if strings.ToLower(system) == "auto" {
		// TODO 根据目录文件自动判断 go:mod.go maven:pom.xml
		return errors.New("system=auto unsupported")
	}
	log.Logger.Debugf("system type %v", project)

	// TODO 完善模板和输出文件配置，以及相对目录问题
	tplPath := filepathAbs("./template/md.tpl")
	outFile := filepathAbs("./out/deps.md")

	if project == "" {
		project = "./"
	}
	project = filepathAbs(project)
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
		deps := make(map[string]*output.Dependency)
		for _, dep := range graphs {
			if _, ok := deps[dep.Requirement.Name]; ok {
				continue
			}

			if info, err := scrape.Info(dep.Requirement); err == nil {
				deps[dep.Requirement.Name] = &output.Dependency{
					System:     dep.Requirement.System,
					Package:    dep.Requirement.Name,
					Version:    dep.Requirement.Version,
					Licenses:   info.Version.Licenses,
					Direct:     dep.Requirement.Direct,
					Advisories: info.Version.Advisories,
				}

				log.Logger.Infof("Pkg: %v, Licenses: %v", info.Package.Name, strings.Join(info.Version.Licenses, ","))
			} else {
				log.Logger.Errorw(err.Error(), "p", dep.Requirement.Name, "v", dep.Requirement.Version)
			}
		}

		// TODO 输出选项及自定义配置
		out := template.NewOutput()
		if err := out.Generate(deps, tplPath, outFile); err != nil {
			log.Logger.Error(err)
		}
	}

	return nil
}
