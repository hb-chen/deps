package deps

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/hb-chen/deps/pkg/log"
	"github.com/hb-chen/deps/pkg/mod"
	req "github.com/imroc/req/v3"
	"golang.org/x/time/rate"
)

type PackageInfo struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Owners  []interface{} `json:"owners"`
	Version struct {
		Version                string        `json:"version"`
		SymbolicVersions       []interface{} `json:"symbolicVersions"`
		RefreshedAt            int           `json:"refreshedAt"`
		IsDefault              bool          `json:"isDefault"`
		License                string        `json:"license"`
		DependencyCount        int           `json:"dependencyCount"`
		DependentCount         int           `json:"dependentCount"`
		DependentCountDirect   int           `json:"dependentCountDirect"`
		DependentCountIndirect int           `json:"dependentCountIndirect"`
		Links                  struct {
			Origins []string `json:"origins"`
			Repo    string   `json:"repo"`
		} `json:"links"`
		Projects []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			ObservedAt  int    `json:"observedAt"`
			Issues      int    `json:"issues"`
			Forks       int    `json:"forks"`
			Stars       int    `json:"stars"`
			Description string `json:"description"`
			License     string `json:"license"`
			Link        string `json:"link"`
			ScorecardV2 struct {
				Date string `json:"date"`
				Repo struct {
					Name   string `json:"name"`
					Commit string `json:"commit"`
				} `json:"repo"`
				Scorecard struct {
					Version string `json:"version"`
					Commit  string `json:"commit"`
				} `json:"scorecard"`
				Check []struct {
					Name          string `json:"name"`
					Documentation struct {
						Short string `json:"short"`
						Url   string `json:"url"`
					} `json:"documentation"`
					Score   int      `json:"score"`
					Reason  string   `json:"reason"`
					Details []string `json:"details"`
				} `json:"check"`
				Metadata []interface{} `json:"metadata"`
				Score    float64       `json:"score"`
			} `json:"scorecardV2"`
		} `json:"projects"`
		Advisories      []*Advisory `json:"advisories"`
		RelatedPackages struct {
		} `json:"relatedPackages"`
	} `json:"version"`
	DefaultVersion string `json:"defaultVersion"`
}

type Advisory struct {
	Source         string   `json:"source"`
	SourceID       string   `json:"sourceID"`
	SourceURL      string   `json:"sourceURL"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	ReferenceURLs  []string `json:"referenceURLs"`
	Severity       string   `json:"severity"`
	GitHubSeverity string   `json:"gitHubSeverity"`
	Aliases        []string `json:"aliases"`
	DisclosedAt    int      `json:"disclosedAt"`
	ObservedAt     int      `json:"observedAt"`
}

type Dependency struct {
	Package    string
	License    string
	Direct     bool
	Advisories []*Advisory
}

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(rate.Every(100*time.Millisecond), 1)
}

func Deps() {
	graphs, err := mod.ModGraph()
	if err != nil {
		log.Logger.Error(err)
	} else {
		deps := make(map[string]*Dependency)
		for _, graph := range graphs {
			if _, ok := deps[graph.Requirement.Path]; ok {
				continue
			}

			limiter.Wait(context.TODO())
			if info, err := info(graph.Requirement.Path); err == nil {
				direct := true
				if graph.Module.Version != "" {
					direct = false
				}
				deps[graph.Requirement.Path] = &Dependency{
					Package:    graph.Requirement.Path,
					License:    info.Version.License,
					Direct:     direct,
					Advisories: info.Version.Advisories,
				}

				log.Logger.Infof("Pkg: %v, License: %v", info.Package.Name, info.Version.License)
			} else {
				log.Logger.Error(err)
			}
		}

		fmt.Println("Dependencies:")
		for _, dep := range deps {
			fmt.Printf("Pkg: %v, License: %v , Direct: %v ,Advisories: %v \n", dep.Package, dep.License, dep.Direct,
				len(dep.Advisories))
		}
	}
}

func info(pkg string) (*PackageInfo, error) {
	reqUrl := fmt.Sprintf("https://deps.dev/_/s/go/p/%s/v/", url.PathEscape(pkg))

	resp, err := req.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	info := PackageInfo{}
	if err := resp.UnmarshalJson(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
