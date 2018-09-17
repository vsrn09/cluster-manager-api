// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x6b\x6f\xe3\x36\xb3\xfe\xee\x5f\x31\xf0\x97\x93\x1c\x24\xd6\x6e\xb2\x3d\xa7\x48\x9a\x83\xba\x4e\xdb\x35\xb2\xb9\x60\x9d\x6e\xd0\x4f\x06\x4d\x8d\x65\x1e\x53\xa4\x4a\x52\x76\xbc\xc5\xfe\xf7\x17\xbc\x48\xa2\x64\x39\x7b\x69\x0b\xbc\x78\x03\xb4\x6b\x89\x33\xa3\xe1\x33\xb7\x87\xb2\x93\x04\x26\xb2\xd8\x29\x96\xad\x0c\x9c\xbd\x7a\xfd\x3d\xcc\x48\xae\x4b\x91\xc1\xec\x7a\x06\x13\x2e\xcb\x14\xee\x88\x61\x1b\x84\x89\xcc\x8b\xd2\x30\x91\xc1\x23\x92\x1c\x48\x69\x56\x52\xe9\xd1\x20\x49\x06\x49\x02\xef\x18\x45\xa1\x31\x85\x52\xa4\xa8\xc0\xac\x10\xc6\x05\xa1\x2b\xac\x56\x4e\xe0\x03\x2a\xcd\xa4\x80\xb3\xd1\x2b\x38\xb2\x02\xc3\xb0\x34\x3c\xbe\xb4\x26\x76\xb2\x84\x9c\xec\x40\x48\x03\xa5\x46\x30\x2b\xa6\x61\xc9\x38\x02\x3e\x53\x2c\x0c\x30\x01\x54\xe6\x05\x67\x44\x50\x84\x2d\x33\x2b\xf7\x9c\x60\xc5\x7a\x02\xbf\x07\x1b\x72\x61\x08\x13\x40\x80\xca\x62\x07\x72\x19\x0b\x02\x31\xc1\x69\xfb\xb7\x32\xa6\xb8\x48\x92\xed\x76\x3b\x22\xce\xe1\x91\x54\x59\xc2\xbd\xa8\x4e\xde\x4d\x27\x3f\xdf\xcd\x7e\x3e\x3d\x1b\xbd\x0a\x4a\xbf\x09\x8e\x5a\x83\xc2\x3f\x4a\xa6\x30\x85\xc5\x0e\x48\x51\x70\x46\xc9\x82\x23\x70\xb2\x05\xa9\x80\x64\x0a\x31\x05\x23\xad\xd3\x5b\xc5\x2c\x6e\x27\xa0\xe5\xd2\x6c\x89\x42\x6b\x26\x65\xda\x28\xb6\x28\x4d\x0b\xb3\xca\x45\xa6\x5b\x02\x52\x00\x11\x30\x1c\xcf\x60\x3a\x1b\xc2\x4f\xe3\xd9\x74\x76\x62\x8d\x3c\x4d\x1f\xdf\xde\xff\xf6\x08\x4f\xe3\xf7\xef\xc7\x77\x8f\xd3\x9f\x67\x70\xff\x1e\x26\xf7\x77\xd7\xd3\xc7\xe9\xfd\xdd\x0c\xee\x7f\x81\xf1\xdd\xef\x70\x33\xbd\xbb\x3e\x01\x64\x66\x85\x0a\xf0\xb9\x50\x76\x07\x52\x01\xb3\x68\x62\xea\xa0\x9b\x21\xb6\x5c\x58\x4a\xef\x92\x2e\x90\xb2\x25\xa3\xc0\x89\xc8\x4a\x92\x21\x64\x72\x83\x4a\xd8\x4c\x28\x50\xe5\x4c\xdb\xa8\x6a\x20\x22\xb5\x66\x38\xcb\x99\x21\xc6\xdd\xda\xdb\xd7\x68\x60\x45\x6e\x19\x5d\x11\xe4\xf0\x01\x05\x7e\x64\x04\x7e\xc8\x37\xfe\xd3\x8f\x59\x4e\x18\x1f\x51\x99\xff\x9f\x95\x1b\x73\xb6\x26\xf0\x8e\x28\x8d\x02\x7e\x20\xf6\x6a\xc4\xdd\x55\x2c\x38\xd0\x3b\x61\xc8\x33\x5c\xc1\xb0\x50\xd2\xc8\xf3\xe1\xe5\x60\x50\x10\xba\xb6\xae\x52\x5e\x6a\x83\x6a\x9e\x13\x41\x32\x54\x73\x52\xb0\xcb\xc1\x80\xe5\x85\x54\x06\x86\x99\x94\x19\xc7\x84\x14\x2c\x21\x42\xc8\xe0\xf6\xc8\x99\x19\x5e\xd6\x62\xee\x9a\x9e\x66\x28\x4e\xf5\x96\x64\x19\xaa\x44\x16\x4e\xb4\x57\x6d\x30\xf0\xab\x70\x94\xa9\x82\x8e\x32\x62\x70\x4b\x76\x7e\x99\xce\x33\x14\xf3\x60\x65\x14\xac\x8c\x64\x81\x82\x14\x6c\x73\x56\xad\x1c\xc3\x15\xfc\x39\x00\x60\x62\x29\x2f\xdc\x27\x00\xc3\x0c\xc7\x0b\x18\x4e\xfc\x96\xe0\xd6\x6f\x09\xc6\x0f\xd3\xe1\xa5\x93\xd8\xf8\x02\xbb\x80\xe1\xe6\xd5\xe8\x6c\xf4\x2a\xdc\xa6\x52\x18\x42\x4d\x65\xc7\xfe\x09\x92\x5b\x53\x55\xad\x4f\xee\x26\x8f\x41\xd8\xfe\x95\x8a\x5f\xc0\xd0\x16\x86\xbe\x48\x92\x8c\x99\x55\xb9\xb0\x58\x27\xda\xcb\x9f\x52\x41\x4d\x12\xa0\x3d\x0d\xd0\x9e\x92\x82\x45\x36\xd0\x06\xe8\x02\x86\x24\xcd\x99\xf8\x31\x56\x1c\x31\x19\xe4\x3e\xd9\x7f\xdc\xff\xf0\xd9\xa0\x12\x84\xcf\x53\x49\x75\xe5\xe8\x5f\x75\x23\x45\x4d\x15\x73\x10\x5f\xc0\xf0\x56\x2a\x04\xb2\x90\xa5\x81\x43\x08\x7e\x1a\x00\x68\xba\xc2\x1c\xf5\x05\xbc\x7d\x7c\x7c\x98\x5d\x76\xef\xd8\x1b\x54\x0a\x5d\xba\x3b\xc3\x50\xf8\xf6\x11\xc9\xff\x6b\x29\x9c\x99\x42\xc9\xb4\xa4\x87\xd6\x3f\x5d\x0e\x06\x1a\xd5\x86\x51\xac\x1d\xf1\xfb\xb5\xf5\xcc\x38\xb7\xfa\x1b\xe6\x3a\x25\xa9\xf2\xd7\xad\xab\x82\xc2\x44\x21\x31\x58\xe9\x1d\xb5\x2e\x6f\x75\x76\x0c\x0a\x4d\xa9\x84\xee\x2c\xbd\xc7\x82\xef\x8e\xa3\x04\xa8\x33\xd4\x55\xc0\x88\x14\x6c\x64\x81\xae\xf2\xae\xf9\x2b\x4a\x03\x17\x30\x74\x35\xb2\x79\x5d\xa1\x3d\x6c\xc9\x2c\x64\xba\xb3\x42\xff\xdd\xdc\xfe\x14\x22\xdc\xda\x98\x42\xa3\x18\x6e\x7c\x9b\xd1\x86\x98\x52\xdb\xd6\x5c\xef\xd2\xb6\x10\x60\x46\xc3\xba\x5c\x20\x95\x62\xc9\x32\xd7\x85\xa8\x14\x02\xa9\x61\x1b\x66\x76\x35\x12\xbf\xa2\xa9\x61\x68\x3e\xb7\x31\x68\xee\x7f\x3b\x00\x19\xbe\x0c\x40\xef\x4e\x53\xe4\x68\xb0\x27\x7e\xd7\x6e\xa1\x76\xbc\x75\xd9\xf6\xbd\xb5\xf4\xed\xee\x07\x4f\xbe\x7a\x07\x75\xac\x08\x70\xa6\x8d\x8d\x53\x50\xd4\x3d\x21\x78\x67\x45\x8e\xda\xd7\x87\x42\x61\xd7\xfe\xee\x70\x24\xd6\xc7\xcf\xef\xa8\x54\xa2\x6a\x92\xae\xb5\xaa\xdc\x95\x66\x68\x0b\xa4\x60\x60\x2b\x33\x0a\xd7\xaf\x68\x02\x6b\x99\x46\xe2\x47\xcd\xed\xbd\x4d\x86\xfb\x7f\xdb\x06\x83\xbb\x3d\x7b\xfb\x34\x18\xe4\xa8\xb5\x9d\x72\xdd\x36\xd0\x34\x94\x3b\x92\x63\x45\x7f\xaa\x2a\x33\x12\x16\xd8\x74\x19\x4c\x9d\xb0\x25\x1b\x22\x73\x93\x01\xae\xe0\xf5\x65\x65\xe1\x71\x15\x64\xed\x28\xaf\xb8\x80\xc3\xc1\x49\xb4\x1e\xfd\x10\xe4\x66\x05\xd2\x46\xe9\x0a\xce\x2e\x0f\x7a\xeb\x80\x8a\x1a\xe0\x0a\x1d\x47\x91\xca\xd1\xc0\xd8\xed\x2d\xd1\xb1\xd3\x96\x77\x39\x86\x68\x89\x18\x6a\x33\xf0\x9d\x48\x72\x90\xeb\xbd\x0d\xa4\x68\x08\xe3\xba\x8b\x44\x50\x05\x85\xba\x90\x42\xa3\xdf\x91\x5f\x9c\x1a\xcc\x6b\xc1\xee\x16\x5a\x0d\xe7\x4b\xd0\xe6\x52\xae\x2d\xd1\x2b\x5e\xc4\x7a\xfc\x34\xb3\xe8\xa4\x28\x0c\x23\xdc\x17\xda\xf8\x69\x16\xdd\x02\xb2\xd5\xde\x9b\x4a\xe5\x63\xa9\x70\x5f\xc9\xde\x6d\xa9\x39\xb1\x2b\x38\xbf\xec\xf3\xb5\x0a\x95\x86\x23\xb2\xd5\x09\x59\xeb\x64\x93\x5b\xa2\x9a\xa0\xa1\xc7\xb1\xcb\x51\x50\xdf\x1c\x40\xa4\x13\xd1\xa9\x6e\xc1\xc1\x84\xef\xfe\x3b\x6d\x30\xdf\x8f\x59\x1c\x81\x6b\x17\xb4\x17\xe3\xd0\xed\x9f\x71\x22\x11\x63\x59\x74\xf4\xec\xff\xd2\x1e\x71\x23\x2d\x35\x30\x4a\xee\xfe\x53\x83\xb1\x3f\x3b\x1a\x60\x26\xb2\xe4\x69\x2b\x24\x0b\xac\xf0\x08\xad\xa0\xaf\x8a\x66\xf5\xb8\xb6\xaa\x71\xcd\x05\x67\xc2\x3c\x3f\x5c\x29\x61\x26\x34\x9e\x7c\x31\xc0\xaf\xbf\x15\xe0\xb3\x7f\x3e\xdb\xeb\x69\xf6\xad\x19\x1f\x94\xde\xf5\x0e\x59\x2c\x6c\xab\x4c\xfb\x7a\xd2\x3e\xd4\xb1\x50\xe3\xcc\x75\xa7\x21\xc5\xfb\x63\x69\xcb\x87\x9e\xf6\xd5\x53\x21\x0d\xa8\x71\x8d\xe9\x56\x7e\xf4\x68\xd7\xf9\x71\xde\xe7\x74\x54\xeb\xff\xde\xae\xf7\xe8\x47\x6c\xd5\xc8\x8a\xac\xda\x8f\x07\xcc\x45\xf2\xdd\xbc\x3a\x3c\x4d\x7b\x1b\x5b\x9d\x9f\xa7\x40\x4b\xa5\x50\x18\x1e\x86\x22\xf3\xb5\x23\x15\xe4\x84\xe8\xcf\x0e\xf8\x8a\x14\xc9\x25\xdc\x94\x0b\x54\x02\x0d\xb6\xb4\xd6\xdf\xeb\x79\x25\xe4\x70\x74\x8b\x52\xa0\x5c\xd6\x5e\xcc\x63\x4a\xd5\x90\x9a\xf0\x08\x5b\xed\xfb\xf4\x61\x8f\x42\x8c\x9f\x66\x6e\xbf\xbe\xf2\xcf\x2f\x0f\x48\xdd\x04\xa9\x50\xe8\x6f\x0e\xc8\x7d\xb8\x7d\x22\x0a\x9d\xa8\x2f\x71\xb8\x82\xff\xad\x08\xd4\x17\x70\x0e\xa6\xe1\xed\xb8\xa9\xd9\x15\xcb\x56\x73\xb2\x21\x8c\x93\x05\xe3\xcc\xec\xe0\x0a\xbe\x6b\x01\xb9\x24\x0b\xc5\x68\x18\xfa\xa5\xee\x70\x2b\x34\x5b\xa9\xd6\xf3\x20\x74\x05\xff\x73\x39\x38\x18\xfd\x0a\x88\x3f\x07\x1d\x10\xaf\x89\x21\x30\x41\x51\x65\xd5\xf8\x69\x66\x6f\xf9\x3b\x90\x12\x43\xe6\xd4\x7f\x8e\xa3\x1c\xb7\x47\xeb\x5d\xc9\xda\x53\xa0\xaf\xfb\xd2\xe8\x73\x5c\x3c\x3f\xfd\x7e\x0f\xcc\x60\x5e\xb7\xec\x07\x15\x12\xba\x54\x98\xda\x22\xb6\xac\x4c\xcb\x52\x51\x6c\x97\xcd\x54\x68\xe3\x5e\xde\x65\x4a\x96\x45\xa7\xc9\x8d\x9f\x66\xd5\xfa\xaf\x76\x19\x58\xb8\x9a\x7b\x69\x1f\xe7\x26\x6a\x8c\xae\xf6\xc0\xa8\xa0\x6c\x83\xd2\x4a\x46\xaf\xa8\x30\x73\x7c\xbc\xd4\xa7\x48\xb4\x39\x7d\x7d\x02\x68\xe8\xe8\xb8\x96\x0c\x31\x0b\x72\x35\x94\x2d\x23\xad\x54\xf8\x28\x05\xea\xc8\xe0\xe2\x04\xaa\xcf\x67\xd4\x7d\xde\xa2\xfd\x9c\x76\x9f\x54\x03\x10\x1e\x19\x5b\x9d\x7b\xab\x35\xfa\x4d\xd6\xfe\x22\x15\x6c\x57\x28\x40\xcb\xdc\xbd\x2b\x15\x99\x06\x9b\xe1\x84\x2b\x24\xe9\xce\x46\xcf\x5a\xed\xc2\xd2\x13\xab\xbd\x5a\xfd\xf0\x30\x01\x96\x9e\xc0\x82\x13\xb1\x76\x47\x70\xfb\xdf\xd0\x5b\xb4\x55\xef\xae\x77\xb2\x1c\x9e\xc0\x92\x71\x8e\x29\xb0\xa5\x7b\x7f\x6b\x1d\xb0\xe9\xf1\xe1\x61\xd2\x45\x72\x53\xd0\x79\xdc\xb5\x2b\x62\x81\xb4\x54\x16\x3f\x17\xe4\xae\x92\x0e\xab\x3e\x05\xbc\xfe\xd9\x65\xd7\xdf\xe9\xf8\x16\x94\xe4\xcd\x3b\xcb\xaa\x7e\x8f\x88\x12\xc7\x55\x41\x69\x2d\x29\x73\xad\x27\x4d\xbb\xcf\x61\x24\x9f\x5b\x0b\x73\xa2\x44\x93\xb1\x0d\xda\x9d\xc4\x85\x14\x97\x4c\xd8\xc3\xb1\xd9\x15\xe8\xde\x5b\x88\x32\x5f\xd8\x26\xb2\xac\xd3\x56\x77\xa1\x6f\x67\x77\x0b\xf5\xda\xbe\xb3\x77\x94\x7f\x37\xe2\x44\x65\x78\x20\x29\x9d\x50\x17\xc8\x5b\x26\x58\x5e\xe6\x7d\x8e\xc0\x51\x8a\x4b\x52\x72\xe3\x6a\xff\x23\x2a\xd9\x98\x64\xc2\x9c\x9f\x41\xce\xc4\xfc\x8f\x92\x08\xe3\x9b\x5a\x1b\xe2\x5b\xf2\xfc\x17\x2c\x93\xe7\xd8\xf2\x79\x74\x7c\x4d\x12\x4b\xa4\xe2\xe1\x33\x7e\x98\xc2\xcc\x1f\xc1\x23\xaa\xd5\x9c\xb5\xe1\xcf\xa0\xe7\x29\x97\x0d\x77\xa5\x5d\x51\xd2\x7d\xbd\x2e\x3d\x5b\x82\x2c\x50\xf9\x31\x65\x0f\x95\xf7\x37\x07\xce\x21\x95\xa9\x9e\x57\x00\x7b\x25\x63\x48\x06\xd2\x33\xbd\x8c\xd9\x13\x65\x21\x35\x33\x52\xed\xba\xb1\xcb\x98\x89\x26\xe9\xeb\xbd\x5c\x5e\x11\xbd\xaa\xb8\x88\xb5\x44\x65\x9e\x33\xd3\x67\xc5\xaf\xec\x45\xab\x67\x9c\x19\x85\xe8\xb6\x4a\x39\x12\xe1\xdb\x86\x9d\x00\xbd\x66\xad\xf0\xdc\x52\x1e\x6c\x0f\xe0\x24\xb1\xdd\xd6\xb1\x2c\x37\x3d\xba\xba\xee\xe6\x3c\xf5\x7a\x6f\x5a\x7a\x1f\x9a\x08\x67\xd2\x8d\x45\xcf\x94\xf2\x82\x71\xdc\xf3\x41\x46\xf8\x7c\xd7\xb2\x33\xf1\x1a\xaa\x99\xac\x91\x1e\xad\x16\xdd\x5c\x8d\xb4\x1e\x38\x31\x36\x72\xc0\x8c\x07\xc1\x0b\xa6\x2e\x7d\x12\x50\xa5\x70\xdf\x65\x44\x8c\xa4\xa2\xff\x95\x62\x0f\x69\xa8\xb6\x14\x25\x85\x5b\xea\xc9\x95\xb0\x9b\x16\x3d\xaa\xb8\x7b\x08\x3a\x6d\x4f\xe7\x32\x7c\xfd\xe2\x1a\xae\x2b\xf8\x68\x50\x47\x0d\x25\x1e\xd6\x2d\x9e\x30\x43\xaa\xd0\xdc\xe0\x6e\xea\x77\x69\xeb\x6a\x4c\x29\xea\x16\xa9\xd3\x4e\x6a\xbe\xc6\xdd\xbc\x43\xa7\x1b\x1b\x5e\xeb\x06\x77\xb5\x1d\x72\xc8\x8e\x5f\xb0\xe6\x5a\x6c\xc1\xda\x7a\xef\xa7\xe8\x61\x13\xf5\x98\xf5\xa7\x83\xbf\x84\x4b\xf7\x34\xd8\x66\x50\x45\xf1\x22\x26\xa4\x28\xfa\xc0\x78\x44\x41\x84\x79\x61\x03\xc6\x0b\x74\x37\xfe\x40\xb4\xde\x4a\x95\xbe\xa0\x59\x54\x22\x31\x51\x72\x01\x28\x17\xf5\x17\x18\x2f\xa1\x1f\x89\x79\xd7\xdf\x54\xa9\x65\xe1\x21\x94\xca\x52\x18\x5f\x75\xee\x9d\x59\xeb\xdd\x54\x98\xe9\x9e\x48\x37\xa4\xed\x08\x9f\x2f\x80\x4b\x92\xc2\x82\x70\xdb\xe9\xd5\x71\x07\x60\x6f\x62\xe6\xbf\xc9\x18\x87\x87\xb4\xa0\x9e\x70\x86\xc2\x4c\x53\x38\x22\x6b\x72\xe1\x90\xbf\x6e\x1d\xaf\xa9\x13\xe8\x83\xdb\xab\xfa\x0c\x0c\xea\x15\x4c\x7d\x16\x7c\x02\x56\x45\xd5\x4f\xa7\x6f\xfa\xe8\xb4\xdb\x75\x97\x43\x06\xdb\x5c\xd2\xaa\x56\xbf\x96\x46\x77\x33\xf0\x10\x91\xae\xbe\x07\xa8\xbe\x0f\x6a\xc5\xca\x1e\x20\x09\x5f\xdb\x7f\x7d\x74\x5c\x4c\xea\xd4\xef\x04\xa7\x79\x6e\x6f\x60\x82\x7b\xd5\xe5\x97\x73\xf2\x9b\xaf\xe1\xe4\x7f\x03\x4b\xea\x3e\x6f\x6f\xd0\x8a\xe8\xa4\xdf\x4b\x19\xdb\x07\xdc\xae\x67\x9e\x5f\xcd\x0c\x11\x29\x51\xe9\xfc\xfa\x6c\xbe\x39\x7b\x99\x65\x9d\xfd\x63\x2c\xeb\xfc\x1f\x63\x59\x6f\x3e\xf7\x25\x41\x74\x30\x6e\x55\xc4\xd7\x34\xdc\x38\x6e\xde\xde\x2d\xa1\x2b\x26\x62\xb3\x91\xe9\x52\xa3\x72\xb1\xb1\x06\x67\xb3\xb7\x71\x2b\x8b\x50\xaf\xc5\x7a\xc9\x91\xd4\xe6\x33\xfa\x4e\xa4\xef\x90\xe0\xbe\xcb\x3f\xa0\xeb\x41\x74\x12\xdd\xb0\x58\xd5\xf5\xf7\xba\xe6\x99\xf5\x11\x43\x0a\xa3\x24\xb7\x24\x41\xe0\x08\x1e\x57\x4c\x83\x90\xa9\xfb\xe9\x86\x14\x7c\x07\x04\x72\xe2\x5f\x22\x2c\xab\x9f\xb2\x20\x4f\xdd\x2f\x3b\x5c\x59\xa4\xa3\x3d\xbe\xd1\x4c\x82\x37\x5d\xbe\xe1\xed\x93\x1c\x75\x41\x6c\xa7\xe0\x52\x64\xcd\x4f\x60\x42\x48\x1c\x6e\xee\xe6\xe1\x17\x49\xa5\x60\x7f\x94\xc8\x77\xc0\x5c\xa0\x97\xde\x51\x41\x32\x4c\x0f\xbd\x39\xf3\x8f\x8c\xdb\x60\x08\xb4\x86\xad\x3b\x07\x5b\x46\xa5\x98\xc6\xbd\xf4\xa8\xfb\xc8\x7e\x7e\xe4\x95\x89\x30\xf3\xff\x15\x00\x00\xff\xff\xc6\x26\x74\x67\x8a\x24\x00\x00")

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

	info := bindataFileInfo{name: "api.proto", size: 9354, mode: os.FileMode(420), modTime: time.Unix(1537145376, 0)}
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
	"api.proto": {apiProto, map[string]*bintree{}},
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
