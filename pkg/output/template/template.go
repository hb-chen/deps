package template

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hb-chen/deps/pkg/log"
	"github.com/hb-chen/deps/pkg/output"
	"github.com/pkg/errors"
)

type tpl struct {
	opts *Options
}

func (t *tpl) Generate(deps map[string]*output.Dependency) error {
	tmpl, err := template.ParseFiles(t.opts.TplFile)
	if err != nil {
		return errors.Wrap(err, "tpl:"+t.opts.TplFile)
	}

	// create dir when not exist
	dir := filepath.Dir(t.opts.OutFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0760); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile(dir, filepath.Base(t.opts.OutFile)+".")
	if err != nil {
		return err
	}
	defer func() {
		if err = tmpfile.Close(); err != nil {
			if e, ok := err.(*os.PathError); !ok || e.Err != os.ErrClosed {
				log.Logger.Info(err)
			}
		}
	}()

	err = tmpl.Execute(tmpfile, deps)
	if err != nil {
		return errors.Wrap(err, "tmpl:"+t.opts.OutFile)
	}

	if err = tmpfile.Chmod(0644); err != nil {
		return err
	}

	// Close the file immediately for platforms (eg. Windows) that cannot move
	// a file while a process is holding a file handle.
	err = tmpfile.Close()
	if err != nil {
		return err
	}

	err = os.Rename(tmpfile.Name(), t.opts.OutFile)
	if err != nil {
		return err
	}

	return nil
}

func NewOutput(opts ...Option) output.Output {
	o := NewOptions(opts...)
	return &tpl{opts: o}
}
