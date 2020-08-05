package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

func IsFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	os.Remove(filename)
	return false
}

func DownloadFile(source string, fileSuffix string, filenameWithSuffix string) (string, error) {

	basePath := "tmp/download/" + filenameWithSuffix
	if FileExist(basePath) {
		return basePath, nil
	}

	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := "tmp" + ".download"
	client := new(http.Client)
	resp, err := client.Get(source)
	if err != nil {
		return "", err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		log.Println("Error: <", err, "> when get file remote size")
		return "", err
	}
	if IsFileExist("tmp", fsize) {
		return "had", nil
	}
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if resp.Body == nil {
		return "", errors.New("body is null")
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, ew := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	if err == nil {
		file.Close()

		filenameWithSuffix := path.Base(source)
		fileWSuffix := path.Ext(filenameWithSuffix)
		filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileWSuffix)

		decodeurl, _ := url.QueryUnescape(filenameOnly)

		newPath := "tmp/download/" + decodeurl + fileSuffix

		log.Println("=== Download file to  " + newPath + "  ===")

		os.Rename(tmpFilePath, newPath)
		return newPath, nil
	}
	return "", err
}
