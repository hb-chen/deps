package template

import (
	"embed"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/hb-chen/deps/pkg/log"
	"github.com/hb-chen/deps/pkg/output"
)

//go:embed tpl/*.tpl
var embedTpl embed.FS

type tpl struct {
	opts *Options
}

func (t *tpl) Generate(deps map[string]*output.Dependency) (err error) {
	log.Logger.Debugf("template file: %v", t.opts.TplFile)
	var tmpl *template.Template
	if ok, file := InternalTpl(t.opts.TplFile); ok {
		pattern, err := io.ReadAll(file)
		if err != nil {
			return errors.Wrap(err, "tpl:"+t.opts.TplFile)
		}

		tmpl = template.Must(template.New(t.opts.TplFile).Parse(string(pattern)))
	} else {
		tmpl, err = template.ParseFiles(t.opts.TplFile)
		if err != nil {
			return errors.Wrap(err, "tpl:"+t.opts.TplFile)
		}
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
				log.Logger.Warn(err)
			}

			err = nil
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
		log.Logger.Warn(err)
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
