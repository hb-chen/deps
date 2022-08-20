package scrape

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hb-chen/deps/pkg/graph"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(rate.Every(100*time.Millisecond), 1)
}

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
		Licenses               []string      `json:"licenses"`
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

func Info(p *graph.Package, info *PackageInfo) error {
	if err := limiter.Wait(context.TODO()); err != nil {
		return err
	}

	reqUrl := fmt.Sprintf("https://deps.dev/_/s/%s/p/%s/v/%s", p.System, url.PathEscape(p.Name), p.Version)

	resp, err := req.Get(reqUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		if err := resp.UnmarshalJson(info); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("get package info status: " + resp.Status)
	}
}
