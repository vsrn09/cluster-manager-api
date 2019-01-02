// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// assets/generated/swagger/api.swagger.json
package swaggerjson

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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xff\x6f\xdb\xb8\x92\xff\xbd\x7f\xc5\xc0\x77\xc0\x65\x81\xc6\xee\xf6\xf0\x80\x87\x1e\x0e\xef\xf2\x92\xbe\x3e\xa3\x6d\x1a\xc4\xfd\x82\x87\xeb\xc2\xa0\xa5\xb1\xcd\x8d\x44\xaa\x24\x65\xaf\xf7\xd0\xff\xfd\xc0\x2f\xb2\x25\x59\xb4\x65\x29\x89\x95\xd6\x0b\x2c\x1a\x4b\xd4\x70\x66\x38\x24\x3f\x9c\x19\x92\xff\xf7\x0c\xa0\x27\x97\x64\x36\x43\xd1\x7b\x05\xbd\x97\xfd\x17\xbd\xe7\xfa\x19\x65\x53\xde\x7b\x05\xfa\x3d\x40\x4f\x51\x15\xa1\x7e\x7f\x19\xa5\x52\xa1\x80\xf7\x84\x91\x19\x0a\xb8\xb8\x19\x9a\xf2\x00\xbd\x05\x0a\x49\x39\xd3\xa5\x16\x2f\xfa\x19\x21\x80\x5e\xc0\x99\x22\x81\x5a\x53\x03\xe8\x31\x12\x1b\x72\x23\x12\xcb\x94\xcd\xe0\xf2\xfa\xf2\xa3\x2b\x0e\xd0\x4b\x45\xa4\x5f\xce\x95\x4a\xe4\xab\xc1\x60\x46\xd5\x3c\x9d\xf4\x03\x1e\x0f\xa4\x2d\x7f\x1e\xb0\x40\x0d\x02\xcb\xcb\x79\x6c\x79\x39\x27\x09\xdd\xd0\xc0\x98\x50\x43\x85\x84\x31\x65\xff\x93\xff\xb0\x4f\x79\xcf\x14\xfb\xfe\x0c\xe0\xbb\x91\x56\x06\x73\x8c\x51\xf6\x5e\xc1\xff\x5a\x9e\x4d\xdd\x99\x00\xfa\x87\xfe\xe2\x37\x53\x36\xe0\x4c\xa6\x85\xc2\x24\x49\x22\x1a\x10\x45\x39\x1b\xfc\x2e\x39\xdb\x94\x4d\x04\x0f\xd3\xa0\x66\x59\xa2\xe6\x72\xa3\xf2\x01\x49\xe8\x60\xf1\x6b\x26\x65\x5e\x7b\x33\xcc\x2b\x53\xb3\x9f\xc6\x31\x11\x2b\x2d\xee\x17\x1a\x45\x20\x50\x09\x8a\x0b\x04\x35\x47\x90\x8a\xa8\x54\x02\x9f\x02\x01\x47\x0c\x08\x0b\x81\x2a\x09\x77\xe9\x04\x03\xce\xa6\x74\x06\x53\x2e\x20\xe0\x8c\x61\xa0\xe8\x82\xaa\xd5\x5a\x95\x00\x3d\x9e\xa0\x30\x2c\x0f\x43\x5d\xc7\x1b\x54\xce\x0e\xf2\x85\x04\xca\x84\x33\x89\xb2\xc0\x1b\x40\xef\xe5\x8b\x17\xa5\x47\x00\xbd\x10\x65\x20\x68\xa2\x9c\xc5\xe4\x08\x59\x89\x74\x83\x90\xad\xcf\x00\x7a\xff\x2e\x70\xaa\xbf\xf8\xb7\x41\x88\x53\xca\xa8\xa6\x20\x33\x2d\x8d\x9d\x2d\x8c\x49\x42\x37\x5c\xde\x62\x12\xad\x7a\x05\x42\xdf\x9f\x55\xfd\xfd\x3d\x27\x4e\x42\x04\x89\x51\xa1\xd8\x34\x9e\xfd\xaf\x24\x48\x66\xca\xe6\xdf\xe7\x3b\x85\xbc\x26\x31\xea\x76\xd0\xad\x92\xb5\x84\xe2\x30\x41\x88\x38\xbf\xc3\x10\xd2\xa4\x5f\x26\x41\xcd\x97\xdf\x52\x14\xab\xf2\x2b\x81\xdf\x52\x2a\x50\x37\xc9\x94\x44\x12\x4b\xaf\xd5\x2a\x31\x8c\x49\x25\x28\x9b\xe5\xc5\xff\xfe\x7c\xbf\x38\x64\x29\xfb\x12\x03\x81\x6a\x7c\x87\xab\x31\x0d\xf7\xc8\xf6\x71\x8e\x30\x32\xe5\xdf\xe2\x6a\x18\x1a\x73\xba\xb8\x19\xc2\x45\x10\xa0\x94\x5d\x14\x8b\x18\xce\xb4\x74\xb5\x45\xb3\xc2\xbc\xc5\xd5\x5a\x3c\xd2\x3d\xf1\x04\xce\x34\xe3\xfb\x65\xba\x35\x05\x3b\x2b\xca\x9f\xa9\xc0\x3e\x49\x92\x7a\xb6\x77\x91\x24\x1d\xb6\x3a\x23\x8b\x42\x46\x98\xaa\x21\xcb\x47\x53\xb0\xdb\x0d\x93\x10\x29\x97\x5c\xd4\x69\x9a\x1b\x57\xb4\xdb\x02\xc9\x74\xb2\xe6\xbc\xe6\x70\x97\xfb\xa2\xab\xb2\x25\x82\x2f\x68\x58\x98\xa8\xab\xc4\xc9\xcf\x4c\xd9\x27\x12\xce\xc8\x52\x0e\x8c\x76\x06\x8b\x78\x49\x04\x0e\x50\x05\xbf\x3c\x9c\x6c\xa5\xb7\xc8\xd2\xb8\x34\xf7\x9a\xe7\x29\x33\x33\x3f\x96\xdb\x08\xcc\xf8\x57\xf1\x50\x4b\xb0\xfd\xd8\x8a\x54\x84\x05\xbf\x6d\xa9\x69\x4a\xd2\x48\xa3\xad\x5c\xad\x95\xd0\x21\xf7\x65\x4f\x91\x59\x19\x34\x64\xc0\x79\xf3\xf1\x6f\xee\xaf\x75\x2b\xf6\x42\x8c\x50\xe1\x6e\x64\x67\xcb\x6c\x90\xdc\x0e\x94\x76\x65\x8a\x3e\x01\xa0\x56\x60\xb4\x2b\x58\xed\xcb\x9c\x28\xa0\x32\x8f\xd5\xfe\x43\x82\xfe\x50\x43\xb6\x10\xa5\x12\x7c\xd5\x99\x6e\x7e\x42\x6b\x27\xb4\x76\x5c\x51\x4e\x68\xad\xcb\x0d\x73\x42\x6b\x27\xb4\x76\xdf\xb2\x3d\x41\xb4\x56\xa3\x09\x02\x12\x45\x13\x12\xdc\xf5\x53\x11\xd5\xb0\xaa\x4f\xb7\xef\x34\x1e\xd0\x5f\x81\xfe\x0c\x14\xef\x8c\x35\xad\x45\xd1\xb5\xa0\x54\xf5\xfa\xc9\xf0\x2a\x33\x2d\xf7\xd9\xe3\xc8\x73\x9f\x40\x3a\xe1\x72\x8f\x83\xd4\xf4\x1a\xa9\x47\x82\x3a\x48\xfa\x52\x20\x79\x12\x48\xba\xc0\xe8\xa3\x20\xe9\x09\x0f\xb7\x8c\xc0\xda\x47\xd5\x9b\x9c\x79\x28\x91\x96\xad\xe3\xbe\x15\xf0\x5e\xce\xea\x88\x7f\x1f\x06\x97\xee\xb1\x37\x12\xfe\x9e\x4a\x05\xe4\x40\xc3\xbb\x30\x9f\x39\x06\xae\x79\x88\xb2\xcb\xd6\x57\xe0\xf6\x67\xb4\xbe\x82\x02\x1e\xdc\xfa\x9e\xe5\xb4\x57\x8e\x16\x0d\x22\x5a\x18\x02\x0f\x08\x19\x11\xd0\xdf\xea\x39\xc0\xd1\x92\xb5\x22\x41\xef\x74\x85\x1d\x36\xce\x22\xa7\x8f\x62\x9d\xa7\x45\xf9\x69\x51\x7e\x5c\x51\x4e\x8b\xf2\x2e\x37\xcc\x69\x51\x7e\x5a\x94\xdf\xb7\x6c\x4f\x70\x51\xfe\xd0\x50\x28\x4d\x66\x82\x84\x78\x28\x1a\x4a\x05\x03\xf7\x29\x70\xd3\xca\xd2\x58\x28\x81\x19\x5d\x20\xab\x81\xde\xdf\xa0\xfa\x64\x09\x38\xce\x87\x6c\xca\x45\x6c\x4a\x74\x1c\x2a\x79\xf9\x7e\x14\xe0\x74\xea\x71\x47\x77\x83\x9d\xb0\xeb\x09\xbb\x1e\x57\x94\x13\x76\xed\x72\xc3\x9c\xb0\xeb\x93\xc0\xae\xcd\x12\x1d\x40\xe9\x67\x4b\x04\x22\x10\x02\xce\xa4\x9e\x59\x29\xb3\x39\xc2\x0e\x13\x3d\xc1\xa8\xc0\x5e\x27\xad\x52\x18\x27\x0a\x14\x5f\x03\xbf\x3a\x4e\xda\x22\x56\xea\x32\xb0\x2b\x72\xfa\x33\x7a\x68\x8b\x1a\x38\x92\x8b\x56\x60\x88\x4c\x51\x12\xc9\x81\x86\x7e\xb9\x75\xc9\x5e\x13\x4d\x93\x90\x28\x04\xb2\x94\x90\x23\x03\xa9\xc4\xd0\x2d\x4e\xea\xd8\xab\xa6\x71\xf1\x65\x74\xb9\xa1\xd0\x6d\xab\xdd\xe6\xf7\xe7\xb4\xdd\x6d\x3d\x74\xc0\x82\xcd\x3a\xa5\x89\x0d\xeb\x0f\x5b\x5b\xb1\x26\xf2\xb4\xec\xb8\xc4\xf1\x4f\x6c\xc9\x25\x4d\x1c\xc7\x96\xe7\x18\xc5\x79\xf3\x3d\x20\x09\x37\xc4\x24\xe2\x2b\x0c\x41\xd3\x80\x60\x4e\x84\xda\x61\xb0\x36\xcb\xf5\x9f\x18\xc5\x97\xe5\x92\x5d\xb3\xd4\x12\xab\x8f\x62\xa2\xdb\x3d\xbe\x4a\xd8\x6c\xfb\xa1\xa2\x51\x84\x02\xe4\x9c\xa7\x51\x08\x13\x04\xca\xa4\x22\x51\x84\x21\x70\xd6\x19\xf4\x9d\x90\xe0\x8e\xcc\x30\x53\x6d\x1d\x20\xbe\x25\x98\x1e\x14\x3b\x23\x50\xd9\xc6\x2b\x9b\x48\x97\x81\x6b\x12\x3f\xd2\x12\xe1\xe4\xcc\x3a\x39\xb3\x4e\xce\xac\xc7\x96\xe5\xe4\xcc\xea\xb2\x40\x3f\xa6\x33\xeb\x14\x16\x3a\x7a\x58\xa8\x76\x26\xf1\x05\xa4\x8c\x7e\x4b\x11\x68\x08\x8a\x03\x65\x21\x0d\xf4\xaa\x4f\xcd\xa9\x7c\xdc\xc4\xe2\x3a\xb0\xe6\x94\xf3\xfd\x13\xe4\x7c\xbb\x25\x02\x9c\x71\x01\x02\xdd\xaf\x5f\x72\x8b\xb7\xaf\xec\xa3\xb6\xce\xa5\x2e\x3c\x41\xbb\xc7\x2f\x0d\x14\x5d\x20\x50\x73\x78\x86\x41\xb6\x73\x22\x81\x44\x02\x49\xb8\x82\x09\x22\xdb\x2c\x03\x97\x54\xcd\xed\x89\x1b\x7a\x04\x2a\xb9\xde\xcb\xcb\xc1\xa1\xad\xff\x49\xac\x07\xcb\xbc\xfe\x8c\x3e\x8b\xb2\x0e\x8e\xe3\xad\xb0\xcb\xc3\x82\xbb\xad\xa9\xdd\x3b\x52\x7e\x0b\xbd\xc9\x52\xd5\x3f\x6e\x95\xec\x9a\x81\x96\x58\xfd\x19\xed\xb3\xa4\x82\xe3\x98\xe7\xe6\x08\xa8\x83\x53\xad\xdc\xa7\x40\x37\xb9\x46\x40\x26\x3c\x55\x40\x12\x0a\x12\xc5\x62\x5f\xae\xd5\x67\x4b\xe1\xe9\x24\x59\x39\x86\x1b\x59\x6b\x93\xc6\x5a\x9f\x76\x95\x63\x6d\x73\xde\x54\x79\xf3\x42\xe1\xf7\xc5\xdb\xd1\x28\xc1\x20\xdf\xb0\xd9\xbc\xce\x27\xbf\x63\xb0\x99\xbe\x34\x4a\x4f\x50\x28\x5a\xd2\x74\x2f\xe7\xe9\x2f\x37\x41\xfd\x0d\x16\x7e\x77\x7f\xe1\x90\xb2\x5c\x50\x41\x71\x98\xa4\x34\x0a\xf3\xf1\xed\x5e\xa5\x4a\x19\x0f\x71\x9c\x70\x1e\x95\xd9\xdb\x81\xec\x37\x75\x5e\xf3\x10\xe1\x86\xf3\xc8\xb8\xbe\xaa\x6b\x20\x61\x38\x0e\x78\xca\x94\xaf\x06\xca\x14\x96\x1d\x85\x3d\x6b\xce\xee\xf5\x7f\xbe\xf4\xd4\x9f\xc6\x13\x7b\x9a\x14\x65\x81\x40\x22\x11\x26\xab\x6a\x2e\x04\xc6\x7c\x81\x0f\xc6\x08\x5b\x73\x12\x62\x15\x27\x95\x03\xc8\x4e\xe3\xfb\xfc\xfe\x0b\x11\xd8\xd6\xfe\xb4\xf2\x99\xd9\x42\xe5\x91\x99\x08\x41\x8a\xe3\x76\x8f\x2a\x8c\xcb\xe5\x0f\xb1\x57\xcb\xf9\x7b\x12\xcc\x29\xb3\x02\xf8\xe0\xfc\x46\x7f\xae\xb4\x84\xe5\x9c\x06\x73\x58\x22\x2c\x09\x33\x49\x09\x24\x34\x6b\xa9\xbd\x86\xec\x1a\xf8\x41\x85\x2d\x37\x98\x95\xf4\xd6\xd4\x7c\x5f\xf2\x5a\x39\x60\x2a\x78\xec\x11\xba\x96\x2d\xf9\x58\x6b\x61\x4b\xf3\x32\xe8\xaa\x3b\x4a\xe8\xc5\x96\xfe\xd8\xb8\x58\x46\xa3\x7f\x3a\x17\xcb\x3e\x91\x0a\xdb\x1a\xdd\x58\x7c\xf1\x76\x64\x30\x29\x0b\xf0\x8d\xe0\x69\xd2\x46\x1e\x87\x69\x9a\xc9\xc3\x72\xfe\x96\x99\xe1\xa4\x7a\xc6\xb2\xf4\x0e\xaf\x23\x93\x12\x74\x69\x38\x1b\x29\xc2\x42\x22\xc2\xf1\xd5\xcb\xf1\xe2\xe5\x73\x40\x15\xf4\x7f\xa9\xae\x32\xa6\x6c\xfc\x2d\x25\x4c\x51\xb5\x7a\x80\x91\xee\x3d\x65\x34\x4e\x63\x70\x23\x1e\x9f\x5a\xa8\xcd\x02\x94\x70\xe6\xbc\x2d\x66\x02\xfa\x13\x05\xf7\xb1\x48\xfe\x78\x50\x16\xc9\x1f\xcd\x58\x7c\x56\x62\xb5\xa2\x39\x4c\x5b\x4b\xb0\x9e\x24\x20\xb6\x79\x08\x0b\xab\x2a\xeb\xed\x30\xe7\x2f\xd6\x9c\xbf\x8c\xae\x88\x22\x97\xc8\x4a\x87\x5d\x1e\x6a\xcb\x2e\x64\xd0\xc4\xd2\xbe\x98\x11\xc8\x12\x80\xb3\x54\x9e\x23\x91\xea\xfc\xd7\x9d\x36\x46\x16\x84\x46\x64\x42\x23\xaa\x56\xe3\x3f\x39\xbb\x8f\x21\xb7\xa6\x0b\xa8\xcc\x78\x9e\x15\x30\xac\xe4\x84\x98\x3c\x87\xec\xef\x97\x81\xf9\x7b\x89\xfa\xef\x70\x5b\x3a\x7f\xdb\xdb\x7a\x74\x3b\x81\x6b\xa8\x3a\xed\x7a\x6f\xc3\xd4\x7d\x0d\x21\xf1\x5f\xfa\x11\x11\x33\x3c\x0d\x1e\x3f\xce\xe0\x71\x23\xdc\x41\xb6\xa9\xc0\x70\x58\xea\x57\x07\x5b\xda\x22\x09\xc6\x34\x6c\x3c\x25\x7e\xbe\xb9\x04\x1a\x3e\x87\x49\x44\xd8\x9d\x99\xeb\xf5\xff\x5f\xf5\x1a\x88\x28\x04\xce\xd0\x3c\x58\xf1\xf4\x6b\xef\x39\x4c\xa9\x49\x5c\xa0\x53\xfd\xc0\xe4\xdc\xfe\xfd\x5f\x1f\x34\x8d\xea\x56\x97\x18\xa4\x42\x8f\x36\x46\x87\x4d\xd9\x1c\x39\x2a\xbb\xa6\x6c\x4a\xe2\xb1\xe0\x11\x8e\x89\x68\x36\xa0\x1a\xdf\xf2\xc5\x7b\xd0\x44\x8c\xc4\xf9\x24\xe3\x33\x22\xd8\x2f\x59\x3b\x4a\xc9\x03\x6a\x16\xed\x61\x58\xcb\x92\xfe\xc1\x05\x2c\xe7\xc8\x40\xf2\xd8\x84\x15\xd8\x4c\x1a\xdd\x65\xbe\x59\xab\xeb\xb0\x68\x3c\x6f\x90\xa1\xa0\xc1\xda\x8b\x97\xb9\x33\xb9\xa4\x8a\x8b\x55\x1b\x93\xb1\x87\x56\x37\x99\x74\x36\xf9\xd7\x9f\x6e\xdf\x6d\xd4\x64\xbc\xcd\x02\x13\xee\x59\xae\x36\xc5\x6c\xf9\x0a\x35\x79\xeb\xa9\xde\x83\x40\x4b\xee\x8a\x0a\x5f\x4b\x0b\xd5\xcd\xa8\x1a\x6f\xbb\x8e\x0e\x33\x33\x45\x66\xc0\x99\x05\xa1\xd4\xea\xcd\xb5\x69\xa5\xf6\x74\x95\x01\x8f\x63\xda\x02\xc7\x13\x39\x5f\xe3\x5e\xaa\xc0\x91\xf3\x56\xa7\x04\xe2\x58\x2a\xa2\x9a\x36\x1b\xaa\xb9\x1e\x19\x05\x30\xae\x4c\xad\x9a\x22\x2c\x89\x84\x20\x42\xc2\x6c\x77\x98\xa4\x34\xf2\x30\x61\xfc\x20\xe3\xb0\x29\x03\x57\x66\xe8\x9a\x5a\x77\x8a\x47\x4c\xde\xaa\x1d\x9d\x55\xe9\x4a\x66\xdc\xa6\x83\x2a\xae\xf5\x9a\xd0\xc8\xe3\x52\x71\x2f\x45\xa3\xfa\x2e\xdd\xc7\xa6\xaa\x6a\xfa\x49\x44\x94\xb6\xf1\x46\xf4\x6f\xdc\xc7\x40\x95\x6d\x26\x5b\x9f\x4d\x1e\x18\x80\x48\x19\xa3\x4c\x9b\xed\xbe\xde\x57\xe5\x0a\x2b\xe6\x6f\xb7\xe8\x7d\xc5\xcc\xa8\xa6\xbd\xc1\x9f\x2f\xe5\x9d\xc5\x4a\x89\x4b\xed\x6a\xf6\xa5\x33\xf9\x5c\x24\x8d\x17\x09\xde\x6c\xa3\x5a\xd3\x96\xfe\x3a\x28\xba\x27\x53\x69\xe7\x46\x33\x61\x69\x7b\xd8\xf2\x75\xec\x32\x83\xf2\x91\x33\xc7\xf0\x02\x34\xdc\x42\xe4\xe9\x71\x59\x96\x47\x53\x27\xf1\x4d\x55\x9a\x48\xce\x51\x7b\x40\x7a\x88\x67\xc9\xb7\x6c\xe1\xc1\xf6\x6d\xbb\xc8\xb3\x78\xf1\x65\x04\xf9\x52\xd5\x5c\x94\x12\xef\x77\xf0\x51\xcb\xb5\xef\xb7\x76\xcd\x8e\x4c\x30\xa0\x53\x77\x71\xc5\x57\x56\xa4\x60\x41\xb8\xd9\x95\xf2\xdf\xf0\x97\xff\xaa\xe6\xd7\xa5\xa8\xb4\x67\x38\xe7\x0e\xae\x9e\x12\x5c\x0a\x44\xe3\x36\xba\xcc\x08\x78\x66\x0d\xf7\x3a\x1f\xaa\x6a\x32\x80\x6f\x9f\x96\xd5\xa2\xef\xf2\x2d\x71\x33\x0a\x13\xce\x35\x3a\xf0\x2d\x00\x2b\x5f\x6f\x7a\x36\x91\x36\x65\x87\x80\x4c\xcd\x28\x37\x4d\xa3\x2c\x49\xa4\x89\xd0\xe5\x00\x4e\x1b\x87\x7e\xd2\x78\xf1\xe3\x4b\xb2\xf4\x78\x2e\x6d\xae\x63\xd3\x8a\x2a\x33\x20\x3d\x63\x5f\x96\x88\xd8\xb4\x2e\x4f\x7a\xa2\x67\x0a\x2e\x65\x09\x36\x9e\x80\xfd\xb9\x83\x47\x99\x13\x2f\xb7\x07\x80\x47\x5b\xcd\x79\xf2\xc3\x7c\x28\x64\x9d\xa3\xd5\x78\x75\x5d\xce\xdc\x6a\xd0\x29\xdd\x18\x74\x85\x8a\xd0\x68\xa8\x30\x6e\xa3\xb8\x86\xb2\x6c\xe4\xd8\x1d\xa3\x6d\x0a\x51\x2a\xae\xde\xf1\x74\x0a\x73\x5b\xd2\x38\x46\x29\xc9\xac\x59\x5d\x17\x61\x68\x26\x17\x12\x55\x64\x33\x14\x6f\x64\xda\xcb\xce\xe6\x82\xa6\xd6\xc8\x2c\x77\xd7\x93\x59\x55\x99\xab\x9e\x6a\x05\x14\x27\x48\x04\x0a\xc5\xef\xb0\xa9\x5f\x7d\xc3\x85\xb9\x8e\x0b\x2c\x45\xb0\x24\x77\xb4\x43\xf3\x39\xdc\x3e\x1a\x59\x2a\xfe\xfe\xb3\xaf\x29\x0e\xec\x43\xa7\xde\xd3\xa9\xde\xf3\x34\x8d\x68\x54\xe6\xda\xa7\xb8\xed\xac\xf2\xde\xe8\xe3\xc5\xc7\x4f\xa3\xf1\xa7\xeb\xd1\xcd\xeb\xcb\xe1\x3f\x86\xaf\xaf\xf2\x49\x50\x37\xb7\x1f\x3e\x0f\x47\xc3\x0f\xd7\xc3\xeb\x37\xf9\xe7\xb7\x9f\xae\xb7\x1e\xbd\xbe\xfc\x70\x7d\x39\x7c\x57\x7a\x3c\xfa\xf8\xe1\xe6\xa6\xf4\xec\xf5\xed\xed\x87\xdb\xfc\x83\xab\xd7\x6f\x6e\x2f\xae\x5e\x5f\x65\x2a\x58\x67\x99\xe5\x73\xd0\x77\x70\xba\xd1\xee\x39\x6c\x17\x7b\x05\xd7\x5c\x81\x44\xf5\x95\xc1\x39\xe4\x45\x7a\x05\x06\x03\xe5\x9e\x98\xa6\xc1\x75\xfa\x79\x71\x89\x4a\x25\x4c\x50\x83\x0b\xe7\x2a\xee\x1b\x82\x4e\x17\x96\x96\xfb\xb1\x93\xcc\x9c\x48\x9b\x08\xec\xc8\xd8\xcb\xf5\x24\x4c\xd3\x28\x5a\x41\x2a\xc9\x24\x42\x47\x7a\xa3\x53\x47\x7e\xf3\xa0\xa2\x0a\xa2\xac\x73\x7b\xc9\xc5\x9d\x26\x48\x4c\x2e\x72\xb4\x72\x5c\x87\x9c\x61\xe6\xf2\x74\xbc\x3c\xd7\x28\x7d\x0e\x44\xba\x25\x76\x86\x9c\x62\x62\x38\x35\xae\xc3\x10\x41\xf2\xa9\xd2\x8b\x31\xcb\x55\xd6\xa4\x96\xa5\xec\x57\x3d\xcd\xd9\x9d\xb0\xa1\xa1\x63\xcc\xc0\x12\x31\x7f\xee\xa4\x10\x93\x95\xd9\xdd\xc8\xac\x7e\x0c\x81\xcc\x6c\x2c\x8d\xec\xd7\x4e\x32\x2e\xad\xd2\xec\x1d\x17\x46\x41\x5a\x21\x1c\x04\x4a\xc5\x05\x9a\x36\x80\x69\xca\x02\x3b\xa2\x50\xb5\xda\x8f\x1e\x2b\x72\x2d\xda\x0c\xe9\x11\x0f\xca\x0e\xf2\xfa\x83\xa1\x59\xa8\x98\x8d\xf2\x5b\x41\xd7\xf2\x58\xd7\xf9\xfc\xba\x2c\x9e\x67\xe3\x56\x0f\x96\x99\x54\x2b\x57\x26\x47\xc8\x13\x52\x2f\x45\x2a\x9b\x8c\xe6\x15\x81\xca\x36\x86\x14\x12\x45\xc6\x41\x39\x3f\xa2\xb6\x26\xaa\xd2\x2c\x76\xfb\x7c\x1e\xc7\xe4\xea\x38\xc4\x9a\x19\x9c\x40\xc9\x53\x11\x6c\x67\x64\x1c\xa4\xae\x8a\xc0\x72\x35\x93\x7f\xff\xd7\x07\xb0\x96\xda\x1d\xf3\xf7\xe4\x60\x3c\xba\xf9\x1f\xcb\x33\xed\xbf\x71\x75\x7d\x10\xbf\x37\xf2\xd3\xd6\x0f\x5d\x50\x40\xe6\x94\xde\xe3\x6a\xcd\x2a\x2d\xfa\x5b\x9f\xb4\x9f\xd3\xaf\x86\x63\x87\x2a\xd6\xca\x3e\x87\x20\x15\x02\x99\x8a\x56\xd6\xdb\xa9\xc1\xd6\x52\x6a\xbc\x14\x13\xe2\xe9\xcf\x77\x7f\x95\xad\x83\xd6\x8b\x4d\xc0\xf3\x6d\x3a\x41\xc1\x50\xa1\xcf\xdd\xdf\x22\xe8\x50\x39\x13\x1d\xe0\xef\x6f\x1d\x80\xa8\x0b\xb3\x2a\xeb\x39\x24\x70\xb0\xaf\xa2\x7d\xb1\x83\x39\x9d\xcd\xc7\xf9\x3c\xba\x87\xf1\xaa\x6f\x87\xf2\x73\xc8\xfa\x9f\x17\x9e\xd5\x3b\x2a\xbd\x12\x18\x4f\xc9\x44\xd0\xa0\xb1\xcd\xd9\xcf\xdd\x28\x58\x0a\x7c\x37\xea\xd3\x9d\x8d\x5d\xf8\xb5\xbc\x24\x32\x3f\x01\xc0\x64\x55\xd8\x98\xec\x19\x6f\x1d\xd6\x68\xe9\x4d\x30\x2e\x22\x7f\xf3\x84\xc6\x0b\x5b\x76\x26\x64\x9c\x41\xb6\x8d\xa9\x6d\xa3\xdd\xcf\x9e\x8a\x80\x33\x25\x78\x34\x4e\x22\xc2\x1e\x78\xc3\xc1\xc3\xec\xae\x08\x78\x9c\x08\x2a\x71\x3f\x98\xd4\x5d\x0f\xc5\x8f\x2d\x23\x49\xe8\x18\x59\x98\x70\xda\x38\xdc\x45\x65\xee\xec\x22\x02\x0b\x12\xa5\x08\x11\xbd\x43\xa0\xc9\xab\x84\x0b\xe5\x12\x03\xdc\x1e\x6b\x02\x0b\x2a\x54\x4a\x22\x18\xde\x0c\xf4\xeb\xaf\xec\x86\x48\xa9\xbb\xa4\x0d\x9c\x00\xfe\xa1\x50\x30\x12\x41\x90\x4a\xc5\x63\x14\xd2\x0d\x5d\x64\x12\xa1\x4b\x0f\x8a\x53\xe6\x4e\x17\xf0\x39\xb3\xeb\xf6\x91\x8a\xab\x68\xbb\x37\xb0\x5d\x1a\xfd\xe6\x47\x87\x6c\xb7\x3a\x5f\xf9\xd0\x6c\xb5\x27\xb4\x5e\xa6\x68\xc1\xbb\xd9\x3c\xcc\x54\x79\xaa\x58\xf7\xb4\x6b\x43\xde\x08\xeb\x0d\xa6\xb9\xc8\xb7\x27\x0d\xbb\x85\xb7\xdb\xe0\x51\x47\x40\xc2\x12\x05\xda\xeb\x05\x1a\xe8\xb7\x9c\xdb\x7a\x74\x7c\x9d\xdf\x18\x64\x0f\x49\x88\x91\x79\x26\x57\x5d\x56\x26\x24\xb8\x9f\x5a\x0d\x25\xbb\x05\x51\x57\x6b\x03\x4f\x49\x12\x39\x3c\xbb\x23\x20\x9b\xf0\x7a\xf3\xfb\xbe\x44\xe2\xfd\x7c\x6e\xf2\x7b\x77\xe5\xa9\x06\xa5\x96\x6c\xa6\x13\x5b\x17\xf3\x6e\x4c\x6d\xb3\x8e\xd9\xae\x27\xa3\x56\x5d\x95\x9e\x11\x9a\x0d\x44\xa6\x26\x22\xd0\x75\x4f\x7b\x25\xc7\x19\x43\xa9\x30\x84\x15\x89\x23\x38\x37\xaf\x3e\x9b\x2a\xfa\xe6\x89\x46\x28\xc8\x94\xf4\xa4\x72\x19\x86\xc7\x09\x59\x45\x9c\x1c\x14\x7a\xcb\x8d\x26\x2b\x85\xbb\x56\xf4\x96\xb6\xee\x07\xd9\x09\x22\x67\xe6\x8c\x53\x2a\x66\x24\x49\x30\x04\xc9\xa3\xd4\xc8\x62\xb6\xa3\x54\xed\xa3\xd8\xdf\xf3\xd5\xd6\xdd\x66\xdd\x1b\x58\x87\xa5\x18\x82\x8d\x5c\xc8\x95\xd4\x80\x78\x17\xe2\x7e\x4c\xbc\x95\x07\xe9\xfb\x91\xd6\xbb\xf2\x4d\x79\x6d\xda\xee\x87\x6b\xb7\xb6\x2b\xa5\x5c\x42\x4a\x23\xbd\xee\xb9\xc4\xa6\x7b\x9a\xbe\x24\xac\x0c\xeb\x5c\xaa\xac\x07\xd5\xb9\x91\xf6\x38\x5b\x0e\xc9\x7a\xa4\xb7\x1b\x6f\x12\x2e\x25\xd5\xa0\x5c\xd0\xd9\x5c\x01\xe3\xcb\x66\xcd\x56\x38\x16\xa3\x7b\x8d\x34\x9c\xe6\xa0\xe1\x92\x48\xf8\xf0\x76\x67\xe3\x8c\x69\xe5\x26\x19\xd8\x85\x2d\xf6\xed\xb5\xa9\xe6\x2c\xdb\x3e\x31\xdc\xe9\xaf\xdd\x0e\xb2\x9b\x5a\x4c\xba\x5d\xe6\x8f\x2c\x80\xfc\x1d\xad\x55\x75\x28\x52\x1b\x77\x42\xf5\xb8\x51\x6f\x2f\x47\xe6\xdd\xdf\x75\x32\xb0\x2f\xeb\x72\x7d\x4c\xef\xac\x61\xed\xfe\xb3\x7b\x0f\x80\x75\x07\x9c\x1d\x53\x42\xfa\x1e\x95\x18\xb4\x31\x42\xa5\x28\x9b\x3d\x80\x5b\xf9\x28\xb9\xec\xcd\xa3\xd6\x36\x8e\xbe\x97\x97\xee\xef\x3f\x78\x22\x81\x9f\xea\x53\xe3\xba\x37\xa0\x3f\xd9\x05\xff\x26\x64\xf4\x8e\x4c\x30\x3a\xca\x7a\x3f\x7f\x00\x08\x81\xc8\xf0\xe1\x5f\xf7\x35\x0f\x94\x19\x3f\x62\x65\x1d\x75\xb5\x75\xb3\xdd\xaf\x0f\xc8\xab\xab\x3a\xa7\xb5\x74\x42\xeb\xd6\xd9\xac\xa5\x53\x59\x2b\xb3\xe0\x4a\x27\xb1\xee\xe3\xbf\x78\xc0\xdb\xb1\x23\xea\x6e\xc2\x1b\x80\x9b\x41\x21\x3b\xe8\xfe\x90\x55\xc1\x43\xcf\xee\xed\x3c\x4b\xd7\x1b\x6f\x92\xbf\x6a\xea\xa9\x9a\xc6\x4d\x47\x09\xf3\x25\x9c\xd1\xe9\x3a\x74\xe5\x6c\xe6\x17\xbb\x3f\xdc\xf2\x72\xf6\x2d\x25\xab\x3e\xe5\x03\xc9\x63\x9c\xa5\x2b\x77\x3c\xa4\x67\xda\xb8\xa7\x9d\xb9\xb6\x0e\x73\x6f\x86\x3d\x6e\xcc\x9e\x2e\x39\x58\xdf\x99\xa5\x38\x9c\x2d\x5e\xf6\x7f\xfd\xb5\xff\x62\xcb\xa3\xb1\x6d\x0e\xe3\x25\x0d\xbd\x3a\xba\x8f\x75\xab\xd3\xd5\xfa\x4e\x99\x73\x5d\xa1\x7b\xfa\xb7\xaf\x6c\x64\xdb\x93\x2a\x98\x93\x05\x16\x0b\xd9\x5c\xf5\x44\xd0\x05\x8d\x70\x86\xf2\x6f\x1e\xf4\xa2\x8b\x8d\xd7\x76\x76\x8c\xd5\xd8\xda\x4c\x5d\xde\x28\x55\xf9\x30\x8c\x8b\x94\x58\x71\xbc\xe1\xfc\x13\x12\x3c\x21\xc1\x12\x87\xf7\x84\x04\x2b\x8f\x67\x3d\x01\xc1\x7b\x03\x82\xbe\x9b\xc1\x8e\x81\x0f\xb2\x99\xda\x1f\x6c\xe8\x62\xd6\x68\x1b\x4d\x77\xdb\x9c\xef\x6f\xef\xae\xff\xd6\xae\x1f\xd8\xce\xee\x3d\x21\xbe\x9d\xb6\x7f\x1e\x5b\x2b\xdf\xd3\xd9\x42\xe6\xee\xcf\xe3\x27\xec\x55\xe6\xa5\x1b\x67\x91\xb4\x8e\x8b\x67\x0e\x75\x5d\xf3\x12\xf3\xbb\xb2\xb8\x67\x31\xf1\x44\x20\x5d\xd5\x55\xc2\x3f\xfe\xb0\xb4\x9d\x8d\xd7\x42\xe6\x54\xa2\x68\xe5\x77\xcb\x08\xf8\x0f\x13\xce\xd9\xd5\x03\x1e\x59\x9c\x5f\x32\x71\x7f\xb6\x4a\x8b\x93\x28\x4d\x22\x05\x17\x35\x79\x68\x70\x88\x46\xc5\x9d\x38\x77\x7f\x95\xeb\xee\xbb\x3e\x93\xcf\xe6\xbb\x82\xc9\x77\xed\x83\x49\x77\x34\x3b\x2a\xa9\x04\xce\xa2\x15\x90\x6c\xab\x25\x9d\x5a\xc3\x9b\x52\x8c\xcc\x6e\x50\xe7\xe6\xeb\x57\xf3\x6c\x9c\x9a\x8f\x99\xe7\x50\xf6\x1d\xe7\xc8\x78\x5c\x0c\x5a\x27\x96\xcd\xb5\x3a\x62\x97\x68\x5a\xfb\x94\x8f\xc2\xae\x02\x77\xd5\x6f\xf6\xcc\x28\xd2\x79\x41\xd7\x37\x15\x64\xd9\x9f\x57\x3c\xc8\x5d\x55\x50\x6a\xad\xf7\x5c\xa0\xdb\x1f\x9e\xa1\xc2\xf7\x56\x52\xb8\xb8\x19\x3a\xbd\xb9\x13\x3d\x7a\x73\xa5\x12\xf9\x6a\x30\x98\x51\x35\x4f\x27\xfd\x80\xc7\x03\x49\x62\x99\xb2\xd9\x79\xc0\x02\x95\xe9\xea\xdc\xe9\xea\x9c\x24\x54\x33\xf5\xfd\xd9\xf7\x67\xff\x1f\x00\x00\xff\xff\xe9\xb1\xb3\x5e\x2c\xad\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 44332, mode: os.FileMode(420), modTime: time.Unix(1546461269, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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
