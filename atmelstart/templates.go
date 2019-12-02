// Code generated by "esc -o templates.go -pkg atmelstart templates"; DO NOT EDIT.

package atmelstart

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/toolchain.cmake": {
		name:    "toolchain.cmake",
		local:   "templates/toolchain.cmake",
		size:    2759,
		modtime: 1575313752,
		compressed: `
H4sIAAAAAAAC/5xWTXObSBA9i1/RZbvKUmlRdit7ySEHDFjLRgIVoMQ6TY2GljxrGLTDkMRF8d+3hg8Z
yUrWlZPE0P36vek3zVzD3PXd0IpdB+69hQsmOAH4QQyu48WGwTL6hCTjgmdlRiT+W3KJyfizG0Ze4MP7
2e8Tw7iGey4SYHl24CnKmbHjIiEHme8lzcZWuCS2DVRmpsgFmki33NwzNrkQ9vBwHjedXogL7v62g9Xm
LDbf/sPyw3ND6HaOAiVnt7DLJWC2xSTBBIrnQmEG37h6zEsFVEAQzYwC1dheWp9cEm2i2F0S31q60CE0
cBEqsAcCXzJsYgfLlbdwQ7ipWqn1pM+YTi/nTKevsh4ejmmdDigVT7l6Hmb2wtuk7qluKK4kfkWhCki5
eOJiD7woSoRvjzxFUFgovXaJTRxuejYktsK5G5N4s3Ihiq3Ys8nCuwutcPOyDf1GhUEQAy3gpDsTkHmu
hvD3nu8QHUtWVvwXWQaOS1ZhMA+tJfjuZzecnACfhh/xU76VVD6/Eb/jDIG/2LwJfsdTHE/etVKoenxj
Hc+3F2vHfXudA2VPdI9v3SbL/mTNe3jjGiyVYQqRolI1AJBwiUzlsjOJFS/dBYliK4yJ42lvdYZbh6Hr
x2ThRc2L1jELLp5QQsEkP3RcFp7/yQ1JZIfeKoarm+oMsX5XVbM2L2rS6vpKC7fzLMvF0V5ARdL4ECXs
UrovOqXBchn45G7tLRyiS5H7hTWP4MrM1GOZbcHM2KH8WFUze7U+Q24g2XR6Ac/Wp6eHuqm61ea5BjMo
wNztSsEUz4VZYPOr1xKq6OA5S3OxNxlN0wLA3L8H8wtNUzAdQqpq5uBXzrCuCbmavJoHJ5y6qdDSGfAZ
sKzBLFTycS/KDx/g6tK0uIB4lPh/mNPpH3+CuRO5qR4l0qSgOzQLRRVnRbsuleLtP/zO8NBswJHGhbY1
BNwHl3T2ON/q84bWYH5JfzN1UanMvczLA5hp1q2iSPo1szggKz4KKvJZ87cL2bNBY+Kb6sSW2heaa15K
hqBPbgH4XUnKFCawk3kGt3vG3i3pE+q3tzMj5YUaW6uV6zswdHQUrENbn72FGxkAAFVlgqRijzBr8e81
fF03L39wHOr66piLIqlro+HnCZaWCR6PKP91mt2Y0RVf0ezKOFz+Cs2MMpmPqWo6RWiSEPyOrFR0qz8a
VO5REUEznBjGSFsB013zDDfV4G09w3Q3aSO2XFyM2HIxBNFD9jie7jzfCjct05uqr1EPEH8S3hfUM03r
u4agVIdSFYDprrHHrFk/FTc+pdd8UOd+8/sDf2j80TWE2KhTjzjA1zxJh3eQ+QGl4lic11iFwcoNY8+N
IFjHq3Xc3jNOFHcSnIFpVA6885LK+wnRE2i1dWW6MDKw3CsO67uFZ5/JHPrrhUNz/6EKYctFr1NvIisL
lWeE5VlGRTI2RqP21gDnpYIobkeDMRrpSWH5zrGH/dUFzEDjU/nc7YPutG5D3/XaGE0MA0XyM6tO/gsA
AP//FM69IscKAAA=
`,
	},

	"/templates": {
		name:  "templates",
		local: `templates`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"templates": {
		_escData["/templates/toolchain.cmake"],
	},
}
