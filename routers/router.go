// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"EasyTutor/controllers/admincontroller"
	"EasyTutor/controllers/commentcontroller"
	"EasyTutor/controllers/offercontroller"
	"EasyTutor/controllers/requestcontroller"
	"EasyTutor/controllers/searchcontroller"
	"EasyTutor/controllers/storagecontroller"
	"EasyTutor/controllers/teachercontroller"
	"EasyTutor/controllers/usercontroller"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1/easy-tutor",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&usercontroller.UserController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&commentcontroller.CommentController{},
			),
		),
		beego.NSNamespace("/request",
			beego.NSInclude(
				&requestcontroller.RequestController{},
			),
		),
		beego.NSNamespace("/teacher",
			beego.NSInclude(
				&teachercontroller.TeacherController{},
			),
		),
		beego.NSNamespace("/offer",
			beego.NSInclude(
				&offercontroller.OfferController{},
			),
		),
		beego.NSNamespace("/storage",
			beego.NSInclude(
				&storagecontroller.StorageController{},
			),
		),
		beego.NSNamespace("/search",
			beego.NSInclude(
				&searchcontroller.SearchController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&admincontroller.AdminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
