# Deps

- Go module graph
  - `TODO` Support more language: `maven` `npm`
- Dependencies license, advisories...
  - [deps.dev](https://deps.dev/)
- https://github.com/ferstl/depgraph-maven-plugin
- https://stackoverflow.com/questions/25997519/how-to-view-the-dependency-tree-of-a-given-npm-module

## Install

```shell
go install github.com/hb-chen/deps@latest
```

## Usage

```shell
Usage:
  deps [flags]

Flags:
  -h, --help             help for deps
  -p, --project string   Project path (default "", use "pwd")
  -s, --system string    System type:auto, mod, maven (default "auto")
  -v, --version          version for deps
```

### Golang
```shell
# Run
cd {golang project path}
deps -s mod
# Or run with -p
deps -s mod -p {project path}

# Output
Dependencies:
Dependencies:
Pkg: golang.org/x/tools, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/kr/text, Licenses: MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/mod, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
...
```

### Maven

> 依赖 [github.com/ferstl/depgraph-maven-plugin](https://github.com/ferstl/depgraph-maven-plugin)

```shell
# Run
cd {golang project path}
deps -s maven
# Or run with -p
# deps -s go -p {project path}
deps -s maven -p ./example/java/

# Output
Dependencies:
Dependencies:
Pkg: org.mockito:mockito-junit-jupiter, Licenses: MIT , Direct: false ,Advisories: 0 
Pkg: org.ow2.asm:asm, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: jakarta.xml.bind:jakarta.xml.bind-api, Licenses: non-standard , Direct: false ,Advisories: 0 
Pkg: org.apache.logging.log4j:log4j-api, Licenses: Apache-2.0 , Direct: false ,Advisories: 0 
...
```

## TODO

- npm
- 取本地 mod/pom 中的 licenses 信息
- 支持模板输出，自定义格式
- 支持自定义 `mod` 或 `maven` 的 `cmd`
- 通过 Project 目录文件自动识别语言及包管理工具


## deps.dev API

- Go
  - https://deps.dev/_/s/go/p/github.com%2Fhb-chen%2Fdeps/v
  - https://deps.dev/_/s/go/p/github.com%2Fhb-chen%2Fdeps/versions
  - https://deps.dev/_/s/go/p/github.com%2Fhb-chen%2Fdeps/events
  - https://deps.dev/_/s/go/p/github.com%2Fhb-chen%2Fdeps/v/v0.0.0-20220424095100-bb6c886cdd4d/dependencies
  - https://deps.dev/_/s/go/p/github.com%2Fhb-chen%2Fdeps/v/v0.0.0-20220424095100-bb6c886cdd4d/dependents
- Java
  - https://deps.dev/_/s/maven/p/org.springframework%3Aspring-core/v/5.3.22
  - https://deps.dev/_/s/maven/p/org.springframework%3Aspring-core/versions
  - https://deps.dev/_/s/maven/p/org.springframework%3Aspring-core/events
  - https://deps.dev/_/s/maven/p/org.springframework%3Aspring-core/v/5.3.22/dependencies
  - https://deps.dev/_/s/maven/p/org.springframework%3Aspring-core/v/5.3.22/dependents
