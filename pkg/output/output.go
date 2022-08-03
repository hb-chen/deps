package output

import (
	"fmt"
	"net/url"

	"github.com/hb-chen/deps/pkg/scrape"
)

type Dependency struct {
	System     string
	Package    string
	Version    string
	Licenses   []string
	Direct     bool
	Advisories []*scrape.Advisory
}

func (d *Dependency) AdvisoriesURL() string {
	advUrl := fmt.Sprintf("https://deps.dev/%s/%s/%s", d.System, url.PathEscape(d.Package), d.Version)
	return advUrl
}

type Output interface {
	Generate(inputs map[string]*Dependency, tplPath, outFile string) error
}
