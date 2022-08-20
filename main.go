package main

import "github.com/hb-chen/deps/cmd"

//go:generate go-bindata -o ./pkg/output/template/bindata.go -pkg template -prefix template template

func main() {
	cmd.Execute()
}
