package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_fileDoesExist(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {t.Errorf("save() error = %v", err); return}

	defer cleanUpTempDir(dir)

	file := filepath.Join(dir, "somefile.txt")

	application.SelectedFile = file

	err = ExecuteSave()
	if err != nil {t.Errorf("save() error = %v", err); return}

	doesExist, err := fileDoesExist(application.SelectedFile)
	if err != nil {t.Errorf("save() error = %v", err); return}

	if doesExist != true {
		t.Errorf("save() error = %v", err)
	}

}
