// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 15, 13, 25, 19, 480256713, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 12, 15, 13, 26, 52, 438845483, time.UTC),
			uncompressedSize: 28115,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x5b\x73\xe3\xb6\xb9\xef\xe7\x57\xd0\x7e\xf0\x10\x11\xad\x23\xd2\x5a\xc5\x11\x05\xed\x6c\x1a\xbb\x75\x67\xeb\xcd\xac\x9d\x74\x5a\x8f\x47\x03\x8b\x90\x8d\x53\x9a\x74\x41\x68\x5d\x57\xd6\xf9\xed\x67\x70\x25\x48\x82\xa2\xa4\x75\xb6\x79\x38\x2f\x89\x45\xe2\xbb\xe0\xc3\x77\x07\x88\xfd\x82\xa8\xc7\x20\x7b\x7d\x5d\xad\x63\xd6\xff\x09\x2f\x30\xa5\x38\xf9\x8c\xd9\x17\x94\xc2\xd5\xe5\xa7\xcb\xb3\xf1\x20\xd0\x8f\xc7\x61\x70\x46\x69\x4e\xc7\x11\x1f\x7c\xbe\x4c\x17\x24\x7d\xc4\x19\xfb\x15\xa5\x4b\x5c\x8e\x46\x0c\xf1\x91\x59\xf2\x69\x71\xc5\x28\x46\x8f\xe3\x48\xc1\x9d\x08\x38\x92\xe2\xab\x97\x82\xe1\xc7\x1a\x19\xfe\xe2\x22\x5b\xe4\xe3\x30\xf8\x89\x50\x3c\x67\x39\x7d\xa9\x81\xe6\xf4\x11\x31\xb8\xfa\xb0\x64\xf9\x78\x10\xfc\xf9\xea\xd3\xe5\x38\x0c\xfe\xf6\xe1\x2f\x1f\xc7\x51\xf0\x19\x3d\xcb\x51\x67\xd9\x3c\x4f\x48\x76\x0f\x57\x3f\xbe\x30\x5c\x8c\x07\xc1\x2f\xd7\xe7\xe1\x68\x1c\x4a\x00\xc1\xfc\x07\x7a\x5f\x18\xc2\x7f\xa5\x84\x61\xfe\x64\x1c\x06\x9f\x31\x4a\xc4\x9f\x51\xf0\x07\x94\xcd\x71\x2a\x7e\x9c\x18\xe6\xc4\xcf\x61\xf0\x91\x14\x4c\xfc\xf9\x4e\x50\xe4\x2c\xc2\xc5\x32\x9b\x33\x92\x67\x3e\x58\xb1\x07\x52\xf4\xef\xee\x60\xb6\x4c\xd3\x58\xfd\x98\x3d\xe5\x05\x1c\x98\xe1\xfd\x27\x9a\xb3\x9c\xbd\x3c\xe1\xfe\x6c\x46\x32\xc2\x4a\x78\x16\x60\x83\x42\x40\x31\x8d\x03\xe2\x98\x62\xb6\xa4\x99\xc7\x1f\x94\xb8\xee\x31\xfb\x9c\xe7\xec\x43\x51\xe3\x04\x07\x14\xac\x24\x84\x4f\x5f\x5f\x33\xfc\xec\x29\x10\xa0\xa8\xfa\xb8\x4f\x31\x4a\x2e\x32\x76\x12\xf9\xb8\xff\x94\x17\x44\xce\x01\xf4\xec\x5f\x01\x06\x2e\xce\x1f\x71\x51\xa0\x7b\x6c\xb1\x0e\x56\x5c\xa5\x30\x54\x0c\xf7\x67\xb3\x7c\xb1\x28\x30\xf3\xad\xf9\x04\x43\xa0\xa7\x81\xdf\x97\x03\x0b\x46\x49\x76\x6f\x0f\xec\xe1\x80\x81\x31\x17\x62\x49\xbc\x60\x88\xb2\xda\x34\x19\x58\x31\xf9\xe2\xd3\xdd\xff\xe0\x39\xf3\x43\x8b\x5b\x94\x24\x7f\x69\xb0\x29\x24\xcc\x5f\x9d\x13\x9c\x26\x9f\x24\x8b\x83\x00\x07\x03\x0b\x12\x67\x49\x93\x90\x9a\x1e\x7f\xa9\x88\x95\x93\xe1\xa0\xa5\xd6\x6c\xaf\x10\x25\xcc\x1b\x69\x85\x85\xd0\xa8\x86\x8b\x31\xa7\x7e\x94\x03\xf7\x53\x12\xe7\x6c\x0a\x4c\x09\x4a\x6d\x91\x28\xd7\xb3\x9d\x9e\x30\xa3\x27\x9c\x91\x5f\x48\xc6\x46\xc3\x8a\xa6\x30\x30\xd6\x23\xe6\x14\x23\x86\x3f\xe6\xd9\xbd\x3f\x50\xeb\x69\xb1\x24\xd4\xc4\x25\x8b\x16\x25\xb2\x40\x51\x92\x5c\xd5\xa6\x51\x53\xa4\x0b\xc1\x17\xd7\x23\x56\x67\xa3\x8e\x0c\x67\x49\x0b\x17\x5d\x1a\xa6\xdd\xf1\xf6\xfa\xa5\x21\xde\x48\xbb\x0c\x3a\xa3\x5b\x4d\x96\x9c\x9a\xa5\x87\xed\xa7\x57\x8e\x59\xfc\x67\xb5\xca\x30\x24\xb4\xa6\x29\x83\x16\x8d\x32\x60\x5f\xaf\x4f\x06\x15\xce\x12\x27\xfd\x6d\x75\xe9\x33\x2e\x9e\xf2\xac\xc0\xbb\xeb\x94\x86\x7c\x63\xdd\x32\x68\x1b\x3a\xd6\x64\x75\xa3\xae\xe9\xe1\x5f\xa7\x73\x8e\x59\x52\x91\xbd\x5c\xbf\x3c\xe1\xba\xfe\xed\x11\xfd\xb4\xfe\x9d\x56\xc3\x1f\x18\xd7\x93\xb2\x3e\xcf\x59\xb6\xe2\x6c\xd7\xa0\x3c\x72\x06\xe5\x65\x26\x57\xb0\xc6\x96\x8e\xc9\x0d\x2e\x2a\xa6\xd0\x5c\xaa\x86\x49\x44\x6e\x31\xa3\x44\xcd\xb7\x2a\x5f\xb1\xd2\xd8\x36\x8f\x53\x7f\x10\xd0\xc0\x29\xa5\x0e\xcc\x5d\xd9\x40\x68\xb2\x81\x06\x0e\xcb\xda\x9c\x53\xec\xb4\x3a\xc4\xd0\x0e\x96\x86\x18\x7a\x2b\xeb\xe2\xa8\x4a\x8b\xaa\xb0\xe1\xb6\x22\xc4\xd0\x9e\x96\x53\xe5\xfa\x8e\xa7\xe2\x6f\x91\x27\x36\x2c\xa5\x3f\x9b\x7d\x11\xc5\x42\xcd\x74\xb8\xf3\x76\xc9\x4f\x70\xf2\x11\x67\xf7\xec\xe1\x0d\x02\x87\x26\x3e\x4b\x71\x56\x0f\x1d\xad\xd4\x3f\x50\x8a\x5e\xbe\x82\x38\x5f\x1b\x21\x03\x81\xc8\x08\x42\xe0\xf6\x41\xff\x6e\xb9\x58\x60\x1a\x34\x1e\xbf\x30\x2c\x75\xbb\xb7\x51\x74\x0c\x04\x5b\xcc\xce\xf2\x01\x7c\x86\xd2\xee\x2b\x0a\xd5\x16\xfe\xf8\x70\x94\x24\x3f\xd6\x34\xa2\x23\x25\x17\x60\x32\x06\x0a\xc8\x5f\x05\x63\x0d\x78\x41\x4f\xbe\xe3\x16\xdc\x4f\xc5\x42\x07\x21\x88\x17\x39\xf5\xb9\x9c\x29\xd4\x4f\x8f\xc3\x98\x4e\xe1\x20\xa6\xc7\xc7\x8a\xb4\x70\x28\xf8\x86\xde\x82\xb5\x16\x36\x37\x62\x85\x0f\x54\x67\xbb\x03\x1b\x81\x35\x75\xee\x3c\xea\x72\xea\x72\x18\x56\x31\xbd\x43\x99\x59\x02\xbd\x55\xb1\x69\x61\x2c\x4b\x4e\x17\x6f\xee\xc2\xb3\x1c\xb9\x67\xf9\x69\xd1\x97\x75\xa0\x8b\x76\x43\xef\x06\x0d\x60\x5e\xd9\xb5\x80\x76\x2d\x45\xd9\xf6\xd8\x7e\x25\x4a\x98\x37\x5a\x08\x0b\xa1\x59\x07\x17\x63\xce\x65\x28\x07\xee\xb7\x0a\xce\xd9\xfc\x67\x33\x71\x27\x4b\x5f\x50\xba\xc4\x7b\xe4\x67\xae\x44\x68\x43\x7e\x56\xeb\x83\x99\x04\xad\x9d\xa7\x5d\xc3\xe0\xe9\x3e\x99\x99\x45\x5e\xd8\x82\x4b\x3d\x1a\x96\x72\x52\x97\xe6\xd7\x97\x28\x55\x64\xbf\x36\xd7\xc4\x95\xd3\x85\x22\xa7\x73\x49\xb6\x0d\x67\x57\x20\x89\x4c\x20\xb1\x80\x71\x96\xb4\x88\xa5\xd3\x0b\xa8\x76\xe0\x0e\x3e\x40\x41\xbc\x95\x07\xd0\xe8\x4a\xfb\x6f\xb0\xe4\xb6\x7e\x35\x6c\x4f\xdb\x6f\xce\x22\x43\x8f\xdf\xa8\x01\xe8\x20\xfe\x84\xec\x2c\xee\x6b\x0a\x9d\x3d\x88\x93\x22\x21\x74\x47\x97\x77\x6a\xb9\xbc\x83\x03\xdb\xbd\x5c\xd4\xbd\x0b\x03\xe3\x05\x4a\x8b\x8a\xbe\x29\x5b\x6e\x2c\x75\x9b\x25\x6b\x30\x94\x24\x97\xd5\x75\xea\x48\xb7\x6c\xc8\x9f\xab\x42\xee\xa8\x96\x6c\xc8\x8b\xaa\x84\x9a\xae\xe3\xd4\x8f\x82\x1e\x0e\x7a\x62\xa2\x55\x70\x6e\x9c\xae\x79\x76\x16\x57\x7a\xef\x60\x87\x0a\x4b\x83\xbc\x55\x99\x65\xf0\x95\xb5\x56\x93\x2b\x77\xc1\xa5\xc7\xed\x59\x75\x39\x66\xf2\xed\x0c\xd4\x45\xfd\xdb\x59\xa8\x8b\xfa\x82\xa4\xb8\xde\xef\xe6\x0c\x90\x6d\xad\x94\xbc\xdf\xe0\x3a\x4b\x1c\x24\x4b\x04\xf5\xcd\x45\x29\x01\x3d\xfc\xdd\xd0\x94\x57\x5b\xb0\xbe\x57\x9d\x7a\xba\x57\x9d\x6a\x58\x90\xa5\x5c\x53\x5f\x5b\x7c\x4c\x09\xb8\xb3\x93\xa9\x80\xee\xe6\x65\x2a\xa0\xe7\xd5\x55\xee\xca\x00\x4a\x58\x99\xb0\x08\xf0\x2d\x0a\xb9\x61\x59\x4f\x0e\xb7\xab\x27\x15\xe9\x8e\x8a\xb2\x2a\xf9\x1d\xb8\x09\x86\x35\x04\xbc\xb6\x74\x2e\xdc\xb6\xf9\xcc\x6e\xfb\x56\x36\xd4\x1b\xe7\x35\xd5\xbd\x2b\x37\x7b\x1b\xf3\x9b\xfd\xf7\xaf\x5a\x66\xb5\x8f\x23\xfb\xba\x3c\xa7\xdc\xaf\x72\xcf\xbf\xa5\xc1\x52\x01\xdf\xd2\xb0\x9a\x81\x5f\x6f\x52\xb5\x92\xee\x52\x29\xbd\x45\xbe\xbd\x3a\x69\x88\x37\x52\x25\x83\xce\xa8\x51\x93\x25\xa7\x0a\xe9\x61\xfb\xa9\x8f\x63\x16\xdf\x4e\x75\x0c\x71\xa1\x16\xcd\xf9\xb6\xa8\x8c\x01\xdb\x59\x5d\x0c\x24\xce\x12\x27\xb9\x6d\x3c\x8f\x3e\x0e\xb2\xeb\x1e\x54\x13\xf6\x0d\xbd\x50\x0d\x71\xc5\x17\xb5\x31\xdc\xea\x91\xaa\x00\xfb\xfb\xa5\xd6\xd9\x7e\xab\xdd\xa8\xfa\xe9\x9d\xb2\xdd\xd1\xcd\xdd\x37\xd9\x91\x72\xf0\x61\x5c\x68\xdb\xb2\xb5\xec\x4a\x39\x50\xed\xbc\x2f\xe5\x94\x57\x27\xf6\xdd\xaa\xad\x1a\x16\xe5\xb4\x37\x4c\xb6\xcb\x26\xcd\x81\xa8\xed\x4d\xd1\x80\xbc\x91\x05\x96\xf8\x8c\xe1\x39\xb8\x72\xda\x9b\x19\xb7\x9f\x99\xb9\x66\xf2\xed\x1c\xb8\x8b\xfa\x5e\x4d\xc3\xfd\x6a\x28\x17\xf9\x85\x3c\x75\xb7\x9b\x5b\x71\xf5\x2c\x9d\x5d\x0e\xe9\x55\x04\x89\xfe\x87\x25\xcb\xdb\xb8\x20\x59\x52\xeb\xb9\x6f\x51\x09\x85\x83\x96\x0e\xb3\xab\xd9\x52\xd3\x63\xe1\x0f\x1c\x3a\xd7\x70\x16\xc3\x9a\xda\xec\x1c\x43\x2b\xa0\x5b\x75\x33\x43\x37\xec\x79\x6d\xa5\x5c\x6e\x29\x92\x6e\xa9\x14\x78\x13\xcd\x45\x4d\xd4\xae\xb6\xcd\x89\x83\x03\x9c\x25\x6e\x79\x75\xf9\x1b\x7d\xe8\x72\x7b\x77\xa3\x21\xde\xc8\xdb\x18\x74\xc6\xd9\x34\x59\x72\xfa\x1a\x3d\x6c\x3f\x57\xe3\x98\xc5\x92\xee\x1c\x26\xf7\x73\x34\x0e\xda\x8c\x3c\xe2\x7c\xb9\xab\x8d\x8d\x36\x6c\xe2\x9c\x44\x4e\x23\x73\xd0\xc6\xfa\xd0\xee\x6f\xe9\x66\xf4\xc9\xe0\xbe\xd8\xba\x6d\xe1\x64\x2f\x77\x67\x39\x9a\x3d\xfc\x9d\xe1\x42\xb8\x95\xa6\xea\xb5\xb8\x1c\x03\x86\x92\xe4\x17\xda\x99\x34\x94\x0e\xc7\x06\xbc\xae\x2f\x7a\xd3\xda\x4f\x22\xcb\xe1\xd8\xb0\x67\x8d\x45\x6b\xf7\x38\x55\xd9\x37\x50\x6d\xe3\xba\x4e\x9c\xae\xcb\x60\xc1\x59\xe2\x14\x5d\x97\xf7\x69\x9c\x0e\xee\x72\x3e\x0a\xe0\x8d\x7c\x8f\xc6\x66\x5c\x4f\x83\x1f\xa7\xe7\x51\xa3\xf6\x73\x3c\xcd\x19\x20\x7a\x5f\xfc\xf6\xd5\x83\x58\x27\x5d\x31\xb8\x99\xf8\x26\x45\x82\x26\x2d\x6c\xaa\x79\x3a\xdc\x5d\x0e\x68\x20\x94\x08\x1d\xdb\xb6\x02\x30\x73\x76\x21\xd9\x3e\xd0\x6b\x40\x9c\x39\x8e\xb3\x77\x29\xb9\x7c\xbb\x5a\xc7\xb8\x7f\x75\xf1\xf7\xb3\x4f\xe7\xb3\xab\x3f\x7d\xfa\x7c\x0d\xa3\xf2\xc1\xc5\xe5\x35\x1c\xc6\xb8\x7f\x7e\xf1\xf1\x6c\x76\xf1\xd3\xd9\xe5\xf5\xc5\xf9\xc5\xd9\xe7\xd9\xc7\xb3\xcb\x3f\x5e\xff\x49\xbc\x2a\x3f\xac\xf8\xe5\xfa\xfc\x74\xf6\xe3\xdf\xae\xcf\xae\xc6\xa1\xfc\xba\x62\x76\x75\xfd\xf9\xe2\xf2\x8f\xe3\x88\xd3\x10\x9e\x1f\x72\x3d\x15\xfa\x28\x4f\x29\x45\x20\xc6\xfd\x45\x9a\x23\xfd\xee\x5c\xfe\x2d\xdf\x2a\x20\x75\x76\xc9\x0c\x1d\x0d\xcb\xa1\xa3\x61\xdb\x50\x52\x7c\x24\x8c\xa5\xf8\x2c\x4b\x08\xca\xa0\x3e\x1e\x15\x8e\x24\x40\xed\xb8\xd4\x4d\x18\x0c\x6e\xf5\x29\x29\x70\x33\xb8\x85\x10\x86\x31\xee\x7f\xcc\x6d\x4f\x56\x9a\x70\x9a\x3f\x43\xf6\x3a\x90\x16\xfc\x40\xee\x1f\x20\x7e\x1d\xac\x15\x80\x6a\x21\xdb\x70\xc6\x58\x3d\x06\xe1\xe0\xe8\x88\x42\x38\x78\xaf\x46\xff\xfd\xec\xf3\xa7\x31\xe7\x47\xfe\x16\xa3\x0d\x2a\x2b\x12\xe7\x6a\xc6\xb6\x39\x2a\x0f\xa0\x79\x9a\x4e\xa7\x03\xd0\x33\x4c\x7d\x37\x8c\x7e\x18\xfe\x30\xfa\x3e\xfa\x61\xe4\x40\x88\xff\xb9\x44\x69\xd5\xb6\x2c\x37\x24\xa6\x08\x19\xff\xdf\xd1\x51\x39\x4d\xc8\xc4\xff\x0d\x3a\xce\x3c\xb4\x98\x1f\x04\x03\x2e\xfd\x1f\x97\x24\x4d\x70\xf5\x2b\x09\xb2\xf0\x0f\x94\x56\x52\x18\x0e\xa2\xe1\x1a\xa7\x05\x56\xbf\xd9\xda\x38\x43\x11\x0d\x7e\x14\x2b\xd1\x47\x69\x9a\xcf\x11\xc3\x3e\x05\x52\xd6\xc5\x13\x9a\x63\x48\xe5\x8f\x47\x92\xa1\x94\xdc\x67\x30\x94\xbf\xbf\x30\x74\x97\x62\xcb\x43\xcb\x07\x33\x92\xcd\x96\x05\x86\x6a\xb9\x48\x71\x89\x0b\x86\x13\x28\x36\x18\xe5\xb3\x5c\xd8\xc7\x4c\x18\xb8\x1e\x27\x81\x0b\x78\x73\xab\x7e\xcb\xcd\x92\x6c\xf9\x38\xc3\x29\x7e\x2c\xf4\xb8\x45\x4e\xe7\x78\x96\xe0\x05\x5a\xa6\xac\x80\x6a\x7f\xd6\xc8\xa0\x9a\x46\xcc\xf1\x4f\x66\xa0\xed\x5b\x1c\x88\x98\x1b\x49\x82\x18\x92\xd2\x69\x2a\x82\xa7\x84\xe8\x86\x44\x45\xa9\xf1\xed\xb0\xe6\x6c\x60\xb1\xbc\x43\x95\xb3\x84\x56\xc8\x68\x3e\x92\x5a\xa7\x9c\x31\x00\x6d\x1c\x08\xe2\x5d\xec\xd7\x4e\x2d\xf6\x8b\x94\xcc\xf1\x1b\xb2\xf1\x44\xf1\x53\xcd\x3a\xc9\xc2\x67\xd3\x8a\x52\xa9\x45\x31\x3a\xc6\xd6\x72\x93\xf0\x7f\x0d\x23\x73\xf4\x84\xe6\x84\xbd\xf8\xe0\xb8\xd4\xcd\x1e\x05\xbd\xf0\x88\x1d\x87\xf1\xf3\x03\x49\x15\xd7\xe2\xcd\x84\xf4\x58\x4f\xed\x35\x16\xb0\x89\x24\xb6\x2c\x40\xb1\x7c\x4f\xf3\xe7\xd2\x1a\x34\x61\xdb\x14\x7a\x0e\x44\xc7\x85\x34\xa6\x27\x94\xf8\xa4\x4d\x04\xa8\xfa\x51\x80\xde\xb9\xe2\x56\x82\x27\x2c\xc6\xbd\x9e\xc9\x56\xfa\xcf\xbc\x82\x13\xe1\xeb\xd8\x9a\x29\x0f\x42\x6e\xe4\x66\x7c\x53\xc7\x2b\xd8\x4a\x5c\xc7\x30\x0c\x58\x0b\xab\x7a\x7c\x38\xda\x88\x2e\x1c\x55\xf0\x45\x9d\xf8\x4e\xa2\x8d\xf8\x74\x9d\xa2\xf0\x0d\x3b\xf1\xd9\x9e\xd9\x81\x4f\x1f\x5e\x53\xf8\x4e\x37\xe3\x53\x81\xb0\x1d\xa3\x1a\xb0\x0b\x8f\x8d\xf8\xe1\xc4\xb9\x35\x9f\xea\x6c\x6c\x13\x1d\x37\x2f\x3f\xe4\x81\x40\xfc\xb4\xd6\x7b\x23\x26\xd7\xfa\x0a\x54\x51\x03\x15\x5f\xeb\x8d\xb8\x5c\x92\x13\xb8\x86\x0d\x5c\x5c\x86\x1b\x71\xb9\x24\x26\x70\x9d\x36\x70\x71\xd9\xb5\xe3\x6a\x5d\x53\x17\x67\x66\x7d\x3b\xf0\x6d\xcb\x9d\x59\xdb\x0d\xf8\x74\x96\x5a\xc9\x78\xb4\x77\x6c\x46\xa8\xd7\x57\x7c\x00\xa9\xa2\x69\x4e\x4a\x6b\xef\x94\xe6\x8c\x13\xeb\xa4\x56\x59\xf6\x1d\xc9\x85\xa3\xdd\xe9\x55\x16\x60\x47\x7a\xbc\x8e\xda\x95\x5e\x65\x81\x36\xd3\x3b\xd0\x49\x99\x4f\x41\x85\xee\x68\xb8\x23\xdd\xa6\xaa\xed\x32\x53\xad\x7b\x7b\xd0\xdc\x61\xb6\x0e\x9a\x3b\xcf\x53\x96\x44\x7b\x92\xd4\x27\x31\x76\xa2\x78\xc5\xe8\x72\xee\xa4\x68\x21\xcf\x44\x82\xb9\x25\xe6\x4c\x65\xa3\xd5\x54\x99\x1d\xc0\x6a\x2a\xb3\x62\x0f\x34\x7f\xf6\x78\xae\x2d\xbe\x40\xf6\x0f\xcf\x53\xc4\x64\x66\x50\x8c\xbd\x42\xf0\xe5\x3d\x2e\x0b\xe6\xdd\x61\x4f\x9e\xbd\x26\xff\xc6\x89\x47\xb2\x94\x64\xb8\x7f\xd8\x4a\x3f\x67\x97\x35\x16\x4a\x09\xea\x64\xb9\x8b\xbc\x4c\xa0\x0d\x59\xc4\xd1\x48\x66\xb2\x5c\x30\x24\x67\xd9\xce\x05\x97\x51\xd3\x93\xc9\x1c\xfc\x86\xdd\x56\x85\xe1\x46\x91\xd7\x94\xa1\x91\x56\x3a\x13\x36\x1b\x57\x35\xdf\x6a\x14\xd3\x14\x32\x3b\x5f\x23\x0b\x9f\x1e\x9d\x44\x51\x18\x45\xef\x86\xdf\x47\x5d\x22\x9a\xa3\x8c\xcb\x82\xd3\xf0\x64\x62\xeb\xdd\xe1\x97\x3c\x4b\xbc\xc8\xbb\x27\xf7\x48\x24\xbd\x5c\x40\x32\xcf\xa4\x93\x49\x18\xcb\x5c\xd1\x5d\x12\x11\x10\x17\xfd\x02\xb3\x9f\x75\xf2\x4b\x8e\x29\x7f\x64\xf2\x77\xcc\x7c\xa6\x7f\x05\xe2\xa5\x92\x47\xd1\xaa\xe2\x0d\x7b\xb2\x23\x8a\xdd\x17\x70\x86\x51\x7b\x89\x8e\x59\xcf\x1e\xdf\xb2\x64\x56\x47\xa5\x49\xd4\x28\xa6\x94\xb5\x5d\xe2\x89\x1a\xaf\xa2\x23\xf0\xe6\x76\xed\xa8\xf9\x58\xdc\x9e\xd9\x2a\xe5\xc2\xb7\x70\xb0\xae\x56\x86\x8c\x2e\x5d\x85\xe1\x16\x3a\x68\x7a\x2d\x2e\x63\xb2\x99\x7f\x7d\x3d\xd8\xc9\xc0\x0c\x62\x6f\x8e\xd2\x14\x27\xde\x33\x61\x0f\xf9\x92\x79\x96\x08\x0f\xc1\xba\x1a\xac\x06\x20\xb6\xba\xf5\x9a\xf3\x58\xe9\x72\x43\x5a\xc7\xa1\x90\x96\x38\xb1\xa6\x8a\x7e\x25\x23\x7a\x0b\xf5\x21\x36\xad\x9e\x3d\x6b\xb0\x3a\xde\x56\x09\xcc\x55\xf0\x03\x38\x78\xcf\x8e\xab\xcf\xc6\x03\xa9\xeb\x05\x8c\xe2\x1a\xec\x71\x43\xf8\x92\xeb\x0c\xfa\xa4\x57\x80\xef\xaa\x3d\xab\x1a\x74\x26\xc7\xe6\x70\x20\xfe\x8f\x60\x69\xec\x31\x1b\x73\xa6\x29\x67\x79\x62\x57\xf9\xea\xb8\x5e\x4c\xb9\x76\x70\xa8\x85\xab\xb0\xb2\x21\x6e\xe8\x2d\xd7\xca\x0c\x42\xbb\xc9\x29\x39\x58\x80\xb2\xa2\x5a\xc2\x1a\xb3\xcb\x49\x16\x2f\x7b\xb5\xa7\xa5\x92\x54\x10\xa1\xde\x12\x1c\xb8\x08\xf4\x96\x00\xac\xe6\x79\xc6\x48\xb6\xc4\x1e\x5b\xaf\x73\x58\x67\xee\x8e\x62\xf4\x8f\xf5\x9a\x2c\xfc\x5c\xad\x8d\x6c\x9e\xb8\xe6\x15\x6f\xac\x7e\x82\xfc\x98\x01\xd9\xad\xa9\xc8\xec\x69\x59\x3c\x54\xad\x1e\xb4\x22\xaa\x11\x0c\x6a\xce\x02\xac\x5d\xbd\x19\xed\xbf\x5b\xda\x29\x24\x23\xc5\x43\xb3\x90\x37\xe7\x79\x55\x9b\x48\x38\xaf\x4a\x31\x1f\xd8\xae\xa9\xd7\xd6\xe0\x14\x5e\x87\x28\xcd\x38\x80\xad\xc3\x3a\x2c\x77\x41\x52\xec\x91\x04\x67\x8c\x2c\x08\xa6\x26\x44\x4b\xbc\xde\x61\x3b\xfd\xb5\xd6\xa1\xa2\x95\xf8\x71\x18\x17\xdc\x00\x0b\x63\x80\x65\xb9\x45\xfa\xf3\x07\x44\xff\x90\x27\xf8\x03\xf3\x0b\x00\xd6\xeb\x6d\xc4\x01\xe2\x5a\x7e\xc4\xca\x45\xb5\xa3\x4d\xa9\x1e\x2d\xae\x90\xe2\x7f\x2e\x09\xc5\x32\x6d\xaa\x35\x51\x2d\x0f\x54\xd7\x44\xb5\x74\xc7\x35\xa5\x3f\x89\x7c\x0a\xe2\x6a\xef\xa4\xb4\x07\xd2\xc3\xe0\x00\x0e\xf8\x82\x1d\x14\xdd\x0b\x82\xd3\x84\xcb\xbd\x77\x68\xe5\x4b\xac\x3d\x35\x29\x4f\xd8\x56\x94\x2d\x20\x8e\x40\x25\x0d\xa4\xde\x2a\xb4\x35\xb1\x12\x46\xd9\x77\xba\xb3\x29\xde\x11\xf1\xa0\x35\xb4\xd4\x99\xa8\x2d\xb9\xb6\xb5\x3a\x79\x60\xef\x2f\x75\x04\x30\xd9\xbb\xbe\x12\x3b\xc5\x8d\x9c\xd4\x23\x59\xc1\x50\x36\xc7\xf9\xc2\xea\x9e\x9b\x9d\x86\xb2\xa1\xcb\xc3\xb1\x0a\x33\x03\xd5\x08\xa3\x13\xa6\x4c\x49\x99\xa7\x5e\x4c\x5b\x4b\xb9\xfb\xe5\x8b\x58\x4c\xde\xbd\x8b\x7e\x18\xbd\xbe\x16\x53\xf8\x6e\x74\x12\x0d\xc0\x8a\xc0\xa2\xc4\x9f\xb9\xc0\xa0\x5f\x4c\x26\xe1\x00\xf4\xb2\x9e\x3f\x7a\xf7\xee\x64\x74\xec\x0b\x2c\xe2\xe1\xb1\x44\xc3\xdd\x21\x99\x84\xd1\x29\x58\x61\xe9\xbf\x88\xf2\x6c\xe2\x45\x34\x18\x5a\x6f\xa6\xd3\xd1\xd1\x49\xf8\x1a\xfe\x10\xd9\x63\x04\x6a\x7b\x50\x18\x1d\x85\xef\x5e\xa3\x68\xa8\x46\x59\x6f\x4e\x8f\xbe\x7f\x8d\x86\x83\x40\x8e\x1a\x9d\xbc\x72\xca\x6b\x1b\x7f\xfd\x99\x79\xb0\x5e\x57\xea\x6a\x9d\x6f\xb5\x7f\xcc\xbc\xd9\x48\x8f\xcd\x91\x74\xfb\x98\xfa\x20\xc8\xad\xe8\x18\x20\x63\x5a\x2a\x6d\x8c\xe9\x04\x57\x42\x23\xba\xc9\x7b\xbd\x5b\xc8\xc3\xf7\xda\x56\xaa\xca\xf9\xf5\x76\xbd\xaa\xef\xa7\x94\xfb\x22\x95\xad\x13\xb3\x07\xd2\x92\x8c\x4b\x2e\x39\x8b\x33\xbd\x8d\xaa\x5b\xc0\x33\x38\xa8\x02\x9a\x84\xd9\xb5\xd3\x21\x77\x2e\xac\x16\x6b\xfd\x33\x7a\x50\x63\xa3\xfd\xfa\x82\x5a\xb5\x21\x98\x6b\x85\xd5\xdc\xb6\x81\x9b\xd9\xb4\x62\xb0\x56\xd9\x91\xae\x1b\x61\xb0\x56\x04\xda\xe9\x6e\x9e\x81\x5a\xfc\x56\x2c\xfa\x3c\x43\xeb\x36\x92\xb5\xf3\x0b\x26\x93\x68\x38\x9d\x46\xc3\x8d\xd8\xc4\xe0\x56\x74\x92\xab\x1b\x76\xdb\xc5\x51\xad\x9b\xe8\x62\x49\x74\x11\x27\x93\x70\x34\x9d\x86\xa3\x4e\x9e\x36\x20\x34\x4c\xbd\x56\x7e\xf6\xc2\xdb\xc9\xe4\xb4\x8b\xd1\x5a\x43\x70\x6b\xbc\xb5\x67\xd1\x2d\x9f\x49\xed\xe1\xc9\x2d\x17\x79\xe7\xd4\x36\xb0\x50\x46\x5f\x06\xa6\xd3\x69\xdd\xb6\x1a\xb3\xa9\xb5\x23\x2b\x66\x26\x77\x37\xeb\x58\x83\xfa\x93\xde\x70\x83\xd1\x95\x9f\x82\x6f\x4b\x47\x9f\x39\xb2\x08\xe9\x47\x9d\x94\x5c\x1d\x5b\xb5\xd5\x7c\x33\x50\x3d\x09\x7b\x2e\x66\x7b\x5d\x6f\x6a\xdf\x0c\x36\x2b\xaa\xab\x85\xab\x09\xd4\xf7\xb0\xdf\x0f\xc6\xa1\x8b\x66\xeb\xf8\x70\xec\xe0\xb1\x37\xac\x73\x39\x1a\x6e\xe2\xd2\xb5\x97\x53\x9e\x64\xd1\x0a\x0a\xf1\x66\x04\x75\xa3\xde\x1d\x43\xb3\x53\xdc\xc4\x10\xd7\xcd\x04\xe2\xe9\xb4\xdd\x00\x0d\x67\xbf\x09\xe2\x66\xab\x79\x5b\xbc\x71\xdd\xb2\xf9\xd3\x70\x14\xd7\x6d\x9b\x3f\xde\x60\xdd\x66\x7a\xbf\x03\x36\x9a\x7d\x70\x47\xea\xca\xb3\x99\xfc\xd9\xd1\x57\xea\x0d\x03\x2c\x8e\x1d\xb4\x5b\xab\x99\xec\x06\x32\xda\xec\x9b\x74\x4a\x87\xb0\x1d\x21\x47\x7b\x5d\x1c\xf5\x29\xad\x5e\x67\xfc\xb5\xf9\x69\xd7\xb1\x0d\xfe\xea\x4c\x4a\xfc\xc2\x5e\x37\xe2\x77\x7a\x8e\x36\xb9\x6e\x70\x1e\xed\x5c\xde\x63\x55\x4e\x5d\x98\xba\xd6\xd5\xf0\xaa\xe4\x10\x93\x6a\x6a\xd2\xdb\xae\x06\xef\xa8\xe5\x4a\xee\x3c\x52\x78\x2c\xcf\xbd\xe2\x21\xa7\xcc\x63\xb9\x37\xcf\x33\x86\x48\xe6\xa1\xcc\xaa\xbe\x75\x7f\x95\xc1\xc3\x43\x3b\x0f\x16\x89\xae\x9b\x05\x99\xf8\xb2\x1e\x94\x45\x51\x7f\x41\xf3\xc7\x3f\xa8\xda\xa3\x12\xce\xd4\x86\x76\xcb\x0c\x29\x28\x3f\xf7\x6c\x95\xab\x3e\xce\xe6\x2e\x97\x8f\xdb\x03\xce\xc4\x7e\x15\x8e\x7c\x0a\xde\xd7\x9f\xf4\xb0\x38\x6a\xdb\x4a\x59\x9c\x8f\x6b\x1c\x3e\x53\x87\x15\x71\xaf\x46\x1b\x83\x98\xbf\x14\xf1\x25\xde\x62\x5e\x45\xbd\xa2\xe4\x1a\xcd\x7a\x8e\x78\x66\x7d\x8b\xdd\x78\x5e\xf0\x65\x93\x05\xe0\x20\x66\x56\x13\xed\xe2\xf2\x5a\xf4\xe3\x21\x84\xe5\xa1\xb4\x7e\x79\x26\xcd\x99\xdb\x96\x27\x5c\x02\xd6\x23\x60\x2d\x0b\xd5\x6c\x42\xa4\xc4\x73\xbb\x7d\x68\xe5\xb1\xbd\x4c\x55\xa9\x68\xc2\x4b\xc3\x55\x0e\x51\x59\x99\x2e\x36\x0d\xe7\x35\xe2\x2a\x87\x3e\x3a\x3a\x09\xc1\x64\x32\x7a\x5d\x1c\x8d\x4e\x4a\xd8\xe5\x46\xd8\xe1\x40\xc1\x86\xef\x78\xd6\x1a\xbd\xfa\x1c\x5a\xa0\x59\x56\xd0\xdc\xb5\xa0\x11\xc0\xdf\x73\xd8\x53\x03\xcb\xd1\x2c\x35\x9a\x3b\x8e\x46\x76\x0c\x75\x95\x5b\xb8\xd5\x3e\x57\xa5\x6e\x7e\x0c\xc5\xc0\xb8\x65\x9c\x9f\x4f\xa7\xbc\x22\x17\x75\x78\xe0\xe7\x47\x7e\x28\xab\xf1\x10\xf4\x54\x41\xbe\xae\x6c\x4a\xb8\x95\x47\x7f\x4a\xef\x4c\x53\xeb\x9a\xc9\xda\xdd\x96\xfe\xe2\x7d\x4b\x3c\x15\x03\xee\x44\x3a\x4b\x71\xb6\x4d\x1e\xdd\xa4\xb3\x89\xe1\x07\x54\xcc\x88\xc3\xc9\xaa\x8e\xcc\xef\xa1\x3b\xb9\x8d\xfb\x14\xbc\xda\x1d\x1b\xdd\xd6\xde\xd6\x77\x6a\x69\xca\xe3\x75\xc6\x93\xd2\x65\x7b\xde\xb8\x4f\xcb\x61\x9e\x67\x05\xf3\x28\x94\x1b\x2a\xfd\x05\xc5\xf8\xdf\xd8\xd7\x27\xcc\x41\xac\x11\x79\xc4\x27\x01\xf7\x46\xc1\x4a\x7e\x1c\x30\xce\x20\x15\x07\xd0\x03\xf9\x55\xd2\x38\x87\xd1\x1a\xae\xd6\xa2\x4b\x5f\x30\x0f\xa9\x53\x92\xaa\x33\xe2\x87\x83\x68\x08\x14\xb9\x05\xcc\x20\x84\xb4\xff\x19\x3d\xbf\x27\x7d\x96\x4b\x4b\xf2\xc1\xf8\xcf\x57\x9f\x2e\xfb\xd2\x73\x92\xc5\x8b\x4f\x34\xc0\x12\xa2\x4a\xa3\xce\x5f\xe8\x37\x77\xf5\x37\x05\xd8\xf0\x05\x93\x8f\x80\xf3\x5b\x23\x1f\x05\xcb\xc6\x9b\x9f\x11\x7b\xf0\x51\x70\xd7\x78\x21\x65\xe3\xa3\x20\x6b\xbc\x92\x9f\x0d\xf9\x28\xc8\x35\x83\x09\x6c\xfb\x3c\x48\x32\xe3\x3a\x86\x5d\x7d\x63\x9d\xb5\xf6\x91\x3e\x4c\x5d\x7e\xe0\xd7\x1c\xe9\xa3\x20\xd1\xe4\x1f\xa0\xeb\xd8\x34\x27\x80\xd4\x7e\x82\xff\x00\xe2\x5f\x4f\xff\x9a\xd3\x7f\x60\x1a\xf5\x0b\x9c\x25\x3e\xaa\x1e\x5c\xf4\x01\x58\x1b\x3d\x28\x64\xa0\x24\x3e\x0b\x0e\x0f\x03\x0c\xd6\x92\x4e\x06\x57\xeb\x52\x59\x72\x9f\x6a\x3d\x20\xb0\xab\xe7\x44\x81\x66\xb6\x80\x5d\xd7\xba\x95\x2a\x91\xc3\x42\x5d\xb8\xe6\x83\xf2\xac\xb0\x0f\xe2\x14\x33\x0f\x89\xff\x2e\xe2\xe2\x99\xb0\xf9\x83\x5f\x94\xf7\xa0\xf9\x00\xac\xe6\xa8\xc0\x5e\xf3\x6a\x2d\x71\x69\xff\xca\x5f\x25\xfc\xff\x68\x0d\xb3\x9b\xfc\x56\x53\xc3\xb0\xbc\x0b\x34\x56\xe8\x78\x4a\xb0\x80\xf6\x75\x96\x3e\x50\x9b\x50\x2d\x14\xe4\x35\xfe\x2b\x7f\x85\xc5\x1f\x2d\x34\xc4\xa8\x2a\x91\xd2\x99\x99\xfb\xe6\x7d\xd0\x45\xcc\xfa\xc7\x07\xfc\x15\xce\x92\x92\x9e\x80\x8b\xd5\x31\x91\x71\xc3\x5f\xfe\x92\xfd\x23\xcb\x9f\x33\x4f\x51\xf2\x28\x9e\x13\xfc\x05\x27\x1e\x0f\x75\x1e\x5d\x66\x8c\x3c\xe2\x43\xd1\x5a\x46\x10\xc2\x65\x96\xe0\x05\xc9\x70\xa2\x1d\xcd\x1a\xf9\x0b\x4b\x63\x90\x3c\x33\x12\x10\xb0\xca\x78\xf9\x25\x05\x8c\x03\x29\x04\x1a\x70\xd6\xc8\xba\x1c\xbf\x28\x63\x89\x0f\xe0\xb4\xe1\xcf\x19\xb0\x06\x2f\x7d\x81\x58\xeb\x0f\xf5\x4b\xfd\xe8\xd2\xbb\x02\x94\xee\x65\x97\x4b\x97\xfd\x1c\x68\xc5\x5a\x5a\xdf\x9c\x5b\x9a\x55\xbb\x88\x57\x2d\xbb\x6b\x95\x35\x02\xbe\xcc\xca\x4b\xff\x4c\xf3\x47\x52\xe0\x3e\xc5\xe2\x13\x06\xf7\xd2\x83\xb5\x9b\x94\xf9\x67\x2a\x6a\xd4\xf4\xf3\x0a\x41\xed\xfb\xb1\xdb\x90\xac\x1e\x93\xe2\xc9\x17\xb6\x0f\xa7\x2b\xdb\x1b\x80\x55\x82\x53\xcc\xb0\x97\xdd\xd0\x5b\x7b\xe3\x85\x2f\x94\x5e\x0b\x0c\xa7\xcc\x27\x3e\x06\x20\x46\x3e\x0d\x8a\xfe\x1d\xc9\x12\x3f\x07\xfa\x2f\x6c\xfe\x5a\x70\xfd\xc3\xff\x7a\xc2\x73\x86\x13\xcf\x52\x62\x6f\x91\x53\xef\x49\x30\xc2\x23\x77\xe2\x25\x6a\x4e\x87\x5c\x1c\x60\xad\xd5\xb9\x4b\x8c\x87\xe7\x88\xa4\x38\xe1\x15\x53\x82\xe7\x79\xc2\xf5\x5b\x2e\xac\xd2\x6f\xfc\xcf\x25\x2e\xd8\x21\x00\xeb\x75\xe9\x17\x29\x9e\x7f\xf1\xf3\x32\xea\xd4\xe3\xa5\x4e\xc1\xad\x88\x99\xd4\x1a\x75\xf6\xe7\x18\xcc\x5c\x50\x6b\x5d\x49\xab\x7e\xc8\xfb\x8e\xfe\x3b\xd2\xd2\x7b\x80\x4a\xec\x94\x0b\x11\x5b\xfe\xf2\x09\x32\x38\xb5\xd3\xd0\x3c\xc1\x3f\xe7\x24\x63\x7e\xbf\xdf\xe7\xe4\x2d\x6e\xe6\x3e\x0d\x56\xfa\x0b\xc3\x31\x81\x77\xea\x03\x40\x2b\x64\x17\xb5\x90\xfd\x2e\x8c\x34\xa5\x0c\x16\xd5\x30\x4b\x41\xeb\x37\x7b\x32\x04\xd7\x3e\xcc\xf3\x0b\x1e\x2d\xb9\x57\xce\x39\xd3\x4c\x5b\x10\x51\x46\x73\xd7\x97\xff\x40\x89\xfb\xe3\x3a\xbf\x08\x08\xcf\xe7\x9f\x94\xeb\x52\x20\xe2\xdf\x32\xd9\x0c\xf1\xe0\x8b\x9c\xe2\x09\xd1\x02\x07\x4f\x75\xd7\xb7\xd6\xb9\x8a\xfb\x13\x3a\xbf\x68\x8d\xd0\xae\xb8\x2b\xbc\x40\xa1\x23\xb4\xf9\x2c\xd6\x11\xa0\x8b\x00\x95\xb9\x90\x33\x40\x17\x20\x2e\x74\x80\x5e\x18\x5b\x5c\xfa\xdc\x23\xd6\xa2\x75\xd1\x88\xd6\x41\x0e\xd6\xf3\x14\x15\x85\x97\xca\xa5\xa5\x4b\x79\x0f\x97\x3c\xbb\x27\xf7\x70\xd1\x23\x36\x7b\x49\x88\x3d\xe8\xae\x98\xbc\xd0\x8f\xae\x15\x82\xd9\xce\x08\xe4\x75\x63\xd4\xf2\xd1\x2f\x8e\x6c\xa0\xa9\x62\x05\x24\x0e\x15\xdb\x7c\x0f\x10\x4f\x07\xdc\x97\xfd\xf0\x6c\xb5\x54\xde\xcd\x57\xfa\x48\x2c\xce\x65\x26\x6d\xcb\x4c\xf4\x32\x57\xee\x5a\x72\x2c\x35\xe1\x7a\xaf\xdd\xa0\x73\xa9\x09\x88\x89\x5e\x6a\xe3\x61\x10\xac\xad\x32\x69\xac\x72\x99\x4c\x77\x04\x3a\x54\x09\x74\xbb\xdd\xeb\xc2\x95\x6f\x73\xb0\x6b\xdc\xee\x61\xfe\x0d\xa3\x5a\x0c\xd2\xcf\x5d\x41\x8f\x8f\x48\x7d\x79\x4f\x9f\x0f\x02\x79\xdb\x84\xf8\x43\xe8\xa3\x5f\x06\xbc\x06\xb9\x2d\xa3\x6b\x3d\x8b\xb0\xa3\xe9\x7a\x43\x2a\x64\x42\x91\x89\x12\x2c\x17\x45\x24\xc9\x16\xf9\xa1\x9d\x8a\xdc\x7f\x85\x9a\xbb\xef\x2c\x92\xea\x57\xbf\x98\xa8\xae\xda\xae\xeb\x87\xbe\x4a\xa5\xcd\xdd\x4f\xff\xaf\xce\x3c\xa5\x32\xff\x0e\x57\x3d\xa7\xd2\x2f\x9c\x49\x15\x1f\xa3\xbf\xd5\xb4\xae\x31\xe4\x13\x5d\xe4\xd4\xe7\x91\x50\x7c\x79\x37\xa1\xfa\xbc\x01\xeb\xc1\xb0\x54\x20\x05\x25\x7a\x9e\x3c\x6b\x96\x16\x42\xb4\x85\x10\x6d\x21\xa4\xb4\x10\xcb\x96\x66\x0e\x5b\xa2\xbf\x5f\x1b\xfa\x82\xa8\x77\x01\x57\x69\x7e\x3f\x2e\x02\x59\x5e\x8f\x69\x20\x5a\xf7\x63\x12\xe8\x98\x3e\xbe\x0b\x28\x46\xc9\x78\x1e\x70\xb8\xf1\x4b\x90\x10\x3a\xbe\x5f\xc7\xf8\x5f\x4f\x39\x65\x9e\x62\xc0\xbb\x88\xff\xeb\xff\x02\x00\x00\xff\xff\x29\x86\xec\x75\xd3\x6d\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
