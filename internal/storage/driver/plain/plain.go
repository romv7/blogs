package plain

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/romv7/blogs/internal/storage/driver"
)

type Plain struct {
	rootPath string
}

var (
	Default = NewPlain(os.Getenv("STORAGE_DIR"))

	ErrAttemptReadDir     = errors.New("attempt to read dir")
	ErrPutOnExistFile     = errors.New("attempt to put a new data on an existing file")
	ErrRemoveNotExistFile = errors.New("attempt to remove a non existing file")
)

func NewPlain(rpath string) *Plain {
	var err error

	if _, err = os.ReadDir(rpath); os.IsNotExist(err) {
		err = os.Mkdir(rpath, os.ModeDir)
	}

	if err != nil {
		log.Panic(err)
	}

	return &Plain{rpath}
}

// Tries to lookup a file from the file system. If it does not found a file specified by
// the key, it returns os.ErrNotExist.
func (p *Plain) Get(key string) (b []byte, err error) {
	path := strings.Join([]string{p.rootPath, key}, "/")

	if finfo, err := os.Stat(path); os.IsNotExist(err) || err != nil {
		return nil, err
	} else if finfo.IsDir() {
		return nil, ErrAttemptReadDir
	} else {
		b, err = os.ReadFile(path)
	}

	return
}

// Saves a file from the directory specified by the `key` argument. If something went wrong
// return the error. The `basename` of the key file is the filename of the file.
func (p *Plain) Put(key string, b []byte) (err error) {
	path := []string{p.rootPath}
	tokens := strings.Split(key, "/")

	for i := range tokens[:len(tokens)-1] {
		path = append(path, tokens[i])

		tpath := strings.Join(path, "/")

		if fstat, err := os.Stat(tpath); err != nil {

			// If it is only an os.ErrNotExist ignore and don't panic.
			if !errors.Is(err, os.ErrNotExist) {
				log.Panic(err)
			}

		} else if fstat.Mode().IsDir() {
			// If the directory already exist check whether it's a directory
			// or a file. If it is a file, return an error.
			continue
		} else if fstat.Mode().IsRegular() {
			return ErrPutOnExistFile
		}

		if err = os.Mkdir(tpath, os.FileMode(0700)); err != nil {
			log.Panic(err)
		}
	}

	path = append(path, tokens[len(tokens)-1])

	if err := os.WriteFile(strings.Join(path, "/"), b, os.FileMode(0644)); err != nil {
		log.Panic(err)
	}
	return
}

// Removes a file by using the provided `key` argument. Returns a os.ErrNotExist if it tries to
// delete a file that does not exist.
func (p *Plain) Remove(key string) (err error) {
	path := strings.Join([]string{p.rootPath, key}, "/")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrRemoveNotExistFile
	}

	if err = os.Remove(path); err != nil {
		log.Panic(err)
	}

	return
}

// Describes what is the content of the `key` argument. Describe() uses an array to denote
// its results, if the`key` however represents a file it will return the key itself, if it doesn't
// exist it will return an os.ErrNotExist.
func (p *Plain) Describe(key string) (res []*driver.PathInfo, err error) {
	path := strings.Join([]string{p.rootPath, key}, "/")
	res = make([]*driver.PathInfo, 0)

	if finfo, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	} else if finfo.Mode().IsRegular() {
		res = append(res, &driver.PathInfo{
			Key: key,
			Sub: nil,
		})

		return res, nil
	} else if err != nil {
		log.Panic(err)
	}

	ents, err := os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}

	for _, ent := range ents {
		finfo, err := ent.Info()
		if err != nil {
			log.Panic(err)
		}

		entKey := key + "/" + ent.Name()

		if finfo.Mode().IsRegular() {
			res = append(res, &driver.PathInfo{
				Key: entKey,
				Sub: nil,
			})
		} else if finfo.IsDir() {
			childs, _ := p.Describe(entKey)

			res = append(res, &driver.PathInfo{
				Key: entKey,
				Sub: childs,
			})
		}
	}

	return
}

// Check whether the `key` argument exists in the `rootPath` directory.
func (p *Plain) Contains(key string) bool {
	path := strings.Join([]string{p.rootPath, key}, "/")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Panic(err)
	}

	return true
}
