# Deps

Deps 是一个帮助获取项目依赖包的版本、开源协议、安全漏洞等清单的工具，数据源来自 Google 的 [deps.dev](https://deps.dev/) 项目，当前支持 Go、Java。

## Feature

- 获取项目依赖包清单，包括以下信息:
  - [x] 包名、版本
  - [x] 开源协议
  - [x] 安全漏洞
  - [ ] 是否直接依赖
  - [ ] 包间依赖拓扑
- 支持语言
  - [x] Go mod
  - [x] Java maven
  - [ ] npm
    - https://stackoverflow.com/questions/25997519/how-to-view-the-dependency-tree-of-a-given-npm-module
- 数据源
  - [x] [deps.dev](https://deps.dev/)
  - [ ] 取本地 mod/pom 中的 licenses 信息
- 其他功能
  - [ ] 支持模板输出，自定义格式
  - [ ] 开源协议友好性标识
  - [ ] 通过 Project 目录文件自动识别语言及包管理工具
  - [ ] License 文件提取
  

<img width="875" alt="image" src="https://user-images.githubusercontent.com/730866/183232928-4d37c224-3252-47d8-ac50-7e848745517a.png">

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
# Markdown
./out/deps.md

# Stdout
Dependencies:
Pkg: golang.org/x/tools, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: github.com/kr/text, Licenses: MIT , Direct: false ,Advisories: 0 
Pkg: golang.org/x/mod, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
...
```

### Maven

> 引入依赖 [github.com/ferstl/depgraph-maven-plugin](https://github.com/ferstl/depgraph-maven-plugin)，能够通过 `mvnw depgraph:graph` OR `mvn depgraph:graph` 生成 `target/dependency-graph.json` 

```shell
# Run
cd {golang project path}
deps -s maven
# Or run with -p
# deps -s go -p {project path}
deps -s maven -p ./example/java/

# Output
# Markdown
./out/deps.md

# Stdout
Dependencies:
Pkg: org.mockito:mockito-junit-jupiter, Licenses: MIT , Direct: false ,Advisories: 0 
Pkg: org.ow2.asm:asm, Licenses: BSD-3-Clause , Direct: false ,Advisories: 0 
Pkg: jakarta.xml.bind:jakarta.xml.bind-api, Licenses: non-standard , Direct: false ,Advisories: 0 
Pkg: org.apache.logging.log4j:log4j-api, Licenses: Apache-2.0 , Direct: false ,Advisories: 0 
...
```

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
