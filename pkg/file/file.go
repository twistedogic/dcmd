package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const WORKSPACE_PREFIX = "/workplace"

func CreateWorkspaceName() string {
	return fmt.Sprintf("%s_%d", WORKSPACE_PREFIX, time.Now().Unix())
}

func isRootPath(fp string) bool {
	return []rune(fp)[0] == '/'
}

func DirPath(fp string) string {
	return filepath.Dir(fp)
}

func GetFilePath(fp string) (string, error) { //TODO: should only return volume
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return "", err
	}
	if fullpath, err := filepath.Abs(fp); err == nil {
		return fullpath, err
	}
	return fp, nil
}

func SyntheticPath(fp string) string {
	return filepath.Join(CreateWorkspaceName(), fp)
}
