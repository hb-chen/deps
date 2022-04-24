# Deps

- Go module graph
  - `TODO` Support more language: `maven` `npm`
- Dependencies license, advisories...
  - [deps.dev](https://deps.dev/)

## Usage
```shell
# Install
go install github.com/hb-chen/deps@latest

# Run
cd {Golang Project Path}
deps

# Output
Dependencies:
Pkg: golang.org/x/time, License: BSD-3-Clause , Direct: true ,Advisories: 0 
Pkg: github.com/benbjohnson/clock, License: MIT , Direct: false ,Advisories: 0 
Pkg: github.com/russross/blackfriday/v2, License: BSD-2-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/hashicorp/errwrap, License: MPL-2.0 , Direct: true ,Advisories: 0 
Pkg: github.com/spf13/pflag, License: BSD-3-Clause , Direct: true ,Advisories: 0 
Pkg: go.uber.org/zap, License: MIT , Direct: true ,Advisories: 0 
Pkg: github.com/stretchr/testify, License: MIT , Direct: false ,Advisories: 0 
Pkg: github.com/pmezard/go-difflib, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/kr/text, License: MIT , Direct: false ,Advisories: 0 
Pkg: github.com/inconshreveable/mousetrap, License: Apache-2.0 , Direct: true ,Advisories: 0 
Pkg: golang.org/x/term, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: golang.org/x/tools, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/stretchr/objx, License: MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/lint, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/davecgh/go-spew, License: ISC , Direct: false ,Advisories: 0 
Pkg: github.com/kr/pretty, License: MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/mod, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: golang.org/x/sync, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/imroc/req/v3, License: MIT , Direct: true ,Advisories: 0 
Pkg: github.com/pkg/errors, License: BSD-2-Clause , Direct: true ,Advisories: 0 
Pkg: go.uber.org/multierr, License: MIT , Direct: true ,Advisories: 0 
Pkg: golang.org/x/net, License: BSD-3-Clause , Direct: true ,Advisories: 0 
Pkg: golang.org/x/text, License: BSD-3-Clause , Direct: true ,Advisories: 0 
Pkg: gopkg.in/yaml.v2, License: Apache-2.0 AND MIT , Direct: false ,Advisories: 0 
Pkg: github.com/yuin/goldmark, License: MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/crypto, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/hashicorp/go-multierror, License: MPL-2.0 , Direct: true ,Advisories: 0 
Pkg: go.uber.org/atomic, License: MIT , Direct: true ,Advisories: 0 
Pkg: github.com/cpuguy83/go-md2man/v2, License: MIT , Direct: false ,Advisories: 0 
Pkg: gopkg.in/yaml.v3, License: Apache-2.0 AND MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/xerrors, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: go.uber.org/goleak, License: MIT , Direct: false ,Advisories: 0 
Pkg: github.com/spf13/cobra, License: Apache-2.0 , Direct: true ,Advisories: 0 
Pkg: golang.org/x/sys, License: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: gopkg.in/check.v1, License: BSD-2-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/kr/pty, License: MIT , Direct: false ,Advisories: 0 
```

## deps.dev API

### Package

https://deps.dev/_/s/go/p/github.com%2Fhb-go%2Fpkg/v/

```json
{
  "package": {
    "system": "GO",
    "name": "github.com/hb-go/pkg"
  },
  "owners": [],
  "version": {
    "version": "v0.0.2",
    "symbolicVersions": [],
    "refreshedAt": 1637377511,
    "isDefault": true,
    "license": "MIT",
    "dependencyCount": 0,
    "dependentCount": 3,
    "dependentCountDirect": 3,
    "dependentCountIndirect": 0,
    "links": {
      "origins": [
        "https://pkg.go.dev/mod/github.com/hb-go/pkg@v0.0.2"
      ],
      "repo": "https://github.com/hb-go/pkg"
    },
    "projects": [
      {
        "type": "GITHUB",
        "name": "hb-go/pkg",
        "observedAt": 1637318255,
        "issues": 2,
        "forks": 1,
        "stars": 4,
        "description": "Go common package. conv、dispatcher、log、pool、rate…",
        "license": "MIT",
        "link": "https://github.com/hb-go/pkg",
        "scorecardV2": {
          "date": "2022-04-11",
          "repo": {
            "name": "github.com/hb-go/pkg",
            "commit": "e76fc4121c84b4d7ff072676f8337984ba5c49d7"
          },
          "scorecard": {
            "version": "4.1.0-112-gf9c2f9d",
            "commit": "f9c2f9d79f7a5bd8863899d548973520c1be4e27"
          },
          "check": [
            {
              "name": "Webhooks",
              "documentation": {
                "short": "This check validate if the webhook defined in the repository have a token configured.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#webhooks"
              },
              "score": -1,
              "reason": "check is not supported for this request: SCORECARD_V6 is not set, not running the Webhook check",
              "details": [
                "Warn: SCORECARD_V6 is not set, not running the Webhook check"
              ]
            },
            {
              "name": "Maintained",
              "documentation": {
                "short": "Determines if the project is \"actively maintained\".",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#maintained"
              },
              "score": 0,
              "reason": "0 commit(s) out of 14 and 0 issue activity out of 2 found in the last 90 days -- score normalized to 0",
              "details": []
            },
            {
              "name": "Code-Review",
              "documentation": {
                "short": "Determines if the project requires code review before pull requests (aka merge requests) are merged.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#code-review"
              },
              "score": 0,
              "reason": "no reviews found",
              "details": [
                "Warn: no reviews found for commit: e76fc4121c84b4d7ff072676f8337984ba5c49d7",
                "Warn: no reviews found for commit: 7fea6be509235850fda5cc2b4f36d0e07c52aeda",
                "Warn: no reviews found for commit: 1493b37868fb617447b4bdb2cf18e07127eadccd",
                "Warn: no reviews found for commit: f95b938800d45ee07a259b6f520096012b4269b4",
                "Warn: no reviews found for commit: c6f38be8e24512262181dbff0fd8c72d6f952a4e",
                "Warn: no reviews found for commit: 60ccdc730ae4836dc00dbe3d6293d82b092bafae",
                "Warn: no reviews found for commit: f24b8bae1eef51d242763f085449709c5d71ac84",
                "Warn: no reviews found for commit: c2d45c2ed032d61d8f33425b68087acc18f7945b",
                "Warn: no reviews found for commit: 346b31e462e2adf0d89e84e74ba5528149be7f4c",
                "Warn: no reviews found for commit: f5e4a11489fb0db6b552eeec21a4f7771acd6e9d",
                "Warn: no reviews found for commit: e1748b361233dfa970e62f6d296baff0ab00f849",
                "Warn: no reviews found for commit: d06d538f80b5fd1d3fbb912477db9d70b4f2ef85",
                "Warn: no reviews found for commit: 120bb6cbc6dab7dcc42342a56a57147d780adc00",
                "Warn: no reviews found for commit: 8fd0981f8aa8e583d5548c6895d912c14fab9512"
              ]
            },
            {
              "name": "CII-Best-Practices",
              "documentation": {
                "short": "Determines if the project has a CII Best Practices Badge.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#cii-best-practices"
              },
              "score": 0,
              "reason": "no badge detected",
              "details": []
            },
            {
              "name": "Vulnerabilities",
              "documentation": {
                "short": "Determines if the project has open, known unfixed vulnerabilities.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#vulnerabilities"
              },
              "score": 10,
              "reason": "no vulnerabilities detected",
              "details": []
            },
            {
              "name": "Signed-Releases",
              "documentation": {
                "short": "Determines if the project cryptographically signs release artifacts.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#signed-releases"
              },
              "score": -1,
              "reason": "no releases found",
              "details": [
                "Warn: no GitHub releases found"
              ]
            },
            {
              "name": "Dependency-Update-Tool",
              "documentation": {
                "short": "Determines if the project uses a dependency update tool.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#dependency-update-tool"
              },
              "score": 0,
              "reason": "no update tool detected",
              "details": [
                "Warn: dependabot config file not detected in source location.\n\t\t\tWe recommend setting this configuration in code so it can be easily verified by others.",
                "Warn: renovatebot config file not detected in source location.\n\t\t\tWe recommend setting this configuration in code so it can be easily verified by others."
              ]
            },
            {
              "name": "License",
              "documentation": {
                "short": "Determines if the project has defined a license.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#license"
              },
              "score": 10,
              "reason": "license file detected",
              "details": [
                "Info: : LICENSE:1"
              ]
            },
            {
              "name": "Token-Permissions",
              "documentation": {
                "short": "Determines if the project's workflows follow the principle of least privilege.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#token-permissions"
              },
              "score": 10,
              "reason": "tokens are read-only in GitHub workflows",
              "details": []
            },
            {
              "name": "Packaging",
              "documentation": {
                "short": "Determines if the project is published as a package that others can easily download, install, easily update, and uninstall.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#packaging"
              },
              "score": -1,
              "reason": "no published package detected",
              "details": [
                "Warn: no GitHub publishing workflow detected"
              ]
            },
            {
              "name": "Dangerous-Workflow",
              "documentation": {
                "short": "Determines if the project's GitHub Action workflows avoid dangerous patterns.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#dangerous-workflow"
              },
              "score": 10,
              "reason": "no dangerous workflow patterns detected",
              "details": []
            },
            {
              "name": "Pinned-Dependencies",
              "documentation": {
                "short": "Determines if the project has declared and pinned its dependencies.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#pinned-dependencies"
              },
              "score": 10,
              "reason": "all dependencies are pinned",
              "details": [
                "Info: GitHub-owned actions are pinned",
                "Info: Third-party actions are pinned",
                "Info: Dockerfile dependencies are pinned",
                "Info: no insecure (not pinned by hash) dependency downloads found in Dockerfiles",
                "Info: no insecure (not pinned by hash) dependency downloads found in shell scripts",
                "Info: no insecure (not pinned by hash) dependency downloads found in GitHub workflows"
              ]
            },
            {
              "name": "Binary-Artifacts",
              "documentation": {
                "short": "Determines if the project has generated executable (binary) artifacts in the source repository.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#binary-artifacts"
              },
              "score": 10,
              "reason": "no binaries found in the repo",
              "details": []
            },
            {
              "name": "Security-Policy",
              "documentation": {
                "short": "Determines if the project has published a security policy.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#security-policy"
              },
              "score": 0,
              "reason": "security policy file not detected",
              "details": []
            },
            {
              "name": "Fuzzing",
              "documentation": {
                "short": "Determines if the project uses fuzzing.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#fuzzing"
              },
              "score": 0,
              "reason": "project is not fuzzed",
              "details": []
            },
            {
              "name": "Branch-Protection",
              "documentation": {
                "short": "Determines if the default and release branches are protected with GitHub's branch protection settings.",
                "url": "https://github.com/ossf/scorecard/blob/f9c2f9d79f7a5bd8863899d548973520c1be4e27/docs/checks.md#branch-protection"
              },
              "score": 0,
              "reason": "branch protection not enabled on development/release branches",
              "details": [
                "Warn: branch protection not enabled for branch 'master'"
              ]
            }
          ],
          "metadata": [],
          "score": 4.8
        }
      }
    ],
    "advisories": [],
    "relatedPackages": {}
  },
  "defaultVersion": "v0.0.2"
}
```

### Versions

https://deps.dev/_/s/go/p/github.com%2Fhb-go%2Fpkg/versions

```json
{
  "package": {
    "system": "GO",
    "name": "github.com/hb-go/pkg"
  },
  "versions": [
    {
      "version": "v0.0.3-0.20211117154348-e76fc4121c84",
      "symbolicVersions": [],
      "dependentCount": 1
    },
    {
      "version": "v0.0.3-0.20210513173633-7fea6be50923",
      "symbolicVersions": [],
      "dependentCount": 1
    },
    {
      "version": "v0.0.2",
      "symbolicVersions": [],
      "isDefault": true,
      "dependentCount": 3
    },
    {
      "version": "v0.0.1",
      "symbolicVersions": [],
      "dependentCount": 0
    }
  ]
}
```

### Events

https://deps.dev/_/s/go/p/github.com%2Fhb-go%2Fpkg/events

```json
{
  "package": {
    "system": "GO",
    "name": "github.com/hb-go/pkg"
  },
  "events": []
}
```

### Dependencies

https://deps.dev/_/s/go/p/github.com%2Fhb-go%2Fpkg/v/v0.0.2/dependencies

```json
{
  "package": {
    "system": "GO",
    "name": "github.com/hb-go/pkg"
  },
  "version": "v0.0.2",
  "dependencyCount": 17,
  "directCount": 7,
  "indirectCount": 10,
  "dependencies": [
    {
      "package": {
        "system": "GO",
        "name": "github.com/hb-go/pkg"
      },
      "version": "v0.0.2",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 0,
      "dependencyCount": 17,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/fatih/color"
      },
      "version": "v1.7.0",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 3,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/go-redis/redis"
      },
      "version": "v6.15.3+incompatible",
      "type": "",
      "owners": [],
      "license": "BSD-2-Clause",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/go-redis/redis_rate"
      },
      "version": "v6.5.0+incompatible",
      "type": "",
      "owners": [],
      "license": "BSD-2-Clause",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 2,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/hashicorp/go-multierror"
      },
      "version": "v1.0.0",
      "type": "",
      "owners": [],
      "license": "MPL-2.0",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 1,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "go.uber.org/zap"
      },
      "version": "v1.10.0",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 2,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "golang.org/x/time"
      },
      "version": "v0.0.0-20210220033141-f8bda1e9f3ba",
      "type": "",
      "owners": [],
      "license": "BSD-3-Clause",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "google.golang.org/grpc"
      },
      "version": "v1.22.1",
      "type": "",
      "owners": [],
      "license": "Apache-2.0",
      "advisories": [],
      "distance": 1,
      "dependencyCount": 5,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/golang/protobuf"
      },
      "version": "v1.2.0",
      "type": "",
      "owners": [],
      "license": "BSD-3-Clause",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/hashicorp/errwrap"
      },
      "version": "v1.0.0",
      "type": "",
      "owners": [],
      "license": "MPL-2.0",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/mattn/go-colorable"
      },
      "version": "v0.1.2",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 2,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/mattn/go-isatty"
      },
      "version": "v0.0.8",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 1,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "go.uber.org/atomic"
      },
      "version": "v1.4.0",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "go.uber.org/multierr"
      },
      "version": "v1.1.0",
      "type": "",
      "owners": [],
      "license": "MIT",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 1,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "golang.org/x/net"
      },
      "version": "v0.0.0-20190620200207-3b0461eec859",
      "type": "",
      "owners": [],
      "license": "BSD-3-Clause",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 1,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "golang.org/x/sys"
      },
      "version": "v0.0.0-20190412213103-97732733099d",
      "type": "",
      "owners": [],
      "license": "BSD-3-Clause",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 0,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "google.golang.org/genproto"
      },
      "version": "v0.0.0-20180817151627-c66870c02cf8",
      "type": "",
      "owners": [],
      "license": "Apache-2.0",
      "advisories": [],
      "distance": 2,
      "dependencyCount": 1,
      "errors": []
    },
    {
      "package": {
        "system": "GO",
        "name": "golang.org/x/text"
      },
      "version": "v0.3.0",
      "type": "",
      "owners": [],
      "license": "BSD-3-Clause",
      "advisories": [
        {
          "source": "OSV",
          "sourceID": "GO-2020-0015",
          "sourceURL": "https://osv.dev/vulnerability/GO-2020-0015",
          "title": "GO-2020-0015",
          "description": "An attacker could provide a single byte to a [`UTF16`] decoder instantiated with\n[`UseBOM`] or [`ExpectBOM`] to trigger an infinite loop if the [`String`] function on\nthe [`Decoder`] is called, or the [`Decoder`] is passed to [`transform.String`].\nIf used to parse user supplied input, this may be used as a denial of service\nvector.\n",
          "referenceURLs": [
            "https://github.com/golang/go/issues/39491",
            "https://github.com/golang/text/commit/23ae387dee1f90d29a23c0e87ee0b46038fbed0e",
            "https://go-review.googlesource.com/c/text/+/238238",
            "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0015.yaml",
            "https://groups.google.com/g/golang-announce/c/bXVeAmGOqz0",
            "https://storage.googleapis.com/go-vulndb/byID/GO-2020-0015.json"
          ],
          "severity": "UNKNOWN",
          "gitHubSeverity": "UNKNOWN",
          "aliases": [
            "CVE-2020-14040"
          ],
          "disclosedAt": 1618401600,
          "observedAt": 1650652211
        },
        {
          "source": "OSV",
          "sourceID": "GO-2021-0113",
          "sourceURL": "https://osv.dev/vulnerability/GO-2021-0113",
          "title": "GO-2021-0113",
          "description": "Due to improper index calculation, an incorrectly formatted language tag can cause Parse\nto panic, due to an out of bounds read. If Parse is used to process untrusted user inputs,\nthis may be used as a vector for a denial of service attack.\n",
          "referenceURLs": [
            "https://go-review.googlesource.com/c/text/+/340830",
            "https://go.googlesource.com/text/+/383b2e75a7a4198c42f8f87833eefb772868a56f",
            "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0113.yaml",
            "https://storage.googleapis.com/go-vulndb/byID/GO-2021-0113.json"
          ],
          "severity": "UNKNOWN",
          "gitHubSeverity": "UNKNOWN",
          "aliases": [
            "CVE-2021-38561"
          ],
          "disclosedAt": 1633521600,
          "observedAt": 1639517407
        }
      ],
      "distance": 3,
      "dependencyCount": 0,
      "errors": []
    }
  ]
}
```

### Dependents

https://deps.dev/_/s/go/p/github.com%2Fhb-go%2Fpkg/v/v0.0.2/dependents

```json
{
  "package": {
    "system": "GO",
    "name": "github.com/hb-go/pkg"
  },
  "version": "v0.0.2",
  "totalCount": 3,
  "directCount": 3,
  "indirectCount": 0,
  "directSample": [
    {
      "package": {
        "system": "GO",
        "name": "github.com/micro-in-cn/starter-kit"
      },
      "version": "v1.18.0"
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/stack-labs/starter-kit"
      },
      "version": "v1.18.1-0.20210219171843-f8996cea7bc2"
    },
    {
      "package": {
        "system": "GO",
        "name": "github.com/hb-go/micro-mesh"
      },
      "version": "v0.0.0-20200506085011-1215b3befb11"
    }
  ],
  "indirectSample": []
}
```
