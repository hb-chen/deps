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

type tpl struct{}

func (*tpl) Generate(deps map[string]*output.Dependency, tplPath, outFile string) error {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return errors.Wrap(err, "tpl:"+tplPath)
	}

	// create dir when not exist
	dir := filepath.Dir(outFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0760); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile(dir, filepath.Base(outFile)+".")
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
		return errors.Wrap(err, "tmpl:"+outFile)
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

	err = os.Rename(tmpfile.Name(), outFile)
	if err != nil {
		return err
	}

	return nil
}

func NewOutput() output.Output {
	return &tpl{}
}
