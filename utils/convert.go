package utils

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ConvertFromCADToPDF(filePath string) string {
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

func ConvertToPDF(filePath string) string {
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

		paths :=  strings.Split(path.Base(filePath), ".")

		var tmp = ""
		for i, n := 0, len(paths) - 1; i < n; i++ {
			tmp += "." + paths[i]
		}
		resultPath := "tmp/convert/" + tmp[1:] + ".pdf"

		return resultPath

	} else {
		return ""
	}
}

func ConvertToImg(filePath string) string {
	fileName := strings.Split(path.Base(filePath), ".")[0]
	//fileDir := path.Dir(filePath)
	fileExt := path.Ext(filePath)
	if fileExt != ".pdf" {
		return ""
	}
	os.Mkdir("tmp/convert/"+fileName, os.ModePerm)
	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		commandName = "cmd"
		params = []string{"/c", "magick", "convert", "-density", "130", filePath, "tmp/convert/" + fileName + "/%d.jpg"}
	} else if runtime.GOOS == "linux" {
		commandName = "convert"
		params = []string{"-density", "130", filePath, "tmp/convert/" + fileName + "/%d.jpg"}
	} else {
		commandName = "convert"
		params = []string{"-density", "130", filePath, "tmp/convert/" + fileName + "/%d.jpg"}
	}
	if _, ok := InteractiveToexec(commandName, params); ok {
		resultPath := "tmp/convert/" + strings.Split(path.Base(filePath), ".")[0]
		if ok, _ := PathExists(resultPath); ok {
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}
