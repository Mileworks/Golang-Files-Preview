package utils

import (
	"path"
	"runtime"
	"strings"
)

func ConvertFromCADToPDF(filePath string) string {
	basePath := "tmp/convert/" + GetFileNameOnly(filePath) + ".dwg.pdf"
	if FileExist(basePath) {
		return basePath
	} else {
		commandName := "java"
		params := []string{"-jar", "cad/cad-preview-addon-1.0-SNAPSHOT.jar", "-s", filePath, "-t", filePath + ".svg"}

		if _, ok := InteractiveToexec(commandName, params); ok {
			resultPath := filePath + ".svg"
			if ok, _ := PathExists(resultPath); ok {
				return ConvertToPDF(resultPath)
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
}

func ConvertToPDF(filePath string) string {

	basePath := "tmp/convert/" + GetFileNameOnly(filePath) + ".pdf"
	if FileExist(basePath) {
		return basePath
	} else {
		commandName := ""
		var params []string
		if runtime.GOOS == "windows" {
			commandName = "cmd"
			params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", "--outdir", "tmp/convert/", filePath}
		} else if runtime.GOOS == "linux" {
			commandName = "libreoffice"
			params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "tmp/convert/", filePath}
		} else { // https://ask.libreoffice.org/en/question/12084/how-to-convert-documents-to-pdf-on-osx/
			commandName = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
			params = []string{"--headless", "--convert-to", "pdf", "--outdir", "tmp/convert/", filePath}
		}

		if _, ok := InteractiveToexec(commandName, params); ok {

			paths := strings.Split(path.Base(filePath), ".")

			var tmp = ""
			for i, n := 0, len(paths)-1; i < n; i++ {
				tmp += "." + paths[i]
			}
			resultPath := "tmp/convert/" + tmp[1:] + ".pdf"
			return resultPath
		} else {
			return ""
		}

	}
}
