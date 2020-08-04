package controllers

import (
	"errors"
	"github.com/mileworks/plm-files-preview/utils"
	"net/url"
	"path"
	"strings"
)

type PLMFileController struct {
	BaseController
}

type AchieveFiles struct {
	Top      string            `json:"top"`
	TreeData map[string]string `json:"treeData"`
}

func (c *PLMFileController) PLMPreview() {

	// 1、拿当前文件下载地址
	downLoadUrl := c.GetString("url")

	_, err := url.ParseRequestURI(downLoadUrl)
	if err != nil {
		panic(errors.New("请输入正确的预览地址"))
	}

	// 2、下载当前的文件，并通过文件后缀进行预览
	fileType, fileSuffix := utils.FileTypeVerify(downLoadUrl)
	local, _ := utils.DownloadFile(downLoadUrl, fileSuffix)

	if fileType == "achieve" {
		utils.UnarchiveFiles(local)
		files, _ := utils.GetFilesFromDirectory(local)

		// process preview url
		baseUrl := "http://" + c.Ctx.Request.Host + "/api/review?file="
		treeMap := make(map[string]string)
		for _, fp := range files {
			reviewUrl := baseUrl + fp
			treeMap[path.Base(fp)] = reviewUrl
		}

		achieveFiles := &AchieveFiles{
			strings.TrimSuffix(path.Base(local), path.Ext(path.Base(local))),
			treeMap}

		c.SuccessJson(achieveFiles)

		return
	}

	var resultPath string
	switch {
	case fileType == "pdf":
		c.Ctx.Output.ContentType("application/x-pdf")
		resultPath = local

		break
	case fileType == "image":
		c.Ctx.Output.ContentType("image/jpeg")
		resultPath = local

		break
	case fileType == "cad":
		c.Ctx.Output.ContentType("application/x-pdf")
		resultPath = utils.ConvertFromCADToPDF(local)

		break
	case fileType == "office":
		c.Ctx.Output.ContentType("application/x-pdf")
		resultPath = utils.ConvertToPDF(local)

		break
	case fileType == "txt":
		c.Ctx.Output.ContentType("text/plain")
		resultPath = local

		break
	case fileType == "video":
		c.Ctx.Output.ContentType("video/mp4")
		resultPath = local

		break
	default:
		panic(errors.New("文件暂时不支持格式"))
	}

	data, _ := utils.File2Bytes(resultPath)
	c.Ctx.Output.Body(data)
}
