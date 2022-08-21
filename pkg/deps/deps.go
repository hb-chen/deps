package deps

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/hb-chen/deps/pkg/graph/maven"
	"github.com/hb-chen/deps/pkg/graph/mod"
	"github.com/hb-chen/deps/pkg/log"
	"github.com/hb-chen/deps/pkg/output"
	"github.com/hb-chen/deps/pkg/output/stdout"
	"github.com/hb-chen/deps/pkg/output/template"
	"github.com/hb-chen/deps/pkg/scrape"
	"github.com/pkg/errors"
)

func filepathAbs(path string) string {
	if p, err := filepath.Abs(path); err != nil {
		log.Logger.Fatal(err)
		return ""
	} else {
		return p
	}
}

func Deps(system, project, tpl, out string, patterns []string) error {
	if strings.ToLower(system) == "auto" {
		// TODO 根据目录文件自动判断 go:mod.go maven:pom.xml
		return errors.New("system=auto unsupported")
	}
	log.Logger.Debugf("system type %v", project)

	tplFile := filepathAbs(tpl)
	outFile := filepathAbs(out)

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

			// skip scrape package
			skipScrape := false
			for _, sp := range patterns {
				matched, err := regexp.MatchString(sp, dep.Requirement.Name)
				if err != nil {
					log.Logger.Error(err)
				} else if matched {
					skipScrape = true
					break
				}
			}

			info := scrape.PackageInfo{}
			if !skipScrape {
				if err := scrape.Info(dep.Requirement, &info); err != nil {
					log.Logger.Errorw(err.Error(), "p", dep.Requirement.Name, "v", dep.Requirement.Version)
				}
			}

			deps[dep.Requirement.Name] = &output.Dependency{
				System:     dep.Requirement.System,
				Package:    dep.Requirement.Name,
				Version:    dep.Requirement.Version,
				Licenses:   info.Version.Licenses,
				Direct:     dep.Requirement.Direct,
				Advisories: info.Version.Advisories,
			}

			log.Logger.Debugf("Pkg: %v, Scraped:%v, Licenses: %v Advisories: %v", dep.Requirement.Name, !skipScrape,
				strings.Join(info.Version.Licenses, ","), len(info.Version.Advisories))
		}

		// 模板输出
		if template.InternalTpl(tpl) {
			// 使用 Template 内置的模板
			tplFile = tpl
		}
		out := template.NewOutput(
			template.WithTplFile(tplFile),
			template.WithOutFile(outFile),
		)
		if err := out.Generate(deps); err != nil {
			log.Logger.Fatal(err)
		}

		// 标准输出
		std := stdout.NewOutput()
		if err := std.Generate(deps); err != nil {
			log.Logger.Error(err)
		}
	}

	return nil
}
