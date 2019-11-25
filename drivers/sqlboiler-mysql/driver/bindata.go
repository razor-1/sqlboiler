// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.go.tpl
// override/templates/singleton/mysql_upsert.go.tpl
// override/templates_test/singleton/mysql_main_test.go.tpl
// override/templates_test/singleton/mysql_suites_test.go.tpl
// override/templates_test/upsert.go.tpl
// DO NOT EDIT!

package driver

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x48\x0e\x7f\xb6\x3f\x05\xd7\xe8\xee\xca\x07\x55\xed\x01\x87\x7b\xe8\x21\x0f\xcd\x9f\x76\x73\x4d\xba\x4e\xdc\x5c\x80\x0b\x82\x62\x22\x51\xce\x20\xe3\x19\x75\x34\x4a\xea\xd3\xe9\xbb\x1f\xc8\x91\x2c\xc9\xb1\x1d\xb7\xdb\x1c\xf6\x29\xd6\x0c\x45\x72\xf8\x23\x7f\xe4\x28\x65\xf9\x12\x5e\x08\x25\x45\x0e\x6f\xf6\x20\x7a\x4b\xbf\x30\x8f\x3e\x89\x1b\x85\xe0\xff\x44\x1f\xc5\x1c\xab\x6a\xc8\xa2\x79\x7c\x8b\x73\xe1\xb7\xe9\x85\x56\x02\xfe\x0b\xd1\xb4\xdd\xe5\x17\x64\x0a\xd1\xdb\x24\x79\xaf\xcc\x8d\x50\xf0\xb2\xaa\x86\xaf\x5e\xc1\x45\x96\xa3\x75\xef\x41\x38\x87\xf3\xcc\xe5\x20\x34\x48\x4d\x6b\x21\x08\x9d\x40\x62\x90\xd7\x8a\x2c\x11\x0e\xc1\x58\x90\x33\x6d\x2c\x82\xd1\x10\x1b\x9d\x2a\x19\xbb\x68\x98\x16\x3a\x86\xc0\xc0\x5f\xca\xd2\xfb\x1f\x5d\x64\x53\xa9\x67\x85\x12\xb6\xaa\xc6\x8d\x95\x80\x9d\xd0\xc6\x41\xf4\xd1\x1c\x18\xed\xf0\xab\xab\xaa\xd8\x7d\x25\x55\xf4\x10\xd5\x8b\x21\x94\x25\xea\x84\x9c\xac\x2d\x1f\x18\x55\xcc\x75\x1e\xd6\xce\xd5\x8f\x70\x63\xa4\x8a\xea\x87\x31\xa0\xb5\xc6\x42\x39\x1c\x58\x74\x85\xd5\x60\x22\x6f\xd8\xdb\xed\xda\xe4\xf7\xde\xa3\x3b\xdc\x0f\xc6\x65\x89\x2a\x47\xf6\x23\x84\x66\xa3\x96\xac\xf7\x75\x52\x55\xe1\x56\x4f\xc6\xc3\x6a\x38\x5c\x3a\x3d\xf4\xe1\xa6\x00\x76\x42\x4e\x3f\x27\x42\xcb\x78\x25\xf8\x93\x3f\x16\x7d\x60\x9d\x39\xad\x71\x00\x76\x86\x63\xf2\xdc\x78\x94\xc3\x81\x4c\xc9\x29\xca\xce\xff\x27\x18\xff\x60\xa3\x3f\xed\x81\x96\x8a\xbc\x18\x64\x14\xa2\x80\xf5\x5d\x5a\x91\x1d\x59\x1b\xa0\xb5\xe3\xf1\x70\x50\xad\x03\x6e\x03\x52\xeb\x80\x82\x22\x97\x7a\x46\xcf\xf8\x15\xe3\xc2\x19\xfb\x2d\x85\xd3\x51\x9d\x7d\x1f\x8a\x93\xc7\xf1\x24\x47\x7c\xec\x8e\x6a\x97\x3a\x51\x7d\x0c\x6d\x2b\x5e\x2f\x75\xde\x7a\x3a\xd6\xbb\x43\xbe\x26\xcf\xba\x79\x45\x6e\x3c\x1f\xac\xf7\xc2\xc2\x7c\x31\x3d\x3b\x59\x1b\xcc\x0b\x2d\xbf\x14\x8d\x55\xd8\x83\xab\xeb\xdc\x59\xa9\x67\x25\xf3\xac\x15\x7a\x86\xf0\x42\x86\xf0\x22\x36\xaa\xc3\xb4\xcd\x0b\x64\x61\x40\x92\x32\x65\x91\xc8\xeb\xa3\xd5\x51\x59\xf2\x8a\xa7\xed\x51\xe8\xe5\x1a\xb7\xea\xdf\x15\x7b\xbb\xcc\x85\xe7\xc8\xb2\x29\x62\x0f\x29\x48\x4c\x5c\xcc\x51\x3b\xe1\xa4\xd1\x90\x1a\x0b\xb7\xe6\x01\x9c\x81\xcc\x9a\x0c\xad\x5a\x40\x91\x63\x1f\x0e\xb6\xd8\x43\x64\xd7\x24\xfd\x73\xe5\xe8\xb2\x4d\xc8\x14\x0c\xec\xb5\xe9\x54\xb7\x0d\xde\xcf\xa3\x8f\xf8\x10\x8c\xca\x32\x9a\xdc\xcd\x3c\x7a\x6f\x40\x1b\x28\xcb\x5e\x23\xa6\x70\xdd\xcb\x04\x13\x0e\x61\xc1\xa7\x1d\x71\xfe\x79\xa4\x09\x48\x45\xd0\x8c\x9c\x9c\x63\xee\xc4\x3c\xfb\xec\xa5\x3e\xdf\xa2\xca\xd0\x8e\x20\x82\xca\x4b\xb7\x35\xf2\x9b\x31\x77\x75\x5a\x75\xab\x29\x31\xfb\x98\x1a\x8b\x3e\xa8\x2c\xb4\x73\x69\x3d\x2e\x9e\xf6\xb4\xe4\xee\xa0\xcd\xc5\xe1\x40\xff\xe7\x10\x53\x51\x28\xc7\x83\xc8\x97\x02\xad\xc4\x3c\xfa\x68\xf4\xbf\xd1\x9a\x7a\x6b\x8a\x04\x6b\x0d\xfa\xa1\x79\xd0\x2d\xec\x75\xa4\x2f\xa5\xbb\xad\x85\x43\x30\x63\x52\xeb\x0b\xe3\x09\xad\x3b\xd6\x29\xeb\xe4\x00\x29\xd4\xc1\x52\xf7\x98\x10\x7d\xbd\x09\xcf\x58\x68\x0a\x96\x87\x00\x1e\xa4\xbb\x05\x01\x8e\x27\x28\x77\x2b\x1c\xd4\xfb\x4d\xed\x50\x1d\x09\x28\x58\x33\xc4\x6c\xb7\x41\xf7\xd5\x2b\xd8\x2f\xa4\x4a\x20\x16\xf1\x2d\xc2\x1d\x2e\x40\xea\x97\x4a\x6a\x84\x62\xa6\xa4\x5a\xc0\x4b\x98\x2f\xf2\x2f\x0a\xee\x73\xc8\xe8\x6f\x66\xcd\x8d\xc2\x79\x3e\x1c\xdc\x14\x29\x85\x20\x77\x76\x2e\xf4\x4c\x21\x35\xb9\xfd\x22\x4d\xd1\x06\x63\xde\x8d\x2e\xad\x74\x38\x65\x12\x0a\x72\x67\x63\xa3\xef\xa3\x63\x67\x44\xd0\xcb\xf3\xe8\x83\xd4\x09\xd1\x1d\x25\xdf\xe7\x10\x62\xd2\xea\xe9\xaa\x2f\x77\x60\x54\xce\x21\x59\xd5\x1d\xf3\x69\xda\xe5\xfd\x85\xc3\xe0\xd7\xe8\xd7\xa7\xdc\xe8\xd3\xc0\x66\x37\xfa\x72\xdf\xe3\xc6\x63\x9d\x9d\xec\xfc\x01\xba\x9a\x94\xdc\xa2\x8a\xb0\x7d\xb3\x07\xb4\x5b\x6f\x8c\x87\x83\x16\xbc\x49\xd1\x80\x77\x53\xa4\x63\x2e\xe5\xb5\x65\xe1\xcb\xf6\x80\xd2\xe5\xb4\x70\xd1\xf9\x89\x89\xef\x48\x13\x27\x50\xe8\xf3\x28\x21\x43\x4f\xbf\x7f\x75\x87\x8b\xeb\x9d\x0d\x5d\x68\xe5\x4d\x0d\x07\xd4\x07\x89\x07\xb8\x26\x7c\xf5\xfc\x54\x1b\xa6\x00\x34\xc3\xa7\x45\x47\x8e\xf4\xd1\x3b\xee\x3c\x51\x9d\x0e\x07\x83\x4d\x1e\xbc\x55\xaa\xa9\xd2\x2d\x52\x6b\x78\x62\x37\x69\x53\xb8\xee\x0b\x6d\x42\xd0\xe3\x78\x38\x18\xd4\xfd\xf0\xcd\xde\x4a\x1d\x5c\x74\x9e\x7e\xc8\x11\x26\x56\xce\x85\x5d\x7c\xc0\x45\x47\x98\x02\xdd\xf0\x92\xb7\xdf\x21\xa5\xa7\xbb\x4c\xa1\x3d\x1f\x99\x86\xa6\x56\x7a\x4e\x08\xb1\x29\x54\xc2\xac\x7f\xc3\x14\x54\x1f\xd7\x13\x14\x28\x99\x73\x0f\x62\x9a\x22\x73\xd0\xa5\x9a\x29\xcd\xd3\xf3\x4c\x21\x75\xff\xc0\xa2\x0b\xdb\x22\xa0\x97\x38\x1b\x22\x62\xe7\x05\xec\x79\xfd\x3e\x9f\xce\x68\xe9\x94\xb8\x39\x48\xa4\x50\x18\xbb\x10\x46\x2b\xae\x8d\x9a\x46\xdc\x74\xe0\x56\xa3\x45\xaf\x01\xf6\x20\x9d\xbb\x68\x9a\x59\xa9\x5d\xca\x08\x8c\xa6\x47\x27\x47\x07\x9f\xe0\xe7\x1c\xde\x9d\xff\x7e\x4a\xe7\x3d\x39\xab\xaa\x15\xdd\x65\x19\x9d\x9f\x55\x15\x5c\xfe\x76\x74\x7e\x04\x3f\xe7\x23\xc6\xc5\x0f\x6a\x79\xf4\x4f\x23\x75\xd0\x9e\xf2\x38\x41\xed\xce\x0a\xe3\x70\xaa\x64\x8c\x8d\xc7\xd1\xc9\x59\x08\xcd\xef\xf3\x33\x4e\xf4\x71\x08\xa3\x70\x34\x6e\xb4\xd5\x0a\x2e\x6f\xd1\xe2\x81\x12\x45\x8e\x8c\x0f\x39\x34\xf2\x07\x3e\xf7\x3f\x5f\x77\x03\xb7\x84\xdd\x1f\xf6\x5e\xa8\x02\x4f\x45\x96\x49\x3d\x0b\xb9\xe0\xda\x86\xb7\x2f\x75\x52\x6f\x6d\x6a\xa0\x9f\x16\x19\x86\x9b\x68\x60\xa9\xb6\x8d\x70\x3d\x24\x74\x9a\x7b\xaf\xbb\x13\x87\x35\xf9\x48\x07\x26\xc1\x3a\x19\x97\xd8\x3c\xb7\xb3\x64\x97\x0c\xae\x71\xb5\xef\x2b\x3b\x5b\xf9\x1e\xcb\x61\x64\xb2\xc6\x94\x21\x3b\xd6\x89\xb4\x18\x53\xde\xfa\x85\x7f\x91\xc4\xef\x69\x60\xa8\xfd\xdc\x0b\xd5\x1b\x2d\x78\x33\x7f\x67\xcd\xbc\x39\x02\x2b\xac\xa9\xb6\x07\xd2\xd8\x53\xa3\xf7\x24\x87\xab\x6b\xa9\x1d\xda\x54\xc4\x58\x56\xcb\x19\x63\x35\x58\x9d\x40\x36\x2f\xb6\xc6\x27\xce\x6e\x36\xdd\xd1\xe1\x4f\x2a\x53\x3f\xa4\x1e\xe2\x4d\x31\x3b\x35\x09\xb2\x56\x2a\x94\x77\x5c\x28\x4a\x07\xed\x3e\xb7\x28\xdb\xe8\xe2\x52\x1d\x3f\x2d\x4d\xd1\x59\x4e\xa6\x2f\x62\xa1\x4f\x44\xee\x3c\xa7\x1f\x1f\x76\x6f\x35\x2b\x3b\xf5\xed\x86\xef\x36\xeb\xb6\x06\x2b\xc3\xbd\x5f\xb5\x98\xf3\xdc\x57\x0f\xaf\x34\x82\xf2\xa8\x1f\x74\x9c\xf6\x3e\x45\x51\x34\x66\x2d\x34\xff\x6f\x7f\xb9\xb6\x10\xf0\x7c\xbb\x45\x51\x7d\xbd\xea\xe9\x5c\xef\xe6\xe7\x26\xe1\xbf\xcd\xc1\xc7\xaf\x7d\xbb\x6b\xcd\xb8\xbd\xa6\x24\xfa\x2d\x82\xae\xb6\x74\xaf\xf5\xec\xb3\xad\x51\xd0\x7c\xb3\xca\xc8\x4b\xc8\x37\x02\x48\x89\xaf\x68\xf5\x10\xa4\x76\x7f\xff\x5b\xcf\x39\xda\xf4\xf3\xef\xa9\xc8\xe0\xea\xba\xa8\x45\x68\xbd\xa1\x3f\x1e\xeb\xfa\x25\xb3\xa5\x66\x96\x9d\x70\x66\x9c\x01\x9e\x52\xea\x1b\xcf\x93\x9e\x7a\x2f\x9b\xd8\xfb\x2c\x89\x3a\x62\x09\x8d\x53\x1b\xc3\x79\x64\xed\x74\xa1\xe3\x77\x42\xaa\xb6\x0c\x8c\xe2\xef\xa5\x3c\xec\x24\xf8\xb5\x29\x82\xc9\x07\x5c\x2c\xef\xca\xaf\x5b\xc8\x56\xbe\x00\xf0\xc7\x29\x6e\xba\x4b\x4d\x3d\xd1\x4f\xd2\x29\x3f\xd3\xd5\xec\xb8\x22\x4d\xb2\x26\xf2\x7e\x78\xd9\xaa\x02\x1e\x00\x63\xa3\x22\x62\xd6\xaa\x0a\xfc\xa9\xfd\xc9\x6a\x9c\x98\x77\x7e\xf9\x65\x73\x84\xff\x4a\xbb\xab\x3b\x57\xaf\xaf\x69\x6f\x3b\x55\x5f\x8d\xda\xb0\x54\xd5\xe8\x7a\x33\x54\xbd\x2b\xe3\x32\x47\x9e\xad\x83\x74\xa7\x94\x1f\x50\x32\x16\x9d\x95\x78\x8f\xcd\xed\x8e\xf9\x39\xdf\x50\x42\x40\xc7\xed\xa5\xfb\xb6\x2e\xb3\x4b\xb7\x0a\xdb\xaa\x1a\xff\x31\xfe\x6f\x06\xab\x1d\x5a\x40\xf7\x04\x9e\x92\x96\x05\xb7\x4a\x8c\x1d\x7a\x63\xed\xe7\xe6\x21\xe8\xdb\x7b\xac\x2e\x9a\xc6\x82\x27\x0c\x6a\x85\x5e\x7f\x97\x34\xd7\xa8\x5c\xc3\x9a\xdf\xaa\xbe\x21\xd4\x1f\x90\x12\x99\xc9\x0a\xfe\x58\x93\xf8\xeb\xc4\xf6\x9c\xa0\xd8\x75\x4b\xe2\xcd\xa3\xdb\xd4\x6e\xd7\xb3\xe6\x1a\xb8\x83\x38\x5f\xfb\x60\xcf\x47\x6a\x67\x03\xcb\xeb\xdf\x60\xcb\x77\xa6\xe5\xbf\x4c\x12\xf3\x36\x75\x68\xbf\xeb\x1b\x53\x4d\x09\x9d\x3e\xce\x4a\x35\x11\x6e\xf7\x5b\xe7\xff\x02\x00\x00\xff\xff\xc0\xad\xaf\xbb\xe9\x1a\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8e\xd3\x30\x10\xc7\xcf\xf6\x53\x0c\x91\x56\x8d\x25\x2b\xcb\x5e\x91\x7a\xd8\xa5\x65\x15\x28\xfd\x2e\x08\x21\x0e\x6e\x3d\x6e\x2d\xa5\x4e\xf1\x47\xa1\x42\x7d\x77\xe4\x24\x6d\xb3\x4b\x41\x1c\xf6\x92\x8c\x3d\x33\x7f\xcf\x6f\x66\x6e\x6f\x61\x19\x74\x21\x17\x3b\x87\xd6\x4f\x02\xda\xc3\xc7\xc3\x6c\x32\xa8\x6f\x1d\x08\x88\x07\xe7\x85\xc7\x2d\x1a\x0f\xce\x5b\x6d\xd6\x10\x5c\xfc\xfa\x0d\x42\xa8\x12\x7b\xc2\x0b\xd8\xd9\x72\xaf\x25\xca\x8c\xaa\x60\x56\xd7\x75\x53\xa9\x05\x48\xab\xf7\x68\x5d\xd6\xd3\xa2\xc0\x95\xe7\xe0\xc5\xb2\xc0\xa1\xd8\x62\xa3\xcf\x21\xec\xa4\xf0\xc8\xe1\xc7\x46\x7b\x2c\xb4\xf3\xf0\xf5\x5b\xed\x63\xa7\x1a\x7e\x51\x72\xf1\x76\xe3\xed\x56\x98\x75\x81\x59\x2e\xd1\xf8\x49\x28\x3d\xce\x0a\xbd\xc2\xf8\x64\x36\x98\x70\x88\xff\xe9\xa4\xa5\xc9\x28\xb9\xbc\x7c\x5d\xe1\x8f\xe4\x73\x02\xa3\x94\x2c\x83\x82\x37\xed\xc4\x47\xf4\x0f\x41\x29\xb4\x29\xa3\x44\xa2\x42\xdb\x72\x8e\xc3\xc9\xb9\x0c\x2a\xa6\xef\x85\x85\x55\x59\x84\xad\x71\x0d\x14\x25\x5a\x41\x81\x26\xbd\xd4\x08\xaf\xba\xf0\x3a\xc2\x92\x53\x68\xb7\x09\x76\xd9\xfb\x52\xb7\x42\x39\x24\x3c\x61\x94\x1c\xe9\x59\xa6\x6e\x23\x83\xee\x49\x43\x6d\x7d\xf6\x6e\x67\xb5\xf1\x2a\xa5\x84\x44\x02\x1e\xff\x49\x3e\x9c\xf5\xa7\x73\xc8\x1f\x87\xa3\x69\x1f\xf2\xe1\x7c\x04\x37\x0e\xd2\x1b\xc7\xe0\xd3\xfd\x60\xd1\x9f\x55\x76\x52\x05\x9f\x7b\x50\x9d\x9a\xb2\x2a\xbb\x05\x5b\x88\x15\x6e\xca\x42\xa2\x75\x55\x13\x17\x0e\x73\x23\xf1\x67\xdb\xc1\x9f\xb1\x72\xb8\xe3\x70\xc7\xa2\x14\xa3\x84\x58\xf4\xc1\x1a\x58\x06\x95\xcd\x2a\xe2\xb4\xa1\x7b\x46\xd1\x40\x9c\x19\xfe\x52\x3c\x8c\x86\xd0\x5b\x8c\x07\xf9\xdb\xfb\x79\x1f\x3e\xf4\xbf\xc0\x62\xdc\x8b\x66\x45\xf5\x04\xaa\xc5\xf4\x62\x48\x71\xe2\xaa\xb4\xa0\x39\xec\xe3\xd6\x58\x61\xd6\xd8\x2c\x7a\x35\x1b\xad\x40\x5f\xa6\x1d\xa9\xb2\xcf\x56\x7b\x7c\x38\x78\x4c\x3b\xbc\x13\x5b\x72\xa4\x84\x7c\x8f\x8b\x29\x9f\x2e\xde\x3f\x36\x76\xcf\x68\x4b\xac\x69\x64\xad\x71\xcd\x93\x40\xb7\x69\x5a\x9a\xfc\x67\x66\x5d\x20\xeb\x34\xd3\xb9\x36\xb6\x23\xfd\x1d\x00\x00\xff\xff\x4c\x0d\x4e\x35\x6a\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xe3\x36\x16\x7e\x96\x7e\xc5\xa9\x81\xe9\x4a\x53\x85\x19\xa0\xc0\x3e\xa4\x10\x82\x89\xe3\x14\x41\x27\x37\x3b\xbb\xc5\xa2\x29\x5a\x46\xa2\x13\x22\x12\xa9\x90\x54\x1c\x6f\x90\xff\xbe\x38\xa4\x2e\x94\x63\x79\x33\xbb\xf3\xd8\xa7\x44\xe2\xc7\x73\xf9\xce\xe1\xe1\x27\x3f\x51\x05\xea\xee\xf9\x6c\xbd\xb8\xfa\xf2\xc0\xd6\x90\x82\x62\x77\xec\xb9\x22\x67\xb5\x36\x53\x59\x56\xbc\x60\xd1\x9f\xd1\x61\x19\x47\x51\x72\x23\xe2\xc3\x1b\xfd\xc3\xf4\xe2\x7c\x71\x3d\xff\x7c\x7a\x7e\x4d\x3e\x1e\x9e\x5c\xcc\x67\xa7\x3f\x9f\xc3\x2f\xb3\x7f\x91\x8f\x87\x37\x22\xfe\xe1\xcf\x38\x0c\xcd\xba\x62\x50\xae\xf5\x63\x71\xcd\xb4\x61\x0a\xb4\x51\x75\x66\xe0\x25\x0c\xf2\xdb\xa9\x14\x02\x3e\xea\xc7\x82\x1c\x1f\x85\xf8\xe2\x9c\x96\x0c\x10\xc2\xc5\x5d\x18\xdc\x4b\x6d\x00\xfa\xe7\x5a\x33\xe5\x3f\x57\x54\x6b\xff\x59\xeb\xa2\x94\x39\xeb\xd7\xa5\xb2\xfb\xb9\x30\x61\x18\xc8\xca\x70\x29\x4e\x78\xd1\x01\xc2\xc0\x30\x6d\x8e\x8f\xac\xd7\xce\xc8\x03\xaf\x16\x57\x5f\xa6\x65\x0e\xb7\x52\x16\xe1\x6b\x18\x2e\x6b\x91\x01\x17\xdc\x44\xb1\x8b\xfb\x8c\x72\x01\x29\x7c\xef\xe5\xf5\xf2\xda\x21\xa3\x12\x3e\x7a\x2b\x31\x68\x66\xea\x2a\x8a\x81\x29\x25\x15\x5a\x40\xae\x99\x52\xee\x45\x18\x06\x4f\xbc\x62\x8a\x2c\x98\x39\x66\x4b\x5a\x17\x26\x9a\xd8\xfd\xa4\x49\x68\x92\xc0\xc4\xa8\x9a\x4d\xe2\x71\x28\xe6\x3a\x49\xe0\xc7\x1f\x3f\xfd\x3d\x0e\xc3\xa0\x24\x0d\x99\x29\xb8\x1d\x3f\x33\xb3\xb0\x19\xb6\x1b\xf2\x5b\x41\x4b\x6b\xb2\x24\x96\xe8\x51\x24\xae\x3a\x9c\x2d\xc0\x28\x0e\x57\x1d\xce\x16\x66\x14\x87\xab\x0d\x0e\x0b\xe4\xe1\x4e\xc5\x30\x1f\x0b\x6a\xab\x3a\x6a\xaf\x65\xc9\xa2\xbd\x8a\x8e\x6e\x40\x8c\x9f\xbe\x57\x72\x6f\xcf\x91\x94\x45\xe7\xe2\x81\x57\xfa\xb1\xc8\xca\x7c\x82\xec\x62\xed\x52\x78\xa2\x05\x25\x47\xec\x8e\x8b\x7f\xd2\x82\xe7\x14\xdb\x2b\x8a\x49\xf3\xc0\xa2\x30\x08\x2c\xc4\x39\x3f\x97\x66\x56\x56\x66\x1d\x39\x1a\x13\x18\xb0\x96\x8c\x82\x91\xfd\x0e\xec\x4a\xd1\x81\xcf\xa5\x89\xec\x3f\xb3\xc7\x9a\x16\x3a\x72\x8c\x26\xf0\xa9\xdb\xe0\x68\xdc\x61\xde\xb5\x49\x87\x6f\x69\x19\xdf\xd0\xb0\xdd\xed\xe8\xd8\x4f\xc2\x20\x26\xd3\x7b\x96\x3d\x44\xc8\x11\x5f\xda\x16\xff\x2e\x05\xc1\x0b\x6c\xfa\x40\x31\x53\x2b\x81\x6f\xc3\xe0\x35\x0c\x83\xfd\x7d\x98\x2a\x46\x0d\x03\x0a\x8a\x8a\x5c\x96\xfc\xdf\x2c\x87\xfc\x16\x30\x06\x62\x4d\x14\x4c\x44\x7e\x51\x63\x48\x53\xf8\x64\xcd\x6d\xd4\xba\xb3\x40\x16\x86\xde\x16\xcc\x2d\x74\x19\xc6\xce\x67\x13\x55\x0a\x25\x29\xe9\x03\xbb\xe8\x66\x42\x14\xff\x34\x1e\xaf\x54\x9a\xfc\xaa\x68\x15\x31\x85\x85\xcb\x64\x5d\xe4\xe2\x6f\x06\xd0\x04\xb8\xb9\x02\x4b\x5e\xd8\x76\x6a\xbc\x7c\x37\x68\x2b\x34\xe7\xb9\xce\x95\xac\xae\x6d\xf0\x5b\xdc\x0e\x78\x0a\x5e\x87\x3b\x33\x4b\xd8\xbb\xf7\x86\x41\x90\xd7\x65\x85\x21\x1c\xa4\xc0\x9e\x59\x46\xa6\xb2\x2c\xa9\xc8\x9b\xce\xc6\xd5\x49\x82\x21\xb9\x71\xa2\x1d\x17\x09\x4c\xf6\xf6\x84\xdc\xcb\xa9\xa1\x6e\xb9\x25\x31\x70\x11\x8c\x5b\x1c\xb3\x86\xa6\x6e\xa9\x66\x76\xdd\x2b\x28\xc6\xa8\x12\x58\xa1\x39\x2e\xc9\x25\xaf\x58\x14\xf7\x71\x2f\x4c\x8e\x39\x1e\xa4\xf0\xfd\xed\xda\x30\x4d\x8e\xea\xe5\xd2\x8e\x5b\x2f\x94\x71\x50\x6f\x88\x2c\x4c\x2e\x6b\x1c\x37\xab\xe1\x4b\x47\xed\xc0\x5d\xe8\x1b\x47\x8c\x1d\xf7\x82\xad\x4e\x7e\x61\xeb\x63\xa6\x8d\x92\x6b\xa6\x22\xef\xba\x4c\x40\xc5\x9b\x9b\x9c\xe1\x8d\x20\x43\xbf\x9e\x7d\x14\x54\x99\xdd\xe5\xdc\x68\xc1\x25\xe5\x05\xcb\xc1\x48\xd0\xb8\x17\xba\x62\x42\xe6\xaa\x81\xad\x38\x6c\x1e\x3f\xb6\x6f\xe2\x6e\xc3\xd5\xb6\xc4\x7e\xa5\x7c\xab\xa3\x65\x69\xc8\xa5\xe2\xc2\x14\x02\x3d\xc4\x9b\xef\x06\xd5\x68\x66\x50\x14\xc7\xef\x8c\x71\x45\xb9\x81\xa5\x54\xa3\xac\x84\x41\xf0\x07\x36\x02\x99\x16\x52\xb3\x28\x86\xfd\x7d\xf8\xbc\x44\x75\xd2\x9e\x16\xae\x21\x97\x82\x25\x90\x21\x02\xcc\x3d\x83\x95\xe2\x86\x01\x13\x39\xc8\xa5\x7d\x51\xf1\x8a\x85\xdb\x19\xfe\x5f\xf3\xde\x68\x96\xff\x33\xf3\xcd\x5e\xc0\xc4\x1b\x23\x82\x17\x3b\xf4\x8a\x2e\xce\x64\xce\x22\x4f\x4c\xc5\xcd\x5f\x4c\x43\xaf\xb8\xc9\xee\xc1\xae\xbe\x84\x41\x46\x35\x6b\xf4\xc9\x41\x3f\x35\x27\xf3\xd9\xd5\x3f\x4e\xe7\xb3\xe3\x49\x8b\x58\xd2\x42\x0f\x21\xc7\xa7\x8b\xcf\x47\x5f\x2c\xa4\x19\x18\xfe\xea\xe5\x7c\x76\x32\x9b\x3b\x0b\x3b\xc4\xd5\x70\xd4\x78\x61\x36\x76\x90\xde\x45\x85\xfc\x2e\x23\x1c\x43\x0d\x7c\x0f\xe7\x75\xfa\x41\xdb\x71\xd4\x4b\xc3\x78\xdc\xd1\xe6\x7d\xd1\xcb\x39\x53\x56\x09\x34\x03\x88\xcb\xda\xf0\x82\x5c\xb3\xb2\xb2\xb0\x09\x8a\x37\x67\xbf\xbd\x21\x76\xdd\x8c\xa3\x95\x75\x9d\xb1\xf5\xb2\xd1\xd7\xd3\x4b\x74\x6d\x09\x0e\x83\x3f\x92\xa6\x1d\xa5\xc6\x93\x6e\x1a\x0d\xe1\x1c\x4b\x4d\x4e\x35\xde\xe6\xcf\x5c\x1b\xdb\x82\xee\x6e\xb2\x36\x52\xc0\x2a\x86\xc1\x2b\xb0\x42\x33\xf8\x8a\x38\xed\x8d\x08\x42\x1a\x9c\x0f\x06\xca\x4e\x33\x62\x80\x58\x81\x93\xaa\xe9\x70\xcb\xd5\xe4\xb7\xac\xe0\x4c\x98\xdf\x11\xd2\x2f\x2f\x9b\x55\xdc\x9c\x7e\xd0\x37\xc2\x16\xa7\x09\xfe\x2d\x0c\xb5\x4d\xfa\x21\x6f\x60\xf8\xb4\x15\x86\x02\xab\xb7\x86\x4f\xb1\x27\x2d\x50\x8c\xc6\x98\xa3\x13\x15\x5b\xbc\x50\xad\x57\x52\xe5\xbd\x09\xbb\x05\x53\xdb\x82\xd6\xba\xd8\xc3\x83\xd1\xa3\xbb\xc3\xd4\x2a\xa5\xd8\xb9\x77\x94\x0f\x7d\x76\xfc\x54\x4a\x1a\x99\xc9\x22\x35\x59\xb5\x8b\xc6\x6e\xc0\xfd\xc5\xe4\x57\x30\xe9\x1f\x78\x6c\xfa\xb2\x22\x56\x2b\xc6\xfd\x7c\xc4\x77\xcd\xe5\x30\x3e\x11\x86\x62\xac\x9f\x07\x38\x7a\xf1\x3c\xfa\x93\xa7\x39\xbf\xad\x0a\x82\x0f\xfa\xa7\x37\x4a\xa8\x75\x5e\x12\x55\x8b\x69\x99\x47\xfa\xb1\x68\x75\xf6\x64\x47\x1c\xbe\x9c\xdc\x1d\x05\x22\xfb\x18\xf0\x80\xe3\x1c\xd0\xdf\x34\x1a\xc3\xa8\xca\xe5\x4a\xf8\xb1\xf0\xa5\xd5\x90\xf6\x7b\xff\xed\x3c\x69\x97\x3a\xc6\xff\xab\x88\x3e\xf8\x7a\x15\xed\x5d\x7e\x52\x93\x39\x2b\xe5\x13\xb6\xd2\xbb\x46\x7f\x4b\x00\x0a\xc1\xa4\xbd\x55\x9b\xab\x26\x01\xaa\xee\x34\x10\x42\xda\x9b\xb2\xcb\xda\x2e\xa4\x40\xab\x8a\x89\x3c\xfa\xed\x77\x07\x78\xd9\x94\xc7\xaf\xce\x04\x21\x04\x1b\x30\xdb\xa2\xac\x1b\x8f\x1e\x0e\x61\x9d\x30\x75\x76\x35\x39\x67\xab\x39\xa3\x39\x53\x2e\x52\xb4\xa6\x9d\xe8\xdd\x26\x9f\xf5\xb8\xb2\xce\x7c\xb9\xec\x4c\x74\x2f\xdd\xdd\xe2\x36\x87\x5e\x3d\x70\x79\x5e\x8b\xb7\xa5\xf0\xf5\x4d\x7b\xa1\xa9\x5a\x08\x2e\xee\x0e\x26\x1d\x9b\x2e\xb7\x78\x08\x77\xae\x7d\x15\xb4\xb1\xba\xa1\x91\x36\xbf\x30\xdf\x23\x76\x32\x29\xb0\x55\xa3\xe6\x67\xa8\xc4\x95\x2f\x1e\xef\xda\x8d\xa6\x4d\xac\x79\xeb\x6e\xf8\xb3\x4e\xd0\x23\x1a\xce\x1e\x0b\x72\x51\x31\xd1\x7f\x28\xe5\x8a\x3f\x31\x45\xec\x47\xc4\x51\xcd\x8b\xfc\xaa\x66\x6a\xdd\x24\xd4\xfe\x4e\xe0\xc6\xe4\xf0\x74\xb6\xd3\xbc\x1d\xd7\xcd\x78\xf4\x86\xe2\xb0\x06\x3d\x11\xc9\x1b\x76\x86\x89\xbc\x86\xff\x09\x00\x00\xff\xff\xdd\x6c\xf4\x1c\x09\x14\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\xd8\x5d\x50\x0b\x85\xd9\x5e\x53\xf8\xe0\x7c\x1c\x82\xb6\x86\x1b\xcb\xe7\x82\x91\x46\x0e\x61\x9a\x54\xc9\x51\x6d\x57\xe0\x7f\x2f\x28\xd9\x8e\x13\x3b\x6d\x0e\xed\x21\x07\x7f\x70\xf0\x66\xde\x7b\x33\x1c\xb6\xed\x19\xfc\x2d\xb5\x92\x1e\x2e\x86\x20\x46\xf1\x1f\x7a\x91\xcb\x7b\x8d\xd0\xff\x88\xb1\x5c\x62\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xb3\x7a\xa2\x1b\x27\x75\x08\xb3\xda\xa3\x23\x4e\xf0\x5f\x04\x28\x33\x17\x79\x0a\x2d\x4b\x48\x4c\xa4\x93\x5a\xa3\xe6\x29\x63\x89\xaa\x40\xa3\xe1\xfb\x02\xd7\x76\x65\xa6\xca\xcc\x1b\x2d\x5d\x08\x23\xad\xaf\xac\x6e\x96\xc6\xa7\x30\x1c\xfe\x0c\x39\x71\x6a\x29\xdd\xe6\x03\x6e\xf6\x09\x2d\x4b\x12\x12\xd3\x85\xaa\xf9\x20\x7e\xd7\xca\xcc\x81\x3a\x1b\x2b\x45\x0f\x60\x8d\xde\x40\xdd\xe7\xc1\x02\x37\x50\xf4\x99\x83\x94\x25\x61\xaf\x6c\xb9\x99\x7e\xfe\x78\xe0\xef\x91\x72\x66\xd4\xd7\x06\x0f\xf5\xfd\xff\x4b\x4e\x63\xa1\xe9\xd2\x76\x64\x40\x16\x0a\x6b\x2a\xad\x0a\x02\x6b\x7a\x6e\x96\x78\xc4\x32\xb6\xdf\x49\x53\xda\xa5\xfa\x8e\x62\x8c\xab\x29\x62\xc9\x53\x96\x7c\x93\x0e\xd0\x75\x1f\xeb\x58\x72\x7e\x0e\x23\x22\x5c\xd6\x04\xf4\x80\x70\x3b\x9e\xde\xdc\xe5\xe0\x55\x89\x60\x2b\x90\x06\x66\x93\x18\x61\x89\x8d\x15\x4f\x5a\x69\x7b\xbf\xb1\xe8\x21\xe7\x94\x5c\x53\x10\x8f\x62\x32\xf8\xd7\x66\xf0\x42\xf3\xaf\x2f\xf3\x4d\x8d\x3e\x83\x4a\x6a\x8f\xe9\xfb\xae\xd0\x5f\x43\x30\x4a\x6f\x3b\x72\x13\xa5\x56\x7c\x30\x33\x5d\x2f\xc8\x3e\xb2\x9c\x56\x04\xbe\xe3\xbe\x80\x7f\xfc\x20\x8b\xf5\xb6\x8d\x69\x5b\x55\x81\xb1\x04\x62\x6c\xaf\xac\x21\x5c\x53\x08\x05\xad\xa3\xb5\xa2\x3f\x8b\x4b\x59\x2c\xe6\xce\x36\xa6\xe4\x69\xdb\xa2\x29\x43\x60\x49\x0f\xf9\xd4\x78\xca\xd7\xbc\xab\x72\x58\xe1\x28\x70\x6f\x95\x16\x97\x38\x57\xa6\xab\xa1\x3d\x1e\xc6\xf2\x35\x2f\x68\x9d\x45\x83\x3b\x86\x57\x81\x52\x96\x94\x58\xa1\x83\xb8\x39\x3c\x85\x16\xbe\xc0\x10\x68\x2d\xee\xac\xd6\xf7\xb2\x58\xf0\x14\x42\x1c\xf1\x7e\x18\x56\x6c\x17\xe9\x25\xe3\x71\x28\x68\x4a\x38\x0b\x01\xe2\xa9\xe3\xbf\x35\x15\x3a\x9e\x3e\x3d\xbd\x6e\x2e\x4d\x47\x77\x7a\x28\x47\xd3\x28\x6c\x63\xa8\x0b\x3c\xbb\x5a\xbb\x57\x80\xa7\xe2\x2a\x62\x5e\x29\xff\xd1\xf9\xb1\x4a\xbe\xa3\x8d\x90\x8e\x38\x82\xde\x3d\x81\x0c\x56\xd2\xc4\x35\x42\x70\x58\x58\x57\x66\x30\xb7\x74\x31\xc8\x7a\xfc\x56\xf4\xb3\x7d\x99\x4d\xae\x47\xf9\xcd\xa9\x7d\xf9\x6d\x1b\xf1\x22\xec\xe8\xd5\x12\x42\xfc\xd1\xf5\x79\x7b\xf7\xea\x8d\x5c\xab\xc0\x7e\x04\x00\x00\xff\xff\x63\xfd\x7b\x9f\x38\x07\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,
	"templates/singleton/mysql_upsert.go.tpl": templatesSingletonMysql_upsertGoTpl,
	"templates_test/singleton/mysql_main_test.go.tpl": templates_testSingletonMysql_main_testGoTpl,
	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,
	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl": &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

