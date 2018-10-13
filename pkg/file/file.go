package file

import (
	"os"
	"path/filepath"
)

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
	return filepath.Join("/workspace", fp)
}
