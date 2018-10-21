package file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// get absolute filepath, based on built executable file
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// get filepath base name
func Basename(fp string) string {
	return path.Base(fp)
}

// get filepath dir name
func Dir(fp string) string {
	return path.Dir(fp)
}

// mkdir dir if not exist
func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// Unlink delete file
func Unlink(fp string) error {
	if !IsExist(fp) {
		return nil
	}
	return os.Remove(fp)
}

// list dirs under dirPath
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

func ReadBytes(cpath string) ([]byte, error) {
	if !IsExist(cpath) {
		return nil, fmt.Errorf("%s not exists", cpath)
	}

	if !IsFile(cpath) {
		return nil, fmt.Errorf("%s not file", cpath)
	}

	return ioutil.ReadFile(cpath)
}

func ReadYaml(cpath string, cptr interface{}) error {
	bs, err := ReadBytes(cpath)
	if err != nil {
		return fmt.Errorf("cannot read %s: %s", cpath, err.Error())
	}

	err = yaml.Unmarshal(bs, cptr)
	if err != nil {
		return fmt.Errorf("cannot parse %s: %s", cpath, err.Error())
	}

	return nil
}

func ReadJson(cpath string, cptr interface{}) error {
	bs, err := ReadBytes(cpath)
	if err != nil {
		return fmt.Errorf("cannot read %s: %s", cpath, err.Error())
	}

	err = json.Unmarshal(bs, cptr)
	if err != nil {
		return fmt.Errorf("cannot parse %s: %s", cpath, err.Error())
	}

	return nil
}

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

func WriteString(filePath string, s string) (int, error) {
	return WriteBytes(filePath, []byte(s))
}

func Download(toFile, url string) error {
	out, err := os.Create(toFile)
	if err != nil {
		return err
	}

	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.Body == nil {
		return fmt.Errorf("%s response body is nil", url)
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func MD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := md5.New()

	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
