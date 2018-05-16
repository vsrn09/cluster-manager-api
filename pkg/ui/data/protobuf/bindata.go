// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x18\xf8\xe5\x92\x43\x6d\xb5\xe9\x1e\x70\x48\x10\xec\xf9\xdc\xde\xd6\x68\x9b\x04\x55\xb6\xc1\x3e\x05\x34\x35\x91\x78\xa1\x48\x2e\x49\xd9\xf5\x2d\xfa\xdd\x0f\xfc\x23\x59\x94\xe5\xf4\x2e\x5b\x60\x57\x0f\xad\x45\x0e\x7f\x9c\xdf\x70\xfe\x51\xc9\x32\x58\x4a\xb5\xd3\xac\xac\x2c\x9c\xbd\x7c\xf5\x77\xc8\x49\x6d\x1a\x51\x42\xfe\x26\x87\x25\x97\x4d\x01\x57\xc4\xb2\x0d\xc2\x52\xd6\xaa\xb1\x4c\x94\x70\x8b\xa4\x06\xd2\xd8\x4a\x6a\x33\x9f\x64\xd9\x24\xcb\xe0\x03\xa3\x28\x0c\x16\xd0\x88\x02\x35\xd8\x0a\x61\xa1\x08\xad\xb0\x9d\x79\x01\x9f\x51\x1b\x26\x05\x9c\xcd\x5f\xc2\x89\x13\x98\xc6\xa9\xe9\xe9\x85\x83\xd8\xc9\x06\x6a\xb2\x03\x21\x2d\x34\x06\xc1\x56\xcc\xc0\x03\xe3\x08\xf8\x85\xa2\xb2\xc0\x04\x50\x59\x2b\xce\x88\xa0\x08\x5b\x66\x2b\xbf\x4f\x44\x71\x9a\xc0\x2f\x11\x43\xae\x2d\x61\x02\x08\x50\xa9\x76\x20\x1f\xfa\x82\x40\x6c\x54\xda\x3d\x95\xb5\xea\x3c\xcb\xb6\xdb\xed\x9c\x78\x85\xe7\x52\x97\x19\x0f\xa2\x26\xfb\xb0\x5a\xbe\xbd\xca\xdf\xce\xce\xe6\x2f\xe3\xa2\x9f\x05\x47\x63\x40\xe3\xaf\x0d\xd3\x58\xc0\x7a\x07\x44\x29\xce\x28\x59\x73\x04\x4e\xb6\x20\x35\x90\x52\x23\x16\x60\xa5\x53\x7a\xab\x99\xb3\xdb\x0b\x30\xf2\xc1\x6e\x89\x46\x07\x53\x30\x63\x35\x5b\x37\x36\xb1\x59\xab\x22\x33\x89\x80\x14\x40\x04\x4c\x17\x39\xac\xf2\x29\xfc\x73\x91\xaf\xf2\x17\x0e\xe4\x6e\x75\xfb\xee\xfa\xe7\x5b\xb8\x5b\x7c\xfa\xb4\xb8\xba\x5d\xbd\xcd\xe1\xfa\x13\x2c\xaf\xaf\xde\xac\x6e\x57\xd7\x57\x39\x5c\xff\x0b\x16\x57\xbf\xc0\xfb\xd5\xd5\x9b\x17\x80\xcc\x56\xa8\x01\xbf\x28\xed\x18\x48\x0d\xcc\x59\x13\x0b\x6f\xba\x1c\x31\x51\xe1\x41\x06\x95\x8c\x42\xca\x1e\x18\x05\x4e\x44\xd9\x90\x12\xa1\x94\x1b\xd4\xc2\x79\x82\x42\x5d\x33\xe3\x4e\xd5\x00\x11\x85\x83\xe1\xac\x66\x96\x58\x3f\x74\xc0\x6b\x3e\x99\x98\x9d\xb0\xe4\x0b\x5c\xc2\x54\x69\x69\xe5\xeb\xe9\xc5\x64\xa2\x08\x7d\x74\xc0\x94\x37\xc6\xa2\xbe\xaf\x89\x20\x25\xea\x7b\xa2\xd8\xc5\x64\xc2\x6a\x25\xb5\x85\x69\x29\x65\xc9\x31\x23\x8a\x65\x44\x08\x19\x37\x99\x7b\x98\xe9\x45\x27\xe6\xdf\xe9\xac\x44\x31\x33\x5b\x52\x96\xa8\x33\xa9\xbc\xe8\xe8\xb2\xc9\x24\xcc\xc2\x49\xa9\x15\x9d\x97\xc4\xe2\x96\xec\xc2\x34\xbd\x2f\x51\xdc\x47\x94\x79\x44\x99\x4b\x85\x82\x28\xb6\x39\x6b\x67\x4e\xe1\x12\x7e\x9b\x00\x30\xf1\x20\xcf\xfd\x2f\x00\xcb\x2c\xc7\x73\x98\x2e\x03\x25\xf8\x18\x28\xc1\xe2\x66\x35\xbd\xf0\x12\x9b\x10\x0e\xe7\x30\xdd\xbc\x9c\xbf\x9a\xbf\x8c\xc3\x54\x0a\x4b\xa8\x6d\x71\xdc\x23\x48\xed\xa0\x3e\x32\x5a\x11\xe4\xf0\x19\x05\xfe\x87\x91\x28\xef\x9e\x46\xf3\x73\x98\x3a\x4f\x36\xe7\x59\x56\x32\x5b\x35\xeb\x39\x95\x75\xb6\x39\x10\xc5\x9a\x30\x27\x5c\xc7\xa9\x7f\x94\x6e\xc0\x09\x47\xa1\xaf\xee\x3f\xff\x0f\x7e\xb1\xa8\x05\xe1\xf7\x85\xa4\xa6\xd5\xe7\xe8\x56\x26\xe4\x8d\x19\x15\xd4\x66\xf1\x20\x67\xf1\x20\x67\x44\xb1\x08\x5f\xa0\xa1\x9a\x79\x4b\x3a\x4a\x52\x23\x90\xb5\x6c\x2c\x1c\x33\xd4\xd7\x09\x80\xa1\x15\xd6\x68\xce\xe1\xdd\xed\xed\xcd\xc5\x70\x20\x77\x23\x54\x0a\xd3\xf8\xa1\x69\x8c\x46\xb7\x45\xf6\x6f\x23\x85\x87\x51\x5a\x16\x0d\x3d\x36\xff\xf5\x62\x32\x31\xa8\x37\x8c\x62\xa7\x48\xe0\xeb\x02\xc3\x45\x09\xc2\x3b\xe4\x5c\xc2\x9d\xd4\xbc\x80\x3c\xca\xce\x60\xcb\x38\x07\x8d\x0a\x89\x05\x02\x2e\xea\x7d\x8a\xb4\xd2\xbb\xbd\x3b\x39\xb7\xf5\x86\x15\x58\x78\x3c\xad\x68\x40\x0a\x40\x27\xfb\xdf\x1f\x4d\x79\x0a\x1a\x6d\xa3\x85\xe9\x8f\x7f\x42\xc5\x77\xa7\x3d\x77\xe8\xfc\xd5\xc7\xc3\x9c\x28\x36\x77\xe7\xd1\x7a\xa1\x7b\x94\x34\x16\xce\x61\xea\x83\x65\xf3\x2a\xab\x1c\xda\xd6\xa1\x4d\xa3\xc4\x5a\x16\xbb\x73\x98\xfe\x75\xba\x3f\xf4\x60\xeb\x3e\x65\x25\x0b\xa0\xb2\x11\x16\x34\x1a\x25\x5d\x00\x03\xdc\x05\xc6\xee\xbd\xd8\x27\x61\xd1\xd4\x6b\xd4\x2e\xd3\x2a\x59\x18\x97\xf5\x5a\xfe\x46\x11\x3a\x62\x84\x9f\xd0\xde\xc8\x62\xe9\xd1\x4f\x7a\x2f\xa9\x19\x7a\x13\xcf\xb1\xc3\xb8\x35\x4a\xb4\x4a\x16\x9e\xd8\x34\x11\x74\x46\x81\xbd\x55\xc6\x2c\xe3\xd9\x7b\x32\xbe\x96\x91\x36\x67\x75\xbc\x96\x1a\x89\xc5\xd6\x89\x4e\x92\xd7\x94\x5b\x32\xf5\x3b\xd8\x35\x09\xb9\xa8\xcf\xf3\x88\x69\xb4\x9a\xe1\x26\x14\x02\x63\x89\x6d\x8c\x3b\xd2\x8e\xa5\x4b\xf2\xc0\xac\x81\xc7\x66\x8d\x54\x8a\x07\x56\xfa\x3a\x41\xa5\x10\x48\x2d\xdb\x30\xbb\xeb\x9f\x70\x67\x86\xfd\xef\x83\xf3\xfd\xdd\x06\x28\xf1\x69\x03\x8c\x32\x2d\x90\xa3\xc5\x91\xf3\x7b\xe3\x27\x3a\xc5\x93\xd7\x54\xf7\x64\xea\xf9\xea\x47\x4d\xfe\x6f\x06\x4c\x18\x4b\x38\x87\x13\xa9\x41\x63\x7c\x3b\x05\xcb\x38\xef\xd1\xb9\x69\x5d\xf5\xd6\x8f\xc3\xc9\x60\x20\xa5\x34\x98\xfc\x7e\x21\x17\xb4\x7a\x9e\x53\x1e\x21\x5a\x21\xaf\x81\x56\x44\xdb\x56\xfa\xd6\x35\x8c\x3e\x21\xaf\xd1\x15\x1a\xab\x1b\xea\x5b\x57\xe6\x5d\xd8\x89\x42\x45\x0c\x10\xae\x91\x14\x3b\x58\x23\x0a\x28\x50\x71\xb9\xc3\x5e\x2a\x33\x2e\x69\xbb\xcc\xd5\x19\x71\x15\xf6\x7c\x87\xbc\x5e\x7a\x94\x93\xe1\x48\x6a\xc6\xe1\xec\x77\x0b\x6e\xc7\xf9\x79\x46\x8c\x5e\xd6\xb1\x1d\x58\x6f\xef\xf9\x3d\x92\x83\x81\x31\xef\xff\x0e\x14\x0f\xfd\x3f\x65\xd9\xd1\xf9\x3a\x99\xf8\x43\x4e\x4b\xb1\xeb\xc3\xd1\xd8\x49\x8d\xc6\xb8\x1e\x32\x29\xa8\x71\x2b\xd7\x48\x8b\x32\x54\xe3\x4b\x78\x75\xd1\x83\x6a\x0b\x9b\xab\xd8\x3d\xd8\x11\x38\xcf\x30\x05\x6c\x85\x5a\xcc\xf6\x3d\x2d\x67\xfb\x4e\xe2\xaa\x2b\x87\x56\xc2\x03\x5a\x1a\x1c\xae\x2b\xb3\xad\xdc\x07\x24\x1b\x04\xac\x95\xdd\x39\xc9\x5f\x1b\xd4\x3b\x70\x21\xd0\xd5\x53\x33\xe4\x15\x60\x9f\x50\xa4\xaf\xbe\x53\xe5\x1b\xf5\xda\x05\x5b\xba\xe3\xa9\x5f\xca\x84\x7d\x7d\x16\xd6\x0c\x37\x1b\x96\xba\x94\x77\x7b\x09\x6b\x2b\x89\x95\x2e\x4a\xbb\x4a\x1a\x3b\x83\x83\x93\xea\x42\xbb\x6b\x21\x74\x77\x23\xf1\x2d\x9c\x97\x48\xb6\xbe\x89\x72\xb9\x42\xba\x5f\x74\x09\x67\xc7\xb5\x1d\x18\xe7\xae\x42\x7f\x53\x92\xda\x5f\x46\xfb\x6a\x6f\x89\xe9\x2b\xed\x6e\x7f\xfe\x9e\xda\xba\x61\x88\x49\xc9\x41\x3e\x1e\x10\x70\x05\xf5\xc0\x0e\x25\x0a\xd4\x7b\x26\xd1\x00\xb1\xf6\x0e\x95\x4e\xca\xe8\xff\x62\x5f\x2e\xe5\xa3\xbb\x60\xaa\xa3\x71\x70\x08\x3d\x30\xc6\xca\x24\xb8\xd1\x53\xcc\xce\x58\xac\x8f\xd2\xbd\xab\x88\x75\x97\xd8\xb4\x8f\xe8\xe1\x1c\x23\x3b\xb2\xbe\xd7\x6b\x58\xd9\xb6\x1a\x6d\x87\x3d\x02\xd7\x93\xbf\x84\xd7\x09\xc9\x61\x35\xef\x1f\xf9\x7e\xc3\x88\xf9\x17\x13\x2c\x65\x65\x28\x26\x72\xf7\x4d\x23\x1e\xb6\x04\xfb\x1d\x96\xb2\xe1\x45\x62\xca\xb6\x4a\xb9\x84\x7c\xd4\x92\x79\x62\xbd\xbe\x9b\x3d\xed\x2a\x87\x75\xbe\xa7\x4a\xeb\x23\xa1\x27\x30\x95\x57\x6d\x8d\x6d\xa9\xf5\x5f\x1c\xfa\x7b\xb4\x1a\xf7\x15\xeb\x25\xb3\xe3\x30\x4c\x1c\x4b\x54\xfb\xb3\xf6\xdf\x85\x84\xf4\x14\x3d\x92\x2f\x00\xe1\xdb\x49\xa8\xfc\x19\x34\xaa\xd4\xa4\x70\x67\xd1\xc7\x8b\x57\xe8\x70\xcc\xa9\xbb\x46\x9d\xba\xfe\x6e\xb6\x65\x45\x3b\xfa\x63\x67\xdb\xa0\x31\x73\x6d\xc1\x06\x53\x51\x52\xd4\x4c\x80\xd2\x6c\xc3\x38\x96\x68\x7e\xdc\x9f\x50\xfb\x99\xc2\xcb\x5d\xc2\x0f\x87\x26\x71\x3a\x38\x7f\xb2\x3d\xa3\xf8\xcf\x43\x56\x46\xe0\x68\xdf\x70\x73\xc4\xa2\x65\xe4\x27\xef\xf7\x49\x17\x2e\xe1\x6f\x4f\x1d\xeb\x30\x71\x91\xc0\x5e\xaa\x98\x53\xc0\x34\x94\xa2\x31\x0f\x0d\x7f\x3a\x56\x23\xbe\x81\x2d\x6a\x84\x92\x6d\x50\x8c\x57\xba\xd4\xcb\x46\xfa\xa0\xef\xec\x66\xb1\x6f\x35\x68\xdd\xc5\x3a\xd4\xbd\x9f\x5c\xce\x64\x34\x4c\xe5\x61\xa6\xdd\xa5\xef\x59\xa1\x8f\xc9\x47\x96\xee\xbb\x9c\xd0\x16\x0e\x13\xc5\x68\x07\xf7\x07\xda\xf9\xb0\x15\xfb\xd3\x99\xf9\xaa\xed\x98\xdb\x5d\x46\x0d\x3b\xd6\x36\xfe\x81\x76\x1d\xa5\x38\x5a\x14\x22\xef\x2e\x34\x9f\x68\xc0\x46\x56\xb7\x99\x2a\xcd\x71\xe3\xb9\x6c\x54\xc3\xbd\xc3\x8e\x6a\x27\x7a\x1d\x40\xe8\xf1\x6b\x14\xf6\xc9\x96\x6a\xb8\xbc\x6b\x4c\xc3\x7a\x3f\xde\xfb\x54\x36\x48\xbd\xe3\xa9\x3c\xa9\xa2\x5e\x5b\x8d\x4a\x1a\x66\xa5\x4e\x8a\xa7\x1b\x4d\xb2\xf6\xe1\x42\x71\xc4\x9d\x7e\x48\xd7\x10\x8d\xd1\x63\xc2\x27\xe7\x13\x81\xc6\xe5\xd3\x1d\xa9\x39\xcc\xfc\xd4\x67\xc2\x1b\x34\x73\x3f\x42\xa5\xb0\x28\x6c\x6c\x66\x5b\xd3\x7b\x81\x83\x54\x7b\xbc\xa9\x1c\x3d\x82\xae\xd3\x9c\x01\x6d\xb4\x46\x61\x79\xec\x0d\x99\x01\xb2\xf5\xdf\xdb\x6b\x42\x0e\x3a\xf7\x83\x36\xf1\x23\x21\xf9\xb7\x7a\x5c\x27\xe3\x55\x71\x88\xc9\x11\x38\x84\xc5\xdd\x37\x01\x16\x77\x61\xbd\x53\x6c\x18\xa3\xe3\x3b\xfd\x76\xa0\xe3\xe2\x66\x05\x28\x0a\x25\x59\xea\x6a\xed\xd8\x01\xb3\xc6\xa0\xf6\x94\x63\x13\xd9\xa2\xf4\x17\x77\x32\x43\x52\xd7\x8b\xc6\x56\xf0\x88\xbb\xee\x2f\x13\x63\x7b\x4b\xd2\xd8\xea\xde\x49\x3d\xc9\xaa\xa5\x9f\x90\xd2\x58\x3a\x47\x77\xf0\x8b\xbb\x3c\xf5\xd7\x32\xc4\x66\xca\x27\x47\xaa\xd1\xbe\xc7\xdd\xaa\x08\xab\x6e\x56\xb0\xf0\xf9\x2a\x69\xd0\xbc\x94\xd3\xe9\x9e\x15\x07\xb4\x02\x46\x58\xf5\x3e\x92\x73\x38\xe4\x18\x4e\x98\xe8\x51\xfc\x6f\x00\x00\x00\xff\xff\xff\x01\xcf\x00\xc4\x1b\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 7108, mode: os.FileMode(420), modTime: time.Unix(1526491145, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
