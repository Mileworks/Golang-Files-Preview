package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/mileworks/plm-files-preview/utils"
	"net/url"
	"path"
	"strings"
)

type PreviewController struct {
	beego.Controller
}

func (c *PreviewController) Preview() {
	c.Data["plm"] = "PLM 文件预览服务"

	previewUrl := c.GetString("previewUrl")

	_, err := url.ParseRequestURI(previewUrl)
	if err != nil {
		panic(errors.New("请输入正确的预览url 地址"))
	}

	fileType, fileSuffix, filenameWithSuffix := utils.FileTypeVerify(previewUrl)
	host := c.Ctx.Request.Host

	switch {
	case fileType == "pdf":
		c.Data["url"] = previewUrl
		c.TplName = "preview.tpl"
		break

	case fileType == "image":
		c.Data["url"] = previewUrl
		c.TplName = "image.tpl"
		break

	case fileType == "cad":

		local, _ := utils.DownloadFile(previewUrl, fileSuffix, filenameWithSuffix)
		resultPath := utils.ConvertFromCADToPDF(local)

		c.Data["url"] = "http://" + host + "/api/getfile?file=" + resultPath
		c.TplName = "preview.tpl"

		break
	case fileType == "office":

		local, _ := utils.DownloadFile(previewUrl, fileSuffix, filenameWithSuffix)
		resultPath := utils.ConvertToPDF(local)

		c.Data["url"] = "http://" + host + "/api/getfile?file=" + resultPath
		c.TplName = "preview.tpl"

		break
	case fileType == "achieve":

		local, _ := utils.DownloadFile(previewUrl, fileSuffix, filenameWithSuffix)
		utils.UnarchiveFiles(local)

		files, base := utils.GetFilesFromDirectory(local)

		// process preview url
		uri := "http://" + c.Ctx.Request.Host
		baseUrl := uri + "/api/preview?previewUrl="
		treeMap := make(map[string]string)
		for _, fp := range files {
			reviewUrl := uri + "/api/review?file=" + fp
			treeMap[path.Base(fp)] = baseUrl + reviewUrl
		}

		c.Data["base"] = base
		c.Data["top"] = strings.TrimSuffix(path.Base(local), path.Ext(path.Base(local)))
		c.Data["treeData"] = treeMap
		c.TplName = "achieve.tpl"

		break
	case fileType == "txt":
		local, _ := utils.DownloadFile(previewUrl, fileSuffix, filenameWithSuffix)

		c.Data["url"] = "http://" + host + "/api/getfile?file=" + local
		c.TplName = "preview.tpl"
		break

	case fileType == "video":
		c.Data["url"] = previewUrl
		c.TplName = "video.tpl"
		break

	default:
		panic(errors.New("文件暂时不支持格式"))
	}

}

// 2 method below only for test
func (c *PreviewController) GetFile() {
	fName := c.GetString("file")

	c.Ctx.Output.ContentType("application/x-pdf")
	data, _ := utils.File2Bytes(fName)
	c.Ctx.Output.Body(data)

}

func (c *PreviewController) AchieveFileForReview() {
	fName := c.GetString("file")

	fileType, _, _ := utils.FileTypeVerify(fName)

	switch {
	case fileType == "pdf":
		c.Ctx.Output.ContentType("application/x-pdf")
		break

	case fileType == "image":
		c.Ctx.Output.ContentType("image/jpeg")
		break

	case fileType == "cad":
		c.Ctx.Output.ContentType("application/x-pdf")
		break

	case fileType == "office":
		c.Ctx.Output.ContentType("application/x-pdf")
		break

	case fileType == "txt":
		c.Ctx.Output.ContentType("text/plain")
		break

	case fileType == "video":
		c.Ctx.Output.ContentType("video/mp4")
		break

	default:
		panic(errors.New("文件暂时不支持格式"))
	}

	data, _ := utils.File2Bytes(fName)
	c.Ctx.Output.Body(data)

}
