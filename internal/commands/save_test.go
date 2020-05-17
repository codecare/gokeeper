package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExecuteSave(t *testing.T) {

	dir, err := ioutil.TempDir("", "example")
	if err != nil {t.Errorf("save() error = %v", err); return}

	defer cleanUpTempDir(dir)

	application.SelectedFile = filepath.Join(dir, "test.vault")
	err = ExecuteSave()
	if err != nil {t.Errorf("save() error = %v", err); return}
}

func cleanUpTempDir(dir string) {
	if closeErr := os.RemoveAll(dir); closeErr != nil { panic(closeErr) }
}
