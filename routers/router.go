package routers

import (
	"github.com/astaxie/beego"
	"github.com/mileworks/plm-files-preview/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	//beego.Router("/api/preview", &controllers.PreviewController{}, "get:Preview;post:PreviewFile")
	beego.Router("/api/preview", &controllers.PreviewController{}, "get:Preview")
	beego.Router("/api/getfile", &controllers.PreviewController{}, "get:GetFile")
	beego.Router("/api/review", &controllers.PreviewController{}, "get:AchieveFileForReview")


	//beego.Router("/api/create", &controllers.PreviewController{}, "post:CreatePreview")
	//beego.Router("/api/update", &controllers.PreviewController{}, "put:UpdatePreview")
	//beego.Router("/api/delete", &controllers.PreviewController{}, "delete:DeletePreview")

}
