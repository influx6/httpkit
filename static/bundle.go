package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct {
	path string
	root string
	data []byte
}

var (
	assets = map[string][]string{

		".tml": []string{ // all .tml assets.

			"http-api-json.tml",

			"http-api-mock.tml",

			"http-api-readme.tml",

			"http-api-test.tml",

			"http-api.tml",
		},
	}

	assetFiles = map[string]fileData{

		"http-api-json.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x4f\x6b\xbb\x30\x18\xc7\xcf\x06\xf2\x1e\x9e\x9f\x87\x1f\x0a\x25\xde\x37\x7a\x18\xbb\x8d\xd1\x0e\x4a\x4f\x63\xd0\x34\x7d\xec\xec\x6a\x2c\x31\xee\x0f\x92\xf7\x3e\xa2\x8e\x69\xa7\x68\x43\x2f\x82\x3c\xf1\xf3\xfd\x9a\x4f\x08\x25\x51\x04\x87\x3c\x93\x10\x27\x9f\xba\x50\x98\x03\x63\x8c\x92\x77\xae\x20\xa0\x04\xca\x92\xad\xb4\x2a\x84\x66\xcb\xed\x01\x85\x66\x0b\x9e\x62\xf5\x30\xe6\x61\xb5\x5c\xc0\x1c\x36\x65\x09\x29\x3f\x55\x6f\xcd\x62\xf0\x2d\xd2\x07\xdf\x07\x63\x36\x94\x8c\x80\xee\x15\x72\x8d\x7d\xb8\x7a\x72\x27\x74\x92\xc9\x0b\xa1\xeb\xd3\x6e\x00\x5a\x4f\xfa\xa1\xa1\xe5\x46\x11\x3c\x66\x7c\xd7\xaa\xa5\x50\x17\x4a\xe6\xc0\x41\xe2\x07\x24\x32\xd7\x5c\x0a\x84\x2c\x06\x6e\x4b\xb4\x6b\xb2\x27\x2e\xde\xf8\x1e\x8d\x61\xe7\x93\x56\x49\x63\x18\x25\x71\x21\xc5\x59\x52\x20\x32\xa9\x51\x6a\xc8\xb5\x4a\xe4\x3e\x84\xc0\x11\x3f\x03\x54\x2a\x53\x21\x94\x94\x78\xd6\x26\x1e\x31\x75\xed\x6a\x37\xc5\x4b\x62\x8b\x84\x9b\x79\x75\x5e\xd8\x5a\xa6\x5c\xe5\xaf\xfc\x18\x3c\xbf\x6c\xbf\x34\xfe\x34\x0f\x67\xf0\xdf\x66\x85\xb7\xd5\xf2\x7f\x73\x90\xc9\xb1\x6a\xe1\xd5\xbb\xe8\x5a\xa2\xac\xff\x89\x12\xaf\xee\xd3\xd0\x6c\xd6\xcc\x66\x50\x62\x5a\xf2\x5a\xfa\x47\xe4\xb5\x8f\x43\xb7\x4b\x67\x32\x28\xef\x37\xa9\x57\x9e\x0b\x7e\x50\x9e\x0b\xec\xea\xf2\x5c\x4a\x5c\x26\x6f\xf4\xd6\x19\x51\xda\x7c\xdc\xe9\xf6\x17\xd8\xd5\x38\x96\xd9\x2b\x77\x6a\xd0\xa0\xd0\xa9\x80\xab\x4b\x9c\x1a\x3c\x49\xdc\x77\x00\x00\x00\xff\xff\x0a\x95\x1d\x8a\x46\x06\x00\x00"),
			path: "http-api-json.tml",
			root: "http-api-json.tml",
		},

		"http-api-mock.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x97\xcf\x6f\xdb\x36\x14\xc7\xef\x06\xfc\x3f\xbc\x04\xd8\x20\x15\x82\x7c\x4f\x91\x43\xd6\x14\x43\x80\x35\x09\x9a\xf5\x34\x0c\x03\x2d\x3e\x59\x5c\x24\x52\x25\x29\x39\x81\xe0\xff\x7d\xe0\x2f\x59\x96\x1c\xa7\x4a\x7b\x58\x2f\x71\x2c\x92\xef\xc7\xf7\x7d\xde\xa3\xb5\x5a\xc1\x27\x91\x3d\x5e\xdd\xdf\xdc\xd5\x28\x89\x16\x12\x28\xe6\x8c\xa3\x02\xc2\x41\x69\xd9\x64\xba\x91\x08\xdb\x82\x65\x05\xb0\xaa\x2e\xb1\x42\xae\x15\xe8\x02\x61\x78\x2a\x17\x12\x6a\x29\x5a\x46\x19\xdf\xc0\x72\xb1\x5a\x01\x81\x4a\x64\x8f\xd0\x28\xb2\x41\x60\x1c\x34\x2a\xad\xec\xc6\x46\x21\x6c\x99\x2e\xac\x95\xae\xfb\xda\x08\x8d\x90\x3e\x58\x6f\xe9\xdd\xfa\x5f\xcc\x74\x7a\x4b\x2a\xdc\xed\xa0\xd0\xba\x36\x8e\xd2\xe5\x42\x3f\xd7\x38\x89\xd6\x85\xd8\x2d\x17\x00\x00\x4a\x0b\x89\xf0\x4e\x3d\xf3\xac\x22\x75\xfa\x89\xd4\xcb\xc5\x6e\xb9\xb0\xd1\xdc\xe2\x76\x7c\x56\xa2\x6e\x24\x57\x40\x80\xe3\x16\x18\x57\x9a\xf0\x0c\x41\xe4\x40\xc6\x7e\xd2\xe5\x22\x6f\x78\x76\xc4\x4a\x14\xc3\xbb\xb1\x61\x1f\x8e\xb3\x0f\xbf\x8e\x96\xfd\x6a\x1f\xf0\x85\x71\x1f\x0d\x82\x8e\x13\xb7\x63\xb7\x8f\xfe\x1a\x4b\xd4\xe8\x15\x46\x27\xbf\xb0\xf6\x98\xe0\xa0\x05\x48\xac\x44\x8b\x40\x60\xc3\x5a\x53\x01\x89\x99\x90\x14\x18\x45\xae\x59\xce\x90\xc2\xfa\x19\x6e\xae\x43\x1e\x51\x35\x89\x3a\xf6\x4e\xa2\x4c\x3f\x41\x26\xb8\xc6\x27\x9d\x7e\x70\x9f\x09\xd4\xcd\xba\x64\xd9\xcd\xb5\x11\x9c\xf1\x4d\x0c\x28\xe5\x3e\x53\x96\xc3\x3f\x09\x14\x44\xc1\xc5\x25\x54\xa9\x4d\x2b\xfd\x43\x10\x1a\x85\x73\xf1\x7b\x38\x33\xeb\x83\xe4\xbd\x3c\x79\xa5\xd3\x8f\xc6\x58\x1e\x9d\x7f\x76\x51\x53\x81\x0a\xb8\xd0\x80\x4f\xcc\x30\x63\x59\x61\x14\x7e\xf9\x7a\xbe\x8f\x24\xee\x45\x72\xff\x04\xb7\x3e\x8b\xd1\x36\xef\x8c\xb3\x72\x2f\xea\xef\xa8\xaf\xca\x72\x80\x81\x2a\x99\xaf\x7f\x59\x02\x69\x09\x2b\xc9\xba\xc4\xa0\xa5\xc8\xc1\x22\xd8\x75\x01\xd5\x7b\x92\x3d\x92\x0d\xee\x76\xe9\xfe\xd9\x01\xbe\x27\xe5\x76\xee\x8f\xcb\x2d\x24\x45\xe9\xb5\xf6\xdf\x7e\x7b\xee\xbf\xd7\xae\xa3\x74\x02\x12\x55\x2d\xb8\x42\x79\x8f\xf2\xde\x3f\x8d\x21\xfa\xeb\xef\x19\x41\x26\xce\x94\x2d\x68\x1c\x0a\xd4\x12\xe9\xf3\x56\x30\xcb\xda\xb8\x1c\x9f\x09\xdf\x60\x64\x54\x88\x1e\x13\x68\x8d\x2f\x94\x39\xc9\xb0\xdb\xc5\xb0\x16\xa2\x1c\x22\xc1\x72\xc0\x12\xab\x04\xc4\xa3\x21\xa9\x4d\xa3\x19\x9e\xe3\xf7\xe6\xd8\xc0\x9a\xab\xbb\xcb\xe1\x12\x48\x5d\x23\xa7\x91\x7f\x90\x58\x47\xf1\x7e\x73\x1f\xf8\x80\x16\x2d\x1b\xf4\x94\xc5\x61\xd9\x2f\xf5\x66\x4a\xe4\xc1\x66\x9c\x4c\xf8\x32\xdb\x25\xc3\xd6\xcc\xd3\xc0\xd1\x9a\x28\xa4\x60\xda\xb6\xe8\x3b\x9a\xf6\x58\xbf\x86\xcc\x37\xb6\xe7\x1c\xe1\x42\xed\xbd\x74\xae\x02\xb9\x68\x38\x7d\xb9\x9d\xfb\xbe\x3f\xb3\x1b\xa7\x6d\x3d\xc3\x7f\xb7\x4b\x7e\xd8\x14\x90\x1f\x07\xfc\x98\x4c\xe6\x21\xb4\x4f\xeb\x10\xa5\xb7\x25\x65\x65\x55\xe9\x2d\x6e\xa7\x49\x55\x44\x67\x85\x1d\x28\xe7\x93\x1c\x3c\x62\x2e\x95\x03\xa4\xbe\xd4\x94\x68\x84\xc6\x7e\xa8\xc9\xb8\xef\x2f\xd5\x0d\x6b\x91\xdb\x6b\xad\x25\x65\x83\x27\xa1\x72\x36\xbf\x89\x2b\xd7\x35\x46\x06\x77\xe8\x2a\x33\xf7\xcf\xa1\x18\x07\x2b\x87\xea\x3a\x41\xba\xe5\xa2\xeb\x8c\xc6\x7e\xeb\x8d\x7a\x60\x15\x2b\x89\x04\x33\x3d\x86\xb3\xe3\xc1\xfc\xed\x99\x3b\x68\xd9\xe1\x3c\xef\x3a\xc0\x52\x61\x7f\x9c\x8b\xed\x90\x82\xe3\x00\xcf\xe3\x02\x5e\x01\x63\x58\xe9\x5b\xa1\x5d\xf7\xac\x18\x6f\x49\xc9\xe8\xb8\xca\xe6\xc3\x47\x8a\x3e\x52\x94\xd2\x84\xea\xea\xfa\x42\x10\xa6\x42\xc9\x3e\xb9\x81\x18\x66\x6c\x4a\x09\x67\x97\x46\x8f\xe3\xd1\xbd\x70\x59\x3a\x85\x7d\x18\xe9\x7d\xaf\xb4\x7f\x32\x9e\x7a\xbd\xdc\x9c\x5a\xb5\x03\x96\x1f\x24\x1a\x2c\x09\xa5\xe1\xd7\x54\xf8\xfd\xc1\xb5\x08\x40\x0e\x38\x75\xce\x4f\x41\xe9\x2c\x1e\x87\x32\x30\xe8\xf6\x1c\x63\xf0\x60\x65\x54\xc9\x37\x4e\x46\x8f\xac\xb3\xfc\x1a\xb2\x78\xa8\xe6\x94\x5b\xdc\xb7\xf6\x84\xde\x11\x13\x99\xf5\x78\x92\x89\x19\x28\xcc\x9f\x60\x3f\x88\x9c\x90\xd5\x51\x84\x9c\xb8\x91\x99\x8b\xe3\xa1\x10\xdb\x5d\xab\xd5\xe9\xde\x18\x4c\x44\x7c\x62\xda\xa0\x36\x23\xd3\x17\xc6\xa7\xb3\x09\x4c\x63\x15\x58\xfd\xf2\x5a\x83\x4e\x69\xf5\xb6\x67\x51\xe7\x3c\xcb\xb7\x0f\xda\xef\xb8\xfe\xdd\xdf\xd5\x0a\xfe\xbc\xbb\xbe\x8b\x28\xb6\x58\x9a\x37\x8d\xf8\xa2\x5f\xb8\x6b\x51\x4a\x46\x11\x8c\x24\xf6\x05\xc4\x26\xcd\xc3\x1d\xbd\x2d\x88\x06\x55\x88\xa6\xa4\x50\xd8\x9f\x5d\xe9\x88\x86\xef\xb8\x47\xcd\x74\xbd\x09\xef\xa1\x48\xcd\x58\xdd\x19\x80\x90\xd3\xdd\x84\xa5\x51\xb7\xf6\x2c\x9d\xec\x29\xbf\x1a\x26\xd9\x1b\x38\xca\xa5\xa8\x06\x1c\x39\x7b\xf6\x16\x08\x1c\xbd\xda\xd4\xff\x87\xa9\x67\xc6\xc7\xcf\x4c\xc3\x7f\x01\x00\x00\xff\xff\x61\x8c\x96\xc7\xdd\x10\x00\x00"),
			path: "http-api-mock.tml",
			root: "http-api-mock.tml",
		},

		"http-api-readme.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\xcf\x6f\x22\x37\x14\xc7\xef\x48\xfc\x0f\xaf\xe1\x02\xab\xec\x90\x56\xdb\x0b\x52\x0f\x59\x96\xa6\xa9\xda\x0d\x5a\xd8\x53\x14\x75\x8c\xfd\x60\x9c\x9d\xb1\xa7\xf6\x1b\x12\x34\x9a\xff\xbd\xb2\xc7\x10\x48\x80\x40\x1a\xb5\x4a\xd5\x1c\x62\xc5\x7e\xfe\xbe\x1f\xfe\x3c\x7b\x52\x96\xd1\x88\x4c\xc1\x29\xba\x9a\xdc\x22\xa7\xe8\x33\xcb\xb0\xaa\xe0\x97\xf1\x78\x08\xe7\xc3\x4b\x68\x36\x7e\xda\xff\xd3\x6c\x34\x1b\xd7\xdf\x5d\x5f\x68\xf8\x82\xb9\x36\x04\x7d\x66\xc4\x4d\x3b\x21\xca\x6d\xaf\xdb\x9d\x69\xe3\xa7\x39\x33\x22\xe2\x3a\xeb\x4e\x98\x98\x61\xb7\x2c\xa3\x21\xe3\xdf\xd8\x0c\x87\x8c\x92\xaa\xea\xec\xd9\x51\xff\xf9\x74\x8b\xf3\xfc\x6c\xfc\xd2\x02\x03\x56\x90\x86\x19\x2a\x34\x8c\x50\x80\xf3\x04\x2c\x97\x30\xd5\x06\x28\x41\xb0\x5e\x02\xe2\xad\x6a\x41\x32\x8e\x9c\xbf\x71\x82\x5e\x16\xef\x73\xe4\x64\xfd\xee\xc2\xa2\x01\xd2\x20\xb3\x3c\xc5\x0c\x15\x01\x53\x02\x72\xa3\xe7\x52\xa0\xb7\x98\x30\xfe\x0d\x95\x00\xa9\x08\xcd\x94\x71\x74\xe6\xc1\x40\x40\xa1\x04\x9a\x54\x2a\x04\x31\xa9\x4d\x18\x27\xa9\x55\xcf\x39\x8c\xe3\x78\xa6\x9b\x0d\x5a\xe4\x08\x1f\x1f\xcb\x94\xcd\x06\x00\xc0\x27\x4c\x91\xb0\xcd\xb5\x22\xbc\xa7\xa8\x5f\x8f\xa7\x2e\x2d\xa9\x66\x1d\x40\x63\xb4\xa9\x4d\x2f\x90\x76\xdb\xb5\x1f\xf2\x0f\x95\xae\xaa\x68\x47\x85\x4f\x6b\xd5\x4e\x2d\xfb\x35\x17\x6c\x77\x04\xa7\x50\x96\x51\x6d\x72\xee\x33\xdb\x94\xdf\x58\xd9\x70\xf2\x38\xf4\xf3\x34\xdd\xed\x63\x39\x4a\x45\xfe\x57\x07\xda\xd7\x37\x47\x65\xe4\x77\xae\xa7\xd5\x37\xb8\x35\xad\xb2\x8c\xea\xa5\x6d\xe9\x6c\xac\x3c\x4a\xe7\x65\x15\xae\x3c\x06\x8e\x86\x4b\x72\xe4\x69\x8b\x35\x79\x53\x9d\xa6\xfa\x4e\xaa\x19\x64\x48\x89\x16\xd6\x13\x8d\x8c\x27\x80\x4a\xe4\x5a\x2a\xf2\x10\xb5\x5a\x21\x15\x28\x61\x78\x35\x1a\x43\x77\x87\xcf\x2e\x54\xce\xba\x05\xbf\x7b\xbd\x1e\xc4\xd3\x42\x71\x68\xbb\x66\x79\xe7\x7a\xea\x7c\x78\xd9\x59\x95\x85\xee\xe1\x9d\xeb\xa5\x82\x64\xba\x2c\x4e\x38\x32\x1f\x6d\xf0\x69\x90\xa3\x9c\x87\x90\x57\xd4\x1b\xe4\xda\x08\xd0\x53\x3f\xbd\xb7\x6e\xe0\xe9\xbf\x4b\x24\x4f\x5c\x3f\x0b\x4c\x25\xce\x51\xf8\x8d\xcd\xc6\xaf\xa3\xab\xcf\xe0\x8f\x48\x91\xeb\x2b\x37\xbb\x6c\xff\x08\xc6\x89\xb4\x70\x27\xd3\x14\xa4\x02\x2a\x8c\x02\x83\x7e\x60\x60\xd0\xba\x26\x96\x73\x77\x01\x30\x2a\x2c\x70\x2d\xd0\xb7\xf9\x7b\x18\xf8\x06\x47\x01\xfd\xa0\x3c\x5e\xe4\xd8\x83\xd0\x92\x2e\xeb\x80\x48\xbd\xfc\xbe\x5e\x66\x79\x9e\x4a\xce\x5c\x0a\xdd\x5b\xab\xd5\xea\xe0\xd6\x04\xbf\xe0\x9f\x05\x5a\x82\x21\x33\x2c\x43\x42\x63\x83\xe8\x3e\xe3\x8f\x5a\x2c\x82\x59\xad\x5b\x96\x90\xb1\xdc\xe7\xbe\x51\x39\x38\x99\x58\xad\x4e\xe0\xe4\xd6\x0f\x55\xb5\x4d\x75\x54\xa7\xdb\xd7\x02\x57\xbe\x7f\x66\x32\x2d\x0c\xf6\xe0\xc7\xb3\xb3\x66\x63\x54\x70\x8e\xd6\xf6\xe0\x87\xb3\xef\xb7\xc7\x65\x73\xad\x2c\xee\x0f\xac\x46\x6c\x77\x48\xad\x16\x5c\x0c\xf6\xd0\xd8\xcb\x8b\x49\x2a\xf9\x1f\x52\x1c\x40\xa5\xbf\xda\x9e\x43\xf2\x02\xc9\x01\x60\x3c\x8f\x0c\x66\x72\xee\xba\xe7\x09\x8b\x5b\xdf\x14\x4f\xe1\xd4\xe8\x6c\x03\xb1\xc0\x53\x90\x71\xa0\x3b\xb6\x6d\x91\x12\x30\xe7\xc2\x55\xa2\xd9\x30\xa1\x5e\x11\x5c\x92\x7b\x30\x1e\x35\x43\xfc\x90\x69\x0c\xf9\x12\x0c\x27\xe0\xcd\xdc\x04\x85\x47\x46\xa0\x22\x39\x5d\xf8\x85\x3a\xee\x7f\x93\x58\x27\xb8\x7e\x4a\x47\x10\xfc\x37\xb1\x3c\x7b\x39\x96\x87\xf5\xcb\x73\x70\x1e\x86\xa4\x7f\xb2\x8e\xa3\x32\x4d\xc3\xc1\xda\x97\x11\xf9\x46\x6e\xb0\x7f\xf6\xfc\xaf\x0f\x06\xe0\x66\x9d\x80\xe1\xd7\xd7\xba\x9e\x96\x9f\x48\xcf\xb1\x50\xdb\x01\x23\xc2\x2c\x77\x5f\x98\x1a\x8a\x30\xf5\xdf\xb8\xae\xe0\x4e\x52\xb2\xa9\xe7\x4f\xc5\x04\x50\x26\x5a\x2c\xde\xe8\x95\xb6\xce\xd8\xfa\x57\xed\xeb\x3f\xca\x1f\x56\x0a\x07\xf2\xff\x80\xf4\xa7\xc1\x6f\x83\xf1\xe0\x95\xa8\x5e\xfe\xeb\x71\xc8\x0d\x27\xbc\xed\xff\xcf\xee\x9b\x7e\x76\x3f\x1c\x77\xed\xc6\x71\xfc\x57\x00\x00\x00\xff\xff\x3a\x63\x49\x7e\x63\x10\x00\x00"),
			path: "http-api-readme.tml",
			root: "http-api-readme.tml",
		},

		"http-api-test.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdf\x4f\xe3\x46\x10\x7e\x2e\x12\xff\xc3\xd4\x12\x27\xfb\x1a\x6d\xda\xd7\x54\x79\xa0\xc0\xdd\xd1\x52\x88\x48\x50\x1f\xab\x8d\x3d\x81\xbd\xdb\x78\xc3\xee\x3a\x80\xac\xfc\xef\xd5\xec\x3a\x0e\xe4\x12\x70\x48\x82\x8e\xc6\xf7\x70\x08\xff\xf8\x66\xe6\xdb\x99\x6f\x76\xc7\x8c\xb9\x86\x70\x7f\x0f\x00\x00\xc7\x98\x5a\x03\x6d\x18\xa2\xd5\x22\x36\xec\x1c\xef\xc2\x38\x33\x56\x0d\x59\xd7\xf2\xf8\xdb\xb1\x30\x23\xc9\x1f\x42\x65\x58\xd7\x26\x2a\xb3\x51\xe4\xdf\x1c\xa3\x36\x42\xa5\xd0\x86\x60\xfc\x5b\xb0\xbf\x17\xed\xef\xed\xef\x0d\xb2\x34\x06\x8d\xd7\xc2\x58\xd4\x97\x2a\xb3\x18\x6a\xfa\x5f\xc3\x8d\xb5\xa3\xcc\x0a\xc9\xdc\x55\xdd\x00\x3e\x12\xf0\x91\xae\xf2\x91\x60\x5f\x7a\xbd\xce\x61\xe7\xb4\x51\xa2\x1a\xab\x45\x7a\xdd\x00\x8d\x46\x65\x3a\xc6\xe2\x42\x04\xf9\xfe\xde\x4f\x1e\x92\x7d\xe1\x69\x22\x31\x0c\x3e\x9f\xf4\x82\xc6\x60\x68\x59\x77\xa4\x45\x6a\x07\x61\xd0\x3c\x30\xcd\x03\x13\x94\x70\x33\x9c\xc8\x19\x66\x9f\xd1\x1e\x4a\x19\x55\xc7\x6a\xb6\x46\x59\x5f\x8a\xf8\x5f\x91\x3c\x0f\xeb\x68\x98\x47\xed\x5c\x74\x57\x76\xf1\x48\x23\xb7\xb8\x18\xee\x6a\x1d\x27\xaf\x46\x89\x07\x9e\x87\x3d\x3e\x39\x3b\xe9\x9d\xac\x81\x7c\x8c\x12\x1d\xf2\x84\xbc\x6e\x36\xa1\x87\xc6\x7a\xaa\xf3\x9c\x75\xad\xce\x62\xcb\x2e\xfa\x5f\x31\xb6\xec\x9c\x0f\x71\x32\x81\x31\x97\x82\xdc\x31\x60\x6f\x10\x34\x25\x21\x8e\xb9\x04\x35\x00\x0e\x4b\x5e\x72\xd0\x1a\x63\xa5\x13\x18\x68\x35\x84\x69\x16\x15\xf9\xf7\xa2\xd5\xd0\xc2\x47\x8b\xc6\x8a\xf4\x9a\xf5\xa2\x9c\x9c\xa5\x8c\x4e\xfa\xd0\x6a\xc3\x39\xde\xfd\xad\xe2\x6f\x87\x9d\xd3\x8b\x11\x6a\x6e\x95\x0e\xfd\x22\x50\xc2\xb6\xda\xa5\x31\x2a\x14\x5f\x3d\x0d\x48\xfa\x44\xa7\xd5\x88\xd3\x27\x5c\xaa\x13\x54\x76\xef\x13\x3e\x4c\x85\xf4\x38\xcd\x26\x5c\x16\x15\x02\x6e\x05\x0c\xdc\x09\x7b\x03\x45\xa1\x5c\x6b\x95\x8d\x18\xad\x4e\xf1\x50\x48\xb8\x8e\xe0\x47\xcc\x07\x79\x6e\xd5\x99\xba\x43\x0d\x0b\x42\x2c\xe2\x0c\xa2\x69\x68\x28\x71\xd8\x00\xd4\x9a\xfc\x3b\x53\x3c\xf1\xd9\xf5\x67\xf7\xe2\x3c\xcc\x73\xf9\x02\xce\xec\xe1\xa2\xf6\xc5\xc0\x61\xfd\xdc\x86\x54\x48\x2a\x49\x28\xfe\x11\xab\x86\x7d\xe2\x42\x62\x12\x06\xdd\x1b\x95\xc9\x04\x6e\xf8\x18\xc1\x64\x71\x8c\xc6\x0c\x32\x29\x1f\x40\x2a\x9e\x60\x02\x84\xd8\x82\x55\xec\x43\x0b\x0e\x7e\xb9\x65\x81\x8b\xa5\x70\x66\xe2\x7f\x78\xd3\x1d\x6e\xcc\x76\x4c\xb3\x19\x9b\xb1\xbd\x27\x1e\x63\x95\x5a\xbc\xb7\x2e\x13\x16\x30\xdd\x86\xa4\x5f\x54\x71\x18\xdb\xfb\x86\xbb\xb5\x21\x02\x0d\x1f\x63\xb2\xac\x3e\xa6\xb5\xb1\x36\x57\x55\xac\x3c\xa2\x45\xe3\x6d\x99\x63\x54\x03\x44\xcc\x25\xde\x66\x68\x6c\x21\xac\xf0\x9d\xb4\x54\x4a\xe3\x32\xef\xa3\x06\xf8\x32\xda\x04\x87\xb1\x5b\x9a\x84\xdc\x26\x17\xb7\xce\xe6\x6a\xf6\x9e\xf0\x6a\xa6\x94\x92\x29\x4f\x2b\x3d\x83\x7a\x96\x77\x24\x12\xac\x8b\x7a\x8c\xd4\x4b\x43\x8d\x86\xe4\xf9\xb6\xbc\x2f\x06\x84\xc3\x8e\x54\x82\x44\x99\x5b\x9f\xae\xe5\x36\x33\x17\x7f\x55\x23\x4f\x63\x8c\x82\x12\xc2\xb8\xd7\x20\x26\xa8\x83\x42\x82\x35\x9a\x91\x4a\x0d\x42\x3f\xb3\x70\xad\x2c\x1c\x24\x2c\x68\x3c\xb1\xd2\x28\x1d\xa8\x4a\x5f\x25\x8b\xb4\x24\x4f\xec\xcc\x87\xfc\x87\x4a\x1e\xd8\x19\xa6\x61\x04\xed\x36\xfc\xfa\x8a\x4c\x29\xfd\x98\x45\x49\x98\xc1\xab\xd2\xe0\x19\xb0\xb9\xb6\xf9\x9a\x9e\x29\x65\xe5\xae\xc9\x97\xf4\xcd\xaa\x4d\xf3\x07\x6e\x99\x7e\xe7\x59\xf7\xcd\xba\x6f\xee\x72\xdf\x9c\x3b\x5f\x10\x8d\xac\xe3\xb6\xf1\xa7\xc7\x75\x27\xad\x3b\xe9\x8e\x74\xd2\x25\xd9\xe1\x15\x66\xae\x99\xba\xcc\x12\x2a\xad\x7c\xfe\x74\x6d\x68\x61\x27\x7d\xd6\xec\xfb\x3a\x81\x6e\xac\x9d\x8e\xb9\x76\x6b\x04\xfd\x07\x4b\xf9\x94\x0d\x06\xa8\xfd\x2d\xb7\x74\xff\x68\x61\xb1\xeb\x26\x3d\xab\xb6\xd7\x0a\x52\xea\xa7\x30\x1b\x38\x83\x7c\x20\x67\x6b\xed\xdc\xa8\x76\x1e\x15\x7e\xbd\x81\x80\x16\xa6\xde\x46\x45\x0b\x63\xff\x6b\x29\xf5\x13\xc5\x39\x29\xcd\xfc\xc5\xad\x09\xa9\x37\xba\x93\xa7\x92\x5a\x46\x6b\x19\xad\x65\xf4\x7d\xc9\xe8\x92\x79\x82\x57\x31\x37\x4f\x28\xe3\x29\x2a\x37\xda\xc6\xd0\x60\x83\x63\x80\x47\x61\xe5\x39\x68\x9e\x26\x9f\x04\xca\x64\x5a\xfd\x10\x50\xb0\x01\x04\x5f\x8d\x4a\x03\x08\x66\x1f\x90\x60\x32\x79\xa2\x64\xd9\x9c\x8e\x3d\x09\xbb\xd5\x06\x02\xa0\x3a\x39\x49\x29\x53\x74\xf8\xa1\x9f\x45\xcc\xff\x12\xba\xd9\xc0\xef\xeb\x11\x84\x0e\x6a\xe9\x81\x7d\x7d\xd2\x5e\x30\xb0\x68\x14\xb0\x48\x77\xaf\xb6\x35\x09\x20\x42\x77\x50\x86\xb7\xa7\xc2\xe7\xea\x48\xa5\x16\x53\xfb\x16\x3a\x5c\x1a\x7b\x1b\x25\x2e\xcd\x55\xdb\x1b\xfa\x6f\xc2\xdf\xcd\xac\x87\x6a\xe5\xaf\xbc\xab\x6c\x0e\xbd\xd5\x9d\xdb\x1c\xd6\x43\xeb\x7a\x68\xfd\x23\x0c\xad\xa7\x7f\x49\x52\xcf\xad\xdf\xc5\xa1\x61\x87\xda\xd5\x7f\x01\x00\x00\xff\xff\x9e\xcc\x69\xd2\xfb\x26\x00\x00"),
			path: "http-api-test.tml",
			root: "http-api-test.tml",
		},

		"http-api.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xe2\x4a\x12\x7e\x26\x52\xfe\x43\x1d\x3f\x1c\xd9\x47\x1e\xa3\x7d\xf5\x51\xa4\x4d\x26\x9c\x9c\xec\xcc\x90\x28\x21\xbb\x5a\xad\x56\x33\x1d\xbb\x80\x9e\x98\x6e\x4f\xbb\x21\x61\x11\xff\x7d\xd5\x37\xdb\x30\x40\x4c\x2e\xb0\x93\x9d\x3c\x40\xe8\x4b\x75\x55\x7d\x55\xf5\x75\xb7\xdd\x6e\xc3\x6c\x16\x5d\x4b\x31\x4e\x64\x74\x71\xfb\x15\x13\x19\x75\xc9\x08\xf5\xc7\x7c\x7e\x42\x92\x3b\x64\x29\xa4\xd8\xa7\x0c\x0b\x20\x0c\x28\x93\x28\xfa\x24\x41\xb8\x1f\xd2\x64\x08\x24\xcb\xf8\x7d\x01\x72\x88\xf0\x67\xaf\x77\x79\x7c\x79\x0e\x92\x43\x4a\x27\x28\xa4\x6e\xed\x53\x46\x32\xe0\x39\x0a\x22\x29\x67\xc0\xfb\x87\x07\xed\xb6\xee\x1a\xd0\x09\x32\x78\x7f\x75\x73\x0a\x02\xbf\x8d\xb1\x90\xd0\xe7\x42\x77\xcd\x66\xdf\xc6\x5c\x22\xac\xd7\x0d\xe4\x34\xc7\x08\x7a\x43\x5a\x00\x2d\x20\x17\x7c\x42\x53\x4c\xe1\x76\xaa\x05\x8c\x0b\x14\x91\x5e\xe9\xaf\x74\x94\x67\x38\x42\x26\x0f\x0f\xd4\x9c\x46\x16\x97\x66\xce\x0e\x0f\x5a\xad\x53\xcc\x50\xa2\x9f\x70\x26\xf1\x41\x46\xef\xcd\x77\x08\x85\x14\x94\x0d\x02\x40\x21\xb8\x50\x03\xcf\x50\xae\x1f\xe5\x57\x0b\x5f\x92\xe4\x8e\x0c\x70\x3e\x8f\x36\x29\x13\x1a\xc1\x81\x92\x7c\x93\xa7\x64\xbd\x0a\xa1\x32\xca\x0c\x39\x4e\x94\x9b\x17\x57\x58\xe8\xa9\xad\x33\x9f\x2f\xea\x7e\x9c\x65\xeb\x57\x70\xdf\x94\x49\xfd\x11\x80\xff\xaf\x7f\x6f\x6f\x92\x9e\x5e\xd9\xf5\x5e\xe0\x4a\xbb\x66\xb3\xc8\x74\xad\xb2\x67\xa1\x67\xc9\x9e\x67\x78\x79\x7e\x78\xa0\x23\x66\xd3\x68\x15\xe4\x9b\xf2\x01\x1f\x72\x5e\xa0\x0e\xc1\x11\xca\x21\x4f\xbf\x0f\xcd\xa1\x94\x39\xdc\x9a\x40\x8b\x1a\xc4\xa4\x5e\x72\x21\x20\xad\xd3\x7e\x53\x92\xc6\x92\x66\xce\x6d\x15\x9a\x36\x5c\x36\x8c\xb0\x31\xbd\x61\x84\x0d\x88\xf5\x23\x9a\xf8\xeb\x0a\x13\x2e\xd2\xa2\x72\x99\xce\x5b\x55\x23\x04\xe6\x02\x0b\x64\xa6\x4c\x08\x2c\x72\xce\x0a\x57\x14\x24\x07\x52\x2f\x0a\x7a\x1d\x92\x65\x20\xac\x3c\xde\xd7\xd3\x96\x9c\xd7\x0c\xf1\x26\x4e\x77\x7a\x17\x7a\x00\x28\xaf\xc3\x25\x19\xa0\x02\x02\xbe\x7c\x2d\x38\x8b\xbd\x9c\x0c\xd0\xfb\xa2\x7a\xae\xac\xf6\x97\x28\x96\x07\x89\xc5\x2e\x33\xbe\xc7\x25\xc9\xdc\x12\xb5\xc1\x52\xb5\x7f\xb6\x36\x3a\xd1\x66\xd4\xf6\xb9\x56\x69\x50\x8a\x73\x88\xb9\x52\x5d\xa1\x62\xed\x34\x41\x3c\xe4\x59\x5a\x54\xb1\x4a\x72\x0a\x43\xc2\xd2\x0c\x45\xa1\x2b\xb4\x89\x68\xca\x06\xba\x78\x6b\x89\x65\x81\x2f\xca\x1a\x5e\xc6\x7d\xd3\x62\x6e\x61\x71\xca\x19\x95\x74\x01\x1e\xa1\x14\x34\x29\xc0\x7e\x47\x9f\xcc\xb7\xea\x32\x0b\x73\xd1\xa4\xac\x57\x0e\xe8\xe2\x3d\x08\x94\x63\xc1\x94\xf1\x0c\xef\xcb\x55\x29\x2b\x24\x61\x89\x62\x10\x65\xe0\x82\x21\xe5\x5a\x44\xc9\x6a\xb7\xad\x3e\xd1\xe1\x41\x7f\xcc\x12\x25\xd4\x1f\x2d\xeb\x18\xba\x54\x6f\xa2\x60\x00\xbf\x39\x3d\xb4\xdd\x46\x45\xf8\xd5\x36\xea\xb6\xd2\x1b\x31\x8c\x42\xd3\xe0\xf4\x8a\xdd\x5a\xba\x7d\x5e\x99\x6b\x0a\x86\x4a\x1e\xa4\x13\x53\xb8\x34\xb2\x2e\xc3\x24\x87\xc4\x0c\x31\xce\x78\x1c\x31\xcd\xad\x5a\xf6\x15\x1f\x4b\x8c\xa1\x3d\xd3\xff\xcc\xdb\x71\x3e\xbe\xcd\x68\xf2\x99\x1a\x17\x7d\xd2\x55\x30\x86\xcb\x8b\xeb\x9e\x6e\x38\xb9\x38\xfd\x67\x0c\x7f\xbb\xbe\xe8\x1a\x11\xda\x75\xbe\x8a\x31\x67\x7b\x00\x8e\x15\xe4\x03\xac\x2b\x3f\xda\x41\x23\x88\x8f\x4a\x87\x77\xf1\xbe\x27\x48\x82\xbe\x22\x96\x14\xfb\x28\x54\xe4\x46\xae\xbb\x33\xa2\xd2\x77\x3f\xce\x59\x9f\xfb\x9e\x5d\xcf\xf2\x89\x17\x84\xae\xfb\x1f\x54\x0e\x8d\xac\x51\xd4\x61\xa9\x1f\x04\x81\xf2\x64\xab\x95\xc8\x87\xe8\x4f\x24\x29\x0a\x3f\x88\xae\x51\xfa\x9e\xd6\x8a\xc9\x77\xbd\x69\x8e\x5e\x08\x1e\xc9\xf3\x8c\x26\x3a\x15\xda\x2a\xff\x3c\x3b\xf3\x11\x55\x4a\x84\x0c\x20\x16\xa9\x74\x49\xa7\x3f\x28\x66\x69\x51\x4e\xd5\x3f\x6d\x50\x78\x63\x91\x79\x31\x28\x05\xaf\x8c\x10\x3f\x88\x6e\xae\x3e\x2a\x08\x29\x1b\xf8\x81\x89\x09\x67\xc8\x84\x08\xa0\x2c\xe1\x23\x15\xe4\x4f\x24\x5b\x23\x89\xf6\x15\x20\x0a\x08\x65\xae\x42\xe1\x14\x13\xae\x3c\xa4\x74\x39\xe1\xe9\xd4\x0f\x82\xc8\xb4\xf9\xbf\xba\x35\x83\xdf\xf5\xac\x5f\x8e\x80\xd1\x0c\xac\x0d\x6b\x7d\xd4\x51\x88\xf7\x7d\xef\x0f\x42\x33\x4c\x55\xb8\xe6\x44\x14\xa8\x3e\xc9\x48\x85\x73\x0a\x63\x91\x45\x7f\x27\xd9\x18\x8b\xe6\x3e\x6b\xb5\x3c\x1d\x4b\x5e\xac\x94\x09\xcb\xc6\x46\xae\xac\x39\xb3\x4c\x54\x14\xc2\x66\x5e\x03\xc8\x55\x0a\x3c\x05\xe8\x94\x48\xe2\xc5\x25\x7a\xe1\x93\xf0\x77\xb4\x14\x3a\xf0\x94\xb2\xae\x8e\x44\x55\xfe\x85\x50\x21\x56\x43\xfb\xc9\xb8\x25\x65\x25\xe2\x62\x0b\xa3\xf7\x8b\x94\xa3\x77\x38\xc5\x4c\x9d\x6c\xb6\x80\xcb\x6a\x68\xfe\x1e\xd3\xd3\x2b\x24\x91\xe3\xc2\x8b\x75\x75\x8e\xae\xf5\x2f\x83\x45\xba\x08\x5f\x95\x74\x4a\xa6\x8a\x24\x7f\xc5\x94\x72\x57\xf5\xdc\x6c\x4b\x8d\xdd\xa5\xbc\xdd\x02\xb7\x16\x36\xdb\xc8\x68\x56\x51\x9d\xd9\xf9\xee\x97\xea\x6e\x9a\x33\x9d\x3b\xd7\xed\x88\xe9\xcc\x72\x7b\x65\xba\x12\xa0\x9d\x30\xdd\xe1\x41\xcb\x80\x74\x7e\x1a\x02\xbf\x73\x19\x73\x42\x06\x7e\x10\x9d\xa1\xb4\x73\xbc\x12\x49\x65\x85\x4a\xaf\x5f\xf8\x9d\x49\x94\x47\xd3\xa4\xcb\xa1\x9c\x5d\xed\x15\x29\xb3\xec\xb4\x6d\xa9\x68\x5e\xc1\x5d\x46\x70\xa1\x43\xc2\x19\x71\x9e\x9a\x95\x51\xa2\x00\xc6\xd5\xf9\x69\xcc\x8c\x5d\x73\xeb\x92\xef\xe8\xff\x29\x77\x07\x2f\x4c\xff\x5b\x95\x23\x25\xad\x0c\xa1\x5b\x9e\x4e\x9b\x7b\xd9\x95\x23\xf5\x87\x42\x18\xf1\x65\xc1\x69\xb5\x6a\xa1\x10\x43\x19\x3b\xff\x07\x7b\x83\x4d\xa6\xaf\x20\x9e\x85\x0d\x43\x55\xc6\xc2\x6a\x22\xfc\x10\x9b\xbd\x5d\xc0\xfd\x6a\xfb\x8b\x6d\xb6\x17\xeb\x2d\x5d\xb9\xef\xe8\x72\x5b\xf0\xd5\x90\xe5\xaa\xa3\x96\x2a\x07\xf8\xab\x66\xd5\xee\xd3\xcc\x55\xd3\x5e\x29\xf9\xb4\xf3\xb1\xd3\xeb\xac\xa7\x61\x77\xc3\xbb\x23\x1a\x36\xcb\x35\xa1\xe1\x47\x04\x96\xae\xdd\xcd\xb1\xf1\x6d\x53\x69\x7f\x24\x9b\x6b\xe2\xd8\xf4\xd5\x4a\xf9\xab\xd7\xec\x2a\xe6\xab\x9a\xfd\x02\x07\x05\x15\x8f\x5b\x5f\x55\xbe\xd4\x79\xf0\x67\x31\x7f\xed\x62\x7e\x86\x72\xaf\x95\xfc\xac\x63\x0e\x57\x57\x9d\xeb\xcb\x8b\xee\x75\xe7\x5d\xfd\x94\xb5\xaa\xb4\xeb\x67\x72\x3b\xaa\xeb\x67\x28\xf7\x7a\xb6\x32\xe0\xfc\xe4\x82\xd7\x38\x56\x69\x35\x36\x9d\xab\xd4\x5c\x2d\x1f\xd3\xd5\xb7\x6a\x36\x12\x6b\xd5\xf6\x25\x6e\xd4\x06\x28\xdf\x76\xbd\x7d\xe4\xb2\xeb\xe2\x43\x08\xa5\xe3\x9f\xcb\x5f\xca\x99\x05\x0a\x4a\x32\xfa\x1f\x4c\x9f\xea\x57\xf3\x78\xd3\xd2\xc0\xbd\xa0\x12\xc5\x8f\xe1\xea\x1f\x91\xda\x2e\x3e\xac\xe2\xb4\x85\x3b\x41\xf3\x24\x7b\x23\x6d\xd9\x69\x24\xcb\x9a\x3c\xae\xb4\x8f\x54\xd7\x93\xd7\xb3\x29\x4b\xbf\x8a\xb1\x3b\xd6\x3a\xce\xb2\x7d\x13\x97\x41\x68\x37\xdc\x35\x21\x02\xb8\x48\x51\x84\xe6\xeb\x64\x6a\x5f\x6f\xd1\xbd\xb4\x0f\x3c\x5d\x49\x6b\xbe\xa7\x87\x7b\xc1\xef\xe0\x78\x4c\x0d\x16\xa9\x70\xc3\x79\x1a\xf9\xf6\x85\x9f\x6a\x4c\x4b\xcf\x82\x23\x3d\x52\xeb\x01\x98\x15\xb8\xd4\xe9\x91\x22\xf1\x6c\x2e\xda\x7c\x7c\x54\x93\x93\xe9\xd3\x74\x39\x99\x6e\xd4\x46\x77\xd7\xf2\x70\x51\x2b\xe5\x3d\x55\x65\xdd\x8b\x22\xea\x77\x4e\x06\xd8\xe5\xd5\x13\x87\xda\x9b\x10\xce\x12\x91\xe7\x1b\xf7\x0a\xcb\x6f\x4a\xd4\x94\x5e\xea\x32\xf4\x7a\xa4\x30\x4b\x38\x9b\x44\xc7\x92\x53\x5f\xe4\xf9\x5a\x3e\xdd\x82\x01\x84\xea\xc1\x49\xf5\x46\x8a\x70\xa6\xb0\xf1\xe8\x16\x05\xa4\x28\x09\xcd\xb6\xb8\x78\x5a\xbe\x71\x5c\xbc\x94\x6b\xf6\x64\xc8\x3e\x07\xd1\x18\x28\xb0\x56\x79\x05\x8e\xe0\xdd\x5f\xea\xa1\x93\x0f\x36\xef\xcd\x96\x9c\xec\x20\x5c\xe1\xdb\x7c\xf0\xa2\xae\xcd\xff\xd7\xdc\x59\x77\xc0\xa2\x1b\xeb\x7b\x3a\xfd\x76\xce\xda\xad\x9d\xad\xd8\xe1\x52\x61\x09\xd7\xa5\xc6\x8b\x6d\xfe\x0c\x67\xed\x69\x03\xb8\xab\x2d\x5e\x83\x77\xb5\xac\xa6\xca\xb7\xb1\x73\xba\x69\xb2\x03\xe2\x6a\x9f\x68\x3b\xea\xef\x61\xc5\x16\x5f\x37\x65\x01\xac\x78\x19\x3d\x4b\x27\x6f\x77\xbb\xf9\x96\x36\x94\xdb\x6f\x1b\xff\x1b\x00\x00\xff\xff\x5f\xbc\xdd\x2d\x19\x2d\x00\x00"),
			path: "http-api.tml",
			root: "http-api.tml",
		},
	}
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
	return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
	reader, size, err := FindFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

	gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error) {
	item, ok := assetFiles[path]
	if !ok {
		return nil, 0, fmt.Errorf("File %q not found in file system", path)
	}

	return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
	body, err := ReadFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error) {
	body, err := ReadFileByte(path, doGzip)
	return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
	body, err := ReadFileByte(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error) {
	reader, _, err := FindFile(path, doGzip)
	if err != nil {
		return nil, err
	}

	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	var bu bytes.Buffer

	_, err = io.Copy(&bu, reader)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
	}

	return bu.Bytes(), nil
}
