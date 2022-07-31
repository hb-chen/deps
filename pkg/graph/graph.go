package graph

const PkgSysMaven = "maven"
const PkgSysGo = "go"

type Dependency struct {
	Package     *Package
	Requirement *Package
}

type Package struct {
	System  string
	Name    string
	Version string
	Direct  bool
}

type Graph interface {
	Graphs() ([]*Dependency, error)
}
