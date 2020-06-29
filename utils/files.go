package utils

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/**
 * 预览文件相关处理
 */

var (
	AllOfficeEtx  = []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}
	AllImageEtx   = []string{".jpg", ".png", ".gif", ".bmp", ".heic", ".tiff"}
	AllCADEtx     = []string{".dwg", ".dxf"}
	AllAchieveEtx = []string{".tar.gz", ".tar.bzip2", ".tar.xz", ".zip", ".rar", ".tar", ".7z", "br", ".bz2", ".lz4", ".sz", ".xz", ".zstd"}
	AllTxtEtx     = []string{".txt", ".java", ".php", ".py", ".md", ".js", ".css"}
)

func FileTypeVerify(url string) (string, string) {
	if strings.Contains(url, ".pdf") {
		return "pdf", ".pdf"
	}

	for _, x := range AllOfficeEtx {
		if strings.Contains(url, x) {
			return "office", x
		}
	}

	for _, x := range AllImageEtx {
		if strings.Contains(url, x) {
			return "image", x
		}
	}

	for _, x := range AllCADEtx {
		if strings.Contains(url, x) {
			return "cad", x
		}
	}

	for _, x := range AllAchieveEtx {
		if strings.Contains(url, x) {
			return "achieve", x
		}
	}

	for _, x := range AllTxtEtx {
		if strings.Contains(url, x) {
			return "txt", x
		}
	}

	return "", ""

}

func File2Bytes(filename string) ([]byte, error) {

	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

func UnarchiveFiles(file string) {
	log.Println("=== Achieve from " + file + "  ===")
	decompressName := strings.TrimSuffix(path.Base(file), path.Ext(path.Base(file)))
	archiver.Unarchive(file, "tmp/decompress/"+decompressName)
}

func GetFilesFromDirectory(source string, uri string) (map[string]string, string) {

	decompressName := strings.TrimSuffix(path.Base(source), path.Ext(path.Base(source)))
	base := "tmp/decompress/" + decompressName + "/" + decompressName
	if IsDir(base) {
		base = "tmp/decompress/" + decompressName + "/" + decompressName + "/*"
	}else{
		base = "tmp/decompress/" + decompressName + "/*"
	}
	files, _ := filepath.Glob(base)
	baseUrl := uri + "/api/preview?previewUrl="

	treeMap := make(map[string]string)
	for _, fp := range files {
		reviewUrl := uri + "/api/review?file=" + fp
		treeMap[path.Base(fp)] = baseUrl + reviewUrl
	}

	return treeMap, base[:len(base)-2]
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}