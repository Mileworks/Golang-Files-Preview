package utils

import (
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
	AllImageEtx   = []string{".jpeg", ".jpg", ".png", ".gif", ".bmp", ".heic", ".tiff"}
	AllCADEtx     = []string{".dwg", ".dxf"}
	AllAchieveEtx = []string{".tar.gz", ".tar.bzip2", ".tar.xz", ".zip", ".rar", ".tar", ".7z", "br", ".bz2", ".lz4", ".sz", ".xz", ".zstd"}
	AllTxtEtx     = []string{".txt", ".java", ".php", ".py", ".md", ".js", ".css", ".xml", ".log"}
	AllVideoEtx   = []string{".mp4", ".webm", ".ogg"}
)

func FileTypeVerify(url string) (string, string, string) {
	filenameWithSuffix := path.Base(url)   //获取文件名带后缀
	filesuffix := path.Ext(url)            //文件后缀
	if strings.Contains(filesuffix, "?") { //单独测试过
		filesuffix = filesuffix[0:strings.Index(filesuffix, "?")]
		filenameWithSuffix = filenameWithSuffix[0:strings.Index(filenameWithSuffix, "?")]
	}

	if strings.Contains(url, ".pdf") {
		return "pdf", ".pdf", filenameWithSuffix
	}

	for _, x := range AllOfficeEtx {
		if filesuffix == x {
			return "office", x, filenameWithSuffix
		}
	}

	for _, x := range AllImageEtx {
		if strings.Contains(url, x) {
			return "image", x, filenameWithSuffix
		}
	}

	for _, x := range AllCADEtx {
		if strings.Contains(url, x) {
			return "cad", x, filenameWithSuffix
		}
	}

	for _, x := range AllAchieveEtx {
		if strings.Contains(url, x) {
			return "achieve", x, filenameWithSuffix
		}
	}

	for _, x := range AllTxtEtx {
		if strings.Contains(url, x) {
			return "txt", x, filenameWithSuffix
		}
	}

	for _, x := range AllVideoEtx {
		if strings.Contains(url, x) {
			return "video", x, filenameWithSuffix
		}
	}

	return "", "", filenameWithSuffix

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
	log.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

func UnarchiveFiles(file string) {
	log.Println("=== Achieve from " + file + "  ===")
	decompressName := strings.TrimSuffix(path.Base(file), path.Ext(path.Base(file)))
	archiver.Unarchive(file, "tmp/decompress/"+decompressName)
}

func GetFilesFromDirectory(source string) ([]string, string) {

	decompressName := strings.TrimSuffix(path.Base(source), path.Ext(path.Base(source)))
	base := "tmp/decompress/" + decompressName

	files, _ := filepath.Glob(filepath.Join(base, "*"))
	for i := range files {
		// __MACOSX 目录 过滤掉
		if strings.Index(files[i], "__MACOSX") == -1 {
			base = files[i]
			break
		}
	}

	files, _ = filepath.Glob(filepath.Join(base, "*.*"))
	// Mac 过滤__MACOSX 目录 和.DS_Store 文件
	var files_result []string
	for i := range files {
		if strings.Index(files[i], "__MACOSX") == -1 && strings.Index(files[i], ".DS_Store") == -1 {
			// 清空字符串名中空格并重命名
			os.Rename(files[i], strings.Replace(files[i], " ", "", -1))
			files[i] = strings.Replace(files[i], " ", "", -1)

			files_result = append(files_result, files[i])
		}
	}

	return files_result, base[:len(base)-2]
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func GetFileNameOnly(url string) string {

	filenameWithSuffix := path.Base(url)                               //获取文件名带后缀
	fileSuffix := path.Ext(filenameWithSuffix)                         //获取文件后缀
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名

	return filenameOnly
}
