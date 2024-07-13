package template

import (
	"io/fs"
	"path"
)

func InternalTpl(tpl string) (bool, fs.File) {
	f, err := embedTpl.Open(path.Join("tpl", tpl))
	return err == nil, f
}
