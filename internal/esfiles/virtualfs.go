package esfiles

import (
	"archive/zip"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var yarnVPathRegExp = regexp.MustCompile(`(?:__virtual__|\$\$virtual)(?:/|\\)[^/|\\]*(?:/|\\)([0-9]*)`)

var zipVPathRegExp = regexp.MustCompile(`\.zip[/|\\]`)

func resolveVirtualPath(vpath string) (path string) {
	path = vpath

	var vSec = yarnVPathRegExp.FindString(vpath)
	if len(vSec) > 0 {
		var n, _ = strconv.ParseUint(yarnVPathRegExp.FindStringSubmatch(vSec)[1], 10, 64)

		var paths = yarnVPathRegExp.Split(vpath, 2)
		var basePath = paths[0]
		for i := uint64(0); i < n; i++ {
			basePath = filepath.Join(basePath, "..")
		}
		var pkgPath = paths[1]

		path = filepath.Join(basePath, pkgPath)
	}

	return
}

func vreadFile(path string) (data []byte, err error) {
	path = resolveVirtualPath(path)

	var vSec = zipVPathRegExp.FindStringIndex(path)

	if vSec != nil {
		var zipPathStart = vSec[1]

		var pathToZip = path[:zipPathStart-1]
		var pathInZip = normalizeZipPath(path[zipPathStart:])

		var zipReader *zip.ReadCloser
		var destFile fs.File
		var destFileInfo fs.FileInfo

		zipReader, err = zip.OpenReader(pathToZip)
		if err != nil {
			return
		}
		defer zipReader.Close()

		destFile, err = zipReader.Open(pathInZip)
		if err != nil {
			return
		}
		defer destFile.Close()

		destFileInfo, err = destFile.Stat()
		if err != nil {
			return
		}

		var buffer = make([]byte, destFileInfo.Size())
		_, err = destFile.Read(buffer)
		if err != nil {
			return
		}
		data = buffer
	} else {
		data, err = os.ReadFile(path)
	}

	return
}
